package method

import(
    //"log"
    "errors"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    //"github.com/wangkangda/zerochaincode/zeroTrans/zklib/transaction"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib/model"
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

func (q *Query) Execute()([]byte, error){
    var res []byte
    var err error
    switch q.function{
    case model.ConstAmount:
        var amount int
        amount, err = q.ctx.GetAmount( q.key )
        res = []byte(strconv.Itoa(amount))
    case model.ConstMerkle:
        var m string
        m, err = q.ctx.GetMerkle()
        res = []byte(m)
    case model.ConstMerkleSize:
        var ms int
        ms, err = q.ctx.GetMerkleSize()
        res = []byte(strconv.Itoa(ms))
    default:
        err = errors.New("Not Such Query Object")
    }
    return res, err
}
