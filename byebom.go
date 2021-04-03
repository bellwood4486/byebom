package byebom

import (
	"bufio"
	"fmt"
	"io"
)

// Normalize normalizes the contents of the io.Reader.
func Normalize(r io.Reader) io.Reader {
	if r == nil {
		return r
	}

	br := bufio.NewReader(r)
	bs, err := br.Peek(3)
	if err != nil {
		return br
	}
	if bs[0] == 0xEF && bs[1] == 0xBB && bs[2] == 0xBF {
		if _, err := br.Discard(3); err != nil {
			panic(fmt.Sprintf("byebom: %v", err))
		}
	}
	return br
}
