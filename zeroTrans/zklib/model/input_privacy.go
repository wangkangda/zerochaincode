package model

import(
    "fmt"
    "strings"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

type PrivacyInput struct{
    pour        zklib.Pour
    sn          string
}

func (i *PrivacyInput)GetType()int{
    return PrivacyTransaction
}

func (i *PrivacyInput)Prepare(ctx Context){
}

func (i *PrivacyInput)Verify(ctx Context)bool{
    //if snset.contains( i.sn ) return false
    return i.pour.Verify(ctx.params, ctx.merkle)
}

func (i *PrivacyInput)Execute(ctx Context)error{
    //snset.add( i.sn )
    return nil
}

func (i *PrivacyInput)String()(string){
    res := fmt.Sprintf("%v\n%v", i.pour.String(), i.sn)
    return res
}

func (i *PrivacyInput)FromString(istr string)error{
    ostr := strings.Split(istr, "\n")
    i.pour.FromString( ostr[0] )
    i.sn = ostr[1]
    return nil
}

