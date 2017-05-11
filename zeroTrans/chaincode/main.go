package main

import (
    "fmt"
    "log"
    "errors"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"

    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
    "github.com/wangkangda/zerochaincode/zeroTrans/method"
)

type ZeroChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string)([]byte, error){
    //初始化参数

}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
    log.Info( "Invoke %v", function )
    inputs, outputs := zklib.GetTransaction()
    h = method.NewTransaction(stub, inputs, outputs)
    return h.Execute()
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
    log.Info( "Query %v", function)
    h = method.NewQuery(stub, function, args)
    return h.Execute()
}

