package utils

import (
	"net/url"
	"testing"
	"time"
)

func TestConvertToTime(t *testing.T) {

	s := int64(1601852937109189300)
	//s := int64(1597744469875374950)
	//s := int64(1597742179739)
	//s := int64(1597744254)
	tt, err := ConvertToTime(s)
	if nil != err {
		t.Fatalf("%v", err.Error())
		return
	}
	t.Logf("%s", tt.Format(time.RFC3339Nano))
}

func TestQueryEscapePwd(t *testing.T) {
	pwd := url.QueryEscape("Cc*LiveManager!2020")
	username := url.QueryEscape("hd_live")
	t.Logf("username: %s  ==  password: %s\n", username, pwd)
}
