#include "gui.h"

extern "C" {

//User Interface
void imshow(const char *winname, InputArray* inputArray){
	std::string s = std::string(winname);
	cv::imshow(s, *static_cast<cv::_InputArray*>(inputArray));
}

int  waitkey(int delay){
	return cv::waitKey(delay);
}

//Reading and Writing Images and Video
Mat* imread(const char* filename){
	std::string s = std::string(filename);
	return (cv::Mat*) new cv::Mat(cv::imread(s));
}







} // extern "C"
