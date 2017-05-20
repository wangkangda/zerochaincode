package storage

import(
    "os"
    "bufio"
    "log"
    "encoding/json"
    "github.com/wangkangda/zerochaincode/zeroTrans/zklib"
)

const(
    datapath = "storage.dat"
    backpath = "backup.dat"
)

var(
    AddressList      map[string]*zklib.Address
)

type PackedData struct{
    AddressList     map[string]string   `json:"address_list"`
}

func GetStorage() error{
    initf := false
    f, err := os.Open(datapath)
    if err != nil {
        f, err = os.Create(datapath)
        if err != nil{
            log.Printf("Error in creating file:%v", err.Error())
            return err
        }
        initf = true
    }
    defer f.Close()
    buf := bufio.NewReader(f)
    var allStorage PackedData
    json.Unmarshal( []byte(GetLines(buf, initf)), &allStorage )
    AddressList = make(map[string]*zklib.Address)
    for addr, obj := range allStorage.AddressList{
        AddressList[addr] = &zklib.Address{}
        AddressList[addr].FromString( obj )
    }
    return nil
}

func SaveStorage()error{
    fdone := false
    f, err := os.Create(backpath)
    if err != nil{
        log.Printf("Error in creating file:%v", err)
        return err
    }
    defer func(){
        f.Close()
        if fdone {
            err := os.Rename(backpath, datapath)
            if err != nil{
                log.Printf("Error while save data: %v\n", err)
            }
        }
    }()
    var allStorage PackedData
    allStorage.AddressList = make( map[string]string )
    for addr, obj := range AddressList{
        //log.Printf("Save %v: %v\n", addr, obj)
        if obj != nil{
            allStorage.AddressList[addr] = obj.String()
        }
    }
    resbyte, err := json.Marshal(allStorage)
    if err != nil{
        return err
    }
    log.Printf("Get Json %v len\n", len(resbyte))
    f.WriteString( string( resbyte ) )
    f.WriteString( "\n" )
    log.Println("Write All Storage")
    fdone = true
    return nil
}

func GetLines( buf *bufio.Reader, empty bool )string{
    var res string
    var err error
    if !empty{
        res, err = buf.ReadString('\n')
        if err != nil{
            log.Printf("Error in creating file:%v", err.Error())
            panic( err )
        }
    }
    return res
}
