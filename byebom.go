package byebom

import (
	"fmt"
	"io"

	"golang.org/x/text/encoding"

	goenc "github.com/mattn/go-encoding"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func NewUTF8Reader(r io.Reader) io.Reader {
	vr := transform.NewReader(r, encoding.UTF8Validator)
	return transform.NewReader(vr, unicode.UTF8BOM.NewDecoder())
}

func NewReaderWithoutBOM(r io.Reader, fallbackEncName string) (io.Reader, error) {
	enc := goenc.GetEncoding(fallbackEncName)
	if enc == nil {
		return nil, fmt.Errorf("byebom: unknown encode %v", fallbackEncName)
	}
	return transform.NewReader(r, unicode.BOMOverride(enc.NewDecoder())), nil
}
