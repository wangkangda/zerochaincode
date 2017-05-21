package model

import (
    "fmt"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

type Context struct{
    Stub        shim.ChaincodeStubInterface
    Amount      map[string]int
    Merkle      *zklib.Merkle
    Params      *zklib.Params
    Commitnum   int
}

func (ctx *Context)InitContext()error{
    ctx.Merkle = new(zklib.Merkle)
    ctx.Merkle.GetMerkle()
    ctx.Commitnum = 1
    return nil
}

func (ctx *Context)GetMerkle( )(string, error){
    res, err := ctx.Stub.GetState("merkle")
    return string(res), err
}
func (ctx *Context)GetMerkleSize( )(int, error){
    commitnum, err := ctx.Stub.GetState("commitnum")
    if err != nil{
        return 0, err
    }
    return strconv.Atoi(string(commitnum))
}
func (ctx *Context)GetAmount( addr string )(int, error){
    key := fmt.Sprintf("%v%v", "amount", addr)
    value, err := ctx.Stub.GetState(key)
    if err != nil{
        return 0, err
    }
    return strconv.Atoi(string(value))
}

func (ctx *Context)AddAmount( addr string ){
    if ctx.Amount == nil{
        ctx.Amount = make(map[string]int)
    }
    ctx.Amount[addr] = 0
}
func (ctx *Context)GetContext()error{
    if ctx.Params == nil{
        ctx.Params.GetParams(0)
    }
    if ctx.Amount != nil{
        for addr, _ := range ctx.Amount {
            key := fmt.Sprintf("%v%v", "amount", addr)
            value, err := ctx.Stub.GetState(key)
            if err != nil{
                fmt.Println("Get Error:", err)
                ctx.Amount[addr] = 0
            }else{
                ctx.Amount[addr], _ = strconv.Atoi(string(value))
            }
        }
    }
    if ctx.Merkle == nil{
        merkle, err := ctx.Stub.GetState("merkle")
        if err != nil{
            return err
        }
        ctx.Merkle.FromString( string(merkle) )
    }
    if ctx.Commitnum == 0{
        commitnum, err := ctx.Stub.GetState("commitnum")
        if err != nil{
            return err
        }
        ctx.Commitnum , _ = strconv.Atoi(string(commitnum))
    }
    return nil
}

func (ctx *Context)SaveContext()error{
    if ctx.Params != nil {
        ctx.Params.DelParams()
    }
    if ctx.Amount != nil{
        for addr, value := range ctx.Amount{
            key := fmt.Sprintf("%v%v", "amount", addr)
            err := ctx.Stub.PutState(key, []byte(strconv.Itoa(value)))
            if err != nil{
                return err
            }
        }
    }
    if ctx.Merkle != nil{
        err := ctx.Stub.PutState("merkle", []byte(ctx.Merkle.String()))
        ctx.Merkle.DelMerkle()
        if err != nil{
            return err
        }
    }
    if ctx.Commitnum != 0{
        err := ctx.Stub.PutState("commitnum", []byte(strconv.Itoa(ctx.Commitnum)))
        if err != nil{
            return err
        }
    }
    return nil
}
