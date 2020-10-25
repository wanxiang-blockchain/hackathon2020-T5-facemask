[TOC]

# 区块链浏览器接口文档

## 用户管理

### 系统管理员初始化（只能调用一次）

* DESCRIPTION：初始化系统管理员的账号和密码（只能调用一次）。

* URI： /init

* METHOD：GET

* INPUT：

  * Name（string）[必须]：注册姓名，必须唯一。
  * Password（string）[必须]：登陆密码。
  * Address (string) [可选]：初始化管理员时为管理员配置链最高管理员权限对应的账号
  * Passphrase (string) [可选]：初始化管理员时为管理员配置链最高管理员权限账号对应的密码

* OUTPUT：

  ```json 

  {
  "code":200,
  "msg":"注册成功",
  "data":null
  }

  ```
  
* EXAMPLE：

  ```shell
  curl http://localhost:9999/init?Name=admin&Password=admin&Address=0x8d4d2Ed9cA6c6279BaB46Be1624cF7ADbAB89E18&Passphrase=0
  ```

### 用户登陆

* DESCRIPTION：用户登陆。

* URI： /login  

* METHOD：GET

* INPUT：

  * Name（string）[必须]：注册姓名，必须唯一。
  * Password（string）[必须]：登陆密码。

* OUTPUT：

  ```json
  成功:
{
  "code": 200,
  "msg": "登陆成功",
  "data": [
    "2020-09-16 11:23:39",
    "172.25.1.211",
    "SUPERADMIN"
  ]
}
  失败:
   {
   "code":400,
   "msg":"用户名不存在或密码错误，请重新登陆！",
   "data":
          {
          "code":10007,
          "msg":"password is incorrect"}
          }
   }
  ```

* EXAMPLE：

  ```shell
  curl http://localhost:9999/login?Name=admin&Password=admin
  ```
  
### 添加用户

* DESCRIPTION：超级管理员添加用户并赋予用户角色。

* URI： /superadmin/adduser

* METHOD：GET

* INPUT：

  * Name（string）[必须]：注册姓名，必须唯一。
  * Password（string）[必须]：登陆密码。
  * Role(string)[必须]：用户角色(只能有SUPERADMIN和USER两种角色)
  * Passphrase（string）[必须]：生成keyfile账户的密码。

* OUTPUT：

 ```json
  成功:
     {
     "code":200,
     "msg":"登陆成功",
     "data":"SUPERADMIN"
    
     }
  失败:
     {
     "code":400,
     "msg":"增加用户失败！",
     "data":
           {
           "code":10005,
           "msg":"bad role"
           }
    }
```

* EXAMPLE：

  ```shell
  curl http://localhost:9999/superadmin/adduser?Name=wxbc1&Password=123&Role=SUPERADMIN&Passphrase=0
  ```
  
### 获取全部用户

* DESCRIPTION：超级管理员获取全部用户信息。

* URI： /superadmin/getusers

* METHOD：GET

* INPUT：

  * PageIndex（int64）[必须]：页码。
  * PageSize（string）[必须]：每页数量。

* OUTPUT：

```
成功：
{
  "code": 200,
  "msg": "添加用户成功",
  "data": [
    {
      "id": "5f60933f026420c6fbb52393",
      "name": "admin",
      "password": "密码不可见",
      "address": "123",
      "publicKey": "",
      "privateKey": "",
      "role": "SUPERADMIN",
      "registertime": "2020-09-15"
    }
  ]
}

失败：
{
  "code": 400,
  "msg": "未设置页码或每页大小！",
  "data": {
    "code": 10001,
    "msg": "parameter invalid error"
  }
}
```

* EXAMPLE：

  ```shell
  curl http://localhost:9999/superadmin/getusers?PageIndex=1&PageSize=100
  ```

### 更新用户信息

* DESCRIPTION：超级管理员更新用户信息（用户角色以及密码等）。

* URI： /superadmin/update

* METHOD：GET

* INPUT：

  * Name（string）[必须]：需要更新信息的用户的姓名，必须唯一。
  * NewPassword（string）[必须]：更新后的密码。
  * Role（string）[必须]：更新后的角色

* OUTPUT：

```
{
  "code": 200,
  "msg": "更新成功",
  "data": null
}

失败：
{
  "code": 400,
  "msg": "更新失败！",
  "data": {
    "code": 10005,
    "msg": "bad role"
  }
}
```

