package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ensureFiles(t *testing.T, rootDir string, files ...string) {
	for _, f := range files {
		p := rootify(rootDir, f)
		_, err := os.Stat(p)
		assert.Nil(t, err, p)
	}
}

func TestEnsureRoot(t *testing.T) {
	require := require.New(t)

	// setup temp dir for test
	tmpDir, err := ioutil.TempDir("", "config-test")
	require.Nil(err)
	defer os.RemoveAll(tmpDir)

	// create root dir
	EnsureRoot(tmpDir)

	WriteConfigFile(tmpDir, DefaultConfig())

	// make sure config is set properly
	data, err := ioutil.ReadFile(filepath.Join(tmpDir, defaultConfigFilePath))
	require.Nil(err)

	checkConfig(t, string(data))

	ensureFiles(t, tmpDir, "data")
}

func TestEnsureTestRoot(t *testing.T) {
	require := require.New(t)

	testName := "ensureTestRoot"

	// create root dir
	cfg := ResetTestRoot(testName)
	defer os.RemoveAll(cfg.RootDir)
	rootDir := cfg.RootDir

	// make sure config is set properly
	data, err := ioutil.ReadFile(filepath.Join(rootDir, defaultConfigFilePath))
	require.Nil(err)

	checkConfig(t, string(data))

	// TODO: make sure the cfg returned and testconfig are the same!
	baseConfig := DefaultBaseConfig()
	pvConfig := DefaultPrivValidatorConfig()
	ensureFiles(t, rootDir, defaultDataDir, baseConfig.Genesis, pvConfig.Key, pvConfig.State)
}

func checkConfig(t *testing.T, configFile string) {
	t.Helper()
	// list of words we expect in the config
	var elems = []string{
		"moniker",
		"seeds",
		"proxy-app",
		"blocksync",
		"create-empty-blocks",
		"peer",
		"timeout",
		"broadcast",
		"send",
		"addr",
		"wal",
		"propose",
		"max",
		"genesis",
	}
	for _, e := range elems {
		if !strings.Contains(configFile, e) {
			t.Errorf("config file was expected to contain %s but did not", e)
		}
	}
}
