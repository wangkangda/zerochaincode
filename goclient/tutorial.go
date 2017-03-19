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
import "strconv"

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

func getCounter( params []string )int{
    queReq := ReqQuery( params[0], "counter", nil )
    resp := httpPostForm( queReq )
    data := getResp( resp )
    res, err := strconv.Atoi( string(data) )
    check(err)
    return res
}
func getCommitment( params []string , index int )string{
    si := strconv.Itoa( index )
    queReq := ReqQuery( params[0], "commitment", []string{si} )
    resp := httpPostForm( queReq )
    return getResp( resp )
}
func getAmount( params []string, address string ) int{
    queReq := ReqQuery( params[0], "amount", []string{address} )
    resp := httpPostForm( queReq )
    data := getResp( resp )
    res, err := strconv.Atoi( string(data) )
    check(err)
    return res
}
func transfer( params[]string, fromuser string, touser string, amount int ) string{
    traReq := ReqTransfer( params[0], fromuser, touser, amount )
    resp := httpPostForm( traReq )
    return getResp( resp )
}
func mint( params[]string, fromuser string )(int, string){
    //mint
    p1 := C.CString( params[1] )
    defer C.CCStrDel( p1 )
    fmt.Println("Get params", C.GoString(p1) )
    oParams := C.CCParamsLoad( p1 )
    defer C.CCParamsDel( oParams )
    pricoin := C.CCPricoinGen2( oParams )
    defer C.CCStrDel( pricoin )
    sPricoin := C.GoString( pricoin )
    oPricoin := C.CCPricoinLoad( oParams, pricoin )
    defer C.CCPricoinDel( oPricoin )
    commint := C.CCPubcoinGen( oParams, oPricoin )
    defer C.CCStrDel( commint )

    mintReq := ReqMint( params[0], fromuser, C.GoString(commint) )
    resp := httpPostForm( mintReq )
    fmt.Println(resp)
    mintid := getCounter( params ) - 1
    return mintid, sPricoin
}
func getWitness( params []string, mintid int )string{
    p := C.CString( params[1] )
    defer C.CCStrDel( p )
    oParams := C.CCParamsLoad( p )
    mintnum := getCounter( params )
    for index := len(params)-3; index<mintnum; index++{
        params = append(params, getCommitment(params, index))
    }
    sAccum := C.CString( params[2] )
    oAccum := C.CCAccumLoad( oParams, sAccum )
    for i:=0; i<mintnum; i++{
        if( i==mintid ){
            continue
        }
        C.CCStrDel( sAccum )
        ptmp := C.CString( params[3+i] )
        sAccum = C.CCAccumCal( oParams, oAccum, ptmp )
        C.CCAccumDel( oAccum )
        oAccum = C.CCAccumLoad( oParams, sAccum )
        C.CCStrDel( ptmp )
    }
    C.CCAccumDel( oAccum )
    res := C.GoString( sAccum )
    C.CCStrDel( sAccum )
    return res
}
func spend( params []string, accum string, pricoin string, recvUser string)string{
    p1 := C.CString( params[1] )
    defer C.CCStrDel( p1 )
    oParams := C.CCParamsLoad( p1 )
    defer C.CCParamsDel( oParams )
    sPricoin := C.CString( pricoin )
    defer C.CCStrDel( sPricoin )
    oPricoin := C.CCPricoinLoad( oParams, sPricoin )
    defer C.CCPricoinDel( oPricoin )
    oAccum := C.CCAccumLoad( oParams, C.CString(accum) )
    defer C.CCAccumDel( oAccum )

    coinspend := C.CCSpendGen( oParams, oPricoin, oAccum, C.CString(recvUser) )
    defer C.CCStrDel( coinspend )
    spendReq := ReqSpend( params[0], C.GoString(coinspend), "testuser2" )
    resp := httpPostForm( spendReq )
    sn := getResp(resp)
    fmt.Println( "Spend pricoin sucess! SerialNum: ", sn )
    return sn
}

func Init( params []string )[]string{
    depReq := ReqDeploy()
    resp := httpPostForm( depReq )
    dat := getResp( resp )
    params = append( params, string(dat) )

    queReq := ReqQuery( params[0], "params", nil )
    resp = httpPostForm( queReq )
    dat1 := getResp( resp )
    params = append( params, string(dat1) )

    acuReq := ReqQuery( params[0], "origin_accum", nil )
    resp = httpPostForm( acuReq )
    dat2 := getResp(resp)
    params = append( params, string(dat2 ) )

    return params
}

func Coinbase( params []string, user string )string{
    baseReq := ReqCoinbase( params[0], user, 100 )
    resp := httpPostForm( baseReq )
    data := getResp(resp)
    return data
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

