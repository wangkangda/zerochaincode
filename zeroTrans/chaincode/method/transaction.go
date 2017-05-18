package method

import(
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib/model"
)

type Transaction struct{
    stub    shim.ChaincodeStubInterface
    ctx     model.Context
    inputs  []*model.Input
    outputs []*model.Output
}

func NewTransaction( s shim.ChaincodeStubInterface, args []string ) *Transaction{
    t := new( Transaction )
    t.inputs = make([]*model.Input, 2)
    t.outputs = make([]*model.Output, 2)
    for i=0; i<2; i++{
        t.inputs[i] = new(model.Input)
        t.inputs[i].FromString( args[i] )
        t.inputs[i].Prepare(t.ctx)
        t.outputs[i] = new(model.Output)
        t.outputs[i].FromString( args[2+i])
        t.outputs[i].Prepare(t.ctx)
    }
    t.ctx.GetContext()
    return t
}

func (t *Transaction) Execute()([]byte, error){
    for i=0; i<2; i++{
        if !t.inputs[i].Verify(t.ctx) {
            return errors.New("Input %v Error", i)
        }
        if !t.outputs[i].Verify(t.ctx) {
            return errors.New("Output %v Error", i)
        }
    }
    for i=0; i<2; i++{
        _, err := t.inputs[i].Execute(t.ctx)
        if err != nil {
            return err
        }
        _, err := t.outputs[i].Execute(t.ctx)
        if err != nil {
            return err
        }
    }
    return nil, nil
}


    var err error
    for _, o := range t.outputs {
        o.Execute()
    }
    return nil, nil
}
