package tst

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	. "loger"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func CreateTestAccount() string {
	t := time.Now().UnixNano()
	name := "go" + strconv.FormatInt(t, 10)
	return name
}

func CreateMd5str(clearStr string) string {
	h := md5.New()
	h.Write([]byte(clearStr))
	return hex.EncodeToString(h.Sum(nil))
}

func RandNum(l int) string {
	var rn, r string
	var ir int64
	b := make([]byte, l)
	for j, _ := range b {
		b[j] = '0'
	}
	switch l {
	case 4:
		ir = rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(9999)
	case 6:
		ir = rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(999999)
	}
	r = strconv.FormatInt(ir, 10)
	for i := l - 1; i >= 0 && (i-l+len(r)) >= 0; i-- {
		b[i] = r[i-l+len(r)]
	}
	rn = string(b[:])
	return rn
}

func baseUrl(act, jsid string) (u string) {
	u = "http://" + Env.ServerIp1 + ":" + Env.ServerPort1 + act + ";jsessionid=" + jsid + "?"
	return u
}

func CreateAGBHeader(salerId, salerKey, salerVer string) http.Header {
	h := make(http.Header)
	timestamp := time.Now().Format("20060102150405")
	headerId := timestamp + RandNum(4)
	term_no := timestamp + RandNum(6)
	t := headerId + timestamp + term_no + salerKey
	h.Add("id", headerId)
	h.Add("reqtimestamp", timestamp)
	h.Add("terminal_no", term_no)
	h.Add("sign", CreateMd5str(t))
	h.Add("sp_id", salerId)
	h.Add("access_version", salerVer)
	return h
}

func httpreq(serverip, serverport, url string) (r string, err error) {
	var hc http.Client
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		ErrMsg("Init req faild")
		return "Init req faild", err
	}
	req.Header = CreateAGBHeader(Env.SalerId1, Env.SalerKey1, Env.AppVer1)
	resp, err := hc.Do(req)
	if err != nil {
		ErrMsg("Do req failed")
		return "Do req failed", err
	}
	return getRes(resp)
}

func getRes(resp *http.Response) (d string, err error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		ErrMsg("Read res fail")
		return "Read res fail", err
	}
	d = string(b)
	return
}
