package routers

type fileIn struct {
	//Content []byte `json:"content"`
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
}

type fileDel struct {
	UUID int `json:"uuid"`
}

//func saveFile(c *gin.Context) {
//	id := uuid.New()
//	externalUuid := uuid.New()
//	///////////
//
//	file, err := c.FormFile("file")
//	if err != nil {
//		log.ErrorLogger.Println(err)
//	}
//
//	err = c.SaveUploadedFile(file, "files/"+file.Filename)
//	if err != nil {
//		log.ErrorLogger.Println(err)
//	}
//	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
//	///////////////
//
//	var data fileIn
//	if err := c.BindJSON(&data); err != nil {
//		log.ErrorLogger.Println(err)
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": fmt.Sprintf("%v", err),
//		})
//		return
//	} else {
//		resp, err := utils.Conn.Insert("file", []interface{}{id, data.Name, location, data.UserID, externalUuid})
//
//		if err != nil {
//			log.ErrorLogger.Println("Error", err)
//			log.ErrorLogger.Println("Code", resp.Code)
//			c.IndentedJSON(http.StatusInternalServerError, gin.H{
//				"error": fmt.Sprintf("%v", err),
//			})
//			return
//		}
//
//		newData := resp.Tuples()
//
//		if len(newData) == 0 {
//			log.ErrorLogger.Println("Почему-то не получили данные при сохранении юзера в БД")
//			c.IndentedJSON(http.StatusInternalServerError, gin.H{
//				"error": fmt.Sprintf("%v", err),
//			})
//			return
//		}
//
//		c.IndentedJSON(http.StatusOK, gin.H{
//			"data": userFull{UserID: newData[0][0].(uint64), Name: newData[0][1].(string), OrgID: newData[0][2].(uint64)},
//		})
//	}
//}
//
//func getFile(c *gin.Context) {
//	uuid := c.Query("uuid")
//
//	// Loop over the list of albums, looking for
//	// an album whose ID value matches the parameter.
//	for _, a := range albums {
//		if a.ID == uuid {
//			c.IndentedJSON(http.StatusOK, a)
//			return
//		}
//	}
//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
//}
//
//func deleteFile(c *gin.Context) {
//	var data fileDel
//	if err := c.BindJSON(&data); err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusBadRequest, gin.H{
//			"error": fmt.Sprintf("%v", err),
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"data": data,
//		})
//	}
//}
