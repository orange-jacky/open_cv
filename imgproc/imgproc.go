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
func CalcHist(images *Mat, nimages int, channels []int, mask *InputArray, hist *OutputArray, dims int, histSize []int, ranges [][]float32) {
	c_ranges := make([]*C.float, len(ranges))
	c_ranges_ptr := (**C.float)(unsafe.Pointer(&c_ranges[0]))
	for i, range_i := range ranges {
		c_ranges[i] = (*C.float)(unsafe.Pointer(&range_i[0]))
	}

	C.calcHist(images.Pointer(),
		C.int(nimages),
		(*C.int)(unsafe.Pointer(&channels[0])),
		mask.Pointer(), hist.Pointer(),
		C.int(dims),
		(*C.int)(unsafe.Pointer(&histSize[0])),
		c_ranges_ptr)
}

func CalcHistWithSparseMat(images *Mat, nimages int, channels []int, mask *InputArray, hist *SparseMat, dims int, histSize []int, ranges [][]float32) {
	c_ranges := make([]*C.float, len(ranges))
	c_ranges_ptr := (**C.float)(unsafe.Pointer(&c_ranges[0]))
	for i, _ := range ranges {
		c_ranges[i] = (*C.float)(&ranges[i][0])
	}

	C.calcHistWithSparseMat(images.Pointer(),
		C.int(nimages),
		(*C.int)(unsafe.Pointer(&channels[0])),
		mask.Pointer(), hist.Poniter(),
		C.int(dims),
		(*C.int)(unsafe.Pointer(&histSize[0])),
		c_ranges_ptr)
}

//Miscellaneous Image Transformations
func CvtColor(src *InputArray, dst *OutputArray, code int) {
	C.cvtColor(src.Pointer(), dst.Pointer(), C.int(code))
}
