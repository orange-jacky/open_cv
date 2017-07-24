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

	c_ranges := C.makeFloatArray(C.int(len(ranges)))
	defer C.freeFloatArray(c_ranges, C.int(len(ranges)))
	for i, range_i := range ranges {
		array := C.makeFloat(C.int(len(range_i)))
		for ii, val := range range_i {
			C.setFloat(array, C.float(val), C.int(ii))
		}
		C.setFloatArray(c_ranges, array, C.int(i))
	}

	C.calcHist(images.Pointer(),
		C.int(nimages),
		(*C.int)(unsafe.Pointer(&channels[0])),
		mask.Pointer(), hist.Pointer(),
		C.int(dims),
		(*C.int)(unsafe.Pointer(&histSize[0])),
		c_ranges)
}

func CalcHistWithSparseMat(images *Mat, nimages int, channels []int, mask *InputArray, hist *SparseMat, dims int, histSize []int, ranges [][]float32) {
	c_ranges := C.makeFloatArray(C.int(len(ranges)))
	defer C.freeFloatArray(c_ranges, C.int(len(ranges)))
	for i, range_i := range ranges {
		array := C.makeFloat(C.int(len(range_i)))
		for ii, val := range range_i {
			C.setFloat(array, C.float(val), C.int(ii))
		}
		C.setFloatArray(c_ranges, array, C.int(i))
	}

	C.calcHistWithSparseMat(images.Pointer(),
		C.int(nimages),
		(*C.int)(unsafe.Pointer(&channels[0])),
		mask.Pointer(), hist.Poniter(),
		C.int(dims),
		(*C.int)(unsafe.Pointer(&histSize[0])),
		c_ranges)
}

//Miscellaneous Image Transformations
func CvtColor(src *InputArray, dst *OutputArray, code int) {
	C.cvtColor(src.Pointer(), dst.Pointer(), C.int(code))
}
