package method

import(
    "fmt"
    "log"
    "github.com/wangkangda/zerochaincode/zeroTrans/client/rpc"
    "github.com/wangkangda/zerochaincode/zeroTrans/client/storage"
)

const(
    ChaincodePath = "github/wangkangda/zerochaincode/zeroTrans/chaincode"
)

func CmdInit()(error){
    req := fmt.Sprintf( rpc.DeployTemplate, ChaincodePath )
    log.Printf("Get Deploy Request:%v\n", req)
    res, err := rpc.JsonSend( []byte(req) )
    if err != nil{
        return err
    }
    storage.ChaincodeId = string(res)
    return nil
}
