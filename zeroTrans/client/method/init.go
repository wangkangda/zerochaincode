package method

import(
    "github.com/wangkangda/zerochaincode/zeroTrans/client/rpc"
    "github.com/wangkangda/zerochaincode/zeroTrans/client/storage"
)

const(
    ChaincodePath = "github/wangkangda/zerochaincode/zeroTrans/chaincode"
)

func CmdInit()(error){
    args := make([]string)
    req := fmt.Sprintf( rpc.DeployTemplate, ChaincodePath )
    log.Printf("Get Deploy Request:%v\n", req)
    res, err := rpc.JsonSend( []byte(req) )
    if err != nil{
        return err
    }
    storage.ChaincodeId = string(res)
}
