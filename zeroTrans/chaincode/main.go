package main

import (
    "fmt"
    "log"
    //"errors"
    //"strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"

    "github.com/wangkangda/zerochaincode/zeroTrans/zklib/transaction"
    "github.com/wangkangda/zerochaincode/zeroTrans/chaincode/method"
)

type ZeroChaincode struct {
}

func (t *ZeroChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string)([]byte, error){
    //初始化参数
    return nil, nil
}

func (t *ZeroChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
    log.Printf( "Invoke %v\n", function )
    inputs, outputs := transaction.GetTransaction( stub, args )
    h := method.NewTransaction(stub, inputs, outputs)
    return h.Execute()
}

func (t *ZeroChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
    log.Printf( "Query %v\n", function)
    h := method.NewQuery(stub, function, args)
    return h.Execute()
}

func main(){
    err := shim.Start(new(ZeroChaincode))
    if err != nil{
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}
