package model

import(
    "fmt"
    "strconv"
    "strings"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

type NormalInput struct{
    sender      zklib.Address
    value       int
    signed      string
}

func (i *NormalInput)GetType()int{
    return NormalTransaction
}

func (i *NormalInput)Prepare(ctx Context){
    ctx.AddAmount( i.sender.String() )
}

func (i *NormalInput)Verify(ctx Context)bool{
    //之后可加入签名验证?
    //verify(signed, i.sender, ctx.signedContent)
    return ctx.Amount[i.sender.String()]>=i.value
}

func (i *NormalInput)Execute(ctx Context)error{
    ctx.Amount[i.sender.String()]-=i.value
    return nil
}

func (i *NormalInput)String()(string){
    return fmt.Sprintf( "%v\n%v\n%v", i.sender.String(), i.value, i.signed )
}

func (i *NormalInput)FromString(istr string)error{
    ostr := strings.Split(istr, "\n")
    i.sender.FromString( ostr[0] )
    var err error
    i.value, err = strconv.Atoi( ostr[1] )
    if err != nil{
        return err
    }
    i.signed = ostr[2]
    return nil
}
