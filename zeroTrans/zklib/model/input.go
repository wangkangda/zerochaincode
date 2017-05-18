package model

import (
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

const(
    NormalTransaction = 0
    PrivacyTransaction = 1
)

type Input interface{
    GetType()int
    Prepare(Context)
    Verify(Context)bool
    Execute(Context)error
    String()(string)
    FromString(string)error
}

