package main

/*
*/
import "C"
import "fmt"
import "errors"
import "strconv"

func Init( params []string ){
    depReq := ReqDeploy()
    resp, err := httpPostForm( depReq )
    check( err )
    params = append( params, string(resp[ 'result' ][ 'message' ]) )

    queReq := ReqQuery( params[0], 'params' )
    resp, err = httpPostForm( depReq )
    check( err )
    params = append( params, string(resp[ 'result' ][ 'message' ] ) )

    acuReq := ReqQuery( params[0], 'accumlator' )
    resp, err = httpPostForm( acuReq )
    check( err )
    params = append( params, string(resp[ 'result' ][ 'message' ] ) )

    return params
}

func Coinbase( params []string ){
    baseReq := ReqCoinbase( params[0], 'testuser1', 100 )
    resp , err := httpPostForm( baseReq )
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

    mintReq := ReqMint( params[0], 'testuser1', C.GoString(commint) )
    resp, err := httpPostForm( mintReq )

    check( err )
}
