#ifndef __GUI_H__
#define __GUI_H__

#include <opencv2/highgui/highgui.hpp>
#include "../core/core.h"

#ifdef __cplusplus
extern "C" {
#endif

//User Interface
void imshow(const char *winname, InputArray* inputArray);
int  waitkey(int delay);

//Reading and Writing Images and Video
Mat* imread(const char* filename);






#ifdef __cplusplus
}
#endif	

#endif