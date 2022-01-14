package main

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"stockexchange/rpc/demo/global"
	"stockexchange/rpc/demo/model"
	"stockexchange/rpc/demo/proto/user"
	"time"
)

// 模拟创建用户
func createUser() {

	_ = global.DB.AutoMigrate(&model.User{})

	options := &password.Options{SaltLen: 8, Iterations: 10, KeyLen: 16, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("123456", options)
	// pbkdf2 是密钥算法
	newPassWord := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	for i := 0; i < 1; i++ {
		user := model.User{
			UserName: fmt.Sprintf("wangqichao%d", i),
			Password: newPassWord,
			Email:    fmt.Sprintf("stimulate%d@gmail.com", i),
			Gender:   "male",
			Role:     1,
		}
		tx := global.DB.Table("user").Save(&user)
		log.Println("result.RowsAffected:", tx.RowsAffected, "result.Error:", tx.Error)
	}
}

// 创建一个RPC客户端
// 原来的程序是 人家一整个service 所以里面有config viper zap

var userClient user.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = user.NewUserClient(conn)
}

func TestCreateUser() {
	for i := 0; i < 1; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &user.CreateUserInfo{
			UserName: fmt.Sprintf("wangqichao1%d", i),
			Email:    fmt.Sprintf("simulate%d@gmail.com", i),
			PassWord: "123456",
			Gender:   "female",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id) // 只要打印了id就代表成功了
	}
}

func TestGetUserByEmail() {
	rsp, err := userClient.GetUserByEmail(context.Background(), &user.EmailRequest{
		Email: "simulate1@gmail.com",
	})
	// TODO 如果没有记录 不应该panic
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Id) // 只要打印了id就代表成功了
}

func TestUpdateUser() {
	rsp, err := userClient.UpdateUser(context.Background(), &user.UpdateUserInfo{
		Id:        8,
		UserName:  "王其超UP",
		PassWord:  "123456",
		Email:     "wangqichao0105@gmail.com",
		Gender:    "female",
		Role:      2,
		IsDeleted: 0,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("success", rsp) // 只要打印了id就代表成功了
}

func TestCheckPassWord() {
	rsp, err := userClient.CheckPassWord(context.Background(), &user.PasswordCheckInfo{
		Password:          "123456",
		EncryptedPassword: "$pbkdf2-sha512$hw4K31j9$2cecaa3813a90c1031c8f0aaf539c6c2",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Success)
}

func main() {
	//fmt.Println("这是demo程序...")
	//initialize.InitDB()
	//initialize.InitLogger()
	//Init()
	//fmt.Println("基础连接成功...")
	// createUser()
	// TestCreateUser()
	// TestGetUserByEmail()
	// TestUpdateUser()
	// TestCheckPassWord()
	// TestGetUserList()
	//TestUpdateUser()
	// gorm验证
	dsn := "root:root@tcp(127.0.0.1:3306)/stockexchange?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&model.UserFav{})
}

// 验证分页
func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &user.PageInfo{
		Pn:    4,
		PSize: 2,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Email, user.UserName, user.PassWord)
	}
}
