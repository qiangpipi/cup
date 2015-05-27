package tst

import (
	. "loger"
)

type LoginNormal struct {
	tc      Testcase
	account string
}

func (ln *LoginNormal) Init() {
	Info("Login-normal initing")
	ln.account = CreateTestAccount()
}

func (t1 *LoginNormal) Execute() {
	Info("Login-normal executing")
	t1.tc.Pass()
	t1.tc.Fail()
}
