package storage

import(
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

var(
    AddresList      map[string]model.Address
)

type Storage struct{
    addressList     map[string]string       `json:"address_list"`
}

func GetStorage() error{
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
    var allStorage Storage
    json.Unmarshal( []byte(GetLines(buf, false)), &allStorage )
    AddressList = make(map[string]string)
    for addr, obj := allStorage.addressList{
        AddressList[addr] = zklib.Address{}
        AddressList[addr].FromString( obj )
    }
}

func SaveStorage(){
    f, err = os.Create(cmd.Datapath)
    if err != nil{
        log.Printf("Error in creating file:%v", err.Error())
        return
    }
    defer f.Close()
    var allStorage Storage
    allStorage.addressList = make(map[string]string)
    for addr, obj := range AddressList{
        allStorage.addressList[ addr ] = obj.String()
    }
    f.WriteString( string(json.Marshal( allStorage ) ) )
    f.WriteString( "\n" )
    log.Println("Write All Storage")
}

func GetLines( buf bufio.Reader, empty bool )string{
    var res string
    if !empty{
        res, err = buf.ReadString("\n")
        if err != nil{
            log.Printf("Error in creating file:%v", err.Error())
            panic( err )
        }
    }
    return res
}
