package zklib
/*
#cgo CFLAGS: -g -Wno-unused-parameter -fPIC -Wno-unused-variable -I libzerocash/libzerocash
#cgo LDFLAGS: -L libzerocash -Wl,-rpath,./libzerocash -lzerocash -flto -DUSE_ASM -DCURVE_ALT_BN128 -L libzerocash/depinst/lib -Wl,-rpath,libzerocash/depinst/lib -L . -lsnark -lgmpxx -lgmp -lboost_system -lcrypto -lcryptopp -lz -ldl -pthread -lboost_program_options -lprocps
#include <stdio.h>
#include <stdlib.h>
#include "Goapi.h"
*/
import "C"
import "unsafe"

type Params struct{
    Ptr     unsafe.Pointer
}
func (p *Params) GetParams( fclient int ){
    p.Ptr = C.CParamsGen( C.int( fclient ) )
}
func (p *Params) DelParams(){
    C.CParamsDel( p.Ptr )
}

type Address struct{
    Ptr     unsafe.Pointer
}
func (a *Address) GetAddress(){
    a.Ptr = C.CAddressGen( )
}
func (a *Address) DelAddress(){
    C.CAddressDel( a.Ptr )
}
func (a *Address) String()string{
    return C.GoString( C.CAddressStr( a.Ptr ) )
}
func (a *Address) FromString( s string ){
    p := C.CString( s )
    defer C.free(unsafe.Pointer(p)) 
    a.Ptr = C.CStrAddress( p )
}

