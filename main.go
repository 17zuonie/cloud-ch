package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"

	"github.com/manifoldco/promptui"
)

var (
	functions = []func(){fixPermission}
	items     = []string{"配置一届的文件夹权限", "新建一届的文件夹", "退出"}
)

func simplePromptValidater(input string) error {
	_, err := strconv.Atoi(input)
	if err != nil {
		return errors.New("Invalid number")
	}
	return nil
}

func simpleIntPrompt(label, def string) int {
	prompt := promptui.Prompt{
		Label:     label,
		Validate:  simplePromptValidater,
		Default:   def,
		AllowEdit: true,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1
	}

	number, err := strconv.Atoi(result)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1
	}
	return number
}

func simpleConfirmPrompt() bool {
	prompt := promptui.Prompt{
		Label:     "你确定要执行吗",
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		return false
	}
	return result == "Y" || result == "y"
}

func system(name string, args ...string) {
	cmd := exec.Command(name, args...)
	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

func getuid(user *user.User) int {
	number, err := strconv.Atoi(user.Uid)
	if err != nil {
		panic(err)
	}
	return number
}

func getgid(user *user.User) int {
	number, err := strconv.Atoi(user.Gid)
	if err != nil {
		panic(err)
	}
	return number
}

func chownR(path string, uid, gid int) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
		}
		return err
	})
}

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
