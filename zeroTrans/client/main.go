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
        fmt.Println(" ZeroChaincode is a anonymous transaction system based on zero knowledge proof ")
        fmt.Println("Usage:")
        fmt.Println("       goclient command [arguments]")
        fmt.Println("")
        fmt.Println("The commands are:")
        fmt.Println("   address")
        fmt.Println("   coingen")
        fmt.Println("   zkproof")
        fmt.Println("   transaction")
        fmt.Println("   query")
        return
    }

    function = os.Args[1]
    switch function{
    case "Init":
        onchain.Init()
    }


}
