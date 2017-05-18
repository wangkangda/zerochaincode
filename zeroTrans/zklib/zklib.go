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
import "fmt"

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
    Addr    Address
    Value   int
    Index   int
}
func (c *Coin) GetCoin( a Address, value int ){
    c.Ptr = C.CCoinGen(a.Ptr, C.int(value))
    c.Addr = a
    c.Value = value
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
    cstr :=  C.CCommitStr( p )
    defer C.free(unsafe.Pointer(cstr))
    res := C.GoString(cstr)
    return res
}

type Merkle struct{
    Ptr     unsafe.Pointer
}
func (m *Merkle) GetMerkle(){
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
    m.Ptr = C.CStrMerkle(p)
}
func (m *Merkle) Insert( s string, idx int ){
    cstr := C.CString( s )
    defer C.free(unsafe.Pointer(cstr))
    p := C.CStrCommit( cstr )
    defer C.CCommitDel( p )
    m.Ptr = C.CMerkleInsert( m.Ptr, p, C.int(idx) )
}

type Pour struct{
    Ptr     unsafe.Pointer
}
func (p *Pour) GetPour(params Params, co1 Coin, co2 Coin,
                            m Merkle, vpub int, cn1 Coin, cn2 Coin){
    p.Ptr = C.CPourGen( params.Ptr,
                        co1.Ptr,    co2.Ptr,
                        co1.Addr.Ptr, co2.Addr.Ptr,
                        C.int(co1.Index),  C.int(co2.Index),
                        m.Ptr,
                        cn1.Ptr,    cn2.Ptr,
                        C.int( vpub ),
                        cn1.Ptr,    cn2.Ptr)
}
func (p *Pour) DelCoin(){
    C.CPourDel( p.Ptr )
}
func (p *Pour) String()string{
    cstr := C.CPourStr( p.Ptr )
    defer C.free(unsafe.Pointer(cstr))
    res := C.GoString( cstr )
    return res
}
func (p *Pour) FromString( s string){
    cstr := C.CString(s)
    defer C.free(unsafe.Pointer(cstr))
    p.Ptr = C.CStrPour(cstr)
}
func (p *Pour) Verify( params Params, m Merkle )bool{
    res := int( C.CPourVerify(params.Ptr, p.Ptr, m.Ptr) )
    return res != 0
}

func TutorialTest(){
    res := C.TutorialTest()
    fmt.Println("Get :", res)
}
