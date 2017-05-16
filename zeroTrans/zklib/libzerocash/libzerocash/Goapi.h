#ifndef GOAPI_H_
#define GOAPI_H_

#ifdef __cplusplus
extern "C"{
#endif
    //Params
    extern void*    CParamsGen();
    extern char*    CParamsStr(void*);
    extern void*    CStrParams(char*);
    extern void     CParamsDel(void*);
    
    //Address
    extern void*    CAddressGen();
    extern char*    CAddressStr(void*);
    extern void*    CStrAddress(char*);
    extern void     CAddressDel(void*);
    
    //Pricoin
    extern void*    CCoinGen(void*, int);
    extern char*    CCoinStr(void*);
    extern void*    CStrCoin(char*);
    extern void     CCoinDel(void*);
    
    //Merkle
    extern void*    CMerkleGen();
    extern char*    CMerkleStr(void*);
    extern void*    CStrMerkle(char*);
    extern void     CMerkleDel(void*);
    extern bool     CMerkleInsert(void*, void*, int);
    
    //Mint
    extern void*    CMintGen(void*);
    extern char*    CMintStr(void*);
    extern void*    CStrMint(char*);
    extern void     CMintDel(void*);
    extern bool     CMintVerify(void*);
    
    //Pour
    extern void*    CPourGen(void*, void*, void*,
                             void*, void*, int, int,
                             void*, void*, void*,
                             int, void*, void* );
    extern char*    CPourStr(void*);
    extern void*    CStrPour(char*);
    extern void     CPourDel(void*);
    extern bool     CPourVerify(void*, void*, void*);
    
    
    extern bool     TutorialTest();
    
#ifdef __cplusplus
}
#endif
#endif