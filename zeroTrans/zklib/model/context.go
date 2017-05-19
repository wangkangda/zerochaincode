package model

import (
    "fmt"
    "strconv"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

type Context struct{
    stub        shim.ChaincodeStubInterface
    amount      map[string]int
    merkle      *zklib.Merkle
    params      *zklib.Params
    commitnum   int
}

func (ctx *Context)InitContext()error{
    ctx.merkle = new(zklib.Merkle)
    ctx.merkle.GetMerkle()
    ctx.commitnum = 1
    return nil
}

func (ctx *Context)GetMerkle( )(string, error){
    res, err := ctx.stub.GetState("merkle")
    return string(res), err
}
func (ctx *Context)GetMerkleSize( addr string )(int, error){
    commitnum, err := ctx.stub.GetState("commitnum")
    if err != nil{
        return 0, err
    }
    return strconv.Atoi(string(commitnum))
}
func (ctx *Context)GetAmount( addr string )(int, error){
    key := fmt.Sprintf("%v%v", "amount", addr)
    value, err := ctx.stub.GetState(key)
    if err != nil{
        return 0, err
    }
    return strconv.Atoi(string(value))
}

func (ctx *Context)AddAmount( addr string ){
    if ctx.amount == nil{
        ctx.amount = make(map[string]int)
    }
    ctx.amount[addr] = 0
}
func (ctx *Context)GetContext()error{
    if ctx.params == nil{
        ctx.params.GetParams(0)
    }
    if ctx.amount != nil{
        for addr, _ := range ctx.amount {
            key := fmt.Sprintf("%v%v", "amount", addr)
            value, err := ctx.stub.GetState(key)
            if err != nil{
                fmt.Println("Get Error:", err)
                ctx.amount[addr] = 0
            }else{
                ctx.amount[addr], _ = strconv.Atoi(string(value))
            }
        }
    }
    if ctx.merkle == nil{
        merkle, err := ctx.stub.GetState("merkle")
        if err != nil{
            return err
        }
        ctx.merkle.FromString( string(merkle) )
    }
    if ctx.commitnum == 0{
        commitnum, err := ctx.stub.GetState("commitnum")
        if err != nil{
            return err
        }
        ctx.commitnum , _ = strconv.Atoi(string(commitnum))
    }
    return nil
}

func (ctx *Context)SaveContext()error{
    if ctx.params != nil {
        ctx.params.DelParams()
    }
    if ctx.amount != nil{
        for addr, value := range ctx.amount{
            key := fmt.Sprintf("%v%v", "amount", addr)
            err := ctx.stub.PutState(key, []byte(strconv.Itoa(value)))
            if err != nil{
                return err
            }
        }
    }
    if ctx.merkle != nil{
        err := ctx.stub.PutState("merkle", []byte(ctx.merkle.String()))
        ctx.merkle.DelMerkle()
        if err != nil{
            return err
        }
    }
    if ctx.commitnum != 0{
        err := ctx.stub.PutState("commitnum", []byte(strconv.Itoa(ctx.commitnum)))
        if err != nil{
            return err
        }
    }
    return nil
}
