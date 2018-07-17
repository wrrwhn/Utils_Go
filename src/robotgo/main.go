package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
)

func main() {

	activeProgress()

}

func activeProgress() {

	list := []string{
		"密码", "password",
		"打印机设置",
		"Microsoft Windows",
		"Microsoft Excel", "Excel(产品激活失败)",
		"Microsoft PowerPoint", 
		"Microsoft Word", 
		"Microsoft Office 激活向导"}

	for _, v := range list {
		pids, err := robotgo.FindIds(v)

		if nil != err {
			fmt.Printf("active[%s].fail: %s\n", v, err.Error())
			continue
		}

		fmt.Printf("FindIds(%s)= %v\n", v, pids)
		if len(pids) > 0 {

			robotgo.ActivePID(pids[0])
			robotgo.KeyTap("escape")
			fmt.Printf("\tactive[%s].esc\n", v)
		}
	}
}

func activeDetail() {

	list := []string{
		"EXCEL.EXE",
		"POWERPOINT.EXE",
		"WORD.EXE"}

	pids, err := robotgo.Pids()
	if nil != err {
		fmt.Printf("active.pids().fail: %s\n", err.Error())
	}
	for _, v := range pids {

		n, err := robotgo.FindName(v)
		if nil != err {
			fmt.Printf(err.Error())
			continue
		}

		// fmt.Println(n, v)
		if contain(list, n) {
			robotgo.ActivePID(v)
			robotgo.KeyTap("escape")
			fmt.Printf("\tactive[%v-%s].esc\n", v, n)
		}
	}
}

func contain(ks []string, k string) bool {

	for _, v := range ks {
		if v == k {
			return true
		}
	}
	return false
}

func find() {

	fmt.Println("Find.*")

	title := "任务管理器"
	hwnd := robotgo.FindWindow(title)
	fmt.Printf("\tfullname[%s]: %d\n", title, hwnd)

	fmt.Printf("\tprogress.*\n")
	names, err := robotgo.FindNames()
	if nil != err {
		fmt.Errorf("\t\tfail to get progress.*: %s\n", err.Error())
	} else {
		for _, name := range names {
			fmt.Printf("\t\t%s\n", name)
		}
	}

	title = "yqjdcyy"
	fmt.Printf("\tid(name(%s*))\n", title)
	ids, err := robotgo.FindIds(title)
	if nil != err {
		fmt.Errorf("\t\tfail to get progress.*: %s\n", err.Error())
	} else {
		for _, id := range ids {
			name, err := robotgo.FindName(id)
			if nil != err {
				fmt.Errorf("\t\tfind[%s] fail: %s\n", name, err.Error())
			} else {
				fmt.Printf("\t\tprogress[%d].name= %s\n", id, name)
			}
		}
	}

	fmt.Printf("\tprogress.*\n")
	ids, err = robotgo.Pids()
	if nil != err {
		fmt.Errorf("\t\tfail to get progress.*: %s\n", err.Error())
	} else {
		for idx, id := range ids {

			if idx >= 3 {
				fmt.Printf("\t\tprogress...")
				break
			}
			name, err := robotgo.FindName(id)
			if nil != err {
				fmt.Errorf("\t\tfind[%s] fail: %s\n", name, err.Error())
			} else {
				fmt.Printf("\t\tprogress[%d].name= %s\n", id, name)
			}
		}
	}

}
