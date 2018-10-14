package goStore

import (
	"errors"
	"runtime"
)
var(
	OutOfRangeError = errors.New("The number of storage is out of range")
)
type Database struct{
	data [][][]byte
	types []DataType
	storeChannel chan [][]byte
	endRunningChannel chan bool
	oneDataLen int
}
func NewDatabase(types []DataType,channelCache int)*Database {
	var base= &Database{}
	base.storeChannel = make(chan [][]byte, channelCache)
	base.endRunningChannel = make(chan bool)
	base.types = types
	base.oneDataLen = len(types)
	return base
}

func(base *Database)running(){
	defer runtime.Goexit()
	for{
		select {
		case <-base.endRunningChannel:
			{
				l:=len(base.storeChannel)
				for i:=0;i<l;i++{
					base.data=append(base.data,<-base.storeChannel)
				}
				return
			}
		case v:=<-base.storeChannel:
			{
				base.data=append(base.data,v)
			}
		}
	}
}
func(base *Database)Start(){
	go base.running()
}
func(base *Database)Stop(){
	base.endRunningChannel<-true
	close(base.endRunningChannel)
	close(base.storeChannel)
}
func(base *Database)Store(values...interface{}){
	if len(values)>base.oneDataLen{
		panic(OutOfRangeError)
	}else{
		var data = make([][]byte,base.oneDataLen)
		for i,v:=range values{
			data[i]=dataTypeToData(base.types[i],v)
		}
		base.storeChannel<-data
	}
}