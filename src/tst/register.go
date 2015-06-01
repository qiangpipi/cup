package tst

import (
	. "loger"
	"strings"
	"time"
)

type RegisterNormal struct {
	tc Testcase
	actionSimpleReg,
	name,
	pwd,
	jsid,
	valid_code string
}

func (rn *RegisterNormal) Init() {
	Info("Register normal initing")
	rn.name = CreateTestAccount()
	rn.actionSimpleReg = action["register"]
	rn.pwd = "111111"
	rn.valid_code = "111111"
	rn.jsid = ""
	rn.tc.Testname = "Register normally"
	rn.tc.starttime = time.Now()
}

func (rn *RegisterNormal) Execute() {
	Info("Register normal executing")
	//////////////////////////
	//Simple register account
	/////////////////////////
	url := baseUrl(rn.actionSimpleReg, rn.jsid)
	params := "name=" + rn.name + "&pwd=" + rn.pwd + "&valid_code=" + rn.valid_code
	url += params
	Debug(url)
	res, err := httpreq(Env.ServerIp1, Env.ServerPort1, url)
	if err != nil {
		rn.tc.Fail(err.Error() + "\n" + res)
	}
	//Res parse
	Debug(res)
	if strings.Contains(res, "\"code\":\"000\"") {
		rn.tc.Pass("")
	} else {
		rn.tc.Fail("Wrong code num and register fail")
	}
}

type RegisterExisting struct {
	tc Testcase
	actionSimpleReg,
	name,
	pwd,
	jsid,
	valid_code string
}

func (re *RegisterExisting) Init() {
	Info("Register normal initing")
	re.name = CreateTestAccount()
	re.actionSimpleReg = action["register"]
	re.pwd = "111111"
	re.valid_code = "111111"
	re.jsid = ""
	re.tc.Testname = "Register normally"
	re.tc.starttime = time.Now()
}

func (re *RegisterExisting) Execute() {
	Info("Register normal executing")
	//////////////////////////
	//Simple register account
	/////////////////////////
	url := baseUrl(re.actionSimpleReg, re.jsid)
	params := "name=" + re.name + "&pwd=" + re.pwd + "&valid_code=" + re.valid_code
	url += params
	Debug(url)
	res, err := httpreq(Env.ServerIp1, Env.ServerPort1, url)
	if err != nil {
		re.tc.Fail(err.Error() + "\n" + res)
	}
	//Res parse
	Debug(res)
	if strings.Contains(res, "\"code\":\"000\"") {
		Debug(res)
		//////////////////////////
		//Register account again
		/////////////////////////
		res, err := httpreq(Env.ServerIp1, Env.ServerPort1, url)
		if err != nil {
			re.tc.Fail(err.Error() + "\n" + res)
		}
		Debug(res)
		if strings.Contains(res, "\"code\":\"030\"") {
			re.tc.Pass("")
		} else {
			re.tc.Fail("Wrong code num for existing account")
		}
	} else {
		re.tc.Fail("Wrong code num and register fail")
	}
}

type RegisterNulUsername struct {
	tc Testcase
	actionSimpleReg,
	name,
	pwd,
	jsid,
	valid_code string
}

func (rnu *RegisterNulUsername) Init() {
	Info("Register normal initing")
	rnu.name = ""
	rnu.actionSimpleReg = action["register"]
	rnu.pwd = "111111"
	rnu.valid_code = "111111"
	rnu.jsid = ""
	rnu.tc.Testname = "Register normally"
	rnu.tc.starttime = time.Now()
}

func (rnu *RegisterNulUsername) Execute() {
	Info("Register normal executing")
	//////////////////////////
	//Simple register account
	/////////////////////////
	url := baseUrl(rnu.actionSimpleReg, rnu.jsid)
	params := "name=" + rnu.name + "&pwd=" + rnu.pwd + "&valid_code=" + rnu.valid_code
	url += params
	Debug(url)
	res, err := httpreq(Env.ServerIp1, Env.ServerPort1, url)
	if err != nil {
		rnu.tc.Fail(err.Error() + "\n" + res)
	}
	//Res parse
	Debug(res)
	if strings.Contains(res, "\"code\":\"031\"") {
		rnu.tc.Pass("")
	} else {
		rnu.tc.Fail("Wrong code num and register fail")
	}
}
