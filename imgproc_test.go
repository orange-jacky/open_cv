package opencv

import (
	//"fmt"
	"testing"
)

func TestExam2(t *testing.T) {
	//例子url:
	//http://docs.opencv.org/2.4.13.2/modules/imgproc/doc/histograms.html

	hsv := NewMat()
	//fmt.Println("hsv", hsv.String())
	src := Imread("testdata/100100.png")
	//fmt.Println("src", src.String())
	CvtColor(NewInputArrayWithMat(src), NewOutputArrayWithMat(hsv), CV_BGR2HSV)
	//fmt.Println("hsv", hsv.String())

	hbins, sbins := 30, 32

	histSize := []int{hbins, sbins}
	hranges := []float32{0, 180}
	sranges := []float32{0, 256}
	ranges := [][]float32{hranges, sranges}
	channels := []int{0, 1}
	mat := NewMat()
	hist := NewMat()
	CalcHist(hsv, 1, channels, NewInputArrayWithMat(mat), NewOutputArrayWithMat(hist),
		2, histSize, ranges)
	maxVal := float64(0)
	MinMaxLocWithInputArray(src, nil, &maxVal)
	//fmt.Println("maxval1", maxVal)

	var scale int = 10
	histImg := NewMatZeros(sbins*scale, hbins*10, CV_8UC3)
	//fmt.Println("histImg", histImg)

	//fmt.Println("hist", histImg)
	for h := 1; h < hbins; h++ {
		for s := 2; s < sbins; s++ {
			binVal := float32(100)
			intensity := CvRound(float64(binVal) * 255 / maxVal)
			Rectangle(histImg,
				NewPoint(h*scale, s*scale),
				NewPoint((h+1)*scale-1, (s+1)*scale-1),
				NewScalarWithAll(float64(intensity)),
			)
		}
	}
	Imshow("source", NewInputArrayWithMat(src))
	Imshow("h-s histogram", NewInputArrayWithMat(histImg))
	Waitkey(0)
}
