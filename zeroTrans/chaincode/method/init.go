package method

import(
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib/model"
)

type InitMethod struct{
    ctx     model.Context
}

func NewInitmethod( s shim.ChaincodeStubInterface ) *InitMethod{
    t := new( Initmethod )
    t.ctx.stub = s
    t.ctx.InitContext()
    return t
}

func (t *InitMethod) Execute()([]byte, error){
    return nil, t.ctx.SaveContext()
}
