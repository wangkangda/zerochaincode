package method

import(
    "fmt"
    "log"
    "errors"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    //"github.com/wangkangda/zerochaincode/zeroTrans/zklib/transaction"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib/common"
)

type Query struct{
    stub    shim.ChaincodeStubInterface
    ctx     model.Context
    function    string
    key         string
}

func NewQuery(stub shim.ChaincodeStubInterface, function string, args []string)*Query{
    q := new(Query)
    q.stub = stub
    if len(args)>=1 {
        q.key = args[0]
    }
    return q
}

func (q *Query) Execute()([]byte, error)
    var res []byte
    var err error
    switch q.function{
    case model.ConstAmount:
        res, err = q.ctx.GetAmount( q.key )
    case model.ConstMerkle:
        res, err = q.ctx.GetMerkle()
    case model.ConstMerkleSize:
        res, err = q.ctx.GetMerkleSize()
    default:
        err = errors.New("Not Such Query Object")
    }
    return res, err
}
