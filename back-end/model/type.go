package model

type User struct {
	ID            string `json:"id" bson:"_id"`
	Name          string `json:"name" bson:"name"`
	Password      string `json:"password" bson:"password"`
	Address       string `json:"address"  bson:"address"`
	PublicKey     string `json:"publicKey"  bson:"publicKey"`
	PrivateKey    string `json:"privateKey"  bson:"privateKey"`
	Passphrase    string `json:"passphrase"  bson:"passphrase"`
	Role          string `json:"role"  bson:"role"`
	RegisterTime  string `json:"registertime"  bson:"registertime"`
	LastLoginTime string `json:"lastlogintime"  bson:"lastlogintime"`
	LastLoginIp   string `json:"lastloginip"  bson:"lastloginip"`
}

type Node struct {
	ID            string `json:"id" bson:"_id"`
	NodePublicKey string `json:"nodepublickey" bson:"nodepublickey"`
	MonitorIp     string `json:"monitorip" bson:"monitorip"`
	MonitorPort   string `json:"monitorport" bson:"monitorport"`
	Groupid       string `json:"groupid" bson:"groupid"`
	Chainid       string `json:"chainid" bson:"chainid"`
	NodeIp        string `json:"nodeip" bson:"nodeip"`
	NodePort      uint32 `json:"nodeport" bson:"nodeport"`
	Rpcport       string `json:"rpcport" bson:"rpcport"`
	Wsport        string `json:"wsport" bson:"wsport"`
	Dashport      string `json:"dashport" bson:"dashport"`
	Password      string `json:"password" bson:"password"`
	CreatorEnode  string `json:"creatorenode" bson:"creatorenode"`
	Bootnodes     string `json:"bootnodes" bson:"bootnodes"`
	NodeName      string `json:"nodename" bson:"nodename"`
	Status        uint32 `json:"status" bson:"status"`
}

type Page struct {
	PageIndex int64 `json:"pageindex" bson:"pageindex"`
	PageSize  int64 `json:"pagesize" bson:"pagesize"`
}

type Platonecli struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

type SystemConfig struct {
	BlockGasLimit             string `json:"blockGasLimit" bson:"blockGasLimit"`
	TxGasLimit                string `json:"txGasLimit" bson:"txGasLimit"`
	IsUseGas                  string `json:"isUseGas" bson:"isUseGas"`
	IsApproveDeployedContract string `json:"isApproveDploeyedContract" bson:"isApproveDploeyedContract"`
	IsCheckDeployPermission   string `json:"isCheckDeployPermission" bson:"isCheckDeployPermission"`
	IsProduceEmptyBlock       string `json:"isProduceEmptyBlock" bson:"isProduceEmptyBlock"`
	GasContractName           string `json:"gasContractName" bson:"gasContractName"`
	Sender                    string `json:"sender" bson:"sender"`
	Endpoint                  string `json:"endpoint" bson:"endpoint"`
	Passphrase                string `json:"passphrase" bson:"passphrase"`
}

type SystemConfigBody struct {
	BlockGasLimit             string `json:"blockGasLimit, omitempty" bson:"blockGasLimit"`
	TxGasLimit                string `json:"txGasLimit, omitempty" bson:"txGasLimit"`
	IsUseGas                  string `json:"isUseGas, omitempty" bson:"isUseGas"`
	IsApproveDeployedContract string `json:"isApproveDeployedContract, omitempty" bson:"isApproveDploeyedContract"`
	IsCheckDeployPermission   string `json:"isCheckDeployPermission, omitempty" bson:"isCheckDeployPermission"`
	IsProduceEmptyBlock       string `json:"isProduceEmptyBlock, omitempty" bson:"isProduceEmptyBlock"`
	GasContractName           string `json:"gasContractName, omitempty" bson:"gasContractName"`
}

type TxStruct struct {
	From string `json:"from" bson:"from"`
}

type RpcStruct struct {
	Endpoint   string `json:"endPoint" bson:"endPoint"`
	Passphrase string `json:"passphrase,omitempty" bson:"passphrase"`
}

type Contract struct {
	//ContractAddr string     `json:"contractAddr" bson:"contractAddr"`
	//Method       string     `json:"method" bson:"method"`
	//Interpreter  string     `json:"interpreter" bson:"interpreter"`
	Data DataStruct `json:"data" bson:"data"`
}

type DataStruct struct {
	Name    string `json:"name" bson:"name"`
	Version string `json:"version" bson:"version"`
	Address string `json:"address" bson:"address"`
}
type CnsRequest struct {
	Name    string `json:"name" bson:"name"`
	Version string `json:"version" bson:"version"`
	Address string `json:"address" bson:"address"`
}

type FireWallRequest struct {
	Address string `json:"address" bson:"address"`
}

type DataVal struct {
	Data string
}

type FireWallRequestBody struct {
	Tx  TxStruct  `json:"tx" bson:"tx"`
	Rpc RpcStruct `json:"rpc" bson:"rpc"`
}

