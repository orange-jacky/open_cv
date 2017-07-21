#ifndef __CORE_H__
#define __CORE_H__

#include <opencv2/core/core.hpp>

#ifdef __cplusplus
extern "C" {
#endif


//basic datatype
typedef void DataType;
typedef void Point_;
typedef void Point3_;
typedef void Size_;
typedef void Rect_;
typedef void RotatedRect;
typedef void TermCriteria;
typedef void Matx;
typedef void Vec;
typedef void Scalar_;
typedef void Range;
typedef void Ptr;
typedef void Mat;
typedef void Mat_;
typedef void InputArray;
typedef void OutputArray;
typedef void NAryMatlerator;
typedef void SparseMat;
typedef void SparseMat_;
typedef void Algorithm;


typedef void Point2f;
typedef void Size2f;
typedef void Rect;
typedef void Scalar;
typedef void Point;
typedef void Size;

//DataType

//Point2f
Point2f* MyPoint2f(float x, float y);
//Size2f
Size2f* MySize2f(float width, float height);

//Point
Point* MyPoint(int x, int y);
//Size
Size* MySize(int x, int y);

//RotatedRect
RotatedRect* MyRotatedRect();
RotatedRect* RotatedRectCenterSize(const Point2f* center, const Size2f* size, float angle);
Rect* RotatedRect_boundingRect(RotatedRect* rect);


//Rect
Rect* MyRect();

//Scalar
Scalar* MyScalar0();
Scalar* MyScalar1(double v0);
Scalar* MyScalar2(double v0, double v1);
Scalar* MyScalar3(double v0, double v1, double v2);

//Mat
Mat* MyMat();
Mat* MyMatWithMat(Mat* mat);
Mat* MyMatRowsColsType(int rows, int cols, int type, const Scalar* s);
int Mat_flags(Mat* mat);
int Mat_dims(Mat* mat);
int Mat_rows(Mat* mat);
int Mat_cols(Mat* mat);
uchar* Mat_data(Mat* mat);
size_t Mat_total(Mat* mat);
size_t Mat_elemSize(Mat* mat);


//InputArray
InputArray * MyInputArray();
InputArray * MyInputArrayWithMat(Mat* mat);

//OutputArray
OutputArray * MyOutputArray();

//SparseMat
SparseMat* MySparseMat();


//Utility and System Functions and Macros
const char* GetBuildInformation();
int GetNumberOfCPUs();

//Drawing functions
void line(Mat* img, Point* p1, Point* p2, const Scalar* color);
void ellipse(Mat* img, Point* p1, Size* axes, double angle, double startAngle, double endAngle, const Scalar *color);
void circle(Mat* img, Point* center, int radius, const Scalar* color);	
void rectangle(Mat* img, Point* p1, Point* p2, const Scalar* color);
void rectangleWithRect(Mat* img, Rect* rect, const Scalar* color);

#ifdef __cplusplus
}
#endif	

#endif
