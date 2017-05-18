package model

import (
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

type PrivacyInput struct{
    pour        Pour
    sn          String
}

func (i *PrivacyInput)GetType(){
    return PrivacyTransaction
}

func (i *PrivacyInput)Prepare(ctx Context){
}

func (i *PrivacyInput)Verify(ctx Context)bool{
    //if snset.contains( i.sn ) return false
    return i.pour.Verify(ctx.Params, ctx.Merkle)
}

func (i *PrivacyInput)Execute(stub.ChaincodeStubInterface)error{
    //snset.add( i.sn )
    return nil
}

func (i *PrivacyInput)String()(string){
    res := fmt.Sprintf("%v\n%v", i.pour.String(), i.sn)
    return res
}

func (i *PrivacyInput)FromString(istr string)error{
    ostr := SplitN(istr, "\n")
    err := i.pour.FromString( ostr[0] )
    if err != nil{
        return err
    }
    i.sn = ostr[1]
    return nil
}

