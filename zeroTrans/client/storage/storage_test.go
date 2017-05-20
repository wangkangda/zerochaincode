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
    fmt.Println("Address:", AddressList["testing2"].String())
    CoinList["testing2"] = make([]*MyCoin, 0)
    c := &MyCoin{}
    c.Coin = new(zklib.Coin)
    c.Coin.GetCoin(*AddressList["testing2"], 50)
    c.Value = 50
    CoinList["testing2"] = append(CoinList["testing2"], c)
    fmt.Println("Coin:", c.String())
    err = SaveStorage()
    fmt.Println(err)
}
