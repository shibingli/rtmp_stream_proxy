package utils

import (
	"encoding/base64"
	"strings"
	"testing"
)

var (
	enTxt = "http://realclouds.eriloan.net/api/v1/realclouds/down/sf"
	deTxt = "C6dvVncG7wbU6Pi16L0eS5ioNhRm_Il9t0MXDM0cWevHbVBlf7qZS5wGgWk4EaEG"
)

func TestDesEnc(t *testing.T) {

	enByte := []byte(enTxt)

	b := DesEnc(GroupFill(enByte))

	if nil == b {
		t.Fatalf("%s", "b nil")
		return
	}

	base64Txt := base64.StdEncoding.EncodeToString(b)
	base64Txt = strings.Replace(base64Txt, "+", "_", -1)
	base64Txt = strings.Replace(base64Txt, "/", "@", -1)

	t.Logf("base64 en: %s", base64Txt)
	if base64Txt == deTxt {
		t.Logf("%s", "OK")
	} else {
		t.Fatalf("%s", "en error")
	}
}

func TestDesDec(t *testing.T) {
	seg := deTxt
	seg = strings.Replace(seg, "_", "+", -1)
	seg = strings.Replace(seg, "@", "/", -1)

	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	t.Logf("base64: %s", seg)

	bd, err := base64.StdEncoding.DecodeString(seg)
	if nil != err {
		t.Fatalf("%v", err)
		return
	}

	b := DesDec(bd)

	if nil == b {
		t.Fatalf("%s", "b nil")
		return
	}

	if enTxt == string(b) {
		t.Logf("base64 dn: %s", string(b))
	} else {
		t.Fatalf("de des: %s", "b nil")
	}

	//base64Txt := base64.StdEncoding.de(b)
	//t.Logf("base64 en: %s", base64Txt)
	//if base64Txt == deTxt {
	//	t.Logf("%s", "OK")
	//}
}
