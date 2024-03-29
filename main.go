﻿package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var console = bufio.NewReader(os.Stdin)

func main() {
	initLog()
	//重定向服务器
	go StartServer()
	fmt.Println("将跳转浏览器授权")
	fmt.Println("请输入App的应用程序(客户端) ID [Client Id]")
	clientId := readLine()
	fmt.Println("请输入App的API权限,并使用空格隔开")
	fmt.Println("如果为空则默认为(offline_access Files.Read Files.Read.All Files.ReadWrite Files.ReadWrite.All)")
	permission := readLine()
	if len(permission) == 0 {
		permission = "offline_access Files.Read Files.Read.All Files.ReadWrite Files.ReadWrite.All"
		fmt.Printf("输入为空，已使用默认设置")
		fmt.Println()
	}
	err := OpenBrowser(BuildMsAuthorizeUrl(clientId, permission))
	if err != nil {
		fmt.Printf("启动浏览器时发生错误\r\n%s", err.Error())
		readLine()
		return
	}
	fmt.Println()
	clientCode := <-CodeCh
	if clientCode != nil {
		fmt.Printf("获取到code(%v)", len(*clientCode))
		fmt.Println()
	} else {
		fmt.Println("获取code失败，请检查client_id和网络设置")
		readLine()
		return
	}
	fmt.Println()
	fmt.Println("请输入App的机密值[Client Secret]")
	clientSecret := readLine()
	fmt.Println()
	tokenData, err := MsTokenRequest(clientId, clientSecret, *clientCode)
	if err != nil {
		fmt.Printf("获取token时发生错误\r\n%s", err.Error())
		_ = readLine()
		return
	}
	fmt.Printf("获取到token(%v)", len(tokenData.AccessToken))
	fmt.Println()
	fmt.Printf("获取到refresh-token(%v)", len(tokenData.RefreshToken))
	fmt.Println()
	saveData := TokenResult{
		ClientId: clientId,
		ClientCode: *clientCode,
		Secret: clientSecret,
		AccessToken: tokenData.AccessToken,
		RefreshToken: tokenData.RefreshToken,
	}
	jsonBytes, err := saveData.Save()
	if err != nil {
		fmt.Println("保存token信息失败，请自行复制获取到的token")
		time.Sleep(time.Second * 3)
		fmt.Println(string(*jsonBytes))
	}else{
		fmt.Println("token信息已保存至工具根目录的result.json文件中")
	}
	fmt.Println("按下回车退出本工具")
	readLine()
}

func readLine() (str string) {
	str, _ = console.ReadString('\n')
	str = strings.TrimSpace(str)
	return
}

func initLog() {
	fmt.Println("Ciallo～(∠・ω< )⌒☆")
	fmt.Println("本工具将帮助你获取微软Graph的API的访问令牌，获取到的结果将保存到工具运行的目录下（默认将选择OneDriver相关权限）")
	fmt.Println()
	fmt.Println("请先保证使用本工具前已经在Microsoft Azure中创建了新的应用，并满足以下条件:")
	fmt.Println("1.重定向URL选择了(Web)类型，并设置值为(http://localhost:11451/auth)")
	fmt.Println("2.在证书和密码选项卡中创建并保存了机密值(Client Secret)")
	fmt.Println("如果需要使用OneDriver相关API(如各种网盘列表)，请在API权限中授权(offline_access Files.Read Files.Read.All Files.ReadWrite Files.ReadWrite.All)")
	fmt.Println("请在确认以上条件后按回车继续")
	_ = readLine()
}
