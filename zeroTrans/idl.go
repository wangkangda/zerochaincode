chaincode{
    type input struct{
       type string  "normal"/"anonymous"
       normal:
        address
        amount
        signature
       anonymous:
        serialNum
        anonyCoin{
            zkProof
            (amount)
            (random)
            (metadata)
        }
    }
    type output struct{
        type string "normal"/"anonymous"
       normal:
        address
        amount
       anonymous:
        commitment
        amount
    }
        
    func transaction( []intput, []output )
    func query( type string, key string )
}
