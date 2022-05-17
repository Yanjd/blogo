package v1

import (
	"blogo/model"
	"blogo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	art, code := model.GetArtInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   art,
		"msg":    errmsg.GetErrMsg(code),
	})
}

func GetArtForCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	arts, code := model.GetArtForCate(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   arts,
		"msg":    errmsg.GetErrMsg(code),
	})
}

func AddArt(c *gin.Context) {
	var art model.Article
	var code int
	_ = c.ShouldBindJSON(&art)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   art,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// ListArts List
func ListArts(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data, code := model.ListArts(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// UpdateArt Update
func UpdateArt(c *gin.Context) {
	var art model.Article
	var code int
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&art)
	model.UpdateArt(id, &art)

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   art,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// DeleteArt Delete
func DeleteArt(c *gin.Context) {
	var code = errmsg.SUCCESS
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
