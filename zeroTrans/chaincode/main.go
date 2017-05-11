package main

import (
    "fmt"
    "log"
    "errors"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"

    "github.com/wangkangda/zerochaincode/zeroTrans/zklib/transaction"
    "github.com/wangkangda/zerochaincode/zeroTrans/method"
)

type ZeroChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string)([]byte, error){
    //初始化参数

}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
    log.Info( "Invoke %v", function )
    inputs, outputs := transaction.GetTransaction( stub, args )
    h = method.NewTransaction(&stub, inputs, outputs)
    return h.Execute()
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
    log.Info( "Query %v", function)
    h = method.NewQuery(&stub, function, args)
    return h.Execute()
}

func main(){
    err := shim.Start(new(SimpleChaincode))
    if err != nil{
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}
