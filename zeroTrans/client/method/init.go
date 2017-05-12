package method

func Init()string, error{
    args := make([]string)
    deploy := chain.NewDeploy( common.ChaincodePath, args )
    return deploy.execute()
}
