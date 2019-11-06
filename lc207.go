package main

import (
	"database/sql"
	"fmt"
	"net/http"
	_ "os"
	"strconv"
	_ "strconv"
)

var m struct {
	http.Client
	sql.DB
}

var annostring = func(str string) func() string {
	return func() string {
		return str
	}
}

var anno = func(str string) string {
	return str
}

func f(num1 string, num2 string) string {
	minL := Min(len(num1), len(num2))
	maxL := Max(len(num1), len(num2))
	res := ""
	car := 0
	ano := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		var tmp int = int((num1[i] - '0') + (num2[j] - '0'))
		res += strconv.Itoa((tmp + car) % 10)
		if ((tmp+car)/10 != 0) {
			car = 1
		} else {
			car = 0
		}
		fmt.Printf("%v ", (tmp+car)%10)
	}
	if len(num1) > minL {
		ano += num1
	} else {
		ano += num2
	}

	for i := maxL - minL - 1; i >= 0; i-- {
		res += strconv.Itoa((int(ano[i]-'0') + car) % 10)
		if ((int(ano[i]-'0')+car)/10 != 0) {
			car = 1
		} else {
			car = 0
		}
	}
	if (car != 0) {
		res += strconv.Itoa(car)
	}
	return res
}

func main() {
	//a := [2][2]int{{0, 1}, {2, 1}}
	//var arr = [5]int{1, 2, 3, 4, 5}
	//sum(arr[:])
	fmt.Printf("\t%-16v %v", "asd","123")
	fmt.Printf("%v", f("999", "999"))
}

func Max(a int, b int) int {
	if (a > b) {
		return a;
	} else {
		return b;
	}
}

func Min(a int, b int) int {
	if (a < b) {
		return a;
	} else {
		return b;
	}
}

func sum(arr []int) int {
	s := 0
	for i := 0; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}

func canFinish(numCourses int, pre [][]int) bool {
	return true
}

func f1(numCourses int, prerequisites [][]int) {
	num := make([][]int, numCourses, numCourses);
	for i := 0; i < numCourses; i++ {
		num[i] = make([]int, numCourses+1)
	}

	for _, v := range num {
		for i := 0; i < len(v); i++ {
			fmt.Printf("%v \n", v[i]);
		}
	}
}
