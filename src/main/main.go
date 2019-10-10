package main

import (
	//pbft "./PBFT_module"
	"fmt"
	"os"
)

func main() {
	//pbft.NewReplica(0, 0, 0, 0)
	//bytes, _ := json.Marshal(nil)
	//println(string(bytes))
	//b1 := []byte{1, 2, 3}
	//b2 := []byte{1, 2, 3}
	//fmt.Println(pbft.BytesEqual(b1, b2))
	//pbft.Test()
	//pbft.GenerateKeyPairAndSave(1024, 1)
	_, err := os.Open("private.pemm")
	if err != nil {
		fmt.Print("wrong")
	}
}
