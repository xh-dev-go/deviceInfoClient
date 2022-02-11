package main

import (
	"fmt"
	"github.com/xh-dev-go/deviceInfoClient/sysMeta"
)

func main() {
	meta, err := sysMeta.GetMetaString()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", meta)
	//fmt.Printf("sdfsdfa")
}
