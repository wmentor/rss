
Simple RSS reader library for Go

# Install

```
go get github.com/wmentor/rss
```

# Usage

```go
package main

import (
  "fmt"

  "github.com/wmentor/rss"
)

func main() {

  list, err := rss.Get("https://morphs.ru/rss.xml")
  if err != nil {
    panic(err)
  }

  for _, rec := range list {
    fmt.Printf("Title=%s\nLink=%s\nDecription=%s\n\n", rec.Title, rec.Link, rec.Description)
  }
}
```
