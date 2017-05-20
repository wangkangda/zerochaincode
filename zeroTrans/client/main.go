package main

import(
    "os"
    "fmt"
    "bufio"
    "strings"
    "github.com/wangkangda/zerochaincode/zeroTrans/client/method"
    "github.com/wangkangda/zerochaincode/zeroTrans/client/storage"
)

func main(){
    storage.GetStorage()
    defer storage.SaveStorage()

    reader := bufio.NewReader( os.Stdin )
    for true {
        fmt.Println("Please Input Command:")
        input, err := reader.ReadBytes('\n')
        if err != nil{
            fmt.Printf("Get %v while input\n", err)
            break
        }
        cmd := strings.Split(string(input[0:len(input)-1]), " ")
        if len(cmd)==0 {
            continue
        }
        function := cmd[0]
        fmt.Println("Get cmd:", function)
        switch function{
        case "init":
            //部署chaincode
            err = method.CmdInit()
            if err != nil{
                fmt.Printf("Get %v while execute Init cmd\n", err)
            }
        case "address":
            //新建/重命名地址，生成公私钥对
            err = method.CmdAddress( cmd )
            if err != nil{
                fmt.Printf("Get %v while execute Address cmd\n", err)
            }
        case "coin":
            err = method.CmdCoin( cmd )
            if err != nil{
                fmt.Printf("Get %v while execute Coin cmd\n", err)
            }
        case "quit":
            fmt.Println("Get quit command")
            return
        default:
            fmt.Println("Not implement function")
            continue
        }
    }
}
