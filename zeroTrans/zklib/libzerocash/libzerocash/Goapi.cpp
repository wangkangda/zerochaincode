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

#define default_tree_depth 5

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

char* ss2str( stringstream &s ){
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

stringstream str2ss( char *s ){
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
    stringstream stream( vc );
    return stream;
}

//Params
void*    CParamsGen(){
    ZerocashParams *p = new ZerocashParams(default_tree_depth);
    p->getProvingKey(1);
    return (void*)p;
}
char*    CParamsStr(void *p){
    stringstream ss;
    ss << *(ZerocashParms*)p;
    //CDataStream stream(SER_NETWORK, 7002);
    stringstream stream;
    stream << *(ZerocashParams*)p;
    return data2str( stream );
}
void*    CStrParams(char* cstr){
    CDataStream stream = str2data( cstr );
    ZerocashParams *p = new ZerocashParams(default_tree_depth);
    //stream >> *p;
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
bool    CMerkleInsert(void* p, void* coin, int nowidx){
    IncrementalMerkleTree *merkle = (IncrementalMerkleTree*)p;
    vector<bool> temp_comVal(cm_size*8);
    convertBytesVectorToVector((*(Coin*)coin).getCoinCommitment().getCommitmentValue(), temp_comVal);
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
bool     CMintVerify(void* mint){
    return (*(MintTransaction*)mint).verify();
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
bool     CPourVerify(void* params, void* pour, void* tree){
    std::vector<unsigned char> pubkeyHash(sig_pk_size, 'a');
    IncrementalMerkleTree &merkleTree = *(IncrementalMerkleTree*)tree;
    vector<bool> root_bv(root_size * 8);
    merkleTree.getRootValue(root_bv);
    vector<unsigned char> rt(root_size);
    libzerocash::convertVectorToBytesVector(root_bv, rt);
    return ((PourTransaction*)pour)->verify(
                        *(ZerocashParams*)params, pubkeyHash, rt);
}


bool TutorialTest() {
    size_t tree_depth = default_tree_depth;
    cout << "\nSIMPLE TRANSACTION TEST\n" << endl;
    
    //libzerocash::timer_start("Param Generation");
    libzerocash::ZerocashParams p(tree_depth);
    p.getProvingKey(1); // Strangely enough, this is how we trigger parameter generation.
    //libzerocash::timer_stop("Param Generation");
    
    vector<libzerocash::Coin> coins(5);
    vector<libzerocash::Address> addrs(5);
    
    cout << "Creating Addresses and Coins...\n" << endl;
    for(size_t i = 0; i < coins.size(); i++) {
        addrs.at(i) = libzerocash::Address();
        coins.at(i) = libzerocash::Coin(addrs.at(i).getPublicAddress(), i);
    }
    cout << "Successfully created address and coins.\n" << endl;
    
    cout << "Creating a Mint Transaction...\n" << endl;
    libzerocash::MintTransaction minttx(coins.at(0));
    cout << "Successfully created a Mint Transaction.\n" << endl;
    
    cout << "Serializing a mint transaction...\n" << endl;
    CDataStream serializedMintTx(SER_NETWORK, 7002);
    serializedMintTx << minttx;
    cout << "Successfully serialized a mint transaction.\n" << endl;
    
    libzerocash::MintTransaction minttxNew;
    serializedMintTx >> minttxNew;
    cout << "Successfully deserialized a mint transaction.\n" << endl;
    
    cout << "Verifying a Mint Transaction...\n" << endl;
    bool minttx_res = minttxNew.verify();
    
    vector<std::vector<bool>> coinValues(5);
    vector<bool> temp_comVal(cm_size * 8);
    for(size_t i = 0; i < coinValues.size(); i++) {
        libzerocash::convertBytesVectorToVector(coins.at(i).getCoinCommitment().getCommitmentValue(), temp_comVal);
        coinValues.at(i) = temp_comVal;
    }
    
    cout << "Creating Merkle Tree...\n" << endl;
    libzerocash::IncrementalMerkleTree merkleTree(coinValues, tree_depth);
    cout << "Successfully created Merkle Tree.\n" << endl;
    
    cout << "Creating Witness 1...\n" << endl;
    merkle_authentication_path witness_1(tree_depth);
    if (merkleTree.getWitness(ConvertIntToVector(1), witness_1) == false) {
        cout << "Could not get witness" << endl;
        return false;
    }
    cout << "Successfully created Witness 1.\n" << endl;
    
    cout << "Creating Witness 2...\n" << endl;
    merkle_authentication_path witness_2(tree_depth);
    if (merkleTree.getWitness(ConvertIntToVector(3), witness_2) == false) {
        cout << "Could not get witness" << endl;
    }
    cout << "Successfully created Witness 2.\n" << endl;
    
    cout << "Creating coins to spend...\n" << endl;
    libzerocash::Address newAddress3;
    libzerocash::PublicAddress pubAddress3 = newAddress3.getPublicAddress();
    
    libzerocash::Address newAddress4;
    libzerocash::PublicAddress pubAddress4 = newAddress4.getPublicAddress();
    
    libzerocash::Coin c_1_new(pubAddress3, 2);
    libzerocash::Coin c_2_new(pubAddress4, 2);
    cout << "Successfully created coins to spend.\n" << endl;
    
    vector<bool> root_bv(root_size * 8);
    merkleTree.getRootValue(root_bv);
    vector<unsigned char> rt(root_size);
    libzerocash::convertVectorToBytesVector(root_bv, rt);
    
    
    vector<unsigned char> as(sig_pk_size, 'a');
    
    cout << "Creating a pour transaction...\n" << endl;
    libzerocash::PourTransaction pourtx(1, p,
                                        rt,
                                        coins.at(1), coins.at(3),
                                        addrs.at(1), addrs.at(3),
                                        1, 3,
                                        witness_1, witness_2,
                                        pubAddress3, pubAddress4,
                                        0,
                                        as,
                                        c_1_new, c_2_new);
    cout << "Successfully created a pour transaction.\n" << endl;
    
    cout << "Serializing a pour transaction...\n" << endl;
    CDataStream serializedPourTx(SER_NETWORK, 7002);
    serializedPourTx << pourtx;
    cout << "Successfully serialized a pour transaction.\n" << endl;
    
    libzerocash::PourTransaction pourtxNew;
    serializedPourTx >> pourtxNew;
    cout << "Successfully deserialized a pour transaction.\n" << endl;
    
    std::vector<unsigned char> pubkeyHash(sig_pk_size, 'a');
    
    cout << "Verifying a pour transaction...\n" << endl;
    bool pourtx_res = pourtxNew.verify(p, pubkeyHash, rt);
    
    return (minttx_res && pourtx_res);
}
