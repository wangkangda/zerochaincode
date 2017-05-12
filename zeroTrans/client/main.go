package main

import(
    "os"
    "fmt"
    "github/wangkangda/zerochaincode/zeroTrans/client/onchain"
    "github/wangkangda/zerochaincode/zeroTrans/client/storage"
)

func main(){
    storage.GetStorage()
    defer storage.SaveStorage()

    arg_num :=  len(os.Args)
    if arg_num <= 1{
        fmt.Println("Please Input Command!")
    }

    function = os.Args[1]
    switch function{
    case "Init":
        onchain.Init()
    }


}
