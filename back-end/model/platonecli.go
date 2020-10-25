package model

import (
	"bytes"
	"context"
	"data-manager/config"
	"data-manager/exterror"
	webCtx "data-manager/web/context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var DefaultPlatonecliModel = &Platonecli{}

func (p *Platonecli) GetAccount(ctx *webCtx.Context, name string, pass string) (string, error) {

	client := &http.Client{}
	data := "passphrase=" + pass
	logrus.Debug(data)
	url := config.Config.HttpConf.RestServer + "/accounts"
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "", nil
	}
	userAddr := &UserAddress{}
	json.Unmarshal(body, userAddr)

	fmt.Println(userAddr)
	result := userAddr.Address

	collection := ctx.DBCtx.Collection(dbUser)
	ctxBd, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{}
	filter["name"] = name

	field := bson.M{"address": userAddr.Address, "passphrase": pass}
	var res User
	err = collection.FindOne(ctxBd, filter).Decode(&res)
	if res.Address != "" {
		err := exterror.ErrUserHasRegistered
		return "", err
	}
	if nil != err {
		return "", exterror.ErrUserNotExist
	}

	_, err = collection.UpdateOne(ctxBd, filter, bson.M{"$set": field})

	return result, nil
}

func (p *Platonecli) GetSystemConfig(ctx *webCtx.Context) (string, error) {

	data := "endPoint=" + config.Config.HttpConf.GetEndpoint()
	blockGasLimit, err := http.NewRequest("GET", config.Config.HttpConf.RestServer+"/sysConfig/block-gas-limit", strings.NewReader(data))
	txGasLimit, err := http.NewRequest("GET", config.Config.HttpConf.RestServer+"/sysConfig/tx-gas-limit", strings.NewReader(data))
	isUseGas, err := http.NewRequest("GET", config.Config.HttpConf.RestServer+"/sysConfig/is-tx-use-gas", strings.NewReader(data))
	isApproveDeployedContract, err := http.NewRequest("GET", config.Config.HttpConf.RestServer+"/sysConfig/is-approve-deployed-contract", strings.NewReader(data))
	checkDeployPermission, err := http.NewRequest("GET", config.Config.HttpConf.RestServer+"/sysConfig/check-contract-deploy-permission", strings.NewReader(data))
	isProduceEmptyBlock, err := http.NewRequest("GET", config.Config.HttpConf.RestServer+"/sysConfig/is-produce-empty-block", strings.NewReader(data))
	gasContractName, err := http.NewRequest("GET", config.Config.HttpConf.RestServer+"/sysConfig/gas-contract-name", strings.NewReader(data))

	systemConfig := &SystemConfigBody{}
	systemConfig.BlockGasLimit, err = GetHttpResponse(blockGasLimit)
	if err != nil {
		logrus.Debug("get blockGasLimit failed")
		return "", err
	}
	systemConfig.TxGasLimit, err = GetHttpResponse(txGasLimit)
	if err != nil {
		logrus.Debug("get txGasLimit failed")
		return "", err
	}
	systemConfig.IsUseGas, err = GetHttpResponse(isUseGas)
	if err != nil {
		logrus.Debug("get isUseGas failed")
		return "", err
	}
	systemConfig.IsApproveDeployedContract, err = GetHttpResponse(isApproveDeployedContract)
	if err != nil {
		logrus.Debug("get isApproveDeployedContract failed")
		return "", err
	}

	systemConfig.IsCheckDeployPermission, err = GetHttpResponse(checkDeployPermission)
	if err != nil {
		logrus.Debug("get checkDeployPermission failed")
		return "", err
	}
	systemConfig.IsProduceEmptyBlock, err = GetHttpResponse(isProduceEmptyBlock)
	if err != nil {
		logrus.Debug("get isProduceEmptyBlock failed")
		return "", err
	}
	systemConfig.GasContractName, err = GetHttpResponse(gasContractName)
	if err != nil {
		logrus.Debug("get gasContractName failed")
		return "", err
	}
	logrus.Debug(systemConfig)
	//systemConfig.gasContractName = 	strings.Replace(systemConfig.gasContractName, "\u0000", "", -1)

	resultJson, err := json.Marshal(*systemConfig)

	return string(resultJson), nil
}

