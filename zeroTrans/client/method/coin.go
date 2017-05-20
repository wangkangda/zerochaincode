package method

import (
    "log"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

type Coin struct{
    coin        zklib.Coin
    commitid    int
    value       int
}

func CmdCoin( cmd []string)error{
    name, value := "", 0
    if len(cmd)!=3 {
        return errors.New("Error for parameter number")
    }
    name = cmd[1]
    value = cmd[2]
    addr, exist := storage.AddressList[ name ]
    if !exist {
        return errors.New("No such Address Name. Please Generate Address First.")
    }
    _, exist := storage.CoinList[ name ]
    if !exist {
        storage.CoinList[name] = make([]*Coin, 0)
    }
    c := &Coin{}
    c.GetCoin( addr, value )
    c.value = value
    c.commitid = len(storage.CoinList[name])
    storage.CoinList[name] = append(storage.CoinList[name], c)
    log.Printf("Get a new coin No.%v for user [%v]", c.commitid, addr)
}
