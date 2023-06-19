package v1

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"webServer/middleware/webjwt"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) { // 注册
	var requestUser = models.Regist{}
	c.Bind(&requestUser) // 前端转入

	userID := 10086 //自动生成并存放到数据库
	registTime := time.Now().Format("2006-01-02 15:04:05")

	name := requestUser.UserName
	telephone := requestUser.PhoneNumber
	password := requestUser.Password

	//数据验证
	fmt.Println(telephone, "手机号码长度", len(telephone))
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": nil,
			"msg":  "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": nil,
			"msg":  "密码不能少于6位",
		})
		return
	}
	if len(name) == 0 {
		var letters = []byte("abcdefghijklmnopqrstuvwxyz") // 随机生成用户名
		result := make([]byte, 10)

		for i := range result {
			result[i] = letters[rand.Intn(len(letters))]
		}
		name = string(result)
	}

	//判断手机号码是否存在
	if models.IsTelephoneExists(telephone) { // 在数据库查找手机号码是否存在
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": nil,
			"msg":  "用户已经存在",
		})
		return
	}

	//密码加密
	// hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 		"code": 500,
	// 		"data": nil,
	// 		"msg":  "加密失败",
	// 	})
	// 	return
	// }

	//把上述的数据存入数据库，从而创建新用户
	models.CreateUser(requestUser, userID, registTime)

	//返回结果
	//发放token
	token, err := webjwt.ReleaseToken(requestUser.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 500,
			"data": nil,
			"msg":  "系统异常",
		})
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": token,
		"msg":  "注册成功",
	})
}

func Login(c *gin.Context) {
	var requestUser = models.Regist{}
	c.Bind(&requestUser)

	//获取参数
	telephone := requestUser.PhoneNumber
	password := requestUser.Password
	//数据验证
	fmt.Println(telephone, "手机号码长度", len(telephone))
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": nil,
			"msg":  "手机号必须为11位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"data": nil,
			"msg":  "密码不能少于6位",
		})
		return
	}

	//判断手机号码是否存在
	if !models.IsTelephoneExists(telephone) { // 在数据库查找手机号码是否存在
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "用户不存在",
		})
		return
	}

	// //判断密码是否正确
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"code": 400,
	// 		"data": nil,
	// 		"msg":  "密码错误",
	// 	})
	// 	return
	// }

	if !models.SecretCorrect(telephone, password) { // 在数据库查找手机号码是否存在
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 400,
			"data": nil,
			"msg":  "密码错误", // 前端提示信息
		})
		return
	}

	//发放token
	token, err := webjwt.ReleaseToken(requestUser.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 500,
			"data": nil,
			"msg":  "系统异常",
		})
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": token,
		"msg":  "登录成功",
	})
}
