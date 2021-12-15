package main

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc"
	"log"
	"stockexchange/rpc/demo/global"
	"stockexchange/rpc/demo/initialize"
	"stockexchange/rpc/demo/model"
	"stockexchange/rpc/demo/proto/user"
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
		Id:        3, // 必须传进去id
		UserName:  "xiaowang1",
		PassWord:  "$pbkdf2-sha512$cuoH74L1$8eb35bfe2aac5c0aa00534eaa2341c8e",
		Email:     "simulate1@gmai3.com",
		Gender:    "male",
		IsDeleted: 0, // 我没传进去role 所以你不能更改
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
	fmt.Println("这是demo程序...")
	initialize.InitDB()
	initialize.InitLogger()
	Init()
	fmt.Println("基础连接成功...")
	// createUser()
	// TestCreateUser()
	// TestGetUserByEmail()
	// TestUpdateUser()
	// TestCheckPassWord()
	TestGetUserList()
}

// 验证分页
func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &user.PageInfo{
		Pn:    1,
		PSize: 2,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Email, user.UserName, user.PassWord)
	}
}
