package method

import(
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/wangkangda/zerochaincode/zeroTrans/chaincode/zklib/model"
)

type InitMethod struct{
    ctx     model.Context
}

func NewInitMethod( s shim.ChaincodeStubInterface ) *InitMethod{
    t := new( InitMethod )
    t.ctx.Stub = s
    t.ctx.InitContext()
    return t
}

func (t *InitMethod) Execute()([]byte, error){
    return nil, t.ctx.SaveContext()
}
