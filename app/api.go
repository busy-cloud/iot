package app

import (
	"archive/zip"
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/log"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"time"
)

func init() {

	now := time.Now()

	api.Register("GET", "iot/app/list", func(ctx *gin.Context) {
		var as []*Manifest

		entries, err := os.ReadDir(APP_PATH)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			ext := filepath.Ext(entry.Name())
			if ext != APP_EXT {
				continue
			}
			app, err := ReadManifest(entry.Name())
			if err != nil {
				log.Error(err)
				continue
			}
			as = append(as, app)
		}

		api.OK(ctx, as)
	})

	api.Register("GET", "iot/app/:app/icon", func(ctx *gin.Context) {
		reader, err := zip.OpenReader(filepath.Join(APP_PATH, ctx.Param("app")+APP_EXT))
		if err != nil {
			_ = ctx.Error(err)
			return
		}
		defer reader.Close()

		//ctx.Writer.WriteHeader(http.StatusNotModified)

		file, err := reader.Open(APP_ICON)
		if err != nil {
			//return nil, err
			//return icon, nil //使用默认图片
			ctx.Header("Last-Modified", now.UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
			ctx.Header("Content-Type", "image/png")
			_, _ = ctx.Writer.Write(icon)
			return
		}
		defer file.Close()

		st, _ := file.Stat()
		buf, err := io.ReadAll(file)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.Header("Last-Modified", st.ModTime().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
		ctx.Header("Content-Type", "image/png")
		_, _ = ctx.Writer.Write(buf)
	})
}
