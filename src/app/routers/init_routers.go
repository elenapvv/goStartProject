package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() {
	router := gin.Default()

	// для организаций
	router.POST("/org", saveOrg)
	router.GET("/org_list", getOrgList)
	router.GET("/org_by_org_id", getOrgByID)
	router.DELETE("/org", deleteOrg)

	// для юзеров
	router.POST("/create_user", createUser)
	router.GET("/user_info", getUserInfo)
	router.DELETE("/user", deleteUser)
	router.PATCH("/user", changeUser)
	//
	//// для файлов
	//router.POST("/file", saveFile)
	//router.GET("/file", getFile)
	//router.DELETE("/file", deleteFile)

	router.Run(":8080")
}
