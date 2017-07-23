#include "core.h"

extern "C" {


//Point2f
Point2f* MyPoint2f(float x, float y){
	return (cv::Point2f*) new cv::Point2f(x, y);
}
//Size2f
Size2f* MySize2f(float width, float height){
	return (cv::Size2f*) new cv::Size2f(width, height);
}

//Point
Point* MyPoint(int x, int y){
	return (cv::Point*) new cv::Point(x, y);
}

//Size
Size* MySize(int x, int y){
	return (cv::Size*) new cv::Size(x, y);
}

//RotatedRect
RotatedRect* MyRotatedRect(){
	return (cv::RotatedRect*) new cv::RotatedRect();
}

RotatedRect* RotatedRectCenterSize(const Point2f* center, const Size2f* size, float angle){
	return (cv::RotatedRect*) new cv::RotatedRect(*static_cast<cv::Point2f const*>(center),
		*static_cast<cv::Size2f const*>(size), angle);
}

Rect* RotatedRect_boundingRect(RotatedRect* rect){
	return (cv::Rect*) new cv::Rect(static_cast<cv::RotatedRect*>(rect)->boundingRect());
}

//Rect
Rect* MyRect(){
	return (cv::Rect*) new cv::Rect();
}


//Scalar
Scalar* MyScalar0(){
	return (cv::Scalar*) new cv::Scalar();
}
Scalar* MyScalar1(double v0){
	return (cv::Scalar*) new cv::Scalar(v0);
}
Scalar* MyScalar2(double v0, double v1){
	return (cv::Scalar*) new cv::Scalar(v0, v1);
}
Scalar* MyScalar3(double v0, double v1, double v2){
	return (cv::Scalar*) new cv::Scalar(v0, v1, v2);
}

//Mat
Mat* MyMat(){
	return  (cv::Mat*) new cv::Mat();
}
Mat* MyMatWithMat(Mat* mat){
	return (cv::Mat*) new cv::Mat(*static_cast<const cv::Mat*>(mat));
}
Mat* MyMatRowsColsType(int rows, int cols, int type, const Scalar* s){
	return (cv::Mat*) new cv::Mat(rows, cols, type, *static_cast< const cv::Scalar*>(s));
}
Mat* MyMatZeros(int rows, int cols, int type){
	return (cv::Mat*) new cv::Mat(cv::Mat::zeros(rows, cols, type));
}
Mat* MyMatZerosWithSize(Size* size, int type){
	return (cv::Mat*) new cv::Mat(cv::Mat::zeros(*static_cast<cv::Size*>(size), type));
}


int Mat_flags(Mat* mat){
	return static_cast<cv::Mat*>(mat)->flags;
}
int Mat_dims(Mat* mat){
	return static_cast<cv::Mat*>(mat)->dims;
}
int Mat_rows(Mat* mat){
	return static_cast<cv::Mat*>(mat)->rows;
}
int Mat_cols(Mat* mat){
	return static_cast<cv::Mat*>(mat)->cols;
}
uchar* Mat_data(Mat* mat){
	return (uchar* )static_cast<cv::Mat*>(mat)->data;
}
size_t Mat_total(Mat* mat){
	return static_cast<cv::Mat*>(mat)->total();
}
size_t Mat_elemSize(Mat* mat){
	return static_cast<cv::Mat*>(mat)->elemSize();
}

//InputArray
InputArray * MyInputArray(){
	return (cv::_InputArray*) new cv::_InputArray();
}
InputArray * MyInputArrayWithMat(Mat* mat){
	return (cv::_InputArray*) new cv::_InputArray(*static_cast<cv::Mat*>(mat));
}

//OutputArray
OutputArray * MyOutputArray(){
	return (cv::_InputArray*) new cv::_InputArray();
}
OutputArray * MyOutputArrayWithMat(Mat* mat){
	return (cv::_OutputArray*) new cv::_OutputArray(*static_cast<cv::Mat*>(mat));
}

//SparseMat
SparseMat* MySparseMat(){
	return (cv::SparseMat*) new cv::SparseMat();
}
SparseMat* MySparseMatWithMat(Mat* mat){
	return (cv::SparseMat*) new cv::SparseMat(*static_cast<cv::Mat*>(mat));
}

//Operations on Arrays
void minMaxLoc(const SparseMat* a, double* minVal, double* maxVal, int* minIdx, int* maxIdx){
	cv::minMaxLoc(*static_cast<const cv::SparseMat*>(a), NULL, NULL, NULL, NULL);
}

//Drawing functions
void line(Mat* img, Point* p1, Point* p2, const Scalar* color){
	cv::line(*static_cast<cv::Mat*>(img), 
		*static_cast<cv::Point*>(p1), *static_cast<cv::Point*>(p2),
		*static_cast<const cv::Scalar*>(color));
}
void ellipse(Mat* img, Point* center, Size* axes, double angle, double startAngle, double endAngle, const Scalar *color){
	cv::ellipse(*static_cast<cv::Mat*>(img), 
		*static_cast<cv::Point*>(center),
		*static_cast<cv::Size*>(axes),
		angle, startAngle, endAngle,
		*static_cast<const cv::Scalar*>(color));
}
void circle(Mat* img, Point* center, int radius, const Scalar* color){
	cv::circle(*static_cast<cv::Mat*>(img),
		*static_cast<cv::Point*>(center),
		radius,
		*static_cast<const cv::Scalar*>(color));
}	
void rectangle(Mat* img, Point* p1, Point* p2, const Scalar* color){
	cv::rectangle(*static_cast<cv::Mat*>(img),
		*static_cast<cv::Point*>(p1), *static_cast<cv::Point*>(p2),
		*static_cast<const cv::Scalar*>(color));
}

void rectangleWithRect(Mat* img, Rect* rect, const Scalar* color){
	cv::rectangle(*static_cast<cv::Mat*>(img),
		*static_cast<cv::Rect*>(rect),
		*static_cast<const cv::Scalar*>(color));
}



//Utility and System Functions and Macros
const char* GetBuildInformation(){
	std::string  str = cv::getBuildInformation();
	return str.c_str();
}
int GetNumberOfCPUs(){
	return cv::getNumberOfCPUs();
}




} // extern "C"
