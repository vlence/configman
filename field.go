package configman

import (
	"fmt"

	"github.com/vlence/gossert"
)

type FieldType uint8

const (
        Uint8   FieldType = 1
        Uint16  FieldType = 2
        Uint32  FieldType = 3
        Uint64  FieldType = 4
        Int8    FieldType = 5
        Int16   FieldType = 6
        Int32   FieldType = 7
        Int64   FieldType = 8
        Float32 FieldType = 9
        Float64 FieldType = 10
        Bool    FieldType = 11
)

type Field struct {
        ftype FieldType
        name  string
        value any
}

func NewField(n string, t FieldType) *Field {
        return &Field{name: n, ftype: t}
}

func (f *Field) Name() string {
        return f.name
}

func (f *Field) IsUint8() bool {
        return f.ftype == Uint8
}

func (f *Field) Uint8() uint8 {
        v, ok := f.value.(uint8)
        gossert.Ok(f.IsUint8() && ok, fmt.Sprintf("%s is not a Uint8 field", f.name))
        return v
}

func (f *Field) Set(v any) bool {
        switch f.ftype {
        case Uint8, Uint16:
                f.value = v
                return true
        default:
                return false
        }
}
