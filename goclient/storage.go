package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
	"strings"
)

func checkerr( err error ) {
	if err != nil{
		fmt.Println(err)
		panic( err )
	}
}
func getData(filepath string) ([]string, error){
	f, err := os.Open(filepath)
	checkerr(err)
	defer f.Close()
	buf := bufio.NewReader(f)
	params := make([]string, 0, 100)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF{
			break;
		}
		checkerr( err )
		line = strings.TrimSpace(line)
		fmt.Println("line:", line)
		params = append(params, line ) 
	}
	return params, nil
}

func saveData(filepath string, params []string) error{
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()
	fmt.Println("write %s lines", len(params))
	for i:=0; i<len(params); i++{
		f.WriteString( params[i] )
		f.WriteString( "\n" )
	}
	return nil
}
