//nolint: gosec
package app

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"sync"
)

// State is the application state.
type State struct {
	sync.RWMutex
	Height uint64
	Values map[string]string
	Hash   []byte

	// private fields aren't marshaled to disk.
	file            string
	persistInterval uint64
	initialHeight   uint64
}

// NewState creates a new state.
func NewState(file string, persistInterval uint64) (*State, error) {
	state := &State{
		Values:          make(map[string]string),
		file:            file,
		persistInterval: persistInterval,
	}
	state.Hash = hashItems(state.Values)
	err := state.load()
	switch {
	case errors.Is(err, os.ErrNotExist):
	case err != nil:
		return nil, err
	}
	return state, nil
}

// load loads state from disk. It does not take out a lock, since it is called
// during construction.
func (s *State) load() error {
	bz, err := ioutil.ReadFile(s.file)
	if err != nil {
		return fmt.Errorf("failed to read state from %q: %w", s.file, err)
	}
	err = json.Unmarshal(bz, s)
	if err != nil {
		return fmt.Errorf("invalid state data in %q: %w", s.file, err)
	}
	return nil
}

// save saves the state to disk. It does not take out a lock since it is called
// internally by Commit which does lock.
func (s *State) save() error {
	bz, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("failed to marshal state: %w", err)
	}
	// We write the state to a separate file and move it to the destination, to
	// make it atomic.
	newFile := fmt.Sprintf("%v.new", s.file)
	err = ioutil.WriteFile(newFile, bz, 0644)
	if err != nil {
		return fmt.Errorf("failed to write state to %q: %w", s.file, err)
	}
	return os.Rename(newFile, s.file)
}

// Export exports key/value pairs as JSON, used for state sync snapshots.
func (s *State) Export() ([]byte, error) {
	s.RLock()
	defer s.RUnlock()
	return json.Marshal(s.Values)
}

// Import imports key/value pairs from JSON bytes, used for InitChain.AppStateBytes and
// state sync snapshots. It also saves the state once imported.
func (s *State) Import(height uint64, jsonBytes []byte) error {
	s.Lock()
	defer s.Unlock()
	values := map[string]string{}
	err := json.Unmarshal(jsonBytes, &values)
	if err != nil {
		return fmt.Errorf("failed to decode imported JSON data: %w", err)
	}
	s.Height = height
	s.Values = values
	s.Hash = hashItems(values)
	return s.save()
}

// Get fetches a value. A missing value is returned as an empty string.
func (s *State) Get(key string) string {
	s.RLock()
	defer s.RUnlock()
	return s.Values[key]
}

// Set sets a value. Setting an empty value is equivalent to deleting it.
func (s *State) Set(key, value string) {
	s.Lock()
	defer s.Unlock()
	if value == "" {
		delete(s.Values, key)
	} else {
		s.Values[key] = value
	}
}

// Commit commits the current state.
func (s *State) Commit() (uint64, []byte, error) {
	s.Lock()
	defer s.Unlock()
	s.Hash = hashItems(s.Values)
	switch {
	case s.Height > 0:
		s.Height++
	case s.initialHeight > 0:
		s.Height = s.initialHeight
	default:
		s.Height = 1
	}
	if s.persistInterval > 0 && s.Height%s.persistInterval == 0 {
		err := s.save()
		if err != nil {
			return 0, nil, err
		}
	}
	return s.Height, s.Hash, nil
}

// hashItems hashes a set of key/value items.
func hashItems(items map[string]string) []byte {
	keys := make([]string, 0, len(items))
	for key := range items {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	hasher := sha256.New()
	for _, key := range keys {
		_, _ = hasher.Write([]byte(key))
		_, _ = hasher.Write([]byte{0})
		_, _ = hasher.Write([]byte(items[key]))
		_, _ = hasher.Write([]byte{0})
	}
	return hasher.Sum(nil)
}
