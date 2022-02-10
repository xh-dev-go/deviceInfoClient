package main

import (
	"deviceInfoClient/sysMeta"
	"fmt"
)

func main() {
	meta, err := sysMeta.GetMetaString()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s",meta)
	//fmt.Printf("sdfsdfa")
}
