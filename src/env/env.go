package env

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Env struct {
	ServerIp1,
	ServerPort1,
	ServerIp2,
	ServerPort2,
	SalerId1,
	SalerKey1,
	AppVer1,
	SalerId2,
	SalerKey2,
	AppVer2 string
}

func NewEnv() Env {
	var env Env
	envbuf := ReadFromFile("common.conf", "conf")
	json.Unmarshal(envbuf, &env)
	return env
}

func ReadFromFile(filename string, foldername string) (buf []byte) {
	pwd := os.Getenv("PWD")
	fullpath := pwd + "/" + foldername + "/" + filename
	buf, err := ioutil.ReadFile(fullpath)
	if err != nil {
		fmt.Println("Please check the data file format")
	}
	return buf
}

func ReadFileLines(filename string, foldername string) (buf []string) {
	pwd := os.Getenv("PWD")
	fullpath := pwd + "/" + foldername + "/" + filename
	f, err := os.Open(fullpath)
	if err != nil {
		fmt.Println("Please check the data file format")
	}
	b := bufio.NewReader(f)
	for line, err := b.ReadString('\n'); err == nil; line, err = b.ReadString('\n') {
		line = strings.TrimSpace(line)
		buf = append(buf, line)
	}
	return
}
