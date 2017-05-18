package model

import (
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

type NormalInput struct{
    sender      Address
    value       int
    signed      string
}

func (i *NormalInput)GetType(){
    return NormalTransaction
}

func (i *NormalInput)Prepare(ctx Context){
    ctx.AddAmount( sender )
}

func (i *NormalInput)Verify(ctx Context)bool{
    //之后可加入签名验证?
    //verify(signed, sender, ctx.signedContent)
    return ctx.amount[sender]>=i.value
}

func (i *NormalInput)Execute(ctx Context)error{
    ctx.amount[sender]-=i.value
}

func (i *NormalInput)String()(string){
    return fmt.Sprintf( "%v\n%v\n%v", i.sender.String(), i.value, i.signed )
}

func (i *NormalInput)FromString(istr string)error{
    ostr := SplitN(istr, "\n")
    err := i.sender.FromString( ostr[0] )
    if err != nil{
        return err
    }
    i.value, err = strconv.Atoi( ostr[1] )
    if err != nil{
        return err
    }
    i.signed = ostr[2]
}
