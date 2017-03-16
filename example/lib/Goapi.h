#ifndef GOAPI_H_
#define GOAPI_H_

#ifdef __cplusplus
extern "C"{
#endif

// print hello
extern void TestGoapi();
extern char* GoCoinGeneration(char *);
extern void GoCoinDestroy(char *);

extern void* GoParamGen( char* );
extern void GoParamDel( void* );
extern void* GoPriCoinGen( void* );
extern void GoPriCoinDel( void* );
extern void* GoAccumGen( void*, int );
extern void GoAccumDel( void* );
extern void* GoCoinSpendGen( void*, void*, void* );
extern void GoCoinSpendDel( void* );
//extern bool GoCoinSpendVer( void*, void*, void* );

extern char* CCParamsGen();
extern void* CCParamsLoad( char* );
extern void CCParamsDel( void* );

extern char* CCAccumGen( void* );
extern void* CCAccumLoad( void*, char* );
extern char* CCAccumCal( void*, void*, char* );
extern void CCAccumDel( void* );

extern char* CCSpendVerify( void*, char*, char*, void* );
extern void CCBignumDel( void* );

extern void* CCPricoinGen( void* );
extern char* CCPricoinGen2( void* );
extern void* CCPricoinLoad( void*, char* );
extern void CCPricoinDel( void* );
extern char* CCPubcoinGen( void*, void* );
extern char* CCSpendGen( void*, void*, void*, char* );
extern void CCStrDel( char* );
#ifdef __cplusplus
}
#endif
#endif
