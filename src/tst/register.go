package tst

import (
	. "loger"
	"strconv"
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
	rn.tc.Duration = time.Since(rn.tc.starttime).Seconds()
	Info("RegisterNormal took " + strconv.FormatFloat(rn.tc.Duration, 'f', -1, 64) + " seconds")
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
	Info("Register existing initing")
	re.name = CreateTestAccount()
	re.actionSimpleReg = action["register"]
	re.pwd = "111111"
	re.valid_code = "111111"
	re.jsid = ""
	re.tc.Testname = "Register normally"
	re.tc.starttime = time.Now()
}

func (re *RegisterExisting) Execute() {
	Info("Register existing executing")
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
	re.tc.Duration = time.Since(re.tc.starttime).Seconds()
	Info("RegisterExisting took " + strconv.FormatFloat(re.tc.Duration, 'f', -1, 64) + " seconds")
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
	Info("Register null username initing")
	rnu.name = ""
	rnu.actionSimpleReg = action["register"]
	rnu.pwd = "111111"
	rnu.valid_code = "111111"
	rnu.jsid = ""
	rnu.tc.Testname = "Register normally"
	rnu.tc.starttime = time.Now()
}

func (rnu *RegisterNulUsername) Execute() {
	Info("Register null username executing")
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
	rnu.tc.Duration = time.Since(rnu.tc.starttime).Seconds()
	Info("RegisterNulUsername took " + strconv.FormatFloat(rnu.tc.Duration, 'f', -1, 64) + " seconds")
}
