#! /usr/bin/env python3
# -*- coding: utf-8 -*-


import os

ART = '''
   _____ _                 _    _____ _                 _    _       _ 
  / ____| |               | |  / ____| |               | |  | |     (_)
 | |    | | ___  _   _  __| | | |    | |__  _   _ _ __ | |__| |_   _ _ 
 | |    | |/ _ \| | | |/ _` | | |    | '_ \| | | | '_ \|  __  | | | | |
 | |____| | (_) | |_| | (_| | | |____| | | | |_| | | | | |  | | |_| | |
  \_____|_|\___/ \__,_|\__,_|  \_____|_| |_|\__,_|_| |_|_|  |_|\__,_|_|
                                                                       
'''

WELCOME = "=== 欢迎使用云上春晖 OP 工具 ==="

MENU = '''
1 -> 新建一届的文件夹
2 -> 配置一届的文件夹权限
3 -> 
x -> 退出
'''

SUBJECTS = ["语文", "数学", "英语", "物理", "化学", "生物", "政治", "历史", "地理", "技术"]

VOLUME = 1
SHAREDIR = "云上春晖"


def ask_with_default(prompt, default):
    answer = input(f"{prompt} [{default}]: ")
    if len(answer) == 0:
        answer = default
    return answer


def ask_for_start(summary):
    print(summary)
    answer = ""
    while len(answer) == 0 or (answer != "n" and answer != "y"):
        answer = input("确定运行吗 [y/n]: ")
    return answer == "y"


def create_directory():
    pass


def fix_permission():
    grade = int(ask_with_default("操作届数", 22))
    start = int(ask_with_default("起始编号", 1))
    end = int(ask_with_default("结束编号", 16))

    if ask_for_start(f"\t届数: {grade}\n\t起始: {start}\n\t结束: {end}\n\t警告: 当前届所有文件夹原有权限会全部重置"):
        os.chdir(f"/volume{VOLUME}/{SHAREDIR}/20{grade}届")
        os.system(f"synoacltool -enforce-inherit .")
        os.system(f"synoacltool -add . group:class:allow:r-x---a-R-c--:---n")
        for i in range(start, end+1):
            user = f"{grade}{i:0>2d}"
            folder = f"{i:0>2d}"
            os.system(f"chown -R {user} {folder}")
            os.system(f"synoacltool -enforce-inherit {folder}")
            os.system(f"synoacltool -add {folder} user:{user}:allow:rwxpdDaARWcCo:fd--")


FUNCTIONS = [lambda: print("0 是不存在的"), create_directory, fix_permission]

if __name__ == "__main__":
    print(ART)
    print(WELCOME)

    choice = ""
    while True:
        print(MENU)
        choice = input("请告诉我你的选择：")
        if choice.isdecimal():
            FUNCTIONS[int(choice)]()
        elif choice == "x":
            break
        else:
            print("您的输入不正确")

    print("Goodbye!")
