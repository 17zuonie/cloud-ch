package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/manifoldco/promptui"
)

var (
	functions = []func(){fixPermission}
	items     = []string{"配置一届的文件夹权限", "新建一届的文件夹", "退出"}
)

func fixPermission() {
	grade := simpleIntPrompt("操作届数", "22")
	start := simpleIntPrompt("起始编号", "1")
	end := simpleIntPrompt("结束编号", "16")

	if simpleConfirmPrompt() {
		os.Chdir(fmt.Sprintf("/volume1/云上春晖/20%02d届", grade))
		system("synoacltool", "-enforce-inherit", ".")
		system("synoacltool", "-add", ".", "group:class:allow:r-x---a-R-c--:---n")

		for i := start; i <= end; i++ {
			username := fmt.Sprintf("%d%02d", grade, i)

			user, err := user.Lookup(username)
			if err != nil {
				fmt.Println("找不到用户: ", username)
				return
			}

			dir := fmt.Sprintf("%02d", i)

			fmt.Println("正在设定所有者: ", username)
			chownR(dir, getuid(user), getgid(user))
			system("synoacltool", "-enforce-inherit", dir)
			system("synoacltool", "-add", dir, fmt.Sprintf("user:%s:allow:rwxpdDaARWcCo:fd--", username))
			fmt.Println("已完成: ", username)
		}
	} else {
		fmt.Println("操作取消了")
	}
}

func main() {
	banner := `
  _____ _                 _    _____ _                 _    _       _ 
 / ____| |               | |  / ____| |               | |  | |     (_)
| |    | | ___  _   _  __| | | |    | |__  _   _ _ __ | |__| |_   _ _ 
| |    | |/ _ \| | | |/ _  | | |    | '_ \| | | | '_ \|  __  | | | | |
| |____| | (_) | |_| | (_| | | |____| | | | |_| | | | | |  | | |_| | |
 \_____|_|\___/ \__,_|\__,_|  \_____|_| |_|\__,_|_| |_|_|  |_|\__,_|_|`
	fmt.Println(banner)

	if os.Getuid() != 0 {
		fmt.Println("权限不足! 请使用 sudo 运行")
		return
	}

	prompt := promptui.Select{
		Label: "请告诉我你的选择",
		Items: items,
	}

	for {
		index, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if result == "退出" {
			break
		}
		functions[index]()
	}
}
