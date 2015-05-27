package tst

import (
	. "loger"
	"time"
)

type Executer interface {
	Init()
	Execute()
}

type Testcase struct {
	Start    time.Time
	Testname string
	Result   string
	Duration int
}

func (t *Testcase) Pass() {
	t.Result = "PASS"
	Info(t.Testname, "PASSED")
}

func (t *Testcase) Fail() {
	t.Result = "FAIL"
	Info(t.Testname, "FAILED")
}
