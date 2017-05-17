package zklib

import (
    "testing"
)


func TestParams(t * testing.T){
    var p Params
    p.GetParams()
    defer p.DelParams()
    fmt.Println("Get a params pointer ", p.Ptr)
    fmt.Println("Try to get length", len(p.Ptr)
}
