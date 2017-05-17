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
    p := C.CAddressStr( a.Ptr )
    defer C.free(unsafe.Pointer(p))
    res := C.GoString( p )
    return res
}
func (a *Address) FromString( s string ){
    p := C.CString( s )
    defer C.free(unsafe.Pointer(p)) 
    a.Ptr = C.CStrAddress( p )
}

type Coin struct{
    Ptr     unsafe.Pointer
}
func (c *Coin) GenCoin( a Address, value int ){
    c.Ptr = C.CCoinGen(a.Ptr, C.int(value))
}
func (c *Coin) DelCoin(){
    C.CCoinDel( c.Ptr )
}
func (c *Coin) String()string{
    p := C.CCoinStr( c.Ptr )
    defer C.free(unsafe.Pointer(p))
    res := C.GoString( p )
    return res
}
func (c *Coin) FromString( s string){
    p := C.CString(s)
    defer C.free(unsafe.Pointer(p))
    c.Ptr = C.CStrCoin(p)
}
func (c *Coin) GetCommit( )string{
    p := C.CCoinCommit( c.Ptr )
    return C.CCommitStr( p )
}

type Merkle struct{
    Ptr     unsafe.Pointer
}
func (m *Merkle) GenMerkle(){
    m.Ptr = C.CMerkleGen()
}
func (m *Merkle) DelMerkle(){
    C.CMerkleDel( m.Ptr )
}
func (m *Merkle) String()string{
    p := C.CMerkleStr( m.Ptr )
    defer C.free(unsafe.Pointer(p))
    res := C.GoString( p )
    return res
}
func (m *Merkle) FromString( s string){
    p := C.CString(s)
    defer C.free(unsafe.Pointer(p))
    m.Ptr = CStrMerkle(p)
}
func (m *Merkle) Insert( s string, idx int ){
    cstr := C.CString( s )
    defer C.free(unsafe.Pointer(cstr))
    p := CStrCommit( cstr )
    defer CCommitDel( p )
    m.Ptr = C.CMerkleInsert( m.Ptr, p, idx )
}

