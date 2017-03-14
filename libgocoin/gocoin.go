package main
/*
#cgo CFLAGS: -l./lib
#cgo LDFLAGS: -L./lib -lhello -lzerocoin -Wl,-rpath,./lib/
#include "Goapi.h"
#include "hello.h"
*/
import "C"
import "fmt"
//import "unsafe"
import "os"

const TUTORIAL_TEST_MODULUS ="a8852ebf7c49f01cd196e35394f3b74dd86283a07f57e0a262928e7493d4a3961d93d93c90ea3369719641d626d28b9cddc6d9307b9aabdbffc40b6d6da2e329d079b4187ff784b2893d9f53e9ab913a04ff02668114695b07d8ce877c4c8cac1b12b9beff3c51294ebe349eca41c24cd32a6d09dd1579d3947e5c4dcc30b2090b0454edb98c6336e7571db09e0fdafbd68d8f0470223836e90666a5b143b73b9cd71547c917bf24c0efc86af2eba046ed781d9acb05c80f007ef5a0a5dfca23236f37e698e8728def12554bc80f294f71c040a88eff144d130b24211016a97ce0f5fe520f477e555c9997683d762aff8bd1402ae6938dd5c994780b1bf6aa7239e9d8101630ecfeaa730d2bbc97d39beb057f016db2e28bf12fab4989c0170c2593383fd04660b5229adcd8486ba78f6cc1b558bcd92f344100dff239a8c00dbc4c2825277f24bdd04475bcc9a8c39fd895eff97c1967e434effcb9bd394e0577f4cf98c30d9e6b54cd47d6e447dcf34d67e48e4421691dbe4a7d9bd503abb9"

func ZerocoinTutorial(){
    C.TestGoapi();
    C.hello();
    var p = C.GoCoinGeneration( C.CString(TUTORIAL_TEST_MODULUS) )
    fmt.Printf( C.GoString(p) )
    //C.GoCoinDestroy( p )
    //defer C.free(unsafe.Pointer(p))
	C.CCStrDel(p)
}
func tutorial(){
	param1 := C.CCParamsGen()
	fmt.Printf( "len:%d", len(C.GoString(param1)) )
	fmt.Printf( "C_CCParamsGen sucessful !\n" )
	oParam1 := C.CCParamsLoad( param1 )
	fmt.Printf( "C_CCParamsLoad sucessful !\n")
	accum1 := C.CCAccumGen( oParam1 )
	fmt.Printf( "C_CCAccumGen sucessful !\n")
	oAccum1 := C.CCAccumLoad( oParam1, accum1 )
	fmt.Printf( "C_CCAccumLoad sucessful !\n")


	pricoin := C.CCPricoinGen2( oParam1 )
    defer C.CCStrDel( pricoin )
    oPricoin := C.CCPricoinLoad( oParam1, pricoin )
    defer C.CCPricoinDel( oPricoin )
	fmt.Printf("C_CCPricoinGen sucessful !\n")
	pubcoin := C.CCPubcoinGen( oParam1, oPricoin )
	fmt.Printf("C_CCPubcoinGen sucessful !\n")
	
	for i:= 0; i<5; i++ {
		newcoin := C.CCPricoinGen(oParam1)
		newpubcoin := C.CCPubcoinGen( oParam1, newcoin )
		accum2 := C.CCAccumCal( oParam1, oAccum1, newpubcoin )
		C.CCAccumDel( oAccum1 )
		oAccum1 = C.CCAccumLoad( oParam1, accum2 )
		C.CCPricoinDel( newcoin )
		//defer C.free(unsafe.Pointer(newpubcoin))
		//defer C.free(unsafe.Pointer(accum2))
		C.CCStrDel(newpubcoin)
		C.CCStrDel(accum2)
	}
	toaddress := "whereverisendto"
	coinspend := C.CCSpendGen( oParam1, oPricoin, oAccum1, C.CString(toaddress))
	sn := C.CCSpendVerify( oParam1, coinspend, C.CString(toaddress), oAccum1)
	if sn == nil {
		fmt.Printf("Verify failed!!!")
	}else{
		fmt.Printf("Verify succeed!!!\n")
		C.CCBignumDel(sn)
	}

	//defer C.free(unsafe.Pointer(param1))
	//defer C.free(unsafe.Pointer(accum1))
	//defer C.free(unsafe.Pointer(coinspend))
	//defer C.free(unsafe.Pointer(pubcoin))
	C.CCStrDel(param1)
	C.CCStrDel(accum1)
	C.CCStrDel(coinspend)
	C.CCStrDel(pubcoin)
	
	C.CCAccumDel(oAccum1)
	C.CCParamsDel(oParam1)

}


func main(){
    //fmt.Printf("libzerocoin go version tutorial. ")
    //fmt.Printf(TUTORIAL_TEST_MODULUS)
    //ZerocoinTutorial();
	if len(os.Args) < 2 {
		fmt.Printf("need the function to execute!")
		return
	}
	switch os.Args[1] {
	case "tutorial":
		tutorial()
	default:
		ZerocoinTutorial()
	}

}
