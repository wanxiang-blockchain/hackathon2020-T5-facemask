package model

import (
	"context"
	"data-manager/config"
	"data-manager/db"
	"data-manager/exterror"
	webCtx "data-manager/web/context"
	gin_session "data-manager/web/session"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	dbUser = "user"
)

type Person struct {
	Name string
	Age  int
}

var DefaultUserModel = &User{}

//var howieArray1 = GetHowieArray1()

func (this *User) SetSuperAdmin(ctx *webCtx.Context, name string, password string, addr string, passphrase string) error {

	this.ID = primitive.NewObjectID().Hex()
	this.Name = name
	this.Password = password
	this.Role = "SUPERADMIN"
	this.Address = addr
	this.Passphrase = passphrase
	this.RegisterTime = time.Now().Format("2006-01-02 15:04:05")
	this.LastLoginTime = time.Now().Format("2006-01-02 15:04:05")
	filter := bson.M{}
	filter["name"] = this.Name
	checkfilter := bson.M{}
	checkfilter["role"] = "SUPERADMIN"
	collection := ctx.DBCtx.Collection(dbUser)

	var res User
	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctxBd, checkfilter).Decode(&res)
	if nil == err {
		err = exterror.ErrSystemHasInit
		return err
	}
	err = collection.FindOne(ctxBd, filter).Decode(&res)
	if nil != err {
		if _, err := collection.InsertOne(ctxBd, this); err != nil {
			return err
		}
	} else {
		err = exterror.ErrUserHasRegistered
		return err
	}
	return nil

}
func (this *User) Test(ctx *webCtx.Context) error {
	http.HandleFunc("/get_host_info", GetHostInfoHandler)
	return nil

}

func GetHostInfoHandler(w http.ResponseWriter, r *http.Request) {
	bytes, _ := ioutil.ReadAll(r.Body)
	if len(bytes) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request := &Person{}
	err := json.Unmarshal(bytes, &request)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("Request paramter: %+v", request)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"Code": 0, "Message":"Success"}`)
}

func (this *User) UserLogin(ctx *webCtx.Context, name string, password string) (output []string, err error) {

	if ctx.Request.Method == "GET" {
		collection := ctx.DBCtx.Collection(dbUser)
		ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)
		filter := bson.M{}
		filter["name"] = name

		data := bson.M{"lastlogintime": time.Now().Format("2006-01-02 15:04:05"), "lastloginip": ctx.Request.RemoteAddr}
		var res User
		err = collection.FindOne(ctxBd, filter).Decode(&res)

		if nil != err {
			return nil, exterror.ErrUserNotExist
		}

		if res.Password != password {
			return nil, exterror.ErrPasswordWrong
		}

		_, err = collection.UpdateOne(ctxBd, filter, bson.M{"$set": data})
		if err != nil {
			return nil, err
		}
		tmpSD, ok := ctx.Get(gin_session.SessionContextName)

		if !ok {
			panic("session middleware")
		}
		sd := tmpSD.(gin_session.SessionData)
		if res.Role == "SUPERADMIN" {
			sd.Set("isSuperAdminLogin", true)
			sd.Set("isLogin", true)

		} else {
			sd.Set("isSuperAdminLogin", false)
			sd.Set("isLogin", true)
		}
		sd.Set("name", name)
		sd.Save()
		var output []string
		output = append(output, res.LastLoginTime, res.LastLoginIp, res.Role)
		return output, nil
	}
	return nil, nil
}

func (this *User) UserUpdate(ctx *webCtx.Context, name string, newpassword string, role string) (string, error) {
	collection := ctx.DBCtx.Collection(dbUser)
	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)

	filter := bson.M{}
	filter["name"] = name
	var res User
	err := collection.FindOne(ctxBd, filter).Decode(&res)

	if nil != err {
		return "", exterror.ErrUserNotExist
	}

	if role != "" && role != "SUPERADMIN" && role != "USER" {
		err := exterror.ErrBadRole
		return "", err
	}
	if newpassword != "" {
		res.Password = newpassword
	}
	var flag int
	if role == "SUPERADMIN" && res.Role != role {
		res.Role = role
		flag = 1
	}
	if role == "USER" && res.Role != role {
		res.Role = role
		flag = 2
	}
	if _, err := collection.DeleteOne(ctxBd, filter); err != nil {
		return "", err
	}

	if _, err := collection.InsertOne(ctxBd, res); err != nil {
		return "", err
	}

	addr, err := GetUserAddress(ctx)
	if nil != err {
		return "", err
	}
	if flag == 2 {
		result, err := DefaultPlatonecliModel.RoleSet(ctx, "contract-deployer", addr)
		if nil != err {
			return "", err
		}

		return result, nil
	}

	if flag == 1 {
		result, err := DefaultPlatonecliModel.RoleSet(ctx, "chain-admin", addr)
		if nil != err {
			return "", err
		}

		return result, nil
	}

	return "", nil
}

func (this *User) AddUser(ctx *webCtx.Context, name string, password string, role string, passphrase string) (string, error) {

	this.ID = primitive.NewObjectID().Hex()
	this.Name = name
	this.Password = password
	this.Role = role
	this.RegisterTime = time.Now().Format("2006-01-02 15:04:05")
	if this.Role != "SUPERADMIN" && this.Role != "USER" {
		err := exterror.ErrBadRole
		return "", err
	}

	filter := bson.M{}
	filter["name"] = this.Name
	collection := ctx.DBCtx.Collection(dbUser)

	var res User
	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := collection.FindOne(ctxBd, filter).Decode(&res)
	if nil != err {
		if _, err := collection.InsertOne(ctxBd, this); err != nil {
			logrus.Errorln(err)
			return "", err
		}
	} else {
		err = exterror.ErrUserHasRegistered
		return "", err
	}

	addr, err := DefaultPlatonecliModel.GetAccount(ctx, name, passphrase)
	if this.Role == "SUPERADMIN" {
		result, err := DefaultPlatonecliModel.RoleSet(ctx, "chain-admin", addr)
		if nil != err {
			return "", err
		}
		return result, nil
	} else {
		result, err := DefaultPlatonecliModel.RoleSet(ctx, "contract-deployer", addr)
		if nil != err {
			return "", err
		}
		return result, nil
	}
	return "", nil

}

func (this *User) DeleteUser(ctx *webCtx.Context, name string) error {

	this.ID = primitive.NewObjectID().Hex()
	this.Name = name

	filter := bson.M{}
	filter["name"] = this.Name
	collection := ctx.DBCtx.Collection(dbUser)

	var res User
	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)

	err := collection.FindOne(ctxBd, filter).Decode(&res)
	if nil == err {
		if _, err := collection.DeleteOne(ctxBd, filter); err != nil {
			return err
		}
	} else {
		err = exterror.ErrUserNotExist
		return err
	}
	return nil

}

func (this *User) GetUser(ctx *webCtx.Context, pageindex int64, pagesize int64) ([]*User, error) {
	collection := ctx.DBCtx.Collection(dbUser)

	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.D{{}}
	findOps := db.BuildOptionsByQuery(pageindex, pagesize)
	//findOps.SetSort(bson.D{{"height", -1}})

	cur, err := collection.Find(ctxBd, filter, findOps)
	if err != nil {
		return nil, err
	}

	results := make([]*User, 20)
	err = cur.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(results); i++ {
		results[i].Password = "密码不可见"
	}
	return results, nil

}

func (this *User) GetAccount(ctx *webCtx.Context, p string) (string, error) {

	client := &http.Client{}
	data := "passphrase=" + p
	req, err := http.NewRequest("POST", config.Config.HttpConf.RestServer+"/accounts", strings.NewReader(data))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	return string(body), nil
}
