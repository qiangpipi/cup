package main

import (
	"env"
	"flag"
	. "loger"
	"os"
	"tst"
)

//Global var

func main() {
	//Parse params here with flag package
	//Define flags
	D = flag.Bool("d", false, "Debug enable/disable and default is false")
	Info("Debug:", *D)
	path := os.Getenv("PWD")
	filename := path + "/conf"
	_, err := os.Stat(filename)
	if err != nil {
		ErrMsg("Please run the binary under cafepot folder", path)
		os.Exit(1)
	}
	testlist := flag.String("tl", "tl", "Test List file name")
	Debug(testlist)
	testcases := flag.String("tcs", "", "Test cases files and seperated by ','")
	Debug(testcases)
	//Flag parse
	flag.Parse()
	tst.Env = env.NewEnv()
	for _, t := range tst.Testlist {
		t.Init()
		t.Execute()
	}
}
