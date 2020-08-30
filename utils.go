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
