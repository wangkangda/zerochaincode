package main
/*
#cgo CFLAGS: -l./lib
#cgo LDFLAGS: -L./lib -lhello -lzerocoin -Wl,-rpath,./lib/
#include "Goapi.h"
#include "hello.h"
*/
import "C"
import "fmt"
import "os"
import "errors"
import "strconv"
import "encoding/json"

func getResp( resp []byte ) map[string]interface{} {
    var data map[string]interface{}
    if err:=json.Unmarshal( resp, &data ); err!=nil{
        check(err)
    }
    fmt.Println(data)
    return data
}

func Init( params []string ){
    depReq := ReqDeploy()
    resp, err := httpPostForm( depReq )
    check( err )
    dat := getResp( resp )
    params = append( params, string(dat[ 'result' ][ 'message' ]) )

    queReq := ReqQuery( params[0], 'params' )
    resp, err = httpPostForm( depReq )
    check( err )
    dat1 := getResp( resp )
    params = append( params, string(dat1[ 'result' ][ 'message' ] ) )

    acuReq := ReqQuery( params[0], 'accumlator' )
    resp, err = httpPostForm( acuReq )
    check( err )
    dat2 := getResp(resp)
    params = append( params, string(dat2[ 'result' ][ 'message' ] ) )

    return params
}

func Coinbase( params []string ){
    baseReq := ReqCoinbase( params[0], 'testuser1', 100 )
    resp , err := httpPostForm( baseReq )
    data := getResp(resp)
    check( err )
}

func Mint( params []string, ){
    oParams := C.CCParamsLoad( params[1] )
    defer C.CCParamsDel( oParams )
    oAccum := C.CCAccumLoad( params[2] )
    defer C.CCAccumDel( oAccum )
    oPricoin := C.CCPricoinGen( oParams )
    defer C.CCPricoinDel( oPricoin )
    commint := C.CCPubcoinGen( oParams, oPricoin )
    defer C.CCStrDel( commint )
    accum := C.CCAccumCal( oParams, oAccum, commint )
    defer C.CCStrDel( accum )
    params[2] = C.GoString( accum )

    mintReq := ReqMint( params[0], C.GoString(commint), 'No implement')
    resp, err := httpPostForm( mintReq )
    check( err )
    fmt.Println(resp)
}
