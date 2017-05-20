package method

import(
    "errors"
    _ "github.com/wangkangda/zerochaincode/zeroTrans/client/storage"
)

func CmdAddress( cmd ){
    newname, oldname := "", ""
    if len(cmd)==1 || len(cmd)>3{
        return errors.New("Error for parameter number")
    }else if len(cmd)==3{
        oldname = cmd[2]
    }
    newname = cmd[1]
    address, exist := storage.AddressList[ newname ]
    if exist {
        if oldname == ""{
            fmt.Printf("Address [%v] exist:\n", newname)
        }else{
            storage.AddressList[ newname ] = storage.AddressList[ oldname ]
            delete( storage.AddressList, oldname )
            fmt.Printf("Rename Address [%v] to [%v]:\n", oldname, newname)
        }
    }else{
        storage.AddressList[ newname ] = Address{}
        storage.AddressList[ newname ].GetAddress()
        fmt.Printf("Get New Address [%v]:", newname)
    }
    fmt.Println(storage.AddressList[newname].String())
    return nil
}
