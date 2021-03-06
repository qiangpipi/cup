package tst

import (
	. "loger"
	"strconv"
	"strings"
	"time"
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
	ln.tc.starttime = time.Now()
}

func (ln *LoginNormal) Execute() {
	Info("Login-normal executing")
	//////////////////////////
	//Simple register account
	/////////////////////////
	url := baseUrl(ln.actionSimpleReg, ln.jsid)
	params := "name=" + ln.name + "&pwd=" + ln.pwd + "&valid_code=" + ln.valid_code
	url += params
	Debug(url)
	res, err := httpreq(Env.ServerIp1, Env.ServerPort1, url)
	if err != nil {
		ln.tc.Fail(err.Error() + "\n" + res)
	}
	//Res parse
	Debug(res)
	if strings.Contains(res, "\"code\":\"000\"") {
		////////////////////////
		//Login with account
		////////////////////////
		url = baseUrl(ln.actionLogin, ln.jsid)
		params = "name=" + ln.name + "&pwd=" + ln.pwd
		url += params
		Debug(url)
		res, err := httpreq(Env.ServerIp1, Env.ServerPort1, url)
		//Res parse
		if err != nil {
			ln.tc.Fail(err.Error() + "\n" + res)
		}
		Debug(res)
		if strings.Contains(res, "\"code\":\"000\"") {
			//Set ln.jsid
			for _, s := range strings.Split(res, ",") {
				if strings.Contains(s, "jsessionid") {
					ln.jsid = strings.Split(s, "\"")[3]
				}
			}
			url = baseUrl(ln.actionLogout, ln.jsid)
			params = ""
			url += params
			Debug(url)
			////////////////////////
			//Logout with account
			////////////////////////
			res, err := httpreq(Env.ServerIp1, Env.ServerPort1, url)
			//Res parse
			if err != nil {
				ln.tc.Fail(err.Error() + "\n" + res)
			}
			if strings.Contains(res, "\"code\":\"000\"") {
				ln.tc.Pass("")
			} else {
				ln.tc.Fail("Wrong code number for logout fail")
			}
		} else {
			ln.tc.Fail("Wrong code number for login fail")
		}
	} else {
		ln.tc.Fail("Wrong code number for simple register fail")
	}
	ln.tc.Duration = time.Since(ln.tc.starttime).Seconds()
	Info("LoginNormal took " + strconv.FormatFloat(ln.tc.Duration, 'f', -1, 64) + " seconds")
}

type LoginWithoutReg struct {
	tc Testcase
	actionLogin,
	actionLogout,
	name,
	pwd,
	valid_code,
	jsid string
}

func (lwr *LoginWithoutReg) Init() {
	Info("Login without registered initing")
	lwr.name = CreateTestAccount()
	lwr.tc.Testname = "Login without registered"
	lwr.actionLogin = action["login"]
	lwr.actionLogout = action["logout"]
	lwr.pwd = "111111"
	lwr.valid_code = "111111"
	lwr.jsid = ""
	lwr.tc.starttime = time.Now()
}

func (lwr *LoginWithoutReg) Execute() {
	/////////////////////////////////////
	//Login in with account non-existing
	/////////////////////////////////////
	url := baseUrl(lwr.actionLogin, lwr.jsid)
	params := "name=" + lwr.name + "&pwd=" + lwr.pwd
	url += params
	Debug(url)
	res, err := httpreq(Env.ServerIp1, Env.ServerPort1, url)
	//Parse
	if err != nil {
		lwr.tc.Fail(err.Error() + "\n" + res)
	}
	Debug(res)
	if strings.Contains(res, "\"code\":\"035\"") {
		lwr.tc.Pass("")
	} else {
		lwr.tc.Fail("Wrong code number for login without register")
	}
	lwr.tc.Duration = time.Since(lwr.tc.starttime).Seconds()
	Info("LoginWithoutReg took " + strconv.FormatFloat(lwr.tc.Duration, 'f', -1, 64) + " seconds")
}
