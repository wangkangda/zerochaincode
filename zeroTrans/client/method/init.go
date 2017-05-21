package method

import(
    "fmt"
    "log"
    "errors"
    "encoding/json"
    "github.com/wangkangda/zerochaincode/zeroTrans/client/rpc"
    "github.com/wangkangda/zerochaincode/zeroTrans/client/storage"
)

const(
    ChaincodePath = "github.com/wangkangda/zerochaincode/zeroTrans/chaincode"
)

func CmdInit()(error){
    req := fmt.Sprintf( rpc.DeployTemplate, ChaincodePath )
    log.Printf("Get Deploy Request:%v\n", req)
    res, err := rpc.JsonSend( []byte(req) )
    if err != nil{
        return err
    }
    resjson := make(map[string]interface{})
    err = json.Unmarshal( res, &resjson )
    if err != nil{
        return err
    }
    result, ok := resjson["result"].(map[string]interface{})
    if !ok {
        return errors.New("Error Json For No Result Field")
    }
    storage.ChaincodeId, ok = result["message"].(string)
    if !ok {
        return errors.New("Error Json Result For No Message Field")
    }
    return nil
}
