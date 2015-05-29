package tst

import (
	. "loger"
	"strings"
)

type LoginNormal struct {
	tc Testcase
	actionSimpleReg,
	actionLogin,
	actionLogout,
	name,
	pwd,
	valid_code,
	jsid string
}

func (ln *LoginNormal) Init() {
	Info("Login-normal initing")
	ln.name = CreateTestAccount()
	ln.actionSimpleReg = action["register"]
	ln.actionLogin = action["login"]
	ln.actionLogout = action["logout"]
	ln.pwd = "111111"
	ln.valid_code = "111111"
	ln.jsid = ""
	ln.tc.Testname = "Login normally"
}

func (ln *LoginNormal) Execute() {
	Info("Login-normal executing")
	//Simple register account
	url := baseUrl(ln.actionSimpleReg, ln.jsid)
	params := "name=" + ln.name + "&pwd=" + ln.pwd + "&valid_code=" + ln.valid_code
	url += params
	Debug(url)
	res, err := httpreq(Env.ServerIp1, Env.ServerPort1, url)
	if err != nil {
		ln.tc.Fail(err.Error())
	}
	if strings.Contains(res, "\"code\":\"000\"") {
		//Res parse
		//Login with account
		url = baseUrl(ln.actionLogin, ln.jsid)
		params = "name=" + ln.name + "&pwd=" + ln.pwd
		url += params
		Debug(url)
		//res,err := httpreq(Env.ServerIp1,Env.ServerPort1,url)
		//Res parse
		//Set ln.jsid
		//Logout with account
		url = baseUrl(ln.actionLogout, ln.jsid)
		params = ""
		url += params
		Debug(url)
		//res,err := httpreq(Env.ServerIp1,Env.ServerPort1,url)
		//Res parse
	} else {
		ln.tc.Fail("Simple register fail")
	}
}
