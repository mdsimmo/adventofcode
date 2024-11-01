package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func aoc_12_run() {

	file, err := os.Open("aoc_12_in.txt")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	var data interface{}
	decoder.Decode(&data)

	fmt.Println(sum_unknown(data))
	
}

func sum_unknown(value interface{}) int {
	mp, ok := value.(map[string]interface{})
	if ok {
		return sum_map(mp)
	}
	arr, ok := value.([]interface{})
	if ok {
		return sum_array(arr)
	}
	_, ok = value.(string)
	if ok {
		return 0
	}
	num, ok := value.(float64)
	if ok {
		return int(num)
	}
	panic(fmt.Sprintf("Unknown type: %+v, %+v", reflect.TypeOf(value), value))
}

func sum_map(data map[string]interface{}) int {
	sum := 0
	for _, value := range(data) {
		if str, ok := value.(string); ok {
			if str == "red" {
				return 0
			}
		}
		sum += sum_unknown(value)
	}
	return sum
}

func sum_array(data []interface{}) int {
	sum := 0
	for _, value := range(data) {
		sum += sum_unknown(value)
	}
	return sum
}

// This works for part 1, but doesn't do part 2 right
// I don't know why - process looks good, but wrong answer
func aoc_12_run_broken() {
	text_bytes, err := os.ReadFile("aoc_12_in.txt")
	if err != nil {
		panic(err)
	}

	text := string(text_bytes)

	num := ""
	sum := []int64 { 0 }
	discard := []bool { false }
	isarray := []bool { false }
	for i := 0; i < len(text); i++ {
		ch := text[i]

		if i >= 5 && (text[i-5:i] == "\"red\"") {
			if !isarray[len(isarray)-1] {
				discard[len(discard)-1] = true
				fmt.Printf("Discard RED:\n\tSum: \t%s\n\tArr:\t%s\n\tDisc:\t%s\n", 
					strings.ReplaceAll(fmt.Sprintf("%+v", sum), " ", "\t"), 
					strings.ReplaceAll(fmt.Sprintf("%+v", isarray), " ", "\t"), 
					strings.ReplaceAll(fmt.Sprintf("%+v", discard), " ", "\t"))
			} else {
				fmt.Printf("RED ignored\n")
			}
		}

		if ch == '{' || ch == '[' {
			sum = append(sum, 0)
			discard = append(discard, false)
			isarray = append(isarray, ch == '[')
			fmt.Printf("Up one:\n\tSum: \t%s\n\tArr:\t%s\n\tDisc:\t%s\n", 
				strings.ReplaceAll(fmt.Sprintf("%+v", sum), " ", "\t"), 
				strings.ReplaceAll(fmt.Sprintf("%+v", isarray), " ", "\t"), 
				strings.ReplaceAll(fmt.Sprintf("%+v", discard), " ", "\t"))

		}
		if ch == '}' || ch == ']' {
			if !discard[len(discard)-1] {
				sum[len(sum)-2] += sum[len(sum)-1]
			}
			sum = sum[:len(sum)-1]
			discard = discard[:len(discard)-1]
			isarray = isarray[:len(isarray)-1]
			fmt.Printf("Down one:\n\tSum: \t%s\n\tArr:\t%s\n\tDisc:\t%s\n", 
				strings.ReplaceAll(fmt.Sprintf("%+v", sum), " ", "\t"), 
				strings.ReplaceAll(fmt.Sprintf("%+v", isarray), " ", "\t"), 
				strings.ReplaceAll(fmt.Sprintf("%+v", discard), " ", "\t"))
		}

		if (ch >= '0' && ch <= '9') || (ch == '-') {
			num += string(ch)
		} else if len(num) > 0 {
			asnum, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				panic(err)
			}
			sum[len(sum)-1] += asnum
			num = ""
			fmt.Printf("Add num: %d\n\tSum: \t%s\n\tArr:\t%s\n\tDisc:\t%s\n", 
				asnum, 
				strings.ReplaceAll(fmt.Sprintf("%+v", sum), " ", "\t"), 
				strings.ReplaceAll(fmt.Sprintf("%+v", isarray), " ", "\t"), 
				strings.ReplaceAll(fmt.Sprintf("%+v", discard), " ", "\t"))
		}
	}

	fmt.Printf("Sum: %+v", sum)
}
