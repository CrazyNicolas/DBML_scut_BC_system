package main

import (
	//pbft "./PBFT_module"
	"flag"
	"fmt"
	//"bufio"
	//"os"
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
	//pbft.GetRegisterDir()
	//f, err := os.Open("src/main/configure/register.dat")
	//if err != nil {
	//	fmt.Println("err = ", err)
	//	return
	//}
	//
	////关闭文件
	//defer f.Close()
	//
	//r := bufio.NewReader(f)
	//buf, _ := r.ReadString('\n')
	//fmt.Println((buf))
	//buf, _ = r.ReadString('\n')
	//fmt.Println((buf))
	//buf, _ = r.ReadString('\n')
	//fmt.Println((buf))
	//buf, _ = r.ReadString('\n')
	//fmt.Println((buf))

	addr1 := flag.String("addr1", "tcp@localhost:8972", "server1 address")
	fmt.Println(*addr1)
}
