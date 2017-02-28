package main

/*
#cgo CFLAGS: -l/opt/gopath/src/github.com/wangkangda/zerochaincode/chaincode_gocoin/lib
#cgo LDFLAGS: -L/opt/gopath/src/github.com/wangkangda/zerochaincode/chaincode_gocoin/lib -lzerocoin -lboost_system -Wl,-rpath,/opt/gopath/src/github.com/wangkangda/zerochaincode/chaincode_gocoin/lib/
#include "Goapi.h"
*/
import "C"
import "fmt"
import "errors"
import "strconv"

import "github.com/hyperledger/fabric/core/chaincode/shim"

type SimpleChaincode struct{
}

func main(){
  err := shim.Start( new(SimpleChaincode) )
  if err != nil{
    fmt.Printf("Error starting Simple chaincodes", err)
  }
}

//Init param and address
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	//generate params
	var params = C.CCParamsGen()
	stub.PutState( "params", []byte( C.GoString(params) ) )

	//get params obj
	oParams := C.CCParamsLoad( params )
	C.CCStrDel(params)

	//init commit counter
	counter := 0
	stub.PutState( "counter", []byte( strconv.Itoa(counter) ) )

	//generate accumulator
	accum := C.CCAccumGen( oParams )
	stub.PutState( "accumlator", []byte( C.GoString(accum) ) )
	C.CCStrDel(accum)

	//release object params
	//defer C.free( oParams )//??
	C.CCParamsDel( oParams )

	return nil, nil
}

func (t *SimpleChaincode) Transaction(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	var transtype = args[0]
	switch transtype{
		case "coinbase":
			address := args[1]
			amount, err := strconv.Atoi( args[2] )
			err = stub.PutState(address, []byte(strconv.Itoa(amount)))
			if err != nil{
				return nil, err
			}

		case "transfer":
			fromAddress := args[1]
			toAddress := args[2]
			amount, err := strconv.Atoi( args[3] )
			signature := args[4]

			fromAmount, err := stub.GetState( fromAddress )
			if err != nil {
				return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
			}
			if fromAmount == nil {
				return nil, fmt.Errorf("the user (from: %s) not exists!", fromAddress)
			}

			toAmount, err := stub.GetState( toAddress )
			if err != nil {
				return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
			}
			if toAmount == nil {
				return nil, fmt.Errorf("the user (to: %s) not exists!", toAddress)
			}
	  
			//then verify the signature, implement later


			if fromAmount < amount {
				return nil, fmt.Errorf("the amount not enough!")
			}

			fromAmount = fromAmount - amount
			toAmount = toAmount + amount

			err = stub.PutState(fromAddress, []byte(strconv.Itoa(fromAmount)))
			if err != nil{
				return nil, err
			}
			err = stub.PutState(toAddress, []byte(strconv.Itoa(toAmount)))
			if err != nil{
				return nil, err
			}

		case "mint":
			fromAddress := args[1]
			commitment := args[2]

			fromAmount, err := stub.GetState( fromAddress )
			if err != nil {
				return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
			}
			if fromAmount == nil {
				return nil, fmt.Errorf("the user (to: %s) not exists!", fromAddress)
			}

			if fromAmount <= 0 {
				return nil, fmt.Errorf("the amount not enough!")
			}
			fromAmount = fromAmount - 1
			err = stub.PutState(fromAddress, []byte(strconv.Itoa(fromAmount)))
			if err != nil {
				return nil, err
			}

			counter, err := stub.GetState( "counter" )
			if err != nil {
				return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
			}
			
			//save the commitment
			err = stub.PutState("commitment"+strconv.Itoa(counter), []byte(commitment))
			if err != nil {
				return nil, err
			}

			counter = counter + 1
			err = stub.PutState("counter", []byte(strconv.Itoa(counter)))
			if err != nil {
				return nil, err
			}

			//process the accumulator
			accum := string( stub.GetState("accumulator") )
			params := string( stub.GetState("params") )
			oParams := C.CCParamsLoad( C.CString(params) )
			oAccum := C.CCAccumLoad( oParams, C.CString(accum) )
			accum = C.CCAccumCal( oParams, oAccum, C.CString(commitment) )
			err = stub.PutState("accumulator", []byte(C.GoString(accum)))
			if err != nil {
				return nil, err
			}

			//release object
			C.CCParamsDel( oParams )
			C.CCAccumDel( oAccum )
			C.CCStrDel( accum )

			return nil, nil

		case "spend":
			coinspend := args[1]
			toAddress := args[2]
			metadata := args[3]

			accum := string( stub.GetState("accumulator") )
			params := string( stub.GetState("params") )
			oParams := C.CCParamsLoad( C.CString(params) )
			oAccum := C.CCAccumLoad( oParams, C.CString(accum) )

			serialNum := C.CCSpendVerify( oParams, C.CString(coinspend), C.CString(toAddress), oAccum)
			if serialNum == nil {
				return nil, fmt.Errorf("The CoinSpend transaction did not verify!")
			}

			return nil, nil
	}
	return nil, nil
}

func (t *SimpleChaincode) Commitment(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	num,_ := strconv.Atoi( string( args[0] ) )
	counter, err := stub.GetState( "counter" )
	if err != nil {
		return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
	}
	iCounter,_ := strconv.Atoi( string(counter) )

	if num >= iCounter {
		return nil, fmt.Errorf("there is not No.%s commitment", args[0] )
	}
	commit, err := stub.GetState("commitment"+strconv.Itoa(num))
	if err != nil{
		return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
	}

	return commit, nil
}

//Invoke
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
  fmt.Println("invoke is running" + function)
  if function == "init" {
    return t.Init(stub, "init", args)
  }
  if function == "transaction" {
    return t.Transaction(stub, "transaction", args)
  }
  fmt.Println("invoke did not find func: "+function)

  return nil, errors.New("Received unknown function invocation: "+function)
}

//Query
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	fmt.Println("query is running"+function)

	switch function {
	case "amount":
		address := args[0]
		amount, err := stub.GetState(address)
		if err != nil { 
			return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
		}
		if amount == nil {
			return nil, fmt.Errorf("the user (from: %s) not existes!", address)
		}
		//iAmount, _ := strconv.Atoi(string(amount))
		return amount, nil
	case "params":
		params, err := stub.GetState("params")
		if err != nil {
			return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
		}
		return params, nil
	case "accumulator":
		accum, err := stub.GetState("accumulator")
		if err != nil {
			return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
		}
		return accum, nil
	case "counter":
		counter, err := stub.GetState("counter")
		if err != nil {
			return nil, fmt.Errorf("get opeartion failed. Error accessing state: %s", err)
		}
		//iCounter, _ := strconv.Atoi(string(counter))
		return counter, nil
	case "commitment":
		return t.Commitment(stub, "commitment", args)
	}

	fmt.Println("query did not find func: " + function)
	return nil, errors.New("Received unknown function query: "+function)
}


