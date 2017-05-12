package rpc

import(
    "io/ioutil"
    "net/http"
)

const JsonTemplate = `{
    "jsonrpc": "2.0",
    "method": "%s",
    "params":{
        "type": 1,
        "chaincodeID": {
            %s
        },
        "ctorMsg":{
            "function": "%s",
            "args": [%s]
        }   
    },
    "id": 1
}`

const ChainUrl = "http://localhost:7050/chaincode"

func JsonSend( req []byte )([]byte, error){
    req, err := http.NewRequest("POST", ChainUrl, bytes.NewBuffer(req))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return body
}
