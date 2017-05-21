package method

import(
    "fmt"
    "errors"
    "github.com/wangkangda/zerochaincode/zeroTrans/client/storage"
    "github.com/wangkangda/zerochaincode/zeroTrans/chaincode/zklib"
)

func CmdAddress( cmd []string )error{
    name, newname := "", ""
    if len(cmd)==1 || len(cmd)>3{
        return errors.New("Error for parameter number")
    }else if len(cmd)==3{
        newname = cmd[2]
    }
    name = cmd[1]
    _, exist := storage.AddressList[ name ]
    if exist {
        if newname == ""{
            fmt.Printf("Address [%v] exist:\n", name)
        }else{
            storage.AddressList[ newname ] = storage.AddressList[ name ]
            delete( storage.AddressList, name )
            storage.CoinList[ newname ] = storage.CoinList[ name ]
            delete( storage.CoinList, name )
            fmt.Printf("Rename Address [%v] to [%v]:\n", name, newname)
            name = newname
        }
    }else{
        storage.AddressList[ name ] = &zklib.Address{}
        storage.AddressList[ name ].GetAddress()
        fmt.Printf("Get New Address [%v]:", name)
    }
    fmt.Println(storage.AddressList[name].String())
    return nil
}
