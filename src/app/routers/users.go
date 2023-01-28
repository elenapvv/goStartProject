package routers

import (
	log "app/logging"
	"app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tarantool/go-tarantool"
	"net/http"
	"strconv"
)

type userCreate struct {
	Name  string `json:"name"`
	OrgID uint64 `json:"org_id"`
}

type userFull struct {
	UserID uint64 `json:"user_id"`
	Name   string `json:"name"`
	OrgID  uint64 `json:"org_id"`
}

type userDel struct {
	UserID uint64 `json:"user_id"`
}

func createUser(c *gin.Context) {
	var data userCreate

	if err := c.BindJSON(&data); err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	} else {
		resp, err := utils.Conn.Insert("user", []interface{}{nil, data.Name, data.OrgID})

		if err != nil {
			log.ErrorLogger.Println("Error", err)
			log.ErrorLogger.Println("Code", resp.Code)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		newData := resp.Tuples()

		if len(newData) == 0 {
			log.ErrorLogger.Println("Почему-то не получили данные при сохранении юзера в БД")
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"data": userFull{UserID: newData[0][0].(uint64), Name: newData[0][1].(string), OrgID: newData[0][2].(uint64)},
		})
	}

}

func getUserInfo(c *gin.Context) {
	userId := c.Query("user_id")

	ui64, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	resp, err := utils.Conn.Select("user", "primary", 0, 1, tarantool.IterEq, []interface{}{ui64})
	if err != nil {
		log.ErrorLogger.Println("Failed to select: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	if resp.Code != tarantool.OkCode {
		log.ErrorLogger.Println("Select failed: %s", resp.Error)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	newData := resp.Tuples()

	if len(newData) == 0 {
		errMsg := fmt.Sprintf("user not found with id=%v", userId)
		log.ErrorLogger.Println(errMsg)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": errMsg})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": userFull{UserID: newData[0][0].(uint64), Name: newData[0][1].(string), OrgID: newData[0][2].(uint64)},
	})
}

func deleteUser(c *gin.Context) {
	var data userDel

	if err := c.BindJSON(&data); err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	} else {
		resp, err := utils.Conn.Delete("user", "primary", []interface{}{data.UserID})

		if err != nil {
			log.ErrorLogger.Println("Error", err)
			log.ErrorLogger.Println("Code", resp.Code)
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		newData := resp.Tuples()

		if len(newData) == 0 {
			errMsg := fmt.Sprintf("user not found with id=%v", data.UserID)
			log.ErrorLogger.Println(errMsg)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": errMsg})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"data": userFull{UserID: newData[0][0].(uint64), Name: newData[0][1].(string), OrgID: newData[0][2].(uint64)},
		})
	}
}

func changeUser(c *gin.Context) {
	var data userFull

	if err := c.BindJSON(&data); err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	} else {
		resp, err := utils.Conn.Replace("user", []interface{}{data.UserID, data.Name, data.OrgID})

		if err != nil {
			log.ErrorLogger.Println("Error", err)
			log.ErrorLogger.Println("Code", resp.Code)
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		newData := resp.Tuples()

		if len(newData) == 0 {
			errMsg := fmt.Sprintf("user not found with id=%v", data.UserID)
			log.ErrorLogger.Println(errMsg)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": errMsg})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"data": userFull{UserID: newData[0][0].(uint64), Name: newData[0][1].(string), OrgID: newData[0][2].(uint64)},
		})
	}
}
