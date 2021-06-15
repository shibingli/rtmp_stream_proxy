package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"log"
)

const (
	//定义密钥
	key = "74660466"
	//定义向量
	iv = "27342962"
)

func GroupFill(plaintext []byte) []byte {

	//计算分组后需要填充的字节数；
	//如果字节不足8个，需要填充不足的字节数，比如不足字节数为3个，则填充[3,3,3]；
	//如果字节刚好是8的倍数，则需要填充8个字节，如[8,8,8,8,8,8,8,8]
	//这样在解码删除时，就可以根据最后一个数确定删除的字节数
	fillCount := des.BlockSize - len(plaintext)%des.BlockSize
	//将int类型的变量fillCount添加到[]byte
	bt := []byte{byte(fillCount)}
	//将字节重复fillCount次
	fillByte := bytes.Repeat(bt, fillCount)
	//将填充的字节添加到名文当中
	plaintext = append(plaintext, fillByte...)
	return plaintext
}

//DesEnc 加密
func DesEnc(plaintext []byte) []byte {
	var (
		block cipher.Block
		err   error
	)
	//创建des
	if block, err = des.NewCipher([]byte(key)); err != nil {
		log.Fatal(err)
	}
	//创建cbc加密模式(链密码模式)
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	//明文加密，src和dst可指向同一内存地址
	blockMode.CryptBlocks(plaintext, plaintext)
	return plaintext

}

//DesDec 解密
func DesDec(ciptext []byte) []byte {
	var (
		block cipher.Block
		err   error
	)
	//创建des
	if block, err = des.NewCipher([]byte(key)); err != nil {
		log.Fatal(err)
	}
	//创建cbc解密模式
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	//解密
	blockMode.CryptBlocks(ciptext, ciptext)
	count := len(ciptext)
	//减去填充的字节
	ciptext = ciptext[:count-int(ciptext[count-1])]
	return ciptext
}
