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
	starttime time.Time
	Testname  string
	Result    string
	Duration  float64
	Reason    string
}

func (t *Testcase) Pass(reason string) {
	t.Result = "PASS"
	t.Reason = reason
	Info(t.Testname, "PASSED")
}

func (t *Testcase) Fail(reason string) {
	t.Result = "FAIL"
	t.Reason = reason
	Info(t.Testname, "FAILED")
}
