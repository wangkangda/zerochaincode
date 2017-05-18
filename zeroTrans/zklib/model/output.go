package model

import (
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

type Output interface{
    GetType() int
    Prepare(Context)
    Verify(Context)bool
    Execute(Context)error
    String()string
    FromString(string)error
}
