package app

import (
	"archive/zip"
	"bufio"
	"bytes"
	"crypto/ed25519"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/busy-cloud/boat/lib"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/iot/bin"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const RootPath = "apps"
const Extension = ".app"
const ManifestName = "manifest.json"
const IconName = "icon.png"
const ListName = "__LIST__"
const SignName = "__SIGN__"

var pubKey, _ = hex.DecodeString("4f851cec1f93a757037fbb7771aead9a346df9cdd1cf623a8c00b691ac369ed5")

type Manifest struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	Author      string `json:"author,omitempty"`
	Copyright   string `json:"copyright,omitempty"`
	Url         string `json:"url,omitempty"`
}

type App struct {
	Manifest

	zipReader     *zip.ReadCloser
	zipReaderLock sync.Mutex
}

func (a *App) ServeFile(path string, ctx *gin.Context) error {
	//TODO 这里会导致前端加载缓慢。。。
	a.zipReaderLock.Lock()
	defer a.zipReaderLock.Unlock()

	if a.zipReader == nil {
		var err error
		filename := filepath.Join(RootPath, a.Id+Extension)
		a.zipReader, err = zip.OpenReader(filename)
		if err != nil {
			return err
		}
	}

	//打开文件
	file, err := a.zipReader.Open(path)
	if err != nil {
		//查找默认首页
		if errors.Is(err, os.ErrNotExist) && path != "index.html" {
			path = "index.html"
			file, err = a.zipReader.Open(path)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	//TODO 只处理 GET与OPTIONS请求

	//加入日期，方便缓存
	st, _ := file.Stat()
	ctx.Header("Last-Modified", st.ModTime().UTC().Format(gmtFormat))
	ctx.Header("Content-Type", mime.TypeByExtension(filepath.Ext(path)))

	_, err = io.Copy(ctx.Writer, file)
	return err
}

//go:embed icon.png
var defaultIcon []byte

var apps lib.Map[App]

func Load(name string) (*App, error) {
	filename := filepath.Join(RootPath, name+Extension)

	reader, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	//解析说明文件
	mani, err := reader.Open(ManifestName)
	if err != nil {
		return nil, err
	}
	manis, err := ioutil.ReadAll(mani)
	if err != nil {
		return nil, err
	}
	var app App
	err = json.Unmarshal(manis, &app.Manifest)
	if err != nil {
		return nil, err
	}

	//读取签名单
	sign, err := reader.Open(SignName)
	if err != nil {
		return nil, err
	}
	signs, err := io.ReadAll(sign)
	if err != nil {
		return nil, err
	}

	//读取校验单
	list, err := reader.Open(ListName)
	if err != nil {
		return nil, err
	}
	lists, err := io.ReadAll(list)
	if err != nil {
		return nil, err
	}

	//验证签名
	ret := ed25519.Verify(pubKey, lists, signs)
	if !ret {
		return nil, errors.New("invalid signature")
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
			return nil, err
		}
		ss := strings.SplitN(string(line), ":", 2)
		if len(ss) != 2 {
			return nil, errors.New("invalid list file: " + string(line))
		}

		found := false
		//效率有点低，但是也没办法，要不先搞个map索引？？？
		for _, f := range reader.File {
			if f.Name == ss[1] {
				found = true
				b, e := hex.DecodeString(ss[0])
				if e != nil {
					return nil, e
				}
				if bin.ParseUint32(b) != f.CRC32 {
					return nil, errors.New("invalid file:" + ss[1])
				}
				break
			}
		}

		if !found {
			return nil, errors.New("not found file:" + ss[1])
		}
	}

	apps.Store(app.Id, &app)

	return &app, nil
}

func LoadAll() error {
	entries, err := os.ReadDir(RootPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		ext := filepath.Ext(entry.Name())
		if ext != Extension {
			continue
		}
		app := strings.TrimSuffix(entry.Name(), Extension)
		_, err = Load(app)
		if err != nil {
			log.Error(err)
			continue
		}
	}

	return nil
}
