package method

import(
    "log"
    "errors"
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

func (q *Query) Execute()([]byte, error){
    var res []byte
    var err error
    switch q.object{
    case "amount":
        res, err = GetAmount( q.key )
    default:
        err = errors.New("Not Such Query Object")
    }
    return res, err
}

func GetAmount( address string )([]byte, error){
    final_key = fmt.Sprintf("%v%v", amount, address)
    amount, err := stub.GetState(address)
    if err != nil{
        log.Error(err.Error())
    }
    return amount, err
}
