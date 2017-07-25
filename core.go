package opencv

/*
#cgo linux pkg-config: opencv
#cgo darwin pkg-config: opencv
#include "cpp/core/core.h"
#include "stdlib.h"

*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

//const
const (
	CV_8U   = C.CV_8U
	CV_8S   = C.CV_8S
	CV_16U  = C.CV_16U
	CV_16S  = C.CV_16S
	CV_32S  = C.CV_32S
	CV_32F  = C.CV_32F
	CV_64F  = C.CV_64F
	CV_8UC1 = C.CV_8UC1
	CV_8UC2 = C.CV_8UC2
	CV_8UC3 = C.CV_8UC3
	CV_8UC4 = C.CV_8UC4
)

//Point
type Point struct {
	ptr  unsafe.Pointer
	x, y int
}

func newPoint(i unsafe.Pointer, x, y int) *Point {
	p := &Point{i, x, y}
	runtime.SetFinalizer(p, deletePoint)
	return p
}

func deletePoint(p *Point) {
	if p.ptr != nil {
		C.free(p.ptr)
		p.ptr = nil
	}
}

func NewPoint(x, y int) *Point {
	p := C.MyPoint(C.int(x), C.int(y))
	return newPoint(p, x, y)
}

func (p *Point) Pointer() unsafe.Pointer {
	return p.ptr
}

//Point2f
type Point2f struct {
	ptr  unsafe.Pointer
	x, y float32
}

func newPoint2f(i unsafe.Pointer, x, y float32) *Point2f {
	p := &Point2f{i, x, y}
	runtime.SetFinalizer(p, deletePoint2f)
	return p
}

func deletePoint2f(p *Point2f) {
	if p.ptr != nil {
		C.free(p.ptr)
		p.ptr = nil
	}
}

func NewPoint2f(x, y float32) *Point2f {
	p := C.MyPoint2f(C.float(x), C.float(y))
	return newPoint2f(p, x, y)
}

//Size
type Size struct {
	ptr           unsafe.Pointer
	width, height int
}

func newSize(i unsafe.Pointer, width, height int) *Size {
	s := &Size{i, width, height}
	runtime.SetFinalizer(s, deleteSize)
	return s
}
func deleteSize(s *Size) {
	if s.ptr != nil {
		C.free(s.ptr)
		s.ptr = nil
	}
}

func NewSize(width, height int) *Size {
	s := C.MySize(C.int(width), C.int(height))
	return newSize(s, width, height)
}

func (s *Size) Pointer() unsafe.Pointer {
	return s.ptr
}

//Size2f
type Size2f struct {
	ptr           unsafe.Pointer
	width, height float32
}

func newSize2f(i unsafe.Pointer, width, height float32) *Size2f {
	s := &Size2f{i, width, height}
	runtime.SetFinalizer(s, deleteSize2f)
	return s
}
func deleteSize2f(s *Size2f) {
	if s.ptr != nil {
		C.free(s.ptr)
		s.ptr = nil
	}
}

func NewSize2f(width, height float32) *Size2f {
	s := C.MySize2f(C.float(width), C.float(height))
	return newSize2f(s, width, height)
}

//Rect
type Rect struct {
	ptr unsafe.Pointer
}

func newRect(r unsafe.Pointer) *Rect {
	rect := &Rect{r}
	runtime.SetFinalizer(rect, deleteRect)
	return rect
}
func deleteRect(r *Rect) {
	if r.ptr != nil {
		C.free(r.ptr)
		r.ptr = nil
	}
}

func NewRect() *Rect {
	rect := C.MyRect()
	return newRect(rect)
}

func (r *Rect) Pointer() unsafe.Pointer {
	return r.ptr
}

//RotatedRect
type RotatedRect struct {
	ptr unsafe.Pointer
}

func newRotatedRect(i unsafe.Pointer) *RotatedRect {
	r := &RotatedRect{ptr: i}
	runtime.SetFinalizer(r, deleteRotatedRect)
	return r
}
func deleteRotatedRect(r *RotatedRect) {
	if r.ptr != nil {
		C.free(r.ptr)
		r.ptr = nil
	}
}

func NewRotatedRect() *RotatedRect {
	r := C.MyRotatedRect()
	return newRotatedRect(r)
}

func NewRotatedRectCenterSize(center *Point2f, size *Size2f, angle float32) *RotatedRect {
	r := C.RotatedRectCenterSize(center.ptr, size.ptr, C.float(angle))
	return newRotatedRect(r)
}

func (r *RotatedRect) BoundingRect() *Rect {
	rect := C.RotatedRect_boundingRect(r.ptr)
	return newRect(rect)
}

//Scalar
type Scalar struct {
	ptr        unsafe.Pointer
	v0, v1, v2 float64
}

func newScalar(i unsafe.Pointer, v ...float64) *Scalar {
	s := &Scalar{ptr: i}
	switch len(v) {
	case 1:
		s.v0 = v[0]
	case 2:
		s.v0 = v[0]
		s.v1 = v[1]
	case 3:
		s.v0 = v[0]
		s.v1 = v[1]
		s.v2 = v[2]
	}
	runtime.SetFinalizer(s, deleteScalar)
	return s
}
func deleteScalar(s *Scalar) {
	if s.ptr != nil {
		C.free(s.ptr)
		s.ptr = nil
	}
}

func NewScalar() *Scalar {
	s := C.MyScalar0()
	return newScalar(s)
}

func NewScalar1(v0 float64) *Scalar {
	s := C.MyScalar1(C.double(v0))
	return newScalar(s, v0)
}

func NewScalar2(v0, v1 float64) *Scalar {
	s := C.MyScalar2(C.double(v0), C.double(v1))
	return newScalar(s, v0, v1)
}

func NewScalar3(v0, v1, v2 float64) *Scalar {
	s := C.MyScalar3(C.double(v0), C.double(v1), C.double(v2))
	return newScalar(s, v0, v1, v2)
}

func NewScalarWithAll(v float64) *Scalar {
	s := C.MyScalarWithAll(C.double(v))
	return newScalar(s)
}

func (s *Scalar) Pointer() unsafe.Pointer {
	return s.ptr
}

//MatND
type MatND Mat

//Mat
type Mat struct {
	ptr unsafe.Pointer
}

func newMat(i unsafe.Pointer) *Mat {
	mat := &Mat{ptr: i}
	runtime.SetFinalizer(mat, deleteMat)
	return mat
}

func deleteMat(m *Mat) {
	if m.ptr != nil {
		C.free(m.ptr)
		m.ptr = nil
	}
}

func NewMat() *Mat {
	mat := C.MyMat()
	return newMat(mat)
}
func NewMatWithMat(i unsafe.Pointer) *Mat {
	mat := C.MyMatWithMat(i)
	return newMat(mat)
}
func NewMatRowsColsType(rows, cols, typ int, s *Scalar) *Mat {
	mat := C.MyMatRowsColsType(C.int(rows), C.int(cols), C.int(typ), s.ptr)
	return newMat(mat)
}
func NewMatZeros(rows, cols, typ int) *Mat {
	mat := C.MyMatZeros(C.int(rows), C.int(cols), C.int(typ))
	return newMat(mat)
}
func NewMatZerosWithSize(size *Size, typ int) *Mat {
	mat := C.MyMatZerosWithSize(size.Pointer(), C.int(typ))
	return newMat(mat)
}

func (m *Mat) Flags() int {
	return int(C.Mat_flags(m.ptr))
}
func (m *Mat) Dims() int {
	return int(C.Mat_dims(m.ptr))
}
func (m *Mat) Rows() int {
	return int(C.Mat_rows(m.ptr))
}

func (m *Mat) Cols() int {
	return int(C.Mat_cols(m.ptr))
}

func (m *Mat) Data() []byte {
	pointer := unsafe.Pointer(C.Mat_data(m.ptr))
	size := m.Total() * m.ElemSize()
	if pointer == nil {
		return nil
	}
	return C.GoBytes(pointer, C.int(size))
}

func (m *Mat) Total() int {
	return int(C.Mat_total(m.ptr))
}

func (m *Mat) ElemSize() int {
	return int(C.Mat_elemSize(m.ptr))
}

func (m *Mat) Pointer() unsafe.Pointer {
	return m.ptr
}

func (m *Mat) String() (s string) {
	s += fmt.Sprintf("total:%v\n", m.Total())
	s += fmt.Sprintf("elemsize:%v\n", m.ElemSize())
	s += fmt.Sprintf("rows:%v\n", m.Rows())
	s += fmt.Sprintf("cols:%v\n", m.Cols())
	s += fmt.Sprintf("data:%v\n", len(m.Data()))
	return s
}

//InputArray
type InputArray struct {
	ptr unsafe.Pointer
}

func newInputArray(i unsafe.Pointer) *InputArray {
	array := &InputArray{ptr: i}
	runtime.SetFinalizer(array, deleteInputArray)
	return array
}

func deleteInputArray(i *InputArray) {
	if i.ptr != nil {
		C.free(i.ptr)
		i.ptr = nil
	}
}

func NewInputArray() *InputArray {
	i := C.MyInputArray()
	return newInputArray(i)
}

func NewInputArrayWithMat(mat *Mat) *InputArray {
	i := C.MyInputArrayWithMat(mat.Pointer())
	return newInputArray(i)
}

func (i *InputArray) Pointer() unsafe.Pointer {
	return i.ptr
}

//OutputArray
type OutputArray struct {
	ptr unsafe.Pointer
}

func newOutputArray(i unsafe.Pointer) *OutputArray {
	array := &OutputArray{ptr: i}
	runtime.SetFinalizer(array, deleteOutputArray)
	return array
}

func deleteOutputArray(i *OutputArray) {
	if i.ptr != nil {
		C.free(i.ptr)
		i.ptr = nil
	}
}
func NewOutputArray() *OutputArray {
	i := C.MyOutputArray()
	return newOutputArray(i)
}
func NewOutputArrayWithMat(mat *Mat) *OutputArray {
	i := C.MyOutputArrayWithMat(mat.Pointer())
	return newOutputArray(i)
}

func (i *OutputArray) Pointer() unsafe.Pointer {
	return i.ptr
}

//SparseMat
type SparseMat struct {
	ptr unsafe.Pointer
}

func newSparseMat(i unsafe.Pointer) *SparseMat {
	s := &SparseMat{ptr: i}
	runtime.SetFinalizer(s, deleteSparseMat)
	return s
}

func deleteSparseMat(i *SparseMat) {
	if i.ptr != nil {
		C.free(i.ptr)
		i.ptr = nil
	}
}
func NewSparseMat() *SparseMat {
	i := C.MyOutputArray()
	return newSparseMat(i)
}

func NewSparseMatWithMat(mat *Mat) *SparseMat {
	i := C.MySparseMatWithMat(mat.Pointer())
	return newSparseMat(i)
}

func (s *SparseMat) Poniter() unsafe.Pointer {
	return s.ptr
}

//Operations on Arrays
func MinMaxLocWithInputArray(mat *Mat, minVal, maxVal *float64) {
	var minv, maxv C.double
	inputarray := NewInputArrayWithMat(mat)
	C.minMaxLocWithInputArray(inputarray.Pointer(), &minv, &maxv)
	if minVal != nil {
		*minVal = float64(minv)
	}
	if maxVal != nil {
		*maxVal = float64(maxv)
	}
}
func Normalize(src *InputArray, dst *OutputArray) {
	C.normalize(src.Pointer(), dst.Pointer())
}
func Subtract(src1, src2 *InputArray, dst *OutputArray) {
	C.subtract(src1.Pointer(), src2.Pointer(), dst.Pointer())
}

//Drawing functions
func Line(img *Mat, p1, p2 *Point, color *Scalar) {
	C.line(img.Pointer(), p1.Pointer(), p2.Pointer(), color.Pointer())
}

func Ellipse(img *Mat, center *Point, axes *Size, angle, startAngle, endAngle float64, color *Scalar) {
	C.ellipse(img.Pointer(), center.Pointer(), axes.Pointer(),
		C.double(angle), C.double(startAngle), C.double(endAngle),
		color.Pointer())
}

func Circle(img *Mat, center *Point, radius int, color *Scalar) {
	C.circle(img.Pointer(), center.Pointer(), C.int(radius), color.Pointer())
}

func Rectangle(img *Mat, p1, p2 *Point, color *Scalar) {
	C.rectangle(img.Pointer(), p1.Pointer(), p2.Pointer(), color.Pointer())
}

func RectangleWithRect(img *Mat, rect *Rect, color *Scalar) {
	C.rectangleWithRect(img.Pointer(), rect.Pointer(), color.Pointer())
}

//Utility and System Functions and Macros
func GetBuildInformation() string {
	return C.GoString(C.GetBuildInformation())
}

func GetNumberOfCPUS() int {
	return int(C.GetNumberOfCPUs())
}

func CvRound(value float64) int {
	return int(C.cvRound(C.double(value)))
}
