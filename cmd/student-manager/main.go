// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"flag"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/EngineerProOrg/BE-K01/configs"
// 	"github.com/EngineerProOrg/BE-K01/pkg/controller"
// 	"github.com/EngineerProOrg/BE-K01/pkg/service"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-redis/redis/v8"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var (
// 	confPath = flag.String("conf", "files/live.json", "path to config file")
// 	rd = redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6377",
// 		Password: "",
// 		DB:       0,
// 	})
// )

// type RedisHW interface {
// 	login(c *gin.Context)
// 	ping(c *gin.Context)
// 	rateLimit(c *gin.Context)
// }

// func login(c *gin.Context) {
// 	username := c.PostForm("username")
// 	if len(username) == 0 {
// 		fmt.Println("Enter username")
// 		return
// 	}

// 	key := fmt.Sprintf("u_%s", username)
// 	ctx := context.Background()

// 	err := rd.SetNX(ctx, key, username, time.Duration(0)).Err()

// 	if err != nil {
// 		fmt.Println("Fail %d", key)
// 		return
// 	}

// 	fmt.Println("Save session done !!")
// }

// func main() {
// 	config := &configs.StudentManagerConfig{}
// 	jsonFile, err := os.Open("users.json")
// 	// if we os.Open returns an error then handle it
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer jsonFile.Close()
// 	bt, err := io.ReadAll(jsonFile)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	err = json.Unmarshal(bt, config)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	r := gin.Default()
// 	r.GET("/ping", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
// 	})
// 	db, err := gorm.Open(mysql.New(mysql.Config{
// 		DSN:                       "root:123456@tcp(127.0.0.1:3306)/engineerpro?charset=utf8mb4&parseTime=True&loc=Local",
// 		DefaultStringSize:         256,
// 		DisableDatetimePrecision:  true,
// 		DontSupportRenameIndex:    true,
// 		SkipInitializeWithVersion: false,
// 	}), &gorm.Config{
// 		SkipDefaultTransaction: true,
// 	})
// 	if err != nil {
// 		fmt.Println("can not connect to db ", err)
// 		return
// 	}
	
// 	if rd == nil {
// 		return
// 	}
// 	service := service.NewService(db, rd)
// 	controller.MappingService(r, service)
// 	r.Run()
// }
