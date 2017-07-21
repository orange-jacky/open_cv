#include "imgproc.h"

extern "C" {

//Histograms
void calcHist(const Mat* images, int nimages, const int* channels,
			 InputArray *mask, OutputArray *hist, int dims,
			 const int * histSize, const float** ranges){
	cv::calcHist(static_cast<const cv::Mat*>(images), nimages, channels,
				*static_cast<cv::_InputArray*>(mask),*static_cast<cv::_OutputArray*>(hist),
				dims, histSize, ranges);
}

void calcHistWithSparseMat(const Mat* images, int nimages, const int* channels,
			 InputArray *mask, SparseMat *hist, int dims,
			 const int * histSize, const float** ranges){
	cv::calcHist(static_cast<const cv::Mat*>(images), nimages, channels,
				*static_cast<cv::_InputArray*>(mask),*static_cast<cv::SparseMat*>(hist),
				dims, histSize, ranges);
}

} // extern "C"