package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	conn := ConnectDB()
	var circleServer crinterface.ServerInterface = NewDbCircle(conn, NewCircleCache())
	if initCircle, ok := circleServer.(Crinterface.CircleInitInterface); ok {
		if err := initCircle.Init(); err != nil {
			log.Fatalln("初始化不成功", err)
		}
	}
	re := gin.Default()
	pprof.Register(re)

	re.POST("/post", func(c *gin.Context) {
		var cr *apis.Circle
		if err := c.BindJSON(&cr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"错误信息": "状态无法发布" + err.Error(),
			})
			return
		}
		cr.Timestamp = time.Now().Unix()
		cr.Visible = true

		if err := circleServer.PostStatus(cr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"错误信息": "状态显示失败",
			})
			return
		}
		c.JSON(http.StatusOk, gin.H{
			"信息	": "成功",
		})
	})

	r.GET("/list", func(*gin.Context) {
		if list, err := circleServer.ListPost(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"错误信息": "未能更新",
			})
			return

		} else {
			c.JSON(http.StatusOk, list)
		}
	})

	r.DELETE("/delete/personalid", func(
		c *gin.Context) {
		idInString := c.Param("personnalid")
		id, err := strconv.Atoi(idInString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"错误信息": "无效id",
			})
			return
		}
		if err := circleServer.DeletePost(uint64(id)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"错误信息": "无法删除个人id",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"信息": "成功删除信息",
			})
		}
	})
	if err := r.Run(":8081"); err != nil {
		log.Fatalln(err)
	}

}

func ConnectDB() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:87654321@tcp(127.0.0.1.3306") / testdb)
	if err != nil {
		log.Fatalln("链接数据库失败:", err)
	}
	fmt.Println("链接数据库成功")
	return conn

}
