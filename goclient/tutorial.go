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
    //fmt.Println(data)
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

    acuReq := ReqQuery( params[0], "accumulator" )
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

func Tutorial( params []string ){
    //mint
    p1 := C.CString( params[1] )
    //defer C.CCStrDel(p1)
    fmt.Println("Get params", C.GoString(p1) )
    oParams := C.CCParamsLoad( p1 )
    C.CCStrDel
    defer C.CCParamsDel( oParams )
    oPricoin := C.CCPricoinGen( oParams )
    defer C.CCPricoinDel( oPricoin )
    commint := C.CCPubcoinGen( oParams, oPricoin )
    defer C.CCStrDel( commint )

    mintReq := ReqMint( params[0], C.GoString(commint), "No implement")
    resp := httpPostForm( mintReq )
    fmt.Println(resp)
    ccid := getResp( resp )

    //make accumlator
    a1 := C.CString( params[2] )
    defer C.free( unsafe.Pointer(a1) )
    oAccum1 := C.CCAccumLoad( oParams, a1 )
    defer C.CCAccumDel( oAccum1 )
    for i:=0; i<5; i++ {
        newcoin := C.CCPricoinGen( oParams )
        newpubcoin := C.CCPubcoinGen( oParams, newcoin )

        mintReq := ReqMint( params[0], C.GoString(newpubcoin), "No implement")
        resp := httpPostForm( mintReq )
        fmt.Println( resp )

        a2 := C.CCAccumCal( oParams, oAccum1, newpubcoin )
        C.CCAccumDel( oAccum1 )
        oAccum1 = C.CCAccumLoad( oParams, a2 )
        C.CCPricoinDel( newcoin )
        C.CCStrDel(newpubcoin)
        C.CCStrDel(a2)
    }


    //spend
    toaddress := "testuser2"
    coinspend := C.CCSpendGen( oParams, oPricoin, oAccum1, C.CString(toaddress) )
    defer C.CCStrDel( coinspend )
    spendReq := ReqSpend( ccid, C.GoString(coinspend), "testuser2" )
    resp = httpPostForm( spendReq )
    fmt.Println(resp)
    sn := getResp(resp)
    fmt.Println("SN: ", sn )
}

