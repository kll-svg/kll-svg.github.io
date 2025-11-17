package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	//userId :=
	//userName, err := queryDatabase(userId)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(userName)
	//str, err := readFile("D:\\learn_git\\allSubject\\code\\go\\awesomeProject\\error\\test.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	// 打印文件内容
	//fmt.Println(str)
	//res, err := safeAccess([]int{1, 2, 3}, 3)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)
	//deferRes := deferReturn()
	//fmt.Println(deferRes)
	res, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

type BussinessError struct {
	Code    int
	Message string
	Time    time.Time
}

const (
	TimeFmt = "2006-01-02 15:04:05"
)

func (e *BussinessError) Error() string {
	return fmt.Sprintf("业务错误%d:%s, at %s", e.Code, e.Message, e.Time.Format(TimeFmt))
}

func queryDatabase(userId int) (string, error) {
	if userId <= 0 {
		return "", &BussinessError{
			Code:    400,
			Message: "用户ID无效",
			Time:    time.Now(),
		}
	}
	if userId == 999 {
		return "", errors.New("用户ID不存在")
	}
	// 模拟查询数据库
	return fmt.Sprintf("user_%d", userId), nil
}

func readFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("文件%s打开失败%w", fileName, err)
	}
	defer file.Close()
	buf := make([]byte, 1024)
	_, err = file.Read(buf)
	if err != nil {
		return "", fmt.Errorf("文件%s读取失败%w", fileName, err)
	}
	// 打印文件内容
	return string(buf), nil
}

func safeAccess(arr []int, index int) (res int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("数组访问越界panic%v", r)
		}
	}()
	res = arr[index]
	return res, nil
}

func deferReturn() (res int) {
	defer func() {
		res++
	}()
	return 10
}

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("除法出错: %v", r)
		}
	}()

	result = a / b
	return result, nil
}
