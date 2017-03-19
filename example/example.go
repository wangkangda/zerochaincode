package main

/*
#cgo CFLAGS: -l/opt/gopath/src/github.com/wangkangda/zerochaincode/example/lib
#cgo LDFLAGS: -L/opt/gopath/src/github.com/wangkangda/zerochaincode/example/lib -lzerocoin -lboost_system -Wl,-rpath,/opt/gopath/src/github.com/wangkangda/zerochaincode/example/lib/
#include "Goapi.h"
*/
import "C"
import "fmt"
import "errors"
import "strconv"

import "github.com/hyperledger/fabric/core/chaincode/shim"


// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    if len(args) != 0 {
        return nil, fmt.Errorf("Number of parameter error!!!")
    }

	//init commit counter
	counter := 0
	stub.PutState( "counter", []byte( strconv.Itoa(counter) ) )

	//generate params
	var params = C.CCParamsGen()
	stub.PutState( "params", []byte( C.GoString(params) ) )

	//get params obj
	oParams := C.CCParamsLoad( params )
	C.CCStrDel(params)

	//generate accumulator
	accum := C.CCAccumGen( oParams )
	stub.PutState( "accumulator", []byte( C.GoString(accum) ) )
    stub.PutState( "origin_accum", []byte( C.GoString(accum) ) )
	C.CCStrDel(accum)

	//release object params
	C.CCParamsDel( oParams )

	return nil, nil
}

func (t *SimpleChaincode) Transaction(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
    if len(args)<1 {
        return nil, fmt.Errorf("Number of Parameter not enough!")
    }
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

            if len(args)!=5 {
                return nil, fmt.Errorf("Number of Parameter Error!")
            }
			fmt.Println("get signature: %s", signature)
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
                toAmount = []byte("0")
				//return nil, fmt.Errorf("the user (to: %s) not exists!", toAddress)
			}

			//then verify the signature, implement later

			iFromAmount, _ := strconv.Atoi( string(fromAmount) )
			iToAmount, _ := strconv.Atoi( string(toAmount) )

			if iFromAmount < amount {
				return nil, fmt.Errorf("the amount not enough!")
			}

			iFromAmount = iFromAmount - amount
			iToAmount = iToAmount + amount

			err = stub.PutState(fromAddress, []byte(strconv.Itoa(iFromAmount)))
			if err != nil{
				return nil, err
			}
			err = stub.PutState(toAddress, []byte(strconv.Itoa(iToAmount)))
			if err != nil{
				return nil, err
			}

		case "mint":
			fromAddress := args[1]
			commitment := args[2]
            signature := args[3]

            fmt.Println("signature: ", signature)
            if len(args)<4 {
                return nil, fmt.Errorf("Number of Parameter not enough!")
            }

			fromAmount, err := stub.GetState( fromAddress )
			if err != nil {
				return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
			}
			if fromAmount == nil {
				return nil, fmt.Errorf("the user (to: %s) not exists!", fromAddress)
			}

			iFromAmount , _ := strconv.Atoi( string(fromAmount) )

			if iFromAmount <= 0 {
				return nil, fmt.Errorf("the amount not enough!")
			}
			iFromAmount = iFromAmount - 1
			err = stub.PutState(fromAddress, []byte(strconv.Itoa(iFromAmount)))
			if err != nil {
				return nil, err
			}

			counter, err := stub.GetState( "counter" )
			if err != nil {
				return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
			}

			//save the commitment
			iCounter, _ := strconv.Atoi( string(counter) )
			err = stub.PutState("commitment"+strconv.Itoa(iCounter), []byte(commitment))
			if err != nil {
				return nil, err
			}

			err = stub.PutState("counter", []byte(strconv.Itoa(iCounter+1)))
			if err != nil {
				return nil, err
			}

			//process the accumulator
			bAccum, _ := stub.GetState("accumulator")
			bParams, _ := stub.GetState("params")
			accum := string( bAccum )
			params := string( bParams )
			oParams := C.CCParamsLoad( C.CString(params) )
			oAccum := C.CCAccumLoad( oParams, C.CString(accum) )
			csAccum := C.CCAccumCal( oParams, oAccum, C.CString(commitment) )
			err = stub.PutState("accumulator", []byte(C.GoString(csAccum)))
			if err != nil {
				return nil, err
			}

			//release object
			C.CCParamsDel( oParams )
			C.CCAccumDel( oAccum )
			C.CCStrDel( csAccum )

            res := strconv.Itoa(iCounter)
            fmt.Println("counter: ", iCounter)
            fmt.Println("mintid: ", res)
			return []byte(res), nil

		case "spend":
			coinspend := args[1]
			toAddress := args[2]
			metadata := args[3]

			fmt.Println( "get metadata: %s", metadata )

			bAccum, _ := stub.GetState("accumulator")
			bParams, _ := stub.GetState("params")
            accum := string( bAccum )
			params := string( bParams )
			oParams := C.CCParamsLoad( C.CString(params) )
			oAccum := C.CCAccumLoad( oParams, C.CString(accum) )

			serialNum := C.CCSpendVerify( oParams, C.CString(coinspend), C.CString(toAddress), oAccum)
			if serialNum == nil {
				return nil, fmt.Errorf("The CoinSpend transaction did not verify!")
			}

            toAmount, err := stub.GetState( toAddress )
            if err != nil {
                return nil, fmt.Errorf("get operation failed.ERROR %s", err)
            }
            if toAmount == nil {
                toAmount = []byte("0")
            }
            iToAmount, _ := strconv.Atoi( string(toAmount) )
            iToAmount = iToAmount + 1
            err = stub.PutState(toAddress, []byte(strconv.Itoa(iToAmount)))
            if err != nil{
                return nil, err
            }
            sn := C.GoString( serialNum )
            C.CCStrDel( serialNum )

			return []byte(sn), nil
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

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("invoke is running" + function)
    if function == "transaction" {
        return t.Transaction(stub, "transaction", args)
    }
    fmt.Println("invoke did not find func: "+function)

    return nil, errors.New("Received unknown function invocation: "+function)
}


//Query
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error){
	fmt.Println("query is running"+function)

	queryvalue := args[0]
	switch queryvalue {
	case "amount":
		address := args[1]
		amount, err := stub.GetState(address)
		if err != nil { 
			return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
		}
		if amount == nil {
			return nil, fmt.Errorf("the user (from: %s) not existes!", address)
		}
		return amount, nil
	case "params":
		params, err := stub.GetState("params")
		if err != nil {
			return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
		}
		return params, nil
	case "origin_accum":
		accum, err := stub.GetState("origin_accum")
		if err != nil {
			return nil, fmt.Errorf("get operation failed. Error accessing state: %s", err)
		}
		return accum, nil
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
		return counter, nil
	case "commitment":
		return t.Commitment(stub, "commitment", args)
	}

	return nil, errors.New("Received unknown function query: "+function)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
