package transaction

import(
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

func GetTransaction( stub *shim.ChaincodeStubInterface, params []string ) ([]*Input, []*Output){
    return nil, nil
}
