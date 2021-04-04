# byebom

byebom trims Byte Order Mark(BOM) of UTF-8.

## Usage

```go
package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/bellwood4486/byebom"
)

const utf8BOM = "\ufeff"

func main() {
	// with BOM
	withBOM := strings.NewReader(utf8BOM + "a")
	s1, _ := io.ReadAll(byebom.NewUTF8Reader(withBOM))
	fmt.Println(string(s1) == "a") // -> true

	// without BOM
	withoutBOM := strings.NewReader("b")
	s2, _ := io.ReadAll(byebom.NewUTF8Reader(withoutBOM))
	fmt.Println(string(s2) == "b") // -> true

	// error when encoding is not UTF-8
	eucjp := strings.NewReader("\xa5\xa8\xa5\xe9\xa1\xbc")
	_, err := io.ReadAll(byebom.NewUTF8Reader(eucjp))
	fmt.Println(err) // -> encoding: invalid UTF-8
}
```

## LICENSE

MIT
