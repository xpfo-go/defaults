package defaults

import "testing"

func TestStructOpt(t *testing.T) {
	type args struct {
		Float32Val float32  `default:"0.32"`
		Float64Val float64  `default:"0.64"`
		Float32Ptr *float32 `default:"1.32"`
		Float64Ptr *float64 `default:"1.64"`
		StringVal  string   `default:"asd"`
		StringPtr  *string  `default:"qwe"`
		UintVal    uint8    `default:"11"`
		UintPtr    *uint8   `default:"22"`
		IntVal     int8     `default:"27"` // keep default dont out type range
		IntPtr     *int8    `default:"72"`
		// nonsupport type of bool, it is recommended that bool is *bool
		BoolPtr *bool `default:"true"` // support 1,t,T,TRUE,True,true,0,f,F,FALSE,False,false
	}

	value := args{
		Float32Val: 5.32,
	}
	StructOpt().FormatStruct(&value)
	t.Log(value)
}
