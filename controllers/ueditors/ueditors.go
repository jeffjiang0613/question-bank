package ueditors

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

func UploadImage(ctx *gin.Context) {

}

func config(ctx *gin.Context) {
	datas ,err := ioutil.ReadFile("ueditor.json")
	if err != nil {
		println(err.Error())
	}
	ctx.String(http.StatusOK,string(datas))
}

func uploadImage(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("upfile")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	hash := md5.New()
	hash.Write([]byte(time.Now().String()))
	hashId := hex.EncodeToString(hash.Sum(nil))
	filename := hashId + path.Ext(header.Filename)

	err = os.MkdirAll(path.Join("static", "upload"), 0775)
	if err != nil {
		panic(err)
	}

	outFile, err := os.Create(path.Join("static", "upload", filename))
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	io.Copy(outFile, file)

	b, err := json.Marshal(map[string]string{
		"url":      fmt.Sprintf("/static/upload/%s", filename), //保存后的文件路径
		"title":    "",                                         //文件描述，对图片来说在前端会添加到title属性上
		"original": header.Filename,                            //原始文件名
		"state":    "SUCCESS",                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	w.Write(b)
}

func Process(ctx *gin.Context) {
	if ctx.Request.Method == "GET"  && ctx.Query("action") == "config" {
		config(ctx)
	} else if ctx.Request.Method == "POST"  && ctx.Query("action") == "uploadimage" {
		uploadImage(ctx.Writer,ctx.Request)
	}
}
