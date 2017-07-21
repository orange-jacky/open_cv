package gui

/*
#cgo linux pkg-config: opencv
#cgo darwin pkg-config: opencv
#include "gui.h"
#include "stdlib.h"
*/
import "C"
import (
	. "github.com/orange-jacky/open_cv/core"
	//"runtime"
	"unsafe"
)

//User Interface
func Imshow(winname string, inputArray *InputArray) {
	c_str := C.CString(winname)
	defer C.free(unsafe.Pointer(c_str))
	C.imshow(c_str, inputArray.Pointer())
}

func Waitkey(delay int) int {
	return int(C.waitkey(C.int(delay)))
}

//Reading and Writing Images and Video
func Imread(filename string) *Mat {
	c_str := C.CString(filename)
	defer C.free(unsafe.Pointer(c_str))
	ptr := C.imread(c_str)
	return NewMatWithMat(ptr)
}
