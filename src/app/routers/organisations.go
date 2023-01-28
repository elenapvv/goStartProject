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

type orgSave struct {
	Name string `json:"name"`
}

type orgFull struct {
	OrgID  uint64 `json:"org_id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

type orgDel struct {
	OrgID uint64 `json:"org_id"`
}

func saveOrg(c *gin.Context) {
	var data orgSave

	if err := c.BindJSON(&data); err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	} else {
		resp, err := utils.Conn.Insert("organisation", []interface{}{nil, data.Name, true})

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
			log.ErrorLogger.Println("Почему-то не получили данные при сохранении организациии в БД")
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"data": orgFull{OrgID: newData[0][0].(uint64), Name: newData[0][1].(string), Status: newData[0][2].(bool)},
		})
	}
}

func getOrgList(c *gin.Context) {
	spaceName := "organisation"
	indexName := "primary" // "scanner"
	idFn := utils.Conn.Schema.Spaces[spaceName].Fields["id"].Id

	var tuplesPerRequest uint32 = 2
	cursor := []interface{}{}
	var orgData []orgFull

	for {
		resp, err := utils.Conn.Select(spaceName, indexName, 0, tuplesPerRequest, tarantool.IterGt, cursor)
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

		if len(resp.Data) == 0 {
			break
		}

		tuples := resp.Tuples()
		for _, tuple := range tuples {
			orgData = append(orgData, orgFull{OrgID: tuple[0].(uint64), Name: tuple[1].(string), Status: tuple[2].(bool)})
		}

		lastTuple := tuples[len(tuples)-1]
		cursor = []interface{}{lastTuple[idFn]}
	}

	if len(orgData) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "organisations not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": orgData})
}

func getOrgByID(c *gin.Context) {
	orgId := c.Query("org_id")
	ui64, err := strconv.ParseUint(orgId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	resp, err := utils.Conn.Select("organisation", "primary", 0, 1, tarantool.IterEq, []interface{}{ui64})
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
		errMsg := fmt.Sprintf("organisation not found with id=%v", orgId)
		log.ErrorLogger.Println(errMsg)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": errMsg})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"data": orgFull{OrgID: newData[0][0].(uint64), Name: newData[0][1].(string), Status: newData[0][2].(bool)},
	})
}

func deleteOrg(c *gin.Context) {
	var data orgDel

	if err := c.BindJSON(&data); err != nil {
		log.ErrorLogger.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
		return
	} else {
		resp, err := utils.Conn.Update("organisation", "primary", []interface{}{data.OrgID}, []interface{}{[]interface{}{"=", 2, false}})

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
			errMsg := fmt.Sprintf("organisation not found with id=%v", data.OrgID)
			log.ErrorLogger.Println(errMsg)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": errMsg})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"data": orgFull{OrgID: newData[0][0].(uint64), Name: newData[0][1].(string), Status: newData[0][2].(bool)},
		})
	}
}
