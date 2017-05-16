#include <iostream>
#include "Goapi.h"
using namespace std;

int main(){
    //int t = TutorialTest();
    //cout <<"run goapi :"<<t<<endl;
    void *p = CParamsGen(1);
    CParamsDel(p);
    cout<<"seccess !"<<endl;
    return 0;
}