type CnsRequestBody struct {
	Tx       TxStruct  `json:"tx" bson:"tx"`
	Rpc      RpcStruct `json:"rpc" bson:"rpc"`
	Contract Contract  `json:"contract" bson:"contract"`
}

type NodeOperateRequest struct {
	Name       string `json:"name" bson:"name"`
	Status     int32  `json:"status" bson:"status"`
	ExternalIP string `json:"externalIP" bson:"externalIP"`
	InternalIP string `json:"internalIP" bson:"internalIP"`
	PublicKey  string `json:"publicKey" bson:"publicKey"`
	P2pPort    uint32 `json:"p2pPort" bson:"p2pPort"`
}
type NodeDataStruct struct {
	Info InfoStruct `json:"info" bson:"info"`
}

type InfoStruct struct {
	Name       string `json:"name" bson:"name"`
	Status     int32  `json:"status" bson:"status"`
	ExternalIP string `json:"externalIP" bson:"externalIP"`
	InternalIP string `json:"internalIP" bson:"internalIP"`

	PublicKey string `json:"publicKey" bson:"publicKey"`
	P2pPort   uint32 `json:"p2pPort" bson:"p2pPort"`
}

type NodeContract struct {
	//ContractAddr string     `json:"contractAddr" bson:"contractAddr"`
	//Method       string     `json:"method" bson:"method"`
	//Interpreter  string     `json:"interpreter" bson:"interpreter"`
	Data NodeDataStruct `json:"data" bson:"data"`
}
type NodeOperateRequestBody struct {
	Tx       TxStruct     `json:"tx" bson:"tx"`
	Rpc      RpcStruct    `json:"rpc" bson:"rpc"`
	Contract NodeContract `json:"contract" bson:"contract"`
}

type UploadFile struct {
	Name     string `json:"name" bson:"name"`
	FilePath string `json:"filepath" bson:"filepath"`
}

type DeployRequestBody struct {
	Tx       TxStruct       `json:"tx" bson:"tx"`
	Rpc      RpcStruct      `json:"rpc" bson:"rpc"`
	Contract DeployContract `json:"contract" bson:"contract"`
}

type DeployContract struct {
	Interpreter string `json:"interpreter" bson:"interpreter"`
}

type DeployRequest struct {
	CodePath    string `json:"codepath" bson:"codepath"`
	AbiPath     string `json:"abipath" bson:"abipath"`
	Interpreter string `json:"interpreter" bson:"interpreter"`
	Name        string `json:"name" bson:"name"`
}

type UserAddress struct {
	Address string `json:"Address" bson:"Address"`
}

type ContractDeploy struct {
	Tx       TxStruct               `json:"tx" bson:"tx"`
	Rpc      RpcStruct              `json:"rpc" bson:"rpc"`
	Contract ContractDeployContract `json:"contract" bson:"contract"`
}

type ContractDeployContract struct {
	Data AddressDataStruct
}

type AddressDataStruct struct {
	Address string `json:"address" bson:"address"`
}

type RoleSetRequest struct {
	Role    string `json:"role" bson:"role"`
	Address string `json:"address" bson:"address"`
}

type ParamRequest struct {
	Tx       TxStruct             `json:"tx" bson:"tx"`
	Rpc      RpcStruct            `json:"rpc" bson:"rpc"`
	Contract ParamRequestContract `json:"contract" bson:"contract"`
}
type ParamRequestContract struct {
	Data SystemConfigBody `json:"data" bson:"data"`
}

type LoginRequest struct {
	Name     string `json:"name" bson:"name"`
	Password string `json:"password" bson:"password"`
}

type InitRequest struct {
	Name       string `json:"name" bson:"name"`
	Password   string `json:"password" bson:"password"`
	Address    string `json:"address"  bson:"address"`
	Passphrase string `json:"passphrase"  bson:"passphrase"`
}

type AddUserRequest struct {
	Name       string `json:"name" bson:"name"`
	Password   string `json:"password" bson:"password"`
	Role       string `json:"role"  bson:"role"`
	Passphrase string `json:"passphrase"  bson:"passphrase"`
}

type DeleteUserRequest struct {
	Name       string `json:"name" bson:"name"`
	Password   string `json:"password" bson:"password"`
	Role       string `json:"role"  bson:"role"`
	Passphrase string `json:"passphrase"  bson:"passphrase"`
}

type NodeOp struct {
	PublicKey string `json:"nodepublickey" bson:"nodepublickey"`
	Groupid   string `json:"groupid" bson:"groupid"`
}

type DeleteNodeOperateRequest struct {
	Name   string `json:"name" bson:"name"`
	Status int32  `json:"status" bson:"status"`
}

type UpdateUserRequest struct {
	Name        string `json:"name" bson:"name"`
	NewPassword string `json:"newpassword" bson:"newpassword"`
	Role        string `json:"role"  bson:"role"`
}
