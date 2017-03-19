package main

import (
	"bufio"
	"io"
	"os"
	"fmt"
	"strings"
    "strconv"
    //"errors"
)

func checkerr( err error ) {
	if err != nil{
		fmt.Println(err)
		panic( err )
	}
}
func getCommit(filepath string) (map[int]string, error){
    f, err := os.Open(filepath)
    var commitList map[int]string
    commitList = make(map[int]string)
    if err != nil{
        return commitList, err
    }
    defer f.Close()
    buf := bufio.NewReader(f)
    for{
        id, err := buf.ReadString('\n')
        id = strings.TrimSpace(id)
        if err == io.EOF{
            //return nil, errors.new("Commitment File not fix")
            break
        }
        comm, cerr := buf.ReadString('\n')
        comm = strings.TrimSpace(comm)
        index, _ := strconv.Atoi( id )
        commitList[ index ] = comm
        if cerr == io.EOF{
            break
        }
        checkerr( cerr )
    }
    fmt.Println( "get commitment: ", commitList )
    return commitList, nil
}
func getData(filepath string) ([]string, error){
	f, err := os.Open(filepath)
    if err != nil {
        f, err = os.Create(filepath)
    }
	checkerr(err)
	defer f.Close()
	buf := bufio.NewReader(f)
	params := make([]string, 0, 100)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF{
			if len(line)>0 {
				//fmt.Println("line:", line)
				params = append(params, line)
			}
			break
		}
		checkerr( err )
		line = strings.TrimSpace(line)
		fmt.Println("line:", line)
		params = append(params, line )
	}
    //fmt.Println("after read file: ", params)
	return params, nil
}

func saveData(filepath string, params []string) error{
	f, err := os.Create(filepath)
	check(err)
	defer f.Close()
	fmt.Println("write lines:", len(params))
	for i:=0; i<len(params); i++{
		f.WriteString( params[i] )
		f.WriteString( "\n" )
	}
	return nil
}
func saveCommit(filepath string, cl map[int]string ) error{
    f, err := os.Create(filepath)
    check(err)
    defer f.Close()
    var i string
    for index := range cl {
        i = strconv.Itoa(index)
        f.WriteString(i)
        f.WriteString("\n")
        val, _ := cl[index]
        f.WriteString(val)
        f.WriteString("\n")
    }
    return nil
}
