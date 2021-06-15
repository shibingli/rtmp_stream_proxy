package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"unicode"
)

//ToStr *
func ToStr(value interface{}) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', 2, 32)
	case float64:
		s = strconv.FormatFloat(v, 'f', 2, 64)
	case int:
		s = strconv.FormatInt(int64(v), 10)
	case int8:
		s = strconv.FormatInt(int64(v), 10)
	case int16:
		s = strconv.FormatInt(int64(v), 10)
	case int32:
		s = strconv.FormatInt(int64(v), 10)
	case int64:
		s = strconv.FormatInt(int64(v), 10)
	case uint:
		s = strconv.FormatUint(uint64(v), 10)
	case uint8:
		s = strconv.FormatUint(uint64(v), 10)
	case uint16:
		s = strconv.FormatUint(uint64(v), 10)
	case uint32:
		s = strconv.FormatUint(uint64(v), 10)
	case uint64:
		s = strconv.FormatUint(v, 10)
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

//StringUtils *
type StringUtils string

//Bool *
func (s StringUtils) Bool() (bool, error) {
	v, err := strconv.ParseBool(s.String())
	return bool(v), err
}

//Float32 *
func (s StringUtils) Float32() (float32, error) {
	v, err := strconv.ParseFloat(s.String(), 32)
	return float32(v), err
}

//Float64 *
func (s StringUtils) Float64() (float64, error) {
	return strconv.ParseFloat(s.String(), 64)
}

//Int *
func (s StringUtils) Int() (int, error) {
	v, err := strconv.ParseInt(s.String(), 10, 32)
	return int(v), err
}

//Int8 *
func (s StringUtils) Int8() (int8, error) {
	v, err := strconv.ParseInt(s.String(), 10, 8)
	return int8(v), err
}

//Int16 *
func (s StringUtils) Int16() (int16, error) {
	v, err := strconv.ParseInt(s.String(), 10, 16)
	return int16(v), err
}

//Int32 *
func (s StringUtils) Int32() (int32, error) {
	v, err := strconv.ParseInt(s.String(), 10, 32)
	return int32(v), err
}

//Int64 *
func (s StringUtils) Int64() (int64, error) {
	v, err := strconv.ParseInt(s.String(), 10, 64)
	return int64(v), err
}

//Uint *
func (s StringUtils) Uint() (uint, error) {
	v, err := strconv.ParseUint(s.String(), 10, 32)
	return uint(v), err
}

//Uint8 *
func (s StringUtils) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(s.String(), 10, 8)
	return uint8(v), err
}

//Uint16 *
func (s StringUtils) Uint16() (uint16, error) {
	v, err := strconv.ParseUint(s.String(), 10, 16)
	return uint16(v), err
}

//Uint32 *
func (s StringUtils) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(s.String(), 10, 32)
	return uint32(v), err
}

//Uint64 *
func (s StringUtils) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(s.String(), 10, 64)
	return uint64(v), err
}

//ToTitleLower *
func (s StringUtils) ToTitleLower() string {
	str := strings.ToLower(s.String()[:1]) + s.String()[1:]
	return str
}

//ToTitleUpper *
func (s StringUtils) ToTitleUpper() string {
	str := strings.ToUpper(s.String()[:1]) + s.String()[1:]
	return str
}

//Contains *
func (s StringUtils) Contains(sep string) bool {
	index := strings.Index(s.String(), sep)
	return index > -1
}

//String *
func (s StringUtils) String() string {
	return string(s)
}

//MD5 *
func (s StringUtils) MD5() string {
	m := md5.New()
	m.Write([]byte(s.String()))
	return hex.EncodeToString(m.Sum(nil))
}

//SHA1 *
func (s StringUtils) SHA1() string {
	sha := sha1.New()
	sha.Write([]byte(s.String()))
	return hex.EncodeToString(sha.Sum(nil))
}

//SHA256 *
func (s StringUtils) SHA256() string {
	sha := sha256.New()
	sha.Write([]byte(s.String()))
	return hex.EncodeToString(sha.Sum(nil))
}

//SHA512 *
func (s StringUtils) SHA512() string {
	sha := sha512.New()
	sha.Write([]byte(s.String()))
	return hex.EncodeToString(sha.Sum(nil))
}

//HmacSHA1 *
func (s StringUtils) HmacSHA1(key string) string {
	mc := hmac.New(sha1.New, []byte(key))
	mc.Write([]byte(s.String()))
	return hex.EncodeToString(mc.Sum(nil))
}

//HmacSHA256 *
func (s StringUtils) HmacSHA256(key string) string {
	mc := hmac.New(sha256.New, []byte(key))
	mc.Write([]byte(s.String()))
	return hex.EncodeToString(mc.Sum(nil))
}

//HmacSHA512 *
func (s StringUtils) HmacSHA512(key string) string {
	mc := hmac.New(sha512.New, []byte(key))
	mc.Write([]byte(s.String()))
	return hex.EncodeToString(mc.Sum(nil))
}

//StdBase64Encode *
func (s StringUtils) StdBase64Encode() string {
	return base64.StdEncoding.EncodeToString([]byte(s.String()))
}

//URLBase64Encode *
func (s StringUtils) URLBase64Encode() string {
	return base64.URLEncoding.EncodeToString([]byte(s.String()))
}

//Base64Decode *
func (s StringUtils) Base64Decode() (string, error) {

	inputVal := s.String()

	seg, err := url.PathUnescape(inputVal)
	if nil != err {
		return "", err
	}

	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}
	v, err := base64.URLEncoding.DecodeString(seg)
	return string(v), err
}

//Base64DecodeByte *
func (s StringUtils) Base64DecodeByte() ([]byte, error) {

	inputVal := s.String()

	seg, err := url.PathUnescape(inputVal)
	if nil != err {
		return nil, err
	}

	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}
	v, err := base64.URLEncoding.DecodeString(seg)
	return v, err
}

//SBCCommaOrSemiColonToDBC 替换字符串中的全角逗号和分号为半角逗号
func (s StringUtils) SBCCommaOrSemiColonToDBC() string {
	inputVal := s.String()

	sourceRune := []rune(inputVal)

	for i, r := range sourceRune {
		if unicode.IsPunct(r) {
			if '，' == r || ';' == r || '；' == r {
				sourceRune[i] = ','
			}
		}
	}

	return string(sourceRune)
}

//TrimRightSeparator 半角字符转换为全角字符
func (s StringUtils) TrimRightSeparator() string {
	inputVal := s.String()

	inputVal = strings.TrimSuffix(inputVal, "/")
	inputVal = strings.TrimSuffix(inputVal, "\\")

	return inputVal
}