* EXAMPLE：

  ```shell
  curl localhost:9999/superadmin/update?Name=wxbc&NewPassword=1234&Role=SUPERADMIN
  ```


### 删除用户

* DESCRIPTION：超级管理员删除用户。

* URI： /superadmin/deleteuser

* METHOD：GET

* INPUT：

  * Name（string）[必须]：删除用户的姓名，必须唯一。

* OUTPUT：

```
成功：
{
  "code": 200,
  "msg": "删除用户成功",
  "data": null
}

失败：
{
  "code": 400,
  "msg": "删除用户失败！",
  "data": {
    "code": 10006,
    "msg": "user not exist"
  }
}
```

* EXAMPLE：

  ```shell
  curl http://localhost:9999/superadmin/deleteuser?Name=wxbc
  ```

## monitor操作管理

### 新建节点

* DESCRIPTION：在monitor服务器上新建一个节点。

* URI： /superadmin/createnode

* METHOD：GET

* INPUT：
 * MonitorIp（string）[必须]：monitor服务器的ip，必须唯一。
 * MonitorPort(string)[必须]：monitor服务器的端口号，必须唯一。
 * Groupid(string)[非必须]：节点的groupid。
 * Chainid(string)[非必须]：       
 * NodeIp(string)[非必须]：        
 * NodePort(string)[非必须]：      
 * Rpcport(string)[非必须]：      
 * Wsport(string)[非必须]：        
 * Dashport(string)[非必须]：     
 * Password(string)[非必须]：      
 * CreatorEnode(string)[非必须]：  
 * Bootnodes(string)[非必须]：     

* OUTPUT：
```
成功：
{
  "code": 200,
  "msg": "请求成功",
  "data": "e050a96a4d28e38fa8211c052013836fe3daf0482e2719acc195caa558acc6874e1b520269f0e7b138f04cb5c10c062b2a36c5c937fbc89bd1d180758ad6cd85"
}
失败：
{
  "code": 400,
  "msg": "创建节点失败！",
  "data": {
    "code": 10008,
    "msg": "node registered"
  }
}


```
* EXAMPLE：

  ```shell
  curl http://http://localhost:9999/superadmin/createnode?MonitorIp=localhost&MonitorPort=50051
  ```

### 启动节点

* DESCRIPTION：区块链节点启动。

* URI： /superadmin/startnode

* METHOD：GET

* INPUT：
 * PublicKey（string）[必须]：节点对应的公钥，必须唯一。
 * Groupid(string)[非必须]：节点的groupid。

* OUTPUT：

  ```json
成功：
{
  "code": 200,
  "msg": "启动节点成功！",
  "data": null
}
失败：
{
  "code": 400,
  "msg": "启动节点失败！",
  "data": {}
}
  ```

* EXAMPLE：

  ```shell
  curl http://localhost:9999/superadmin/startnode?PublicKey=e050a96a4d28e38fa8211c052013836fe3daf0482e2719acc195caa558acc6874e1b520269f0e7b138f04cb5c10c062b2a36c5c937fbc89bd1d180758ad6cd85&Groupid=0
  ```

### 关闭节点

* DESCRIPTION：区块链节点关闭。

* URI： /superadmin/stopnode

* METHOD：GET

* INPUT：
 * PublicKey（string）[必须]：节点对应的公钥，必须唯一。
 * Groupid (string)[非必须]：节点的groupid。

* OUTPUT：

  ```json
成功：
{
  "code": 200,
  "msg": "关闭节点成功！",
  "data": null
}
失败：
{
  "code": 400,
  "msg": "启动节点失败！",
  "data": {}
}
  ```

* EXAMPLE：

  ```shell
  curl http://localhost:9999/superadmin/stopnode?PublicKey=e050a96a4d28e38fa8211c052013836fe3daf0482e2719acc195caa558acc6874e1b520269f0e7b138f04cb5c10c062b2a36c5c937fbc89bd1d180758ad6cd85&Groupid=0
  ```

### 重启节点

* DESCRIPTION：区块链节点重启。

* URI： /superadmin/restartnode

* METHOD：GET

* INPUT：
 * PublicKey（string）[必须]：节点对应的公钥，必须唯一。
 * Groupid(string)[非必须]：节点的groupid。

* OUTPUT：

```json
成功：
{
  "code": 200,
  "msg": "重启节点成功！",
  "data": null
}
失败：
{
  "code": 400,
  "msg": "重启节点失败！",
  "data": {}
}
```

