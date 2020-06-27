package test

import (
	"blog/common"
	"blog/util"
	"testing"
)

func TestRegexp(t *testing.T) {
	var e = "691129301qq.com"
	t.Log(util.RegexpEmail(e) == nil)

}

func TestMail(t *testing.T) {
	t.Log(common.Mail("691129301@qq.com", "测试", "yoyo", "text"))

}
