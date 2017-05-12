package rpc

type ReqDeploy struct{
    request []byte
}

func NewReqDeploy( chaincode string, args []string )*ReqDeploy{
    res = new(ReqDeploy)
    if args == nil || len(args) == 0 {
        sreq = fmt.Sprintf(jsonTemplate, "deploy", chaincode, "init", "")
        res.request = []byte(sreq)
    }
    return res
}

func (r ReqDeploy *) Execute() (string, error){
    res, err := JsonSend( r.request )
    if res== nil || err != nil{
        return '', err
    }
    return string(res), nil
}
