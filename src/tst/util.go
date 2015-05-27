package tst

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

func CreateTestAccount() string {
	t := time.Now().UnixNano()
	name := "go" + strconv.FormatInt(t, 10)
	return name
}

func CreateMd5str(clearPass string) string {
	h := md5.New()
	h.Write([]byte(clearPass))
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
