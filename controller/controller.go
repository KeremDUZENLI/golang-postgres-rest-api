package controller

import (
	"net/http"
	"strconv"

	"validation/model"

	"github.com/gin-gonic/gin"
)

func BeginDatabase(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.BeginDatabaseWithInfos())
}

func CreateDatabase(c *gin.Context) {
	aStruct := &model.MyStruct{}

	c.ShouldBindJSON(aStruct)
	c.IndentedJSON(http.StatusOK, aStruct.CreateDatabaseInfo())
}

func ReadDatabase(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, model.ReadDatabaseInfos())
}

func UpdateDatabase(c *gin.Context) {
	nStr := c.Param("id")
	n, _ := strconv.Atoi(nStr)

	aStruct := &model.MyStruct{}

	c.ShouldBindJSON(aStruct)

	for _, item := range *model.ReadDatabaseInfos() {
		if item.Id == n {
			c.IndentedJSON(http.StatusOK, item)
			model.DeleteDatabaseInfo(n)
			break
		}
	}

	aStruct.UpdateDatabaseInfo()
	c.IndentedJSON(http.StatusOK, aStruct)
}

func DeleteDatabase(c *gin.Context) {
	nStr := c.Param("id")
	n, _ := strconv.Atoi(nStr)

	for _, item := range *model.ReadDatabaseInfos() {
		if item.Id == n {
			c.IndentedJSON(http.StatusOK, item)
			model.DeleteDatabaseInfo(n)
			break
		}
	}
}
