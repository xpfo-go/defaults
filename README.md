# defaults
Set  default value of golang. 
Set default value for struct tag, Support type of uint, *uint, int, *int, 
string, *string, float, *float, *bool. Nonsupport nest struct. function input params type of pointer.

# install
go get github.com/xpfo-go/defaults

# using defaults

```go
package main

import "github.com/xpfo-go/defaults"

func main(){
	type args struct {
		Float64Ptr *float64 `default:"1.64"`
		StringVal  string   `default:"str"`
	}

	value := args{}

	// default tag
	defaults.StructOpt().FormatStruct(&value)
	
	// if use self-defined tag
	defaults.StructOpt("TagName").FormatStruct(&value)
}
```
