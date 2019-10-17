package PBFT_module

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

/**
江声
该函数用来将任意类型转换为string类型
TODO 还有一些其他类型有待处理
*/
func ToString(args interface{}) string {
	switch args.(type) {
	case int:
		return strconv.FormatInt(int64(args.(int)), 10)
	case int8:
		return strconv.FormatInt(int64(args.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(args.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(args.(int32)), 10)
	case int64:
		return strconv.FormatInt(int64(args.(int64)), 10)
	default:
		bytes, _ := json.Marshal(args)
		return string(bytes)
	}
}

/**
江声
该函数用来判断两个byte数组是否相等，可用于校验数字签名
*/
func BytesEqual(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}
	return true
}

/**
@author：江声
读取文件内容，按行读取
*/
func ReadFileLine(path string) ([]string, error) {
	//打开文件
	var lines []string

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return nil, err
	}

	//关闭文件
	defer f.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)

	for {
		//遇到'\n'结束读取, 但是'\n'也读取进入
		buf, err := r.ReadString('\n')
		lines = append(lines, buf)
		if err != nil {
			if err == io.EOF { //文件已经结束
				break
			}
			fmt.Println("err = ", err)
		}

	}
	return lines, err
}
