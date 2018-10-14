package goStore

import (
	"testing"
	"fmt"
)

func TestSync(t *testing.T) {
	base:=NewDatabase([]DataType{INT8,STRING},5)
	base.Start()
	var n int8 = 110
	var s  = "aaa"
	base.Store(n,s)
	base.Stop()
	fmt.Println(base.data)
}