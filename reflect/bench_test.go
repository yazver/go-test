package main

import (
	"reflect"
	"testing"
	"unsafe"
)

func BenchmarkNewAt(b *testing.B) {
	v := reflect.ValueOf(new(int))
	t := v.Type().Elem()
	ptr := v.Pointer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			_ = reflect.NewAt(t, unsafe.Pointer(ptr)).Elem()
		}
	}
}

func BenchmarkNewAtPtr(b *testing.B) {
	v := reflect.ValueOf(new(int))
	t := v.Type()
	ptr := v.Pointer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			_ = reflect.NewAtPtr(t, unsafe.Pointer(ptr)).Elem()
		}
	}
}

func BenchmarkValueAt(b *testing.B) {
	v := reflect.ValueOf(new(int))
	ptr := v.Pointer()
	t := v.Type().Elem()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			_ = reflect.ValueAt(t, unsafe.Pointer(ptr))
		}
	}
}

// --------------------------------

func BenchmarkNewAtAddr(b *testing.B) {
	v := reflect.ValueOf(new(int))
	t := v.Type().Elem()
	ptr := v.Pointer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			_ = reflect.NewAt(t, unsafe.Pointer(ptr))
		}
	}
}

func BenchmarkNewAtPtrAddr(b *testing.B) {
	v := reflect.ValueOf(new(int))
	t := v.Type()
	ptr := v.Pointer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			_ = reflect.NewAtPtr(t, unsafe.Pointer(ptr))
		}
	}
}

func BenchmarkValueAtAddr(b *testing.B) {
	v := reflect.ValueOf(new(int))
	ptr := v.Pointer()
	t := v.Type().Elem()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			_ = reflect.ValueAt(t, unsafe.Pointer(ptr)).Addr()
		}
	}
}

// ---------------------------------------

func BenchmarkFieldIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := struct {
			I int
		}{
			I: 10,
		}
		dest := src
		dest.I = 0
		vSrc := reflect.ValueOf(&src).Elem()
		vDest := reflect.ValueOf(&dest).Elem()
		for i := 0; i < 1000; i++ {
			vDest.Field(0).Set(vSrc.Field(0))
		}
	}
}

func BenchmarkFieldNewAt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := struct {
			I int
		}{
			I: 10,
		}
		dest := src
		dest.I = 0

		vSrc := reflect.ValueOf(&src)
		srcOffset := vSrc.Elem().Type().Field(0).Offset
		srcFieldType := vSrc.Elem().Type().Field(0).Type
		srcPtr := unsafe.Pointer(vSrc.Pointer())

		vDest := reflect.ValueOf(&dest)
		destOffset := vDest.Elem().Type().Field(0).Offset
		destFieldType := vDest.Elem().Type().Field(0).Type
		destPtr := unsafe.Pointer(vDest.Pointer())

		for i := 0; i < 1000; i++ {
			src := reflect.NewAt(srcFieldType, unsafe.Pointer(uintptr(srcPtr)+srcOffset)).Elem()
			dest := reflect.NewAt(destFieldType, unsafe.Pointer(uintptr(destPtr)+destOffset)).Elem()
			dest.Set(src)
		}
	}
}

func BenchmarkFieldNewAtPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := struct {
			I int
		}{
			I: 10,
		}
		dest := src
		dest.I = 0

		vSrc := reflect.ValueOf(&src)
		srcOffset := vSrc.Elem().Type().Field(0).Offset
		srcFieldType := reflect.PtrTo(vSrc.Elem().Type().Field(0).Type)
		srcPtr := unsafe.Pointer(vSrc.Pointer())

		vDest := reflect.ValueOf(&dest)
		destOffset := vDest.Elem().Type().Field(0).Offset
		destFieldType := reflect.PtrTo(vDest.Elem().Type().Field(0).Type)
		destPtr := unsafe.Pointer(vDest.Pointer())

		for i := 0; i < 1000; i++ {
			src := reflect.NewAtPtr(srcFieldType, unsafe.Pointer(uintptr(srcPtr)+srcOffset)).Elem()
			dest := reflect.NewAtPtr(destFieldType, unsafe.Pointer(uintptr(destPtr)+destOffset)).Elem()
			dest.Set(src)
		}
	}
}

func BenchmarkFieldValueAt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		src := struct {
			I int
		}{
			I: 10,
		}
		dest := src
		dest.I = 0

		vSrc := reflect.ValueOf(&src)
		srcOffset := vSrc.Elem().Type().Field(0).Offset
		srcFieldType := vSrc.Elem().Type().Field(0).Type
		srcPtr := unsafe.Pointer(vSrc.Pointer())
		vDest := reflect.ValueOf(&dest)
		destOffset := vDest.Elem().Type().Field(0).Offset
		destFieldType := vDest.Elem().Type().Field(0).Type
		destPtr := unsafe.Pointer(vDest.Pointer())

		for i := 0; i < 1000; i++ {
			src := reflect.ValueAt(srcFieldType, unsafe.Pointer(uintptr(srcPtr)+srcOffset))
			dest := reflect.ValueAt(destFieldType, unsafe.Pointer(uintptr(destPtr)+destOffset))
			dest.Set(src)
		}
	}
}
