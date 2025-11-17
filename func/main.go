package main

import (
	"errors"
	"fmt"
	"sort"
)

func main() {
	name, age, err := QueryUserInfo(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	counter := CreateCounter()
	for i := 0; i < 5; i++ {
		fmt.Println(counter())
	}
	fmt.Println(name, age)

	strs := []string{"hello", "world", "golang"}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	fmt.Println(strs)

	processor := func(s string) string {
		return s + "1"
	}
	res := ProcessString(strs, processor)
	fmt.Println(res)
	fmt.Println(sum(1, 2, 3, 4, 5))
}

func QueryUserInfo(userId int) (string, int, error) {
	if userId <= 0 {
		return "", 0, errors.New("userId is invalid")
	}
	users := map[int]struct {
		Name string
		Age  int
	}{
		1: {
			Name: "张三",
			Age:  18,
		},
		2: {
			Name: "李四",
			Age:  20,
		},
		3: {
			Name: "王五",
			Age:  22,
		},
	}
	if user, ok := users[userId]; ok {
		return user.Name, user.Age, nil
	}
	return "", 0, errors.New("user not found")
}

func CreateCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

type StringProcessor func(string) string

func ProcessString(str []string, processor StringProcessor) []string {
	var result []string
	for _, s := range str {
		result = append(result, processor(s))
	}
	return result
}

func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
