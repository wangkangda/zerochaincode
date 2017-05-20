package model

import(
    "fmt"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

type PrivacyOutput struct{
    coin    zklib.Coin
}

func (o *PrivacyOutput)GetType()int{
    return PrivacyTransaction
}

func (o *PrivacyOutput)Prepare(ctx Context){
}

func (o *PrivacyOutput)Verify(ctx Context)bool{
    return true
}

func (o *PrivacyOutput)Execute(ctx Context)error{
    ctx.merkle.Insert( o.coin.GetCommit(), ctx.commitnum )
    ctx.commitnum ++
    return nil
}

func (o *PrivacyOutput)String()string{
    return fmt.Sprintf("%v", o.coin.String())
}

func (o *PrivacyOutput)FromString(istr string)error{
    o.coin.FromString(istr)
    return nil
}