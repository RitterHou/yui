package common

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

const magicHex = 0x10151657 // 10月15日 16时57分

var (
	root  string // 当前的工作目录
	Magic []byte // Magic Number
)

func init() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	root = dir

	Magic = []byte(strconv.FormatInt(magicHex, 2))
}

func GetAbsPath(filename string) (absPath string) {
	absPath = path.Join(root, filename)
	return
}

func ReadFile(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func WriteFile(filePath string, data []byte) {
	ioutil.WriteFile(filePath, data, 0644)
}

//func IsCompiled(data []byte) bool {
//	if len(data) < 4 {
//		return false
//	}
//	for i := 0; i < 4; i++ {
//		if data[i] != Magic[i] {
//			return false
//		}
//	}
//	return true
//}
