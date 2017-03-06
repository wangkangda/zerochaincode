package main

import(
    "fmt"
    "io/ioutil"
    "net/http"
    //"net/url"
    //"strings"
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
    metadata := url.Values{"jsonrpc":"2.0","method":"deploy","params": {"type": 1,"chaincodeID":{"path":"github.com/hyperledger/fabric/examples/chaincode/go/chaincode_example02"},"ctorMsg": {"args":{"init", "a", "1000", "b", "2000"}},"id": 1}}

	resp, err := http.PostForm("http://localhost:7050/chaincode",metadata)
	if err != nil {
		// handle error
		fmt.Println("Error")
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
func main(){
	httpGet()
	httpPostForm()
}
