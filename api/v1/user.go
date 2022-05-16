package v1

import (
	"blogo/model"
	"blogo/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// AddUser Add
func AddUser(c *gin.Context) {
	var user model.User
	var code int
	_ = c.ShouldBindJSON(&user)
	if len(user.UserName) == 0 && len(user.Password) == 0 {
		code = errmsg.ErrUserFmtWrong
	} else {
		code = model.CheckUser(user.UserName)
		if code == errmsg.SUCCESS {
			model.CreateUser(&user)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   user,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// FindUser Search
func FindUser(c *gin.Context) {

}

// ListUsers List
func ListUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}

	if pageNum == 0 {
		pageNum = -1
	}

	data := model.ListUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// UpdateUser Update
func UpdateUser(c *gin.Context) {
	var user model.User
	var code int
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&user)
	code = model.CheckUser(user.UserName)
	if code == errmsg.SUCCESS {
		model.UpdateUser(id, &user)
	}

	if code == errmsg.ErrUsernameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   user,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// DeleteUser Delete
func DeleteUser(c *gin.Context) {
	var code = errmsg.SUCCESS
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		code = errmsg.ErrUserFmtWrong
	}
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
