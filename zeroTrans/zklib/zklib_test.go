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
