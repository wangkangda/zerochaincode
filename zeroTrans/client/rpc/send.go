package rpc

import(
    //"fmt"
    "errors"
    "bytes"
    "io/ioutil"
    "net/http"
)

const DeployTemplate = `{
    "jsonrpc": "2.0",
    "method": "deploy",
    "params":{
        "type": 1,
        "chaincodeID": {
            "path":"%s"
        },
        "ctorMsg":{
            "function": "init",
            "args": []
        }   
    },
    "id": 1
}`
const TransTemplate = `{
    "jsonrpc": "2.0",
    "method": "invoke",
    "params":{
        "type": 1,
        "chaincodeID": {
            "name":"%s"
        },
        "ctorMsg":{
            "function": "transaction",
            "args": ["%s"]
        }   
    },
    "id": 1
}`
const QueryTemplate = `{
    "jsonrpc": "2.0",
    "method": "query",
    "params":{
        "type": 1,
        "chaincodeID": {
            "name":"%s"
        },
        "ctorMsg":{
            "function": "%s",
            "args": ["%s"]
        }   
    },
    "id": 1
}`




const ChainUrl = "http://localhost:7050/chaincode"

func JsonSend( jsonreq []byte )([]byte, error){
    req, err := http.NewRequest("POST", ChainUrl, bytes.NewBuffer(jsonreq))
    if err != nil{
        return nil, err
    }
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    return body, err
    
    //fmt.Printf("Request: %v\n", req)
    return nil, errors.New("Not Implement")
}
