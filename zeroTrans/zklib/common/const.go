package common

import(
    "errors"
)

const (
    Amount = "amount"
    Params = "params"
    Commit = "commit"
)

var (
    ErrParams error
)

func init(){
    ErrParams = errors.New("Error in received parameter")
}
