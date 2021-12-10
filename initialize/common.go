package initialize

import (
	"io/ioutil"
)

func readConf(filename string) []byte {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return fileContent
}
