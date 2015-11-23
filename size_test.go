package mmm

import (
	"testing"
	"unsafe"
)

// -----------------------------------------------------------------------------

func TestSize_SizeOf_bool(t *testing.T) {
	var v bool
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for bool")
	}
}

func TestSize_SizeOf_int(t *testing.T) {
	var v int
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for int")
	}
}

func TestSize_SizeOf_int8(t *testing.T) {
	var v int8
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for int8")
	}
}

func TestSize_SizeOf_int16(t *testing.T) {
	var v int16
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for int16")
	}
}

//var i16 int16
//var i32 int32
//var i64 int64
//var ui uint
//var ui8 uint8
//var ui16 uint16
//var ui32 uint32
//var ui64 uint64
//var uiptr uintptr
//var f32 float32
//var f64 float64
//var c64 complex64
//var c128 complex128
//var iarr [42]int
//var ichan chan int
//var itf interface{}
//var imap map[int]int
//var iptr *int
//var islice []int
//var str string
//var srt struct{}
