#include <stdio.h>
#include <vector>
#include "Zerocoin.h"

#define DUMMY_TRANSACTION_HASH  0 // in real life these would be uint256 hashes
#define DUMMY_ACCUMULATOR_ID    0 // in real life these would be uint256 hashes

#define TUTORIAL_TEST_MODULUS   "a8852ebf7c49f01cd196e35394f3b74dd86283a07f57e0a262928e7493d4a3961d93d93c90ea3369719641d626d28b9cddc6d9307b9aabdbffc40b6d6da2e329d079b4187ff784b2893d9f53e9ab913a04ff02668114695b07d8ce877c4c8cac1b12b9beff3c51294ebe349eca41c24cd32a6d09dd1579d3947e5c4dcc30b2090b0454edb98c6336e7571db09e0fdafbd68d8f0470223836e90666a5b143b73b9cd71547c917bf24c0efc86af2eba046ed781d9acb05c80f007ef5a0a5dfca23236f37e698e8728def12554bc80f294f71c040a88eff144d130b24211016a97ce0f5fe520f477e555c9997683d762aff8bd1402ae6938dd5c994780b1bf6aa7239e9d8101630ecfeaa730d2bbc97d39beb057f016db2e28bf12fab4989c0170c2593383fd04660b5229adcd8486ba78f6cc1b558bcd92f344100dff239a8c00dbc4c2825277f24bdd04475bcc9a8c39fd895eff97c1967e434effcb9bd394e0577f4cf98c30d9e6b54cd47d6e447dcf34d67e48e4421691dbe4a7d9bd503abb9"

char * data2cstr( CDataStream &s ){
	std::string sstr = s.str();
	char *data = new char[sstr.length()*2+1];
	char t;
	for(int i=0; i<sstr.length(); i++){
		t = sstr[i];
		data[i*2] = 'a'+( (t>>4) & 0x0f);
		data[i*2+1] = 'a'+(t&0x0f);
		//cout<<(short)t<<' ';
	}
	data[sstr.length()*2]='\0';
	//cout<<endl;
	return data;
}
CDataStream cstr2data( char *s ){
	std::vector<char> vc;
	char *e = s, t1, t2, t;
	while( *e != 0 ) {
		t1 = *e;
		e++;
		t2 = *e;
		e++;
		t = ( ((t1-'a')<<4) | (t2-'a') );
		//cout<<(short)t<<' ';
		vc.push_back( t );
	}
	//cout<<endl;
	CDataStream stream( vc, SER_NETWORK, PROTOCOL_VERSION);
	return stream;
}

void TestGoapi(){
    printf("hello from libzerocoin~\n");
}

char * GoCoinGeneration( char * p){
    Bignum testModulus;
    testModulus.SetHex( std::string(p) );
    libzerocoin::Params* params = new libzerocoin::Params(testModulus);
    cout<<"Successfully loaded parameters."<<endl;
    libzerocoin::PrivateCoin newCoin(params);
    libzerocoin::PublicCoin pubCoin = newCoin.getPublicCoin();
    cout<<"Successfully minted a zerocoin."<<endl;
    CDataStream serializedCoin(SER_NETWORK, PROTOCOL_VERSION);
    serializedCoin << pubCoin;
    std::string sstr = serializedCoin.str();
    char *data = new char[sstr.length()+1];
    sstr.copy(data, sstr.length(), 0 );
    return data;
}
void GoCoinDestroy( char * p ){
    delete []p;
}


void* GoParamGen( char* p ){
    Bignum testModulus;
    testModulus.SetHex( std::string(p) );
    libzerocoin::Params* params = new libzerocoin::Params( testModulus );
    cout<<"Successfully loaded parameters."<<endl;
    return (void *)params;
}
void GoParamDel( void* p ){
    delete (libzerocoin::Params*) p;
}

void* GoPriCoinGen( void *params ){
    libzerocoin::PrivateCoin *newCoin= new libzerocoin::PrivateCoin( (libzerocoin::Params*)params );
    return (void*) newCoin;
}
void GoPriCoinDel( void *coin ){
    delete (libzerocoin::PrivateCoin*)coin;
}

void* GoAccumGen( void *p, int test){
    libzerocoin::Params* params = (libzerocoin::Params*)p;
    libzerocoin::Accumulator *accum = new libzerocoin::Accumulator( params );
    for( int i=0; i<test; i++ ){
        libzerocoin::PrivateCoin testCoin(params);
        (*accum) += testCoin.getPublicCoin();
    }
    return (void *)accum;
}
void GoAccumDel( void *accum ){
    delete (libzerocoin::Accumulator *)accum;
}

void* GoCoinSpendGen( void* p, void* a, void* c ){
    libzerocoin::Params* params = (libzerocoin::Params*)p;
    libzerocoin::Accumulator* accum = (libzerocoin::Accumulator*)a;
    libzerocoin::PrivateCoin* coin = (libzerocoin::PrivateCoin*)c;
    libzerocoin::AccumulatorWitness witness(params, *accum, coin->getPublicCoin());
    (*accum) += coin->getPublicCoin();
    uint256 transactionHash = DUMMY_TRANSACTION_HASH;
    uint256 accumulatorID = DUMMY_ACCUMULATOR_ID;
    libzerocoin::SpendMetaData metaData(accumulatorID, transactionHash);
    libzerocoin::CoinSpend spend(params, *coin, *accum, witness, metaData);
    if( !spend.Verify( *accum, metaData) ){
        cout<<"ERROR: Our new CoinSpend transaction did not verify!"<<endl;
        return NULL;
    }
    CDataStream* serializedCoinSpend=new CDataStream(SER_NETWORK, PROTOCOL_VERSION);
    (*serializedCoinSpend) <<spend;
    cout << "Successfully generated a coin spend transaction."<<endl;
    return (void*) serializedCoinSpend;
}
void GoCoinSpendDel( void* cs ){
    delete (CDataStream*) cs;
}

bool GoCoinSpendVer( void* p, void* cs , void* a){
    libzerocoin::Params* params = (libzerocoin::Params*)p;
    CDataStream* serialzedCoinSpend = (CDataStream*)cs;
    libzerocoin::Accumulator* accum = (libzerocoin::Accumulator*)a;
    libzerocoin::CoinSpend newSpend(params, *serialzedCoinSpend);
    uint256 transactionHash = DUMMY_TRANSACTION_HASH;
    uint256 accumulatorID = DUMMY_ACCUMULATOR_ID;
    libzerocoin::SpendMetaData metaData(accumulatorID, transactionHash);
    if(!newSpend.Verify( *accum, metaData )){
        cout<<"ERROR: The CoinSpend transaction did not verify!"<<endl;
        return false;
    }
    Bignum serialNumber = newSpend.getCoinSerialNumber();
    cout <<"Successfully verified a coin spend transaction."<<endl;
    cout << endl <<"Coin serial number is:"<<endl <<serialNumber <<endl;
    return true;
}

char* CCParamsGen(){
	//generate params and return to chaincode as char array
	Bignum testModulus;
	testModulus.SetHex(std::string(TUTORIAL_TEST_MODULUS));
	libzerocoin::Params* params = new libzerocoin::Params(testModulus);
	CDataStream sParams(SER_NETWORK, PROTOCOL_VERSION);
	sParams << (*params);
	//cout<<"Get serial params:"<<sParams.str().size()<<endl;
	char * res = data2cstr( sParams );

	int t = 0;
	char * r = res;
	while( *r != '\0' ){
		t++;
		r++;
	}
	//std::cout<<"param length:"<<t<<endl;
	return res;
}
void* CCParamsLoad(char *p){
	CDataStream sParams = cstr2data( p );
	//cout<<"Get serial params:"<<sParams.str().size()<<endl;
	Bignum testModulus;
	testModulus.SetHex(std::string(TUTORIAL_TEST_MODULUS));
	libzerocoin::Params* params = new libzerocoin::Params(testModulus);
	//cout<<"Create new params"<<endl;
	sParams >> (*params);
	//cout<<"CCParamsLoad sucessful!"<<endl;
	return (void*)params;
}
void CCParamsDel( void* p ){
	delete (libzerocoin::Params*)p;
}

char* CCAccumGen(void *p){
	libzerocoin::Params* params = (libzerocoin::Params*)p;
	libzerocoin::Accumulator accumulator(params);
	CDataStream sAccum(SER_NETWORK, PROTOCOL_VERSION);
	sAccum << accumulator;
	return data2cstr( sAccum );
}
void* CCAccumLoad(void *p, char *s){
	libzerocoin::Params* params = (libzerocoin::Params*)p;
	CDataStream sAccum = cstr2data( s );
	libzerocoin::Accumulator *accum = new libzerocoin::Accumulator(params);
	sAccum >> (*accum);
	return (void*)accum;
}
char* CCAccumCal(void *p, void *a, char *c){
	libzerocoin::Params *params = (libzerocoin::Params*)p;
	libzerocoin::Accumulator *accum = (libzerocoin::Accumulator*)a;
	CDataStream sPubcoin = cstr2data( c );
	libzerocoin::PublicCoin pubcoin( params, sPubcoin );
	if(!pubcoin.validate()){
		std::cout <<"Error: coin is not valid!"<<endl;
		return NULL;
	}
	cout <<"Deserialized and verified the coin."<<endl;
	*accum += pubcoin;
	CDataStream sAccum(SER_NETWORK, PROTOCOL_VERSION);
	sAccum << (*accum);
	return data2cstr( sAccum );
}
void CCAccumDel( void *p ){
	delete (libzerocoin::Accumulator*)p;
}

char* CCSpendVerify( void* p, char* cs, char* metadata, void* a){
	libzerocoin::Params *params = (libzerocoin::Params*)p;
	libzerocoin::Accumulator *accum = (libzerocoin::Accumulator*)a;
	CDataStream sCoinspend = cstr2data( cs );

	libzerocoin::CoinSpend newSpend(params, sCoinspend);
	//do sha256 for metadata
	uint256 hashAddress = *(uint256*)metadata;
	libzerocoin::SpendMetaData smd(hashAddress, 123);
	if(!newSpend.Verify(*accum, smd)){
		return NULL;
	}
	Bignum *serialNum = new Bignum( newSpend.getCoinSerialNumber() );
	std::cout <<"Successfully verified a coin spend transaction." <<endl;
	std::cout <<"Coin serial number is:" <<(*serialNum) <<endl;
    CDataStream sBn(SER_NETWORK, PROTOCOL_VERSION);
    sBn << (*serialNum);
    return data2cstr(sBn);
}
void CCBignumDel( void* p ){
	delete (Bignum*)p;
}

void* CCPricoinGen( void* p ){
	libzerocoin::Params *params = (libzerocoin::Params*)p;
	libzerocoin::PrivateCoin *pricoin = new libzerocoin::PrivateCoin(params);
	return (void*)pricoin;
}
char* CCPricoinGen2( void* p ){
    libzerocoin::Params *params = (libzerocoin::Params*)p;
    libzerocoin::PrivateCoin pricoin(params);
    CDataStream sComm(SER_NETWORK, PROTOCOL_VERSION);
    sComm << pricoin;
    return data2cstr( sComm );
}
void* CCPricoinLoad( void* p, char* s){
    libzerocoin::Params* params = (libzerocoin::Params*)p;
    CDataStream sComm = cstr2data( s );
    libzerocoin::PrivateCoin *pricoin = new libzerocoin::PrivateCoin(params);
    sComm >> (*pricoin);
    return (void*)pricoin;
}
void CCPricoinDel( void* p ){
	delete (libzerocoin::PrivateCoin*)p;
}
char* CCPubcoinGen( void* p, void* pc ){
	libzerocoin::Params *params = (libzerocoin::Params*)p;
	libzerocoin::PrivateCoin *pricoin = (libzerocoin::PrivateCoin*)pc;
	libzerocoin::PublicCoin pubcoin = pricoin->getPublicCoin();
	CDataStream serializedCoin(SER_NETWORK, PROTOCOL_VERSION);
	serializedCoin << pubcoin;
	return data2cstr( serializedCoin );
}
char* CCSpendGen( void *p, void *pc, void *ac, char *md ){
	libzerocoin::Params *params = (libzerocoin::Params*)p;
	libzerocoin::PrivateCoin *pricoin = (libzerocoin::PrivateCoin*)pc;
	libzerocoin::Accumulator *accum = (libzerocoin::Accumulator*)ac;

	libzerocoin::AccumulatorWitness witness(params, *accum, pricoin->getPublicCoin());
	(*accum) += pricoin->getPublicCoin();
	//there should be a hash function
	uint256 metahash = *(uint256*)md;

	libzerocoin::SpendMetaData spenddata( metahash, metahash );
	libzerocoin::CoinSpend spend(params, *pricoin, *accum, witness, spenddata);
	if(!spend.Verify(*accum, spenddata)){
		std::cout<<"ERROR: Our new CoinSpend transaction did not verify!"<<endl;
		return NULL;
	}
	CDataStream sSpend(SER_NETWORK, PROTOCOL_VERSION);
	sSpend << spend;
	return data2cstr(sSpend);
}
void CCStrDel( char *s ){
	delete[] s;
}

