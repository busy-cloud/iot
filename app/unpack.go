package app

import (
	"archive/zip"
	"bufio"
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"github.com/busy-cloud/iot/bin"
	"io"
	"strings"
)

var pubKey, _ = hex.DecodeString("4f851cec1f93a757037fbb7771aead9a346df9cdd1cf623a8c00b691ac369ed5")

func Verify(filename string) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()

	//读取签名单
	sign, err := reader.Open("__SIGN__")
	if err != nil {
		return err
	}
	signs, err := io.ReadAll(sign)
	if err != nil {
		return err
	}

	//读取校验单
	list, err := reader.Open("__LIST__")
	if err != nil {
		return err
	}
	lists, err := io.ReadAll(list)
	if err != nil {
		return err
	}

	//验证签名
	ret := ed25519.Verify(pubKey, lists, signs)
	if !ret {
		return errors.New("invalid signature")
	}

	//逐行验证文件校验
	rdr := bufio.NewReader(bytes.NewReader(lists))
	for {
		//line, err := rdr.ReadString('\n')
		line, _, err := rdr.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		ss := strings.SplitN(string(line), ":", 2)
		if len(ss) != 2 {
			return errors.New("invalid __LIST__ " + string(line))
		}

		found := false
		//效率有点低，但是也没办法，要不先搞个map索引？？？
		for _, f := range reader.File {
			if f.Name == ss[1] {
				found = true
				b, e := hex.DecodeString(ss[0])
				if e != nil {
					return e
				}
				if bin.ParseUint32(b) != f.CRC32 {
					return errors.New("invalid file:" + ss[1])
				}
				break
			}
		}

		if !found {
			return errors.New("not found file:" + ss[1])
		}
	}

	return nil
}