* EXAMPLE：

  ```shell
  curl http://localhost:9999/superadmin/restartnode?PublicKey=e050a96a4d28e38fa8211c052013836fe3daf0482e2719acc195caa558acc6874e1b520269f0e7b138f04cb5c10c062b2a36c5c937fbc89bd1d180758ad6cd85&Groupid=0
  ```

## 用户对链操作

### 获取系统参数

* DESCRIPTION：用户获取系统参数。

* URI： /systemconfig/getsystemconfig

* METHOD：GET

* INPUT：
 * 无

* OUTPUT：

```json
{
  "code": 200,
  "msg": "获取参数成功",
  "data": "{\"blockGasLimit\":\"10000000000\",\"txGasLimit\":\"1500000000\",\"isUseGas\":\"0\",\"isApproveDploeyedContract\":\"0\",\"isCheckDeployPermission\":\"\",\"isProduceEmptyBlock\":\"0\",\"gasContractName\":\"\\\"\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000\\\\u0000 \\\"\"}"
}
```

* EXAMPLE：

  ```shell
  curl http://localhost:9999/superadmin/systemconfig/getsystemconfig
  ```

### 设置系统参数

* DESCRIPTION：用户设置系统参数。

* URI： /systemconfig/setsystemconfig

* METHOD：GET

* INPUT：
 * BlockGasLimit  (string)[非必须]          
 * TxGasLimit (string)[非必须]              
 * IsUseGas   (string)[非必须]         
 * IsApproveDeployedContract (string)[非必须]
 * IsCheckDeployPermission (string)[非必须]  
 * IsProduceEmptyBlock  (string)[非必须]     
 * GasContractName   (string)[非必须]        
* OUTPUT：

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": "success"
}
```

* EXAMPLE：

  ```shell
  curl http://localhost:9999/superadmin/systemconfig/setsystemconfig?BlockGasLimit=2000000000&IsUseGas=1
  ```

### 注册CNS

* DESCRIPTION：用户为某合约注册CNS。

* URI： /cns/setcnsnameforaddress

* METHOD：GET

* INPUT：
 * Name (string)[必选]: CNS合约名称
 * Version (string)[必选]: CNS合约版本号
 * Address (string)[必选]: CNS合约地址


* OUTPUT：

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": "{\"status\":\"Operation Succeeded\",\"logs\":[\"Event [CNS] Notify: 2 [CNS] not owner of registered contract \"],\"blockNumber\":24}"
}
```

* EXAMPLE：

  ```shell
  curl localhost:9999/superadmin/cns/setcnsnameforaddress?Name=tofu&Version=0.0.0.1&Address=0xc89b196d634c564388d15d64ac87e95974a744e3
  ```

  ### CNS重定向

* DESCRIPTION：用户为某合约CNS进行重定向。

* URI： cns/cnsredirect

* METHOD：GET

* INPUT：
 * Name (string)[必选]: CNS合约名称
 * Version (string)[必选]: CNS合约版本号
* OUTPUT：

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": "{\"status\":\"Operation Succeeded\",\"logs\":[\"Event [CNS] Notify: 4 [CNS] Name or version didn't register before \"],\"blockNumber\":25}"
}
```

* EXAMPLE：

  ```shell
  curl localhost:9999/superadmin/cns/cnsredirect?Name=tofu&Version=0.0.0.1
  ```

  ### 开启合约防火墙

* DESCRIPTION：用户为某合约开启合约防火墙。

* URI： /firewall/open

* METHOD：GET

* INPUT：
 * Address (string)[必选]: 合约地址
* OUTPUT：

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": "{\"status\":\"Operation Succeeded\",\"logs\":[\"Event Notify: 1 FW : error, only contract owner can set firewall setting \"],\"blockNumber\":26}"
}
```

* EXAMPLE：

  ```shell
  curl localhost:9999/superadmin/firewall/open?Address=0xc89b196d634c564388d15d64ac87e95974a744e3
  ```

### 关闭合约防火墙

* DESCRIPTION：用户为某合约关闭合约防火墙。

* URI： /firewall/close

* METHOD：GET

* INPUT：
 * Address (string)[必选]: 合约地址
