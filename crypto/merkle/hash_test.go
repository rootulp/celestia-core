package merkle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Fuzz_innerHash(f *testing.F) {
	// Add some seed corpus
	f.Add([]byte{}, []byte{})
	f.Add([]byte{1, 2, 3}, []byte{4, 5, 6})
	f.Add([]byte{255}, []byte{0})

	f.Fuzz(func(t *testing.T, left, right []byte) {
		v034 := innerHashOld(left, right)
		v038 := innerHash(left, right)

		if !assert.Equal(t, v034, v038) {
			fmt.Printf("left: %#v\nright: %#v\n", left, right)
			fmt.Printf("innerHashOld: %#v\ninnerHash: %#v\n", v034, v038)
			t.Logf("Found mismatch with inputs:\nleft: %#v\nright: %#v\ninnerHashOld: %x\ninnerHash: %x",
				left, right, v034, v038)
		}
	})
}

func Test_innerHash(t *testing.T) {
	left := []byte{0x91, 0xd9, 0x34, 0x34, 0x31}
	right := []byte{0x91, 0xd9, 0x34, 0x34, 0x31, 0x91}

	v034 := innerHashOld(left, right)
	v038 := innerHash(left, right)

	assert.Equal(t, v034, v038)
}