func (p *Platonecli) SetSystemConfig(ctx *webCtx.Context, config SystemConfigBody) (string, error) {
	if config.BlockGasLimit != "" {
		//v := reflect.ValueOf(config)
		//changedConfigTag := v.Type().Field(0).Tag.Get("json")
		_, err := DoSetSystemConfig(ctx, "BlockGasLimit", config)
		if nil != err {
			return "", err
		}
	}
	if config.TxGasLimit != "" {
		//v := reflect.ValueOf(config)
		//changedConfigTag := v.Type().Field(1).Tag.Get("json")
		_, err := DoSetSystemConfig(ctx, "TxGasLimit", config)
		if nil != err {
			return "", err
		}
	}
	if config.IsUseGas != "" {
		//v := reflect.ValueOf(config)
		//changedConfigTag := v.Type().Field(2).Tag.Get("json")
		_, err := DoSetSystemConfig(ctx, "IsUseGas", config)
		if nil != err {
			return "", err
		}
	}
	if config.IsApproveDeployedContract != "" {
		//v := reflect.ValueOf(config)
		//changedConfigTag := v.Type().Field(3).Tag.Get("json")
		_, err := DoSetSystemConfig(ctx, "IsApproveDeployedContract", config)
		if nil != err {
			return "", err
		}
	}
	if config.IsCheckDeployPermission != "" {
		//v := reflect.ValueOf(config)
		//changedConfigTag := v.Type().Field(4).Tag.Get("json")
		_, err := DoSetSystemConfig(ctx, "IsCheckDeployPermission", config)
		if nil != err {
			return "", err
		}
	}
	if config.IsProduceEmptyBlock != "" {
		//v := reflect.ValueOf(config)
		//changedConfigTag := v.Type().Field(5).Tag.Get("json")
		_, err := DoSetSystemConfig(ctx, "IsProduceEmptyBlock", config)
		if nil != err {
			return "", err
		}
	}
	if config.GasContractName != "" {
		//v := reflect.ValueOf(config)
		//changedConfigTag := v.Type().Field(6).Tag.Get("json")
		_, err := DoSetSystemConfig(ctx, "GasContractName", config)
		if nil != err {
			return "", err
		}
	}

	return "success", nil

}

