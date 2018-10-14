package goStore

import (
	"testing"
	"math/rand"
	"time"
	"strconv"
)

func TestGoStore(t *testing.T) {
	base:=NewDatabase([]DataType{INT32,STRING},150)
	for i:=0;i<500;i++{
		go func() {

			var i int32 = rand.Int31()
			var s string = strconv.FormatUint(uint64(time.Now().UnixNano()),10)
			base.Store(i,s)
		}()
	}
}