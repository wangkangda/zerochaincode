#include <stdlib.h>
#include <iostream>

#include "libzerocash/Zerocash.h"
#include "libzerocash/Address.h"
#include "libzerocash/CoinCommitment.h"
#include "libzerocash/Coin.h"
#include "libzerocash/IncrementalMerkleTree.h"
#include "libzerocash/MerkleTree.h"
#include "libzerocash/MintTransaction.h"
#include "libzerocash/PourTransaction.h"
#include "libzerocash/utils/util.h"

using namespace std;
using namespace libzerocash;

#define default_tree_depth 4

vector<bool> ConvertIntToVector(uint64_t val) {
    vector<bool> ret;
    
    for(unsigned int i = 0; i < sizeof(val) * 8; ++i, val >>= 1) {
        ret.push_back(val & 0x01);
    }
    
    reverse(ret.begin(), ret.end());
    return ret;
}


char* data2str( CDataStream &s){
    std::string sstr = s.str();
    char *data = new char[sstr.length()*2+1];
    char t;
    for(int i=0; i<sstr.length(); i++){
        t = sstr[i];
        data[i*2] = 'a'+( (t>>4) & 0x0f);
        data[i*2+1] = 'a'+(t&0x0f);
    }
    data[sstr.length()*2]='\0';
    return data;
}

CDataStream str2data( char *s ){
    std::vector<char> vc;
    char *e = s, t1, t2, t;
    while( *e != 0 ) {
        t1 = *e;
        e++;
        t2 = *e;
        e++;
        t = ( ((t1-'a')<<4) | (t2-'a') );
        vc.push_back( t );
    }
    CDataStream stream( vc, SER_NETWORK, 7002);
    return stream;
}

//Params
void*    CParamsGen(int type=0){
    ZerocashParams *p;
    if (type!=0){
        //For Client: import proving key
        p = new ZerocashParams(default_tree_depth, "newpk", "newvk");
    }else{
        //For Chaincode: only import verify key
        p = new ZerocashParams(default_tree_depth, "", "newvk");
    }

    //p->getProvingKey(1);
    return (void*)p;
}

void     CParamsDel(void *p){
    delete (ZerocashParams*)p;
}

//Address
void*    CAddressGen(){
    Address *a = new Address;
    return (void*)a;
}
char*    CAddressStr(void *a){
    CDataStream stream(SER_NETWORK, 7002);
    stream << *(Address*)a;
    return data2str( stream );
}
void*    CStrAddress(char* cstr){
    CDataStream stream = str2data( cstr );
    Address *p = new Address;
    stream >> *p;
    return (void*)p;
}
void     CAddressDel(void* p){
    delete (Address*)p;
}

//Pricoin
void*    CCoinGen( void* addr, int value){
    Address &a = *(Address*)addr;
    Coin *c = new Coin(a.getPublicAddress(), (size_t)value);
    return (void*) c;
}
char*    CCoinStr(void *p){
    CDataStream stream(SER_NETWORK, 7002);
    stream << *(Coin*)p;
    return data2str( stream );
}
void*    CStrCoin(char* cstr){
    CDataStream stream = str2data( cstr );
    Coin *p = new Coin;
    stream >> *p;
    return (void*)p;
}
void     CCoinDel(void *p){
    delete (Coin*)p;
}
void*   CCoinCommit(void *p){
    Coin *c = (Coin*)p;
    return (void*) &( c->getCoinCommitment() );
}

//Commit
char*    CCommitStr(void *p){
    CDataStream stream(SER_NETWORK, 7002);
    stream << *(CoinCommitment*)p;
    return data2str( stream );
}
void*    CStrCommit(char* cstr){
    CDataStream stream = str2data( cstr );
    CoinCommitment *p = new CoinCommitment;
    stream >> *p;
    return (void*)p;
}
void     CCommitDel(void *p){
    delete (CoinCommitment*)p;
}

//Merkle
void*   CMerkleGen(){
    Address addrs = Address();
    Coin coins = Coin(addrs.getPublicAddress(), 0);
    vector<vector<bool>> coinValues(1);
    vector<bool> temp_comVal(cm_size * 8);
    convertBytesVectorToVector(coins.getCoinCommitment().getCommitmentValue(), temp_comVal);
    coinValues.at(0) = temp_comVal;
    IncrementalMerkleTree *res = new IncrementalMerkleTree(coinValues, default_tree_depth);
    return (void*)res;
}
char*   CMerkleStr(void* p){
    IncrementalMerkleTreeCompact compact;
    ( *(IncrementalMerkleTree*)p ).getCompactRepresentation(compact);
    CDataStream stream(SER_NETWORK, 7002);
    stream << compact;
    return data2str( stream );
}
void*   CStrMerkle(char* cstr){
    CDataStream stream = str2data( cstr );
    IncrementalMerkleTreeCompact compact;
    stream >> compact;
    IncrementalMerkleTree *p = new IncrementalMerkleTree(compact);
    return (void*)p;
}
void    CMerkleDel(void* p){
    delete (IncrementalMerkleTree*)p;
}
void*   CMerkleInsert(void* p, void* commit, int nowidx){
    IncrementalMerkleTree *merkle = (IncrementalMerkleTree*)p;
    vector<bool> temp_comVal(cm_size*8);
    convertBytesVectorToVector((*(CoinCommitment*)commit).getCommitmentValue(), temp_comVal);
    vector<bool> index = ConvertIntToVector((uint64_t)nowidx);
    merkle->insertElement(temp_comVal, index);
    return (void*)merkle;
}

