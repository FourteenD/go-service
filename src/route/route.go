package route

import "github.com/gin-gonic/gin"

func Init()  {
	r := gin.Default()
	Load(r)
	r.Run()
}

func Load(r *gin.Engine)  {
	v1 := r.Group("/v1")
	v1.GET("/home",v1.Handlers...)
}
