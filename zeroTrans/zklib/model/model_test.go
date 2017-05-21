package model

import (
    "testing"
    "fmt"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

func TestModel(t *testing.T){
    fmt.Println("Start Model Test")
    m := &NormalInput{}
    m.sender.GetAddress()
    fmt.Println("Get Normal Input:\n", m.String())
    var n Input
    n = &NormalInput{}
    n.FromString( m.String() )
    fmt.Println("After Serialized\n", n.String())

    no := &NormalOutput{}
    no.amount = 100
    no.receiver.GetAddress()
    fmt.Println("Get Normal Output:\n", no.String())
    var mo Output
    mo = &NormalOutput{}
    mo.FromString(no.String())
    fmt.Println("After Serialized\n", mo.String())

    po := &PrivacyOutput{}
    po.coin.GetCoin( no.receiver, 2 )
    fmt.Println("Get Privacy Output:\n", po.String())
    mo = &PrivacyOutput{}
    mo.FromString(po.String())
    fmt.Println("After Serialized\n", mo.String())
    
    pi := &PrivacyInput{}
    merkle := zklib.Merkle{}
    merkle.GetMerkle()
    params := zklib.Params{}
    params.GetParams(0)
    merkle.Insert( po.coin.GetCommit(), 1 )
    po.coin.Index = 1
    c1 := zklib.Coin{}
    c2 := zklib.Coin{}
    c3 := zklib.Coin{}
    c1.GetCoin( no.receiver, 4 )
    c2.GetCoin( no.receiver, 3 )
    c3.GetCoin( no.receiver, 2 )
    merkle.Insert( c1.GetCommit(), 2 )
    c1.Index = 2
    pi.pour.GetPour(params, po.coin, c1, merkle, 1, c2, c3 )
    //fmt.Println("Get Pravicy Input:\n", pi.String())
    n = &PrivacyInput{}
    n.FromString( pi.String() )
    //fmt.Println("After Serialized\n", n.String())
    c := Context{}
    c.Merkle = &merkle
    c.Params = &params
    fmt.Println("Verify Result:", n.(*PrivacyInput).Verify(c))
}