* OUTPUT：

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": "{\"status\":\"Operation Succeeded\",\"logs\":[\"Event Notify: 1 FW : error, only contract owner can set firewall setting \"],\"blockNumber\":27}"
}
```

* EXAMPLE：

  ```shell
  curl localhost:9999/superadmin/firewall/close?Address=0xc89b196d634c564388d15d64ac87e95974a744e3
  ```

### 添加节点

* DESCRIPTION：添加新节点。

* URI： /node/add

* METHOD：GET

* INPUT：
 * Address (string)[必选]: 节点地址
 * Endpoint (string)[可选]：节点终端号
 * Passphrase (string)[可选]： 节点密码
 * Name  (string)[必选]： 节点名称
 * Status  (unt32)[必选]： 节点状态
 * External IP (string)[必选]： 节点External Ip
 * Internal IP (string)[必选]： 节点Internal Ip
 * PublicKey (string)[必选]：节点公钥
 * P2pPort (string)[必选]：节点p2p端口号
* OUTPUT：

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": "{\"status\":\"Operation Succeeded\",\"logs\":[\"Event Notify: 2 0xe3B73cfb3C978bfB1aB85a939eAc4fCfEfabc837 has no permission to add node info. \"],\"blockNumber\":28}"
}
```

* EXAMPLE：

  ```shell
  curl localhost:9999/superadmin/node/add?Name=node0&Status=1&External IP=127.0.0.1&PublicKey=64a684197dbc77b69f418c511e55adb7a4a532a88d25d8e9d34667141d53790b5ff84ed385e35ade60ea9e610b3ac54499119fd1a9bf1344d319aeceadcb5bb7&P2pPort=8888
  ```

### 删除节点

* DESCRIPTION：删除节点。

* URI： /node/delete

* METHOD：GET

* INPUT：
 * Sender (string)[必选]: 用户请求地址
 * Endpoint (string)[可选]：节点终端号
 * Passphrase (string)[可选]： 节点密码
 * Name  (string)[必选]： 节点名称
 * Status  (unt32)[必选]： 节点状态

* OUTPUT：

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": "{\"status\":\"Operation Succeeded\",\"logs\":[\"Event Notify: 2 0xe3B73cfb3C978bfB1aB85a939eAc4fCfEfabc837 no permission update node. \"],\"blockNumber\":29}"
}
```

* EXAMPLE：

  ```shell
  curl localhost:9999/superadmin/node/delete?Name=node0&Status=2
  ```


  	//wasmpath := "/Users/night/dev/go/src/github.com/PlatONEnetwork/PlatONE-Go/cmd/platonecli/test/test_case/wasm/appDemo.wasm"
	//abipath := "/Users/night/dev/go/src/github.com/PlatONEnetwork/PlatONE-Go/cmd/platonecli/test/test_case/wasm/appDemo.cpp.abi.json"
http://localhost:9999/superadmin/deploycontract?Interpreter=wasm&wasmpath := "/Users/night/dev/go/src/github.com/PlatONEnetwork/PlatONE-Go/cmd/platonecli/test/test_case/wasm/appDemo.wasm&abipath := /Users/night/dev/go/src/github.com/PlatONEnetwork/PlatONE-Go/cmd/platonecli/test/test_case/wasm/appDemo.cpp.abi.json

### 部署合约

* DESCRIPTION：部署合约。

* URI： /superadmin/deploycontract

* METHOD：GET

* INPUT：
 * CodePath (string)[必选]: 合约文件路径
 * AbiPath (string)[可选]：合约abi文件路径
 * Interpreter (string)[可选]： 合约类型


* OUTPUT：

```json

  "code": 200,
  "msg": "操作成功",
  "data": "\"{\\n\\t\\\"status\\\": \\\"Operation Succeeded\\\",\\n\\t\\\"contractAddress\\\": \\\"0x322fbc1a36a8ebd6e38ee229d47e6f3a56ef1609\\\",\\n\\t\\\"blockNumber\\\": 57,\\n\\t\\\"GasUsed\\\": 2689439,\\n\\t\\\"From\\\": \\\"0x8d4d2ed9ca6c6279bab46be1624cf7adbab89e18\\\",\\n\\t\\\"To\\\": \\\"\\\",\\n\\t\\\"TxHash\\\": \\\"\\\"\\n}\""
}
```

* EXAMPLE：

  ```shell
  curl http://localhost:9999/superadmin/deploycontract?Interpreter=wasm&CodePath=/Users/night/dev/go/src/github.com/PlatONEnetwork/PlatONE-Go/cmd/platonecli/test/test_case/wasm/appDemo.wasm&AbiPath=/Users/night/dev/go/src/github.com/PlatONEnetwork/PlatONE-Go/cmd/platonecli/test/test_case/wasm/appDemo.cpp.abi.json
  ```

