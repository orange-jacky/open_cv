package imgproc

/*
#cgo linux pkg-config: opencv
#cgo darwin pkg-config: opencv
#include "imgproc.h"
#include "stdlib.h"
*/
import "C"
import (
	. "github.com/orange-jacky/open_cv/core"
	//"runtime"
	"unsafe"
)

/** Constants for color conversion */
const (
	CV_BGR2GRAY = C.CV_BGR2GRAY
	CV_BGR2HSV  = C.CV_BGR2HSV
)

//Histograms
func CalcHist(images *Mat, nimages int, channels []int, mask *InputArray, hist *OutputArray, dims int, histSize []int, ranges [][]float64) {
	C.calcHist(images.Pointer(),
		C.int(nimages),
		(*C.int)(unsafe.Pointer(&channels[0])),
		mask.Pointer(), hist.Pointer(),
		C.int(dims),
		(*C.int)(unsafe.Pointer(&histSize[0])),
		(**C.float)(unsafe.Pointer(&ranges[0])))
}

func CalcHistWithSparseMat(images *Mat, nimages int, channels []int, mask *InputArray, hist *SparseMat, dims int, histSize []int, ranges [][]float64) {
	C.calcHistWithSparseMat(images.Pointer(),
		C.int(nimages),
		(*C.int)(unsafe.Pointer(&channels[0])),
		mask.Pointer(), hist.Poniter(),
		C.int(dims),
		(*C.int)(unsafe.Pointer(&histSize[0])),
		(**C.float)(unsafe.Pointer(&ranges[0])))
}
