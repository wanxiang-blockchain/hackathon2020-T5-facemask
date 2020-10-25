package main

import (
	"data-manager/config"
	middlewares "data-manager/middleware"
	webContext "data-manager/web/context"
	"data-manager/web/controller"
	"data-manager/web/engine"
	gin_session "data-manager/web/session"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	webC := webContext.New()

	router.NoRoute(engine.NewHandler(webC, controller.UserController.NotFound))

	// ===================== Session Middleware  ========================
	gin_session.InitMgr("memory", "127.0.0.1")
	router.Use(middlewares.SessionMiddleware(gin_session.MgrObj))
	// ===================== Cross Domain Middleware  ========================
	router.Use(middlewares.Cors())

	// ===================== init and login function  ========================
	router.GET("/login", engine.NewHandler(webC, controller.UserController.Login))
	router.GET("/init", engine.NewHandler(webC, controller.UserController.SetSuperAdmin))
	v1 := router.Group("superadmin")
	{
		// ===================== User management ========================
		v1.GET("/adduser", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.UserController.AddUser))
		v1.GET("/deleteuser", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.UserController.DeleteUser))
		v1.GET("/update", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.UserController.Update))
		v1.GET("/getusers", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.UserController.GetUser))

		// ===================== Monitor Node management ========================
		v1.GET("/createnode", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.DefaultMonitorController.NewNode))
		v1.GET("/startnode", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.DefaultMonitorController.StartNode))
		v1.GET("/stopnode", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.DefaultMonitorController.StopNode))
		v1.GET("/restartnode", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.DefaultMonitorController.RestartNode))
		//v1.GET("/account", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.GetAccount))

		// ===================== Chain systemConfig management ========================
		v1.GET("/systemconfig/getsystemconfig", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.GetSystemConfig))
		v1.GET("/systemconfig/setsystemconfig", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.SetSystemConfig))

		// ===================== Chain CNS management ========================
		v1.GET("/cns/setcnsnameforaddress", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.SetCnsName))
		v1.GET("/cns/cnsredirect", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.CnsRedirect))

		// ===================== Chain fireWall management (if open, freeze contract)========================
		v1.GET("/firewall/open", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.OpenFireWall))
		v1.GET("/firewall/close", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.CloseFireWall))

		// ===================== Deploy Contract management (need wasm/evm file and abi file)========================
		v1.POST("/deploycontract", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.DeployContract))

		// ===================== Chain Node management ========================
		v1.GET("/node/add", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.AddNode))
		v1.GET("/node/delete", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.DeleteNode))

		// ===================== Chain Role management ========================
		v1.GET("/role/role-lists/roleset", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.RoleSet))

		v1.GET("/getVal", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.GetVal))
		v1.GET("/setVal", middlewares.SuperAdminAuthMiddleWare, engine.NewHandler(webC, controller.PlatonecliController.SetVal))
	}
	v2 := router.Group("user")
	{
		// ===================== Chain CNS management ========================
		v2.GET("/cns/setcnsnameforaddress", middlewares.AuthMiddleware, engine.NewHandler(webC, controller.PlatonecliController.SetCnsName))
		v2.GET("/cns/cnsredirect", middlewares.AuthMiddleware, engine.NewHandler(webC, controller.PlatonecliController.CnsRedirect))

		// ===================== Chain fireWall management (if open, freeze contract)========================
		v2.GET("/firewall/open", middlewares.AuthMiddleware, engine.NewHandler(webC, controller.PlatonecliController.OpenFireWall))
		v2.GET("/firewall/close", middlewares.AuthMiddleware, engine.NewHandler(webC, controller.PlatonecliController.CloseFireWall))

		// ===================== Deploy Contract management (need wasm/evm file and abi file)========================
		v2.POST("/deploycontract", middlewares.AuthMiddleware, engine.NewHandler(webC, controller.PlatonecliController.DeployContract))

		//todo rpc call
		v2.POST("/rpccall", middlewares.AuthMiddleware, engine.NewHandler(webC, controller.DefaultBlockController.DoSomething))
	}
	router.Run(config.Config.HttpConf.Addr())

}
