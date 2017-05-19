package model

import (
    "testing"
    "fmt"
    //"github.com/wangkangda/zerochaincode/zeroTrans/zklib"
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
}
