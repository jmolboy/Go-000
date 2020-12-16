package routes

import (
	"github.com/gin-gonic/gin"
	"jmolboy/week04/internal/routes/user"
	"strconv"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	gin.DisableConsoleColor()
	//setmode:debug,test,release
	gin.SetMode(gin.ReleaseMode)

	g := gin.Default()

	//g.Use(gin.Logger())
	g.Use(gin.Recovery())
	//管理后台路由组
	admGrp(g)
	return g
}

// 管理后台路由
func admGrp(g *gin.Engine) {
	admRoute := "admi"

	//不需要登录路由
	authGrp := g.Group(admRoute)
	authGrp.GET("/user/list", user.List)
}


func Start() {
	g := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	port := 80
	g.Run(":" + strconv.Itoa(port))
}
