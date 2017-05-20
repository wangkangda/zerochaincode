package storage

import(
    "fmt"
    "testing"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

func TestStorage( t *testing.T ){
    err := GetStorage()
    fmt.Println(err)
    AddressList["testing2"] = &zklib.Address{}
    AddressList["testing2"].GetAddress()
    fmt.Println("Address:", AddressList["testing"].String())
    err = SaveStorage()
    fmt.Println(err)
}
