package method

import(
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib/transaction"
)

type Transatction struct{
    stub *shim.ChaincodeStubInterface, 
    inputs []*transaction.Input,
    outputs []*transaction.Output,
}

func NewTransaction( s *shim.ChaincodeStubInterface, i []*transaction.Input, o []*transaction.Output ) *Transaction{
    t := New( Transaction )
    t.stub = s
    t.inputs = i
    t.outputs = o
    return t
}

func (t *Transaction) Execute()([]byte, err){
    var res []byte
    var err error
    for _, i := range t.inputs {
        err = i.Verify()
        if err != nil{
            return err
        }
    }
    for _, o := range t.outputs {
        err = o.Verify()
        if err != nil{
            return err
        }
    }
    for _, i := range t.inputs {
        i.Execute()
    }
    for _, o := range t.outputs {
        o.Execute()
    }
    return nil, nil
}