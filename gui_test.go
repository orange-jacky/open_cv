package opencv

import (
	//"fmt"
	"testing"
)

func TestExam(t *testing.T) {
	//例子url:
	//http://docs.opencv.org/2.4.13.2/modules/core/doc/basic_structures.html#vec
	mat := NewMatRowsColsType(200, 200, CV_8UC3, NewScalar1(0))
	//fmt.Println(mat.String())
	//1
	Rectangle(mat, NewPoint(100, 100), NewPoint(50, 50), NewScalar3(255, 0, 0))
	//2
	rRect := NewRotatedRectCenterSize(NewPoint2f(100, 100), NewSize2f(100, 50), 30)
	brect := rRect.BoundingRect()
	RectangleWithRect(mat, brect, NewScalar3(0, 255, 0))
	Imshow("test", NewInputArrayWithMat(mat))
	Waitkey(0)
}
