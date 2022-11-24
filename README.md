# Read file go

<!--toc:start-->
- [Read file go](#read-file-go)
  - [How to use?](#how-to-use)
  - [Example](#example)
<!--toc:end-->

## How to use?

The project structure should look like this, write code in `./src`:

```txt
.
├── README.md
└── src
    └── lib
        └── read.go
```

To import it in an application which is in `./src`:

```go
import Read "read/src/lib"
```

## Example

In a file which is located in `./src` lets call it `main.go`:

```go
package main

import (
 "fmt"
 Read "read/src/lib"
)

func main() {
 path := "./testdata/read/1.txt"
 file, err := Read.ReadFile(path)
 if err != nil {
  fmt.Printf("The error of read is: %s\n", err)
 }
 var lineCont []byte
 fmt.Printf(string(rune(len(file))) + "\n")
 for i := 0; i < len(file); i++ {
  if string(file[i]) == "\n" {
   fmt.Printf("{%s}, i={%d} \n", string(lineCont), i)
   lineCont = nil
  } else {
   lineCont = append(lineCont, file[i])
  }
 }
}
```

For this example you need to have a file located in the current `$cwd/testdata/read/1.txt`.
So the tree for the given example would look like this:

```txt
.
├── README.md
├── src
│   └── lib
│       └── read.go
└── testdata
    └── read
        └── 1.txt
```

The content of `1.txt` is: `Hello, World!`. So the expected output is: `Hello, World!`

