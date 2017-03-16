package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
	"bytes"
    "os"
    "strconv"
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
        }   
    },
    "id": 1
}`
func ReqDeploy()([]byte){
    chaincode := `"path":"github.com/wangkangda/zerochaincode/example"`
    jsonreq := fmt.Sprintf(json_temp, "deploy", chaincode, "init", "")
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
    args := fmt.Sprintf(`"mint", "%s", "%s", "not implement"`, sendUsr, commint)
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
func ReqQuery(chaincodeid string, reqvalue string, otherargs []string)[]byte{
    args := fmt.Sprintf(`"%s"`, reqvalue)
    for i:=0; i<len(otherargs); i++ {
        args = fmt.Sprintf(`%s, "%s"`, args, otherargs[i] )
    }
    chaincode := fmt.Sprintf(`"name": "%s"`, chaincodeid)
    jsonreq := fmt.Sprintf(json_temp, "query", chaincode, "query", args)
    return []byte(jsonreq)
}
func httpGet() {
	resp, err := http.Get("http://localhost:7050/network/peers")
	if err != nil {
		// handle error
		fmt.Println("Error\n")
		return
	}
    fmt.Println( resp.Header, resp.ContentLength, resp.Close )
    fmt.Println("transfer", resp.TransferEncoding)
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
    fmt.Println("REQUEST:>", string(jsonStr))

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
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body) )
    return body
}

func testPost() {
    url := "http://localhost:7050/chaincode"
    fmt.Println("URL:>", url)
    chaincode := `"path":"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02"`
    jsonStr := fmt.Sprintf(json_temp, `deploy`, chaincode, `init`, `"a", "1000", "b", "2000"`)
    fmt.Println("deploy req:", jsonStr)
    /*
    jsonStr := []byte(`{
        "jsonrpc":"2.0",
		"method":"deploy",
		"params": {
            "type": 1,
			"chaincodeID":{
                "path":"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02"
            },
			"ctorMsg": {
                "function":"init",
                "args":{"a", "1000", "b", "2000"}
            },
            "id": 1
        }
    }`)*/
    req, err := http.NewRequest("POST", url, bytes.NewBuffer( []byte(jsonStr) ))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    fmt.Println( resp.Header, resp.ContentLength, resp.Close )
    fmt.Println("transfer", resp.TransferEncoding)
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
    arg_num := len(os.Args)
    if arg_num <= 1 {
        testPost()
        httpGet()
    }

    pathfile := `chaincode.dat`
    pricoinfile := `pricoinfile.dat`
    params, err := getData(pathfile)
    check(err)
    defer saveData(pathfile, params)
    pricoins, err := getCommit(pricoinfile)
    check(err)
    defer saveCommit(pricoinfile, pricoins)

    if len(params)==0 {
        fmt.Println("empty storage")
        params = Init(params)
    }
	fmt.Println(len(params))

    //Tutorial( params )
    if os.Args[1] == `coinbase` {
        if arg_num <= 2{
            fmt.Println("argument not enough")
            return
        }
        fmt.Println( Coinbase(params, os.Args[2]) )
        return
    }else if os.Args[1] == `transfer` {
        if arg_num != 4{
            fmt.Println("argument not enough")
            return
        }
        amount, _ := strconv.Atoi( os.Args[3] )
        fmt.Println( transfer(params, os.Args[1], os.Args[2], amount) )
        return
    }else if os.Args[1] == `query` {
        if arg_num != 3{
            fmt.Println("argument not enough")
            return
        }
        fmt.Print("User ", os.Args[2])
        fmt.Println(": ", getAmount(params, os.Args[2]) )
    }else if os.Args[1] == `mint` {
        if arg_num != 3{
            fmt.Println("argument not enough")
            return
        }
        fmt.Print("Mint by ", os.Args[2])
        fmt.Println(": ", mint(params, os.Args[2]))
    }else if os.Args[1] == `spend` {
        if arg_num != 4{
            fmt.Println("argument not enough")
            return
        }
        mintid, _ := strconv.Atoi( os.Args[2] )
        recvuser := os.Args[3]
        fmt.Println("Get witness for Commitment ", mintid)
        witness := getWitness( params, mintid )
        pricoin, exist := pricoins[ mintid ]
        if !exist {
            fmt.Println( "Pricoins not exist" )
            return
        }
        fmt.Print("Spend to ", recvuser )
        fmt.Println(" for coin sn:", spend(params, witness, pricoin, recvuser) )
    }
    //fmt.Println(params)

	//httpPostForm()
}
