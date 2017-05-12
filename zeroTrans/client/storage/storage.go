package storage

import(
    "github.com/wangkangda/zerochaincode/zeroTrans/client/cmd"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib/model"
)

var(
    Paramlist      model.Params
    Keylist    []model.KeyPair
    Commlist     []model.Commit
    Amounts     map[model.Address]int
)

func GetStorage() error{
    initf := false
    var input string
    f, err := os.Open(cmd.Datapath)
    if err != nil {
        f, err = os.Create(cmd.Datapath)
        if err != nil{
            log.Printf("Error in creating file:%v", err.Error())
            return err
        }
        initf = true
    }
    defer f.Close()
    buf := bufio.NewReader(f)
    input = GetLines( buf, initf )
    Paramlist = model.GetParams( input, initf )
    input = GetLines( buf, initf )
    Keylist = model.GetKeyPairs( input, initf )
    input = GetLines( buf, initf )
    Commlist = model.GetCommits( input, initf )
    input = GetLines( buf, initf )
    Amounts = model.GetAmounts( input, initf )
}

func SaveStorage(){
    f, err = os.Create(cmd.Datapath)
    if err != nil{
        log.Printf("Error in creating file:%v", err.Error())
        return
    }
    defer f.Close()
    f.WriteString( model.StrParams( Paramlist ) )
    f.WriteString( model.StrKeyPairs( Keylist) )
    f.WriteString( model.StrCommits( Commlist ) )
    f.WriteString( model.StrAmounts( Amounts ) )
    log.Println("Write All Storage")
}

func GetLines( buf bufio.Reader, empty bool )string{
    var res string
    if !empty{
        res, err = buf.ReadString('\n')
        if err != nil{
            log.Printf("Error in creating file:%v", err.Error())
            panic( err )
        }
    }
    return res
}
