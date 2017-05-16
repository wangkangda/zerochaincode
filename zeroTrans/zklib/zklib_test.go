package zklib

import (
    "testing"
    "fmt"
)


func TestParams(t * testing.T){
    var p Params
    p.GetParams( 0 )
    defer p.DelParams()
    fmt.Println("Get a params pointer ", p.Ptr)
}

func TestAddress(t *testing.T){
    var a Address
    var b Address
    a.GetAddress()
    defer a.DelAddress()
    s1 := a.String()
    fmt.Println("Get a address: ", s1)
    b.FromString(s1)
    defer b.DelAddress()
    s2 := b.String()
    fmt.Println("from string address: ", s2)
    ok := s1==s2
    fmt.Println("Is it ok? : ", ok )
}
