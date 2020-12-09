// +build go1.6

// Code generated by codecgen - DO NOT EDIT.

package db

import (
	"errors"
	codec1978 "github.com/ugorji/go/codec"
	"runtime"
	"strconv"
)

const (
	// ----- content types ----
	codecSelferCcUTF87046 = 1
	codecSelferCcRAW7046  = 255
	// ----- value types used ----
	codecSelferValueTypeArray7046     = 10
	codecSelferValueTypeMap7046       = 9
	codecSelferValueTypeString7046    = 6
	codecSelferValueTypeInt7046       = 2
	codecSelferValueTypeUint7046      = 3
	codecSelferValueTypeFloat7046     = 4
	codecSelferValueTypeNil7046       = 1
	codecSelferBitsize7046            = uint8(32 << (^uint(0) >> 63))
	codecSelferDecContainerLenNil7046 = -2147483648
)

var (
	errCodecSelferOnlyMapOrArrayEncodeToStruct7046 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer7046 struct{}

func codecSelfer7046False() bool { return false }
func codecSelfer7046True() bool  { return true }

func init() {
	if codec1978.GenVersion != 20 {
		_, file, _, _ := runtime.Caller(0)
		ver := strconv.FormatInt(int64(codec1978.GenVersion), 10)
		panic(errors.New("codecgen version mismatch: current: 20, need " + ver + ". Re-generate file: " + file))
	}
}

func (x *Response) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer7046
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yy2arr2 := z.EncBasicHandle().StructToArray
		_ = yy2arr2
		const yyr2 bool = false // struct tag has 'toArray'
		var yyq2 = [4]bool{     // should field at this index be written?
			x.Time != "",       // time
			x.Status != "",     // status
			x.Detail != "",     // detail
			len(x.Result) != 0, // result
		}
		_ = yyq2
		if yyr2 || yy2arr2 {
			z.EncWriteArrayStart(4)
			z.EncWriteArrayElem()
			if yyq2[0] {
				r.EncodeString(string(x.Time))
			} else {
				r.EncodeString("")
			}
			z.EncWriteArrayElem()
			if yyq2[1] {
				r.EncodeString(string(x.Status))
			} else {
				r.EncodeString("")
			}
			z.EncWriteArrayElem()
			if yyq2[2] {
				r.EncodeString(string(x.Detail))
			} else {
				r.EncodeString("")
			}
			z.EncWriteArrayElem()
			if yyq2[3] {
				if x.Result == nil {
					r.EncodeNil()
				} else {
					z.F.EncSliceIntfV(x.Result, e)
				} // end block: if x.Result slice == nil
			} else {
				r.EncodeNil()
			}
			z.EncWriteArrayEnd()
		} else {
			var yynn2 int
			for _, b := range yyq2 {
				if b {
					yynn2++
				}
			}
			z.EncWriteMapStart(yynn2)
			yynn2 = 0
			if yyq2[0] {
				z.EncWriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"time\"")
				} else {
					r.EncodeString(`time`)
				}
				z.EncWriteMapElemValue()
				r.EncodeString(string(x.Time))
			}
			if yyq2[1] {
				z.EncWriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"status\"")
				} else {
					r.EncodeString(`status`)
				}
				z.EncWriteMapElemValue()
				r.EncodeString(string(x.Status))
			}
			if yyq2[2] {
				z.EncWriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"detail\"")
				} else {
					r.EncodeString(`detail`)
				}
				z.EncWriteMapElemValue()
				r.EncodeString(string(x.Detail))
			}
			if yyq2[3] {
				z.EncWriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"result\"")
				} else {
					r.EncodeString(`result`)
				}
				z.EncWriteMapElemValue()
				if x.Result == nil {
					r.EncodeNil()
				} else {
					z.F.EncSliceIntfV(x.Result, e)
				} // end block: if x.Result slice == nil
			}
			z.EncWriteMapEnd()
		}
	}
}

func (x *Response) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer7046
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	yyct2 := r.ContainerType()
	if yyct2 == codecSelferValueTypeNil7046 {
		*(x) = Response{}
	} else if yyct2 == codecSelferValueTypeMap7046 {
		yyl2 := z.DecReadMapStart()
		if yyl2 == 0 {
		} else {
			x.codecDecodeSelfFromMap(yyl2, d)
		}
		z.DecReadMapEnd()
	} else if yyct2 == codecSelferValueTypeArray7046 {
		yyl2 := z.DecReadArrayStart()
		if yyl2 != 0 {
			x.codecDecodeSelfFromArray(yyl2, d)
		}
		z.DecReadArrayEnd()
	} else {
		panic(errCodecSelferOnlyMapOrArrayEncodeToStruct7046)
	}
}

func (x *Response) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer7046
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if z.DecCheckBreak() {
				break
			}
		}
		z.DecReadMapElemKey()
		yys3 := z.StringView(r.DecodeStringAsBytes())
		z.DecReadMapElemValue()
		switch yys3 {
		case "time":
			x.Time = (string)(string(r.DecodeStringAsBytes()))
		case "status":
			x.Status = (string)(string(r.DecodeStringAsBytes()))
		case "detail":
			x.Detail = (string)(string(r.DecodeStringAsBytes()))
		case "result":
			z.F.DecSliceIntfX(&x.Result, d)
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
}

func (x *Response) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer7046
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	var yyj9 int
	var yyb9 bool
	var yyhl9 bool = l >= 0
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = z.DecCheckBreak()
	}
	if yyb9 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Time = (string)(string(r.DecodeStringAsBytes()))
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = z.DecCheckBreak()
	}
	if yyb9 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Status = (string)(string(r.DecodeStringAsBytes()))
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = z.DecCheckBreak()
	}
	if yyb9 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Detail = (string)(string(r.DecodeStringAsBytes()))
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = z.DecCheckBreak()
	}
	if yyb9 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	z.F.DecSliceIntfX(&x.Result, d)
	for {
		yyj9++
		if yyhl9 {
			yyb9 = yyj9 > l
		} else {
			yyb9 = z.DecCheckBreak()
		}
		if yyb9 {
			break
		}
		z.DecReadArrayElem()
		z.DecStructFieldNotFound(yyj9-1, "")
	}
}

func (x *Response) IsCodecEmpty() bool {
	return !(x.Time != "" || x.Status != "" || x.Detail != "" || len(x.Result) != 0 || false)
}

func (x *Dispatch) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer7046
	z, r := codec1978.GenHelper().Encoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yy2arr2 := z.EncBasicHandle().StructToArray
		_ = yy2arr2
		const yyr2 bool = false // struct tag has 'toArray'
		var yyq2 = [3]bool{     // should field at this index be written?
			x.Query != "",   // query
			x.Action != "",  // action
			x.Result != nil, // result
		}
		_ = yyq2
		if yyr2 || yy2arr2 {
			z.EncWriteArrayStart(3)
			z.EncWriteArrayElem()
			if yyq2[0] {
				r.EncodeString(string(x.Query))
			} else {
				r.EncodeString("")
			}
			z.EncWriteArrayElem()
			if yyq2[1] {
				r.EncodeString(string(x.Action))
			} else {
				r.EncodeString("")
			}
			z.EncWriteArrayElem()
			if yyq2[2] {
				z.EncFallback(x.Result)
			} else {
				r.EncodeNil()
			}
			z.EncWriteArrayEnd()
		} else {
			var yynn2 int
			for _, b := range yyq2 {
				if b {
					yynn2++
				}
			}
			z.EncWriteMapStart(yynn2)
			yynn2 = 0
			if yyq2[0] {
				z.EncWriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"query\"")
				} else {
					r.EncodeString(`query`)
				}
				z.EncWriteMapElemValue()
				r.EncodeString(string(x.Query))
			}
			if yyq2[1] {
				z.EncWriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"action\"")
				} else {
					r.EncodeString(`action`)
				}
				z.EncWriteMapElemValue()
				r.EncodeString(string(x.Action))
			}
			if yyq2[2] {
				z.EncWriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"result\"")
				} else {
					r.EncodeString(`result`)
				}
				z.EncWriteMapElemValue()
				z.EncFallback(x.Result)
			}
			z.EncWriteMapEnd()
		}
	}
}

func (x *Dispatch) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer7046
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	yyct2 := r.ContainerType()
	if yyct2 == codecSelferValueTypeNil7046 {
		*(x) = Dispatch{}
	} else if yyct2 == codecSelferValueTypeMap7046 {
		yyl2 := z.DecReadMapStart()
		if yyl2 == 0 {
		} else {
			x.codecDecodeSelfFromMap(yyl2, d)
		}
		z.DecReadMapEnd()
	} else if yyct2 == codecSelferValueTypeArray7046 {
		yyl2 := z.DecReadArrayStart()
		if yyl2 != 0 {
			x.codecDecodeSelfFromArray(yyl2, d)
		}
		z.DecReadArrayEnd()
	} else {
		panic(errCodecSelferOnlyMapOrArrayEncodeToStruct7046)
	}
}

func (x *Dispatch) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer7046
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if z.DecCheckBreak() {
				break
			}
		}
		z.DecReadMapElemKey()
		yys3 := z.StringView(r.DecodeStringAsBytes())
		z.DecReadMapElemValue()
		switch yys3 {
		case "query":
			x.Query = (string)(string(r.DecodeStringAsBytes()))
		case "action":
			x.Action = (string)(string(r.DecodeStringAsBytes()))
		case "result":
			z.DecFallback(&x.Result, true)
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
}

func (x *Dispatch) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer7046
	z, r := codec1978.GenHelper().Decoder(d)
	_, _, _ = h, z, r
	var yyj8 int
	var yyb8 bool
	var yyhl8 bool = l >= 0
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = z.DecCheckBreak()
	}
	if yyb8 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Query = (string)(string(r.DecodeStringAsBytes()))
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = z.DecCheckBreak()
	}
	if yyb8 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	x.Action = (string)(string(r.DecodeStringAsBytes()))
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = z.DecCheckBreak()
	}
	if yyb8 {
		z.DecReadArrayEnd()
		return
	}
	z.DecReadArrayElem()
	z.DecFallback(&x.Result, true)
	for {
		yyj8++
		if yyhl8 {
			yyb8 = yyj8 > l
		} else {
			yyb8 = z.DecCheckBreak()
		}
		if yyb8 {
			break
		}
		z.DecReadArrayElem()
		z.DecStructFieldNotFound(yyj8-1, "")
	}
}

func (x *Dispatch) IsCodecEmpty() bool {
	return !(x.Query != "" || x.Action != "" || x.Result != nil || false)
}