func DoSetSystemConfig(ctx *webCtx.Context, changedConfigField string, systemconfig SystemConfigBody) (string, error) {
	client := &http.Client{}

	var requestBody ParamRequest
	var err error
	requestBody.Tx.From, err = GetUserAddress(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Rpc.Endpoint = config.Config.HttpConf.GetEndpoint()
	requestBody.Rpc.Passphrase, err = GetUserPassphrase(ctx)
	requestBody.Contract.Data = systemconfig
	if nil != err {
		return "", err
	}

	jsonStr, err := json.Marshal(&requestBody)
	url, err := getUrl(changedConfigField)
	if nil != err {
		return "", err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	result := string(body)
	return result, nil
}

func getUrl(arg string) (string, error) {
	switch arg {
	case "BlockGasLimit":
		return config.Config.HttpConf.RestServer + "/sysConfig/block-gas-limit", nil
	case "TxGasLimit":
		return config.Config.HttpConf.RestServer + "/sysConfig/tx-gas-limit", nil
	case "IsUseGas":
		return config.Config.HttpConf.RestServer + "/sysConfig/is-tx-use-gas", nil
	case "IsApproveDeployedContract":
		return config.Config.HttpConf.RestServer + "/sysConfig/is-approve-deployed-contract", nil
	case "IsCheckDeployPermission":
		return config.Config.HttpConf.RestServer + "/sysConfig/check-contract-deploy-permission", nil
	case "IsProduceEmptyBlock":
		return config.Config.HttpConf.RestServer + "/sysConfig/is-produce-empty-block", nil
	case "GasContractName":
		return config.Config.HttpConf.RestServer + "/sysConfig/gas-contract-name", nil
	default:
		return "", exterror.ErrParameterInvalid
	}
}

func GetHttpResponse(req *http.Request) (string, error) {
	client := &http.Client{}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	result := string(body)
	return result, nil
}

func (p *Platonecli) SetCnsName(ctx *webCtx.Context, cnsRequestInfo CnsRequest) (string, error) {

	var requestBody CnsRequestBody
	var err error
	requestBody.Tx.From, err = GetUserAddress(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Rpc.Endpoint = config.Config.HttpConf.GetEndpoint()
	requestBody.Contract.Data.Name = cnsRequestInfo.Name
	requestBody.Rpc.Passphrase, err = GetUserPassphrase(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Contract.Data.Version = cnsRequestInfo.Version
	requestBody.Contract.Data.Address = cnsRequestInfo.Address

	url := config.Config.HttpConf.GetRestServer() + "/cns/components"
	logrus.Debug(requestBody)
	contentType := "application/json;charset=utf-8"
	jsonStr, err := json.Marshal(&requestBody)
	logrus.Debug(jsonStr)
	resp, err := http.Post(url, contentType, bytes.NewBuffer(jsonStr))

	if err != nil {
		// handle error
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err

	}
	fmt.Println(string(body))

	result := string(body)
	return result, nil
}

func (p *Platonecli) CnsRedirect(ctx *webCtx.Context, cnsRequestInfo CnsRequest) (string, error) {

	client := &http.Client{}
	var requestBody CnsRequestBody
	var err error
	requestBody.Tx.From, err = GetUserAddress(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Rpc.Endpoint = config.Config.HttpConf.GetEndpoint()
	requestBody.Rpc.Passphrase, err = GetUserPassphrase(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Contract.Data.Name = cnsRequestInfo.Name
	requestBody.Contract.Data.Version = cnsRequestInfo.Version

	logrus.Debug(requestBody)
	jsonStr, err := json.Marshal(&requestBody)

	url := config.Config.HttpConf.GetRestServer() + "/cns/mappings/" + cnsRequestInfo.Name
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		// handle error
		return "", nil
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "", nil
	}
	fmt.Println(string(body))

	result := string(body)
	return result, nil
}

func (p *Platonecli) OpenFireWall(ctx *webCtx.Context, fireWallRequestInfo FireWallRequest) (string, error) {

	client := &http.Client{}
	var requestBody FireWallRequestBody
	var err error
	requestBody.Tx.From, err = GetUserAddress(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Rpc.Endpoint = config.Config.HttpConf.GetEndpoint()
	requestBody.Rpc.Passphrase, err = GetUserPassphrase(ctx)
	if nil != err {
		return "", err
	}
	logrus.Debug(requestBody)
	jsonStr, err := json.Marshal(&requestBody)

	//var jsonStr = []byte(requestBody)
	url := config.Config.HttpConf.GetRestServer() + "/fw/" + fireWallRequestInfo.Address + "/on"
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		// handle error
		return "", nil
	}

	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return "", nil
	}
	fmt.Println(string(body))

	result := string(body)
	return result, nil
}

func (p *Platonecli) CloseFireWall(ctx *webCtx.Context, fireWallRequestInfo FireWallRequest) (string, error) {

	client := &http.Client{}
	var requestBody FireWallRequestBody
	var err error
	requestBody.Tx.From, err = GetUserAddress(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Rpc.Endpoint = config.Config.HttpConf.GetEndpoint()
	requestBody.Rpc.Passphrase, err = GetUserPassphrase(ctx)
	if nil != err {
		return "", err
	}
	logrus.Debug(requestBody)
	jsonStr, err := json.Marshal(&requestBody)
	url := config.Config.HttpConf.GetRestServer() + "/fw/" + fireWallRequestInfo.Address + "/off"
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", nil
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	fmt.Println(string(body))

	result := string(body)
	return result, nil
}

func (p *Platonecli) AddNode(ctx *webCtx.Context, nodeRequestInfo NodeOperateRequest) (string, error) {

	client := &http.Client{}
	var requestBody NodeOperateRequestBody
	var err error
	requestBody.Tx.From, err = GetUserAddress(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Rpc.Endpoint = config.Config.HttpConf.GetEndpoint()
	requestBody.Rpc.Passphrase, err = GetUserPassphrase(ctx)
	if nil != err {
		return "", err
	}

	requestBody.Contract.Data.Info.Name = nodeRequestInfo.Name
	requestBody.Contract.Data.Info.Status = nodeRequestInfo.Status
	requestBody.Contract.Data.Info.ExternalIP = nodeRequestInfo.ExternalIP
	requestBody.Contract.Data.Info.InternalIP = nodeRequestInfo.InternalIP
	requestBody.Contract.Data.Info.PublicKey = nodeRequestInfo.PublicKey
	requestBody.Contract.Data.Info.P2pPort = nodeRequestInfo.P2pPort

	logrus.Debug(requestBody)
	jsonStr, err := json.Marshal(&requestBody)
	url := config.Config.HttpConf.GetRestServer() + "/node/components"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		// handle error
		return "", nil
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	fmt.Println(string(body))

	result := string(body)
	return result, nil
}

func (p *Platonecli) DeleteNode(ctx *webCtx.Context, nodeRequestInfo DeleteNodeOperateRequest) (string, error) {

	client := &http.Client{}
	var requestBody NodeOperateRequestBody
	var err error
	requestBody.Tx.From, err = GetUserAddress(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Rpc.Endpoint = config.Config.HttpConf.GetEndpoint()
	requestBody.Rpc.Passphrase, err = GetUserPassphrase(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Contract.Data.Info.Status = nodeRequestInfo.Status

	logrus.Debug(requestBody)
	jsonStr, err := json.Marshal(&requestBody)
	url := config.Config.HttpConf.GetRestServer() + "/node/components/" + nodeRequestInfo.Name
	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", nil
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	fmt.Println(string(body))

	result := string(body)
	return result, nil
}

func (p *Platonecli) DeployContract(ctx *webCtx.Context, deployRequestInfo DeployRequest) (string, error) {

	file1 := NewUploadFile("code", deployRequestInfo.CodePath)
	file2 := NewUploadFile("abi", deployRequestInfo.AbiPath)
	client := &http.Client{}

	var requestBody DeployRequestBody
	var err error
	requestBody.Tx.From, err = GetUserAddress(ctx)
	if nil != err {
		return "", err
	}

	requestBody.Rpc.Endpoint = config.Config.HttpConf.GetEndpoint()
	requestBody.Rpc.Passphrase, err = GetUserPassphrase(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Contract.Interpreter = deployRequestInfo.Interpreter
	jsonStr, err := json.Marshal(&requestBody)
	resParam1 := make(map[string]string)
	resParam1["info"] = string(jsonStr)

	body, contentType, err := genMultiPartBody([]*uploadFile{file1, file2}, resParam1)

	url := config.Config.HttpConf.GetRestServer() + "/contracts"
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", nil
	}
	req.Header.Set("content-type", contentType)
	resp, err := client.Do(req)

	defer resp.Body.Close()

	body1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	fmt.Println(string(body1))

	result := string(body1)
	return result, nil
}

func (p *Platonecli) RoleSet(ctx *webCtx.Context, role string, addr string) (string, error) {

	var requestBody ContractDeploy
	var err error
	requestBody.Tx.From, err = GetUserAddress(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Rpc.Endpoint = config.Config.HttpConf.GetEndpoint()
	requestBody.Rpc.Passphrase, err = GetUserPassphrase(ctx)
	if nil != err {
		return "", err
	}
	requestBody.Contract.Data.Address = addr

	logrus.Debug(requestBody)
	jsonStr, err := json.Marshal(&requestBody)
	switch role {
	case "contract-deployer":
		{
			return doRoleSet(jsonStr, role)
		}
	case "group-admin":
		{
			return doRoleSet(jsonStr, role)
		}
	case "node-admin":
		{
			return doRoleSet(jsonStr, role)
		}
	case "contract-admin":
		{
			return doRoleSet(jsonStr, role)
		}
	case "chain-admin":
		{
			return doRoleSet(jsonStr, role)
		}
	default:
		return "", exterror.ErrBadRole
	}

}
