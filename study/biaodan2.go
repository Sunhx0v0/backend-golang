// code from chatgpt

/*
在上面的代码中，我们定义了一个名为User的结构体，它包含了四个字段：Name、Age、Email和Password。这些字段都通过form标签指定了它们对应的表单字段名，并使用了binding标签指定了它们的验证规则。

在路由处理函数中，我们使用了ShouldBindWith方法来将表单数据绑定到User结构体中。如果绑定失败，我们会返回一个错误响应。如果绑定成功，我们就可以使用User结构体中的数据来做一些有用的事情了。

这个示例代码使用了gin框架的默认路由和中间件，同时也使用了gin的表单绑定和验证功能。这个代码可以用于处理用户注册表单的数据验证，你可以根据自己的需求进行修改和扩展。
*/

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type User struct {
	Name     string `form:"name" binding:"required"`
	Age      int    `form:"age" binding:"required,min=18,max=60"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6,max=20"`
}

func main() {
	router := gin.Default()

	router.POST("/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindWith(&user, binding.FormPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Do something with valid user data
		user.Age = 19
		user.Name = "sada"
		user.Email = "1465436241@qq.com"
		user.Password = "123456"
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	router.Run(":8080")
}
