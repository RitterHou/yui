package compiler

import "log"

func Build(source []byte) (target []byte) {
	data := string(source)
	log.Println(data)
	return source
}
