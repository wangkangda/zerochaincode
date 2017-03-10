package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
)
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

func httpPostForm() {
    url := "http://localhost:7050/chaincode"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{
  "jsonrpc": "2.0",
  "method": "deploy",
  "params": {
    "type": 1,
    "chaincodeID":{
        "path":"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02"
    },
    "ctorMsg": {
        "args":["init", "a", "1000", "b", "2000"]
    }
  },
  "id": 1
}`)
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
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

}


func main(){
	httpGet()
	httpPostForm()
}
