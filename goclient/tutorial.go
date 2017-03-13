package main
/*
#cgo CFLAGS: -l./lib
#cgo LDFLAGS: -L./lib -lzerocoin -Wl,-rpath,./lib/
#include <stdlib.h>
#include "Goapi.h"
*/
import "C"
import "fmt"
import "encoding/json"
import "unsafe"

func getResp( resp []byte ) string {
    var data map[string]interface{}
    if err:=json.Unmarshal( resp, &data ); err!=nil{
        check(err)
    }
    fmt.Println(data)
    result := data[ "result" ].(map[string]interface{})
    message := result[ "message" ].(string)
    return message
}

func Init( params []string )[]string{
    depReq := ReqDeploy()
    resp := httpPostForm( depReq )
    dat := getResp( resp )
    params = append( params, string(dat) )

    queReq := ReqQuery( params[0], "params" )
    resp = httpPostForm( queReq )
    dat1 := getResp( resp )
    params = append( params, string(dat1) )

    acuReq := ReqQuery( params[0], "accumlator" )
    resp = httpPostForm( acuReq )
    dat2 := getResp(resp)
    params = append( params, string(dat2 ) )

    return params
}

func Coinbase( params []string ){
    baseReq := ReqCoinbase( params[0], "testuser1", 100 )
    resp := httpPostForm( baseReq )
    data := getResp(resp)
    fmt.Println( data )
}

func Mint( params []string, ){
    p1 := C.CString( params[1] )
    //defer C.free( unsafe.Pointer(p1) )
    oParams := C.CCParamsLoad( p1 )
    defer C.CCParamsDel( oParams )
    oPricoin := C.CCPricoinGen( oParams )
    defer C.CCPricoinDel( oPricoin )
    commint := C.CCPubcoinGen( oParams, oPricoin )
    defer C.CCStrDel( commint )

    mintReq := ReqMint( params[0], C.GoString(commint), "No implement")
    resp := httpPostForm( mintReq )
    fmt.Println(resp)

    defer C.free( unsafe.Pointer(p1) )
}
