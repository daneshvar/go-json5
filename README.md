JSON5 implements encoding and decoding for Go forked from [yosuke-furukawa/json5](https://github.com/yosuke-furukawa/json5)

# Go Modules
```
require (
  github.com/daneshvar/go-json5 v0.1.3
)
```

# Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/daneshvar/go-json5"
)

func main() {
	if f, e := os.Open("./example.json5"); e == nil {
		defer f.Close()
		dec := json5.NewDecoder(f)

		var s interface{} // or struct
		if e := dec.Decode(&s); e == nil {
			fmt.Println(s)
		} else {
			fmt.Println(e)
		}
	} else {
		fmt.Println(e)
	}
}  
```
```js
// This is JSON5 example data
// JSON5 can write comment in your json
[
  {
    name: "Hussein",
    family: "Daneshvar",
    skills: ["C++", "Go", "JavaScript",],
    projects: [{name: "Spring", date: "2015",}, {name: "Hafez", date: "2018",}],
    summery: "I'm a \"C++\" Developer\nand a Go Developer",
  },
  {
    name: "Mehrdad",
    family: "Nourozi",
    skills: ["C#", "JavaScript",],
    projects: [{name: "Bandar", date: "2005"}, {name: "NewSite", date: "2018",}],
    summery: "I'm a \"C#\" Developer",
  },
],
```

# TODO
- Block comment
- Multiline string
- Hexadecimal Notation
- Leading Decimal Point
- Positive Sign
- Single Quotes
