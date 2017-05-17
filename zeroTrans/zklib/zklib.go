package zklib
/*
/*
#cgo CFLAGS: -l./lib
#cgo LDFLAGS: -L./lib -lzerocoin -Wl,-rpath,./lib/
#include "Goapi.h"
*/
import (
    "C"
    "fmt"
)

type Params struct{
    Ptr     unsafe.Pointer
}

func (p *Params) GetParams{
    p.Ptr = C.CParamsGen()
}

func (p *Params) DelParams{
    C.CParamsDel( p.Ptr )
}
