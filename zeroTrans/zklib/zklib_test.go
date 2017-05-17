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
    var a, b Address
    var c, d Coin
    a.GetAddress()
    c.GetCoin( a, 100 )
    defer a.DelAddress()
    defer c.DelCoin()
    s1 := a.String()
    s3 := c.String()
    //fmt.Println("Get a address: ", s1)
    b.FromString(s1)
    d.FromString(s3)
    defer b.DelAddress()
    defer d.DelCoin()
    s2 := b.String()
    s4 := d.String()
    //fmt.Println("from string address: ", s2)
    ok := s1==s2 && s3==s4
    fmt.Println("Is it ok? : ", ok )

    var e Coin
    e.GetCoin( a, 50 )
    defer e.DelCoin()
    var m Merkle
    m.GetMerkle()
    defer m.DelMerkle()
    old := m.String()
    cc1 := c.GetCommit()
    cc2 := e.GetCommit()
    m.Insert(cc1, 1)
    m.Insert(cc2, 2)
    n1 := m.String()
    fmt.Println("Original Merkle:", old)
    fmt.Println("Inserted Merkle:", n1)
    ok = old != n1
    fmt.Println("insert different :", ok)
    var m1 Merkle
    m1.FromString(old)
    defer m1.DelMerkle()
    m1.Insert(cc1, 1)
    m1.Insert(cc2, 2)
    n2 := m.String()
    fmt.Println("Another Merkle:", n2)
    ok = n1==n2
    fmt.Println("merkle ok :", ok)
}

