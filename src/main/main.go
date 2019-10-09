package main

import (
	"encoding/json"
)

func main() {
	//pbft.NewReplica(0, 0, 0, 0)
	bytes, _ := json.Marshal(nil)
	println(string(bytes))
}
