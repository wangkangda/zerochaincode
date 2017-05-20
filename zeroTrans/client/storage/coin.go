package storage

import(
    "log"
    "encoding/json"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

type MyCoin struct{
    Coin        *zklib.Coin
    CoinStr     string      `json:"coin"`
    Commitid    int         `json:"commitid"`
    Value       int         `json:"value"`
}

func (c *MyCoin)String()string{
    coin := MyCoin{}
    coin.CoinStr = c.Coin.String()
    coin.Commitid = c.Commitid
    coin.Value = c.Value
    res, err := json.Marshal(coin)
    if err != nil{
        log.Printf("Get Error While Serialized Coin: %v", err)
    }
    return string(res)
}

func (c *MyCoin)FromString(cstr string)error{
    err :=  json.Unmarshal([]byte(cstr), c)
    if err != nil {
        return err
    }
    c.Coin = new(zklib.Coin)
    c.Coin.FromString( c.CoinStr )
    return nil
}
