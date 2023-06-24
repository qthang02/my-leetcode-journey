package two_pointers

import (
	"fmt"
	"testing"
)

func TestCompress(t *testing.T) {
	tc := []struct {
		input  []byte
		output int
	}{
		{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
		{[]byte{'a'}, 1},
		{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4},
	}

	for _, c := range tc {
		name := fmt.Sprintf("compress(%v)", c.input)

		t.Run(name, func(t *testing.T) {
			got := compress(c.input)
			if got != c.output {
				t.Errorf("got %d, want %d", got, c.output)
			}
		})
	}
}