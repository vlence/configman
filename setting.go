package configman

import (
        "fmt"

        "github.com/vlence/gossert"
)

type ValueType uint8

const (
        Uint8   ValueType = 1
        Uint16  ValueType = 2
        Uint32  ValueType = 3
        Uint64  ValueType = 4
        Int8    ValueType = 5
        Int16   ValueType = 6
        Int32   ValueType = 7
        Int64   ValueType = 8
        Float32 ValueType = 9
        Float64 ValueType = 10
        Bool    ValueType = 11
)

type Setting struct {
        ftype      ValueType
        name       string
        desc       string
        value      any
        deprecated bool
}

func NewSetting(n string, t ValueType) *Setting {
        return &Setting{name: n, ftype: t}
}

func (f *Setting) Name() string {
        return f.name
}

func (f *Setting) IsUint8() bool {
        return f.ftype == Uint8
}

func (f *Setting) Uint8() uint8 {
        v, ok := f.value.(uint8)
        gossert.Ok(f.IsUint8() && ok, fmt.Sprintf("%s is not a Uint8 field", f.name))
        return v
}

func (f *Setting) Set(v any) bool {
        switch f.ftype {
        case Uint8, Uint16:
                f.value = v
                return true
        default:
                return false
        }
}
