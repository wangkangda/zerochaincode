package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
	"bytes"
    //"net/url"
    //"strings"

)
var json_temp = `{
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
        },
        "id": 1
    }
}`
func ReqDeploy()([]byte){
    chaincode := `"path":"github.com/wangkangda/zerochaincode/example"`
    jsonreq := fmt.Sprintf(json_temp, "deploy", chaincode, "init", "")
    fmt.Println("deploy req:", jsonreq)
    return []byte(jsonreq)
}
func ReqCoinbase(chaincodeid string, recvUsr string, amount int)[]byte{
    args := fmt.Sprintf(`"coinbase", "%s", "%s"`, recvUsr, amount)
    chaincode := fmt.Sprintf(`"name": "%s"`, chaincodeid)
    jsonreq := fmt.Sprintf(json_temp, "invoke", chaincode, "transaction", args)
    return []byte(jsonreq)
}
func ReqTransfer(chaincodeid string, sendUsr string, recvUsr string, amount int)([]byte){
    args := fmt.Sprintf(`"transfer", "%s", "%s", "%d", "not implement"`, sendUsr, recvUsr, amount)
    chaincode := fmt.Sprintf(`"name": "%s"`, chaincodeid)
    jsonreq := fmt.Sprintf(json_temp, "invoke", chaincode, "transaction", args)
    return []byte(jsonreq)
}
func ReqMint(chaincodeid string, sendUsr string, commint string)([]byte){
    args := fmt.Sprintf(`"mint", "%s", "%s"`, sendUsr, commint)
    chaincode := fmt.Sprintf(`"name": "%s"`, chaincodeid)
    jsonreq := fmt.Sprintf(json_temp, "invoke", chaincode, "transaction", args)
    return []byte(jsonreq)
}
func ReqSpend(chaincodeid string, coinspend string, recvUsr string)([]byte){
    args := fmt.Sprintf(`"spend", "%s", "%s", "not implement"`, coinspend, recvUsr)
    chaincode := fmt.Sprintf(`"name": "%s"`, chaincodeid)
    jsonreq := fmt.Sprintf(json_temp, "invoke", chaincode, "transaction", args)
    return []byte(jsonreq)
}
func ReqQuery(chaincodeid string, reqvalue string)[]byte{
    args := fmt.Sprintf(`"%s"`, reqvalue)
    chaincode := fmt.Sprintf(`"name": "%s"`, chaincodeid)
    jsonreq := fmt.Sprintf(json_temp, "query", chaincodeid, "query", args)
    return []byte(jsonreq)
}
func httpGet() {
	resp, err := http.Get("http://localhost:7050/network/peers")
	if err != nil {
		// handle error
		fmt.Println("Error\n")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpPostForm(jsonStr []byte) []byte{
    url := "http://localhost:7050/chaincode"
    fmt.Println("URL:>", url)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Body:", resp.Body)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", body)
    return body
}

func testPost() {
/*
			"jsonrpc":"2.0",
			"method":"deploy",
			"params": {
				"type": 1,
				"chaincodeID":{
					"path":"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02"},
				"ctorMsg": {
					"args":{"init", "a", "1000", "b", "2000"}},
		"id": 1}}*/
    url := "http://localhost:7050/chaincode"
    fmt.Println("URL:>", url)
    jsonStr := `{"jsonrpc":"2.0",
			"method":"deploy",
			"params": {
				"type": 1,
				"chaincodeID":{
					"path":"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02"},
				"ctorMsg": {
					"args":{"init", "a", "1000", "b", "2000"}},
		"id": 1}}`
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Body:", resp.Body)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", body)
}

func check(e error){
    if e != nil {
		fmt.Println(e)
        panic(e)
    }
}
func main(){
    testPost()
    /*
    pathfile := `chaincode.dat`
    params, err := getData(pathfile)
    check(err)
    if len(params)==0 {
        fmt.Println("empty storage")
        params = Init(params)
        return
    }
	fmt.Println(len(params))
    //params = append(params, "test test")
	fmt.Println(params)

    err = saveData(pathfile, params)
    check(err)
*/
    //httpGet()
	//httpPostForm()
}
