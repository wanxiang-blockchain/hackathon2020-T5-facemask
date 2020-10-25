package model

import (
	"bytes"
	"context"
	"data-manager/config"
	"data-manager/exterror"
	webCtx "data-manager/web/context"
	gin_session "data-manager/web/session"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func doRoleSet(jsonStr []byte, role string) (string, error) {
	client := &http.Client{}
	url := config.Config.HttpConf.RestServer + "/role/role-lists/" + role
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", nil}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil}
	fmt.Println(string(body))

	result := string(body)
	return result, nil
}

func GetUserAddress(ctx *webCtx.Context) (string, error) {
	tmpSD, _ := ctx.Get(gin_session.SessionContextName)
	sd := tmpSD.(gin_session.SessionData)
	username, err := sd.Get("name")

	collection := ctx.DBCtx.Collection(dbUser)
	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{}
	filter["name"] = username

	var res User
	err = collection.FindOne(ctxBd, filter).Decode(&res)
	if nil != err {
		return "", exterror.ErrUserNotExist
	}
	return res.Address, nil
}


func GetUserPassphrase(ctx *webCtx.Context) (string, error) {
	tmpSD, _ := ctx.Get(gin_session.SessionContextName)
	sd := tmpSD.(gin_session.SessionData)
	username, err := sd.Get("name")

	collection := ctx.DBCtx.Collection(dbUser)
	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{}
	filter["name"] = username

	var res User
	err = collection.FindOne(ctxBd, filter).Decode(&res)
	if nil != err {
		return "", exterror.ErrUserNotExist
	}
	return res.Passphrase, nil
}

func NewUploadFile(name, path string) *uploadFile {
	return &uploadFile{
		name: name,
		path: path,
	}
}

type uploadFile struct {
	path string
	name string
}

func genMultiPartBody(files []*uploadFile, reqParams ...map[string]string) (io.Reader, string, error) {
	body := bytes.NewBufferString("")
	writer := multipart.NewWriter(body)

	// write file to body
	for _, file := range files {
		f, err := os.Open(file.path)
		if err != nil {
			return nil, "", err
		}

		part, err := writer.CreateFormFile(file.name, filepath.Base(file.path))
		if err != nil {
			return nil, "", err
		}

		_, err = io.Copy(part, f)
		if err != nil {
			return nil, "", err
		}

		err = f.Close()
		if err != nil {
			return nil, "", err
		}
	}

	// write key value to body
	for _, category := range reqParams {
		for k, v := range category {
			_ = writer.WriteField(k, v)
		}
	}

	_ = writer.Close()

	return body, writer.FormDataContentType(), nil
}
