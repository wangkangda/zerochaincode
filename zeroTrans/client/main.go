package main

import(
    "os"
    "fmt"
    "bufio"
    "strings"
    //"github/wangkangda/zerochaincode/zeroTrans/client/onchain"
    //"github/wangkangda/zerochaincode/zeroTrans/client/storage"
)

func main(){
    //storage.GetStorage()
    //defer storage.SaveStorage()

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
            continue
        case "quit":
            fmt.Println("Get quit command")
            return
        default:
            fmt.Println("Not implement function")
            continue
        }
    }
}
