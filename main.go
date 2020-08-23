package main

import (
	"errors"
	"fmt"
	"os/exec"
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

func system(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}

func fixPermission() {
	fmt.Println(simpleIntPrompt("操作届数", "22"))
	fmt.Println(simpleIntPrompt("起始编号", "1"))
	fmt.Println(simpleIntPrompt("结束编号", "16"))
	if simpleConfirmPrompt() {

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
