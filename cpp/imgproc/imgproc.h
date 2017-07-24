#ifndef __IMGPROC_H__
#define __IMGPROC_H__

#include <opencv2/imgproc/imgproc.hpp>
#include "../core/core.h"


#ifdef __cplusplus
extern "C" {
#endif


//Histograms
void calcHist(const Mat* images, int nimgages, const int* channels,
			 InputArray *mask, OutputArray *hist, int dims,
			 const int * histSize, const float** ranges);

void calcHistWithSparseMat(const Mat* images, int nimgages, const int* channels,
			 InputArray *mask, SparseMat *hist, int dims,
			 const int * histSize, const float** ranges);


//Miscellaneous Image Transformations
void cvtColor(InputArray* src, OutputArray* dst, int code);



#ifdef __cplusplus
}
#endif	

#endif