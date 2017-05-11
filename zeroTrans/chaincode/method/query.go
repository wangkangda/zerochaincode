package method

import(
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib/transaction"
)

type Query struct{
    stub *shim.ChaincodeStubInterface,
    object string,
    key string
)

func NewQuery(stub *shim.ChaincodeStubInterface, o string, p []string)*Query{
    q = New(Query)
    q.stub = stub
    q.object = o
    if p != nil {
        q.key = p[0]
    }
    return q
}

func (t *SimpleChaincode) Execute()([]byte, error){
    return nil, nil
}