//Mint
void*    CMintGen(void* coin){
    MintTransaction* res = new MintTransaction( *(Coin*)coin );
    return (void*)res;
}
char*    CMintStr(void* p){
    CDataStream stream(SER_NETWORK, 7002);
    stream << *(MintTransaction*)p;
    return data2str( stream );
}
void*    CStrMint(char *cstr){
    CDataStream stream = str2data( cstr );
    MintTransaction *p = new MintTransaction;
    stream >> *p;
    return (void*)p;
}
void     CMintDel(void* mint){
    delete (MintTransaction*)mint;
}
int     CMintVerify(void* mint){
    return (int)(*(MintTransaction*)mint).verify();
}

//Pour
void*    CPourGen(void* params,
                  void* coin1,      void* coin2,
                  void* addr1,      void* addr2,
                  int cidx1,        int cidx2,
                  void* tree,
                  void* paddr1,     void* paddr2,
                  int v_pub,
                  void* cnew1,    void* cnew2){
    PourTransaction *p;
    merkle_authentication_path witness_1(default_tree_depth);
    merkle_authentication_path witness_2(default_tree_depth);
    IncrementalMerkleTree &merkleTree = *(IncrementalMerkleTree*)tree;
    vector<bool> root_bv(root_size * 8);
    merkleTree.getRootValue(root_bv);
    vector<unsigned char> rt(root_size);
    convertVectorToBytesVector(root_bv, rt);
    merkleTree.getWitness(ConvertIntToVector((uint64_t)cidx1), witness_1);
    merkleTree.getWitness(ConvertIntToVector((uint64_t)cidx2), witness_2);
    vector<unsigned char> pubkeyHash(sig_pk_size, 'a');
    p = new PourTransaction(1, *(ZerocashParams*)params, rt,
                            *(Coin*)coin1, *(Coin*)coin2,
                            *(Address*)addr1, *(Address*)addr2,
                            (size_t)cidx1, (size_t)cidx2,
                            witness_1, witness_2,
                            (*(Address*)paddr1).getPublicAddress(), (*(Address*)paddr2).getPublicAddress(),
                            (uint64_t)v_pub, pubkeyHash,
                            *(Coin*)cnew1, *(Coin*)cnew2);
    return p;
}
char*    CPourStr(void* pour){
    CDataStream stream(SER_NETWORK, 7002);
    stream << *(PourTransaction*)pour;
    return data2str( stream );
}
void*    CStrPour(char* cstr){
    CDataStream stream = str2data( cstr );
    PourTransaction *p = new PourTransaction;
    stream >> *p;
    return (void*)p;
}
void     CPourDel(void* pour){
    delete (PourTransaction*)pour;
}
int     CPourVerify(void* params, void* pour, void* tree){
    std::vector<unsigned char> pubkeyHash(sig_pk_size, 'a');
    IncrementalMerkleTree &merkleTree = *(IncrementalMerkleTree*)tree;
    vector<bool> root_bv(root_size * 8);
    merkleTree.getRootValue(root_bv);
    vector<unsigned char> rt(root_size);
    libzerocash::convertVectorToBytesVector(root_bv, rt);
    return (int)((PourTransaction*)pour)->verify(
                        *(ZerocashParams*)params, pubkeyHash, rt);
}


int TutorialTest() {
    size_t tree_depth = default_tree_depth;
    cout << "\nSIMPLE TRANSACTION TEST\n" << endl;
    
    libzerocash::ZerocashParams &p = *(libzerocash::ZerocashParams*)CParamsGen(1);
    cout <<"Get Params From File"<<endl;
    
    vector<libzerocash::Coin*> pcoins(5);
    vector<libzerocash::Address*> paddrs(5);
    
    for(size_t i = 0; i < pcoins.size(); i++) {
        paddrs.at(i) = (Address*)CAddressGen();
        pcoins.at(i) = (Coin*)CCoinGen((void*)(paddrs.at(i)), i );
    }
    cout << "Successfully created address and coins.\n" << endl;

    IncrementalMerkleTree *pmerkle=(IncrementalMerkleTree*)CMerkleGen();
    for(size_t i = 1; i < pcoins.size(); i++) {
        int nowidx = i;
        CMerkleInsert((void*)pmerkle, (void*)&(pcoins.at(i)->getCoinCommitment()),nowidx );
    }
    cout << "Successfully created Merkle Tree.\n" << endl;

    Address *paddr3 = (Address*)CAddressGen();
    Address *paddr4 = (Address*)CAddressGen();
    Coin *pc1 = (Coin*)CCoinGen( (void*)paddr3, 2);
    Coin *pc2 = (Coin*)CCoinGen( (void*)paddr4, 2);
    cout << "Successfully created coins to pour.\n" << endl;
    
    PourTransaction *ppourtx = (PourTransaction*)CPourGen(&p,
                            (void*)pcoins.at(1), (void*)pcoins.at(3),
                            (void*)paddrs.at(1), (void*)paddrs.at(3),
                            1, 3,
                            (void*)pmerkle,
                            (void*)paddr3, (void*)paddr4,
                            0, (void*)pc1, (void*)pc2);

    cout << "Successfully created a pour transaction.\n" << endl;
    cout << "Verifying a pour transaction...\n" << endl;
    int pourtx_res = CPourVerify( (void*)&p, (void*)ppourtx, (void*)pmerkle);

    delete &p;
    for(size_t i = 0; i < pcoins.size(); i++){
        CAddressDel( (void*)paddrs.at(i) );
        CCoinDel( (void*)pcoins.at(i) );
    }
    CCoinDel( (void*)pc1 );
    CCoinDel( (void*)pc2 );
    CAddressDel( (void*)paddr3 );
    CAddressDel( (void*)paddr4 );
    CPourDel( (void*)ppourtx );
    CMerkleDel((void*)pmerkle);
    return pourtx_res;
    return 0;
}
