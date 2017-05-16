package zklib
/*
#cgo CFLAGS: -g -Wno-unused-parameter -fPIC -Wno-unused-variable -I libzerocash/libzerocash
#cgo LDFLAGS: -L libzerocash -Wl,-rpath,./libzerocash -lzerocash -flto -DUSE_ASM -DCURVE_ALT_BN128 -L libzerocash/depinst/lib -Wl,-rpath,libzerocash/depinst/lib -L . -lsnark -lgmpxx -lgmp -lboost_system -lcrypto -lcryptopp -lz -ldl -pthread -lboost_program_options -lprocps
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
