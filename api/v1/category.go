package v1

import (
	"blogo/model"
	"blogo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddCate(c *gin.Context) {
	var cate model.Category
	var code int
	_ = c.ShouldBindJSON(&cate)
	if len(cate.Name) == 0 {
		code = errmsg.ErrCateFmtWrong
	} else {
		code = model.CheckCate(cate.Name)
		if code == errmsg.SUCCESS {
			model.CreateCate(&cate)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   cate,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// ListCate List
func ListCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data := model.ListCate(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// UpdateCate Update
func UpdateCate(c *gin.Context) {
	var cate model.Category
	var code int
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&cate)
	code = model.CheckCate(cate.Name)
	if code == errmsg.SUCCESS {
		model.UpdateCate(id, &cate)
	}

	if code == errmsg.ErrCateNameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   cate,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// DeleteCate Delete
func DeleteCate(c *gin.Context) {
	var code = errmsg.SUCCESS
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		code = errmsg.ErrCateFmtWrong
	}
	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
