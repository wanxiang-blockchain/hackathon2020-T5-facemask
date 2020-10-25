package session

import (
	"github.com/sirupsen/logrus"
)

const (
	SessionCookieName  = "session_id" // session_id在Cookie中对应的key
	SessionContextName = "session"    // session data在gin上下文中对应的key
)

//定义一个全局的Mgr
var (
	MgrObj Mgr
)

//构造一个Mgr
func InitMgr(name string, addr string, option ...string) {

	switch name {
	case "memory": //初始化一个内存版管理者
		MgrObj = NewMemory()
		//case "redis":
		//	MgrObj=NewRedisMgr()
	}
	MgrObj.Init(addr, option...) //初始化mgr
	logrus.Info("memory init success.")

}

//
type SessionData interface {
	GetID() string // 返回自己的ID
	Get(key string) (value interface{}, err error)
	Set(key string, value interface{})
	Del(key string)
	Save() // 保存
}

//不同版本的管理者接口
type Mgr interface {
	Init(addr string, option ...string)
	GetSessionData(sessionId string) (sd SessionData, err error)
	CreatSession() (sessionId string, sd SessionData)
}

//gin框架中间件
//func SessionMiddleware(mgrObj Mgr)gin.HandlerFunc{
//
//	return func(c *gin.Context){
//
//		//1.请求刚过来，从请求的cookie中获取SessionId
//		SessionID,err:=c.Cookie(SessionCookieName)
//		var sd SessionData
//		if err != nil {
//			//1.1 第一次来，没有sessionid，-->给用户建一个sessiondata，分配一个sessionid
//			sd=mgrObj.CreatSession()
//		}else {
//			//1.2  取到sessionid
//			//2. 根据sessionid去大仓库取sessiondata
//			sd,err=mgrObj.GetSessionData(SessionID)
//			if err != nil {
//				//sessionid有误，取不到sessiondata，可能是自己伪造的
//				//重新创建一个sessiondata
//				sd=mgrObj.CreatSession()
//				//更新sessionid
//				SessionID=sd.GetID()//这个sessionid用于回写coookie
//			}
//		}
//		//3. 如何实现让后续所有请求的方法都拿到sessiondata？ 让每个用户的dessiondata都不同
//		//3.利用gin框架的c.Set("session",sessiondata)
//		c.Set(SessionContextName,sd)
//		//回写cookie
//		c.SetCookie(SessionCookieName,SessionID,3600,"/","127.0.0.1",false,true)
//		c.Next()
//	}
//}
