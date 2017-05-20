package method

import (
    "log"
    "errors"
    "strconv"
    "github.com/wangkangda/zerochaincode/zeroTrans/client/storage"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

func CmdCoin( cmd []string)error{
    name, value := "", 0
    if len(cmd)!=3 {
        return errors.New("Error for parameter number")
    }
    name = cmd[1]
    value, err := strconv.Atoi(cmd[2])
    if err != nil{
        return errors.New("Error for parameter value")
    }
    addr, exist := storage.AddressList[ name ]
    if !exist {
        return errors.New("No such Address Name. Please Generate Address First.")
    }
    _, exist = storage.CoinList[ name ]
    if !exist {
        storage.CoinList[name] = make([]*storage.MyCoin, 0)
    }
    c := &storage.MyCoin{}
    c.Coin = &zklib.Coin{}
    c.Coin.GetCoin( *addr, value )
    c.Value = value
    c.Commitid = len(storage.CoinList[name])
    storage.CoinList[name] = append(storage.CoinList[name], c)
    log.Printf("Get a new coin No.%v for user [%v]", c.Commitid, addr)
    return nil
}
