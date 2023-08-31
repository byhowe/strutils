# strutils

`strutils` is a small library written in Go for demonstrating how to write a
package from scratch. It includes small utility involving strings.

---

## Installation

Go to you project root, where `go.mod` exists. Then run the following command in your shell:

``` shell
go get github.com/byhowe/strutils
```

## Usage

``` go
package main

import (
	"fmt"

	"github.com/byhowe/strutils-demo"
)

func main(){
	reversed, err := stringutils.Reverse("byhowe")
	if err != nil {
		log.Fatal(err)
	}    
	fmt.Println(reversed) // ewohyb
}
```

## Contributor(s)

* [Ömer Faruk Çavuş](https://github.com/byhowe)

## License

The project is under MIT license.
