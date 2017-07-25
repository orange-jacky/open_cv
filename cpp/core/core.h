#ifndef __CORE_H__
#define __CORE_H__

#include <opencv2/core/core.hpp>
#include <cv.h>

#ifdef __cplusplus
extern "C" {
#endif

// float** 和 [][]float32相互转化
extern float** makeFloatArray(int size);
extern void freeFloatArray(float **f, int size);
extern void setFloatArray(float **f, float *a, int n);
extern float* makeFloat(int size);
void freeFloat(float*f);
void setFloat(float*f, float v, int n);


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
Scalar* MyScalarWithAll(double v);

//Mat
Mat* MyMat();
Mat* MyMatWithMat(Mat* mat);
Mat* MyMatRowsColsType(int rows, int cols, int type, const Scalar* s);
Mat* MyMatZeros(int rows, int cols, int type);
Mat* MyMatZerosWithSize(Size* size, int type);

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
OutputArray * MyOutputArrayWithMat(Mat* mat);

//SparseMat
SparseMat* MySparseMat();
SparseMat* MySparseMatWithMat(Mat* mat);


//Operations on Arrays
void minMaxLocWithSparseMat(const SparseMat* a, double* minVal, double* maxVal);
void minMaxLocWithInputArray(InputArray *src, double* minVal, double* maxVal);
void normalize(InputArray *src, OutputArray *dst);
void subtract(InputArray *src1, InputArray *src2, OutputArray *dst);

//Drawing functions
void line(Mat* img, Point* p1, Point* p2, const Scalar* color);
void ellipse(Mat* img, Point* p1, Size* axes, double angle, double startAngle, double endAngle, const Scalar *color);
void circle(Mat* img, Point* center, int radius, const Scalar* color);	
void rectangle(Mat* img, Point* p1, Point* p2, const Scalar* color);
void rectangleWithRect(Mat* img, Rect* rect, const Scalar* color);

//Utility and System Functions and Macros
const char* GetBuildInformation();
int GetNumberOfCPUs();


#ifdef __cplusplus
}
#endif	

#endif
