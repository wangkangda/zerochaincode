package goclient

import {
	"bufio"
	"io"
	"os"
	"strings"
}

func checkerr( err error ) {
	if err != nil{
		panic err
	}
}
func getData(filepath string) ([]string, err){
	f, err := os.Open(filepath)
	checkerr(err)
	defer f.Close()
	buf := bufio.NewReader(f)
	params := make([]string, 0, 100)
	for {
		line, err := buf.ReadString('\n')
		params = append(params, line )
		if err == io.EOF{
			break;
		}
		checkerr( err )
	}
	return params, nil
}

func saveData(filepath string, params []string) error{
	f, err := os.OpenFile(filepath, os.O_WRONLY, 0666)
	check(err)
	defer f.Close()
	for i:=0; i<len(params); i++{
		f.WriteString( params[i] )
		f.WriteString( '\n' )
	}
	return nil
}
