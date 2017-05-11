package method

import(
    "log"
    "errors"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib/transaction"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib/common"
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
    case common.Amount:
        res, err = GetAmount( q.key )
    case common.Params:
        res, err = GetData(common.Params)
    case common.Commit:
        res, err = GetCommit( q.key )
    default:
        err = errors.New("Not Such Query Object")
    }
    return res, err
}

func (q *Query)GetAmount( address string )([]byte, error){
    final_key = fmt.Sprintf("%v%v", common.Amount, address)
    return q.GetData( final_key )
}

func (q *Query)GetCommit( idString string)([]byte, error){
    num, err := strconv.Atoi( idString )
    if err != nil{
        return nil, common.ErrParams
    }
    final_key := fmt.Sprintf("%v%v", common.Commit, num)
    return q.GetData( final_key )
}

func (q *Query)GetData( key string )([]byte, error){
    res, err := q.stub.GetState(final_key)
    if err != nil{
        log.Error(err.Error())
    }
    return amount, err
}
