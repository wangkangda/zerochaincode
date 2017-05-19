package model

import(
    "fmt"
    "strconv"
    "strings"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

type NormalOutput struct{
    receiver    zklib.Address
    amount      int
}

func (o *NormalOutput)GetType()int{
    return NormalTransaction
}

func (o *NormalOutput)Prepare(ctx Context){
    ctx.AddAmount( o.receiver.String() )
}

func (o *NormalOutput)Verify(context Context )bool{
    return true
}

func (o *NormalOutput)Execute(context Context)error{
    context.amount[ o.receiver.String() ] += o.amount
    return nil
}

func (o *NormalOutput)String()string{
    return fmt.Sprintf("%v\n%v", o.receiver.String(), o.amount)
}

func (o *NormalOutput)FromString(istr string)error{
    ostr := strings.Split(istr, "\n")
    o.receiver.FromString(ostr[0])
    var err error
    o.amount, err = strconv.Atoi(ostr[1])
    return err
}
