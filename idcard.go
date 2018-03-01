package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var IdcardMap map[string]string

// 初始化程序
func init() {
	idFile, err := os.Open("./data/id.json")
	err = json.NewDecoder(idFile).Decode(&IdcardMap)
	if err != nil {
		fmt.Println(err)
	}
}

// 获取省份
func Province(id string) string {
	return IdcardMap[id[:2]+"0000"]
}

// 获取城市
func City(id string) string {
	return IdcardMap[id[:4]+"00"]
}

// 获取地区
func Zone(id string) string {
	return IdcardMap[id[:6]]
}

// 获取完整省市地区
func Address(id string) string {
	return Province(id) + City(id) + Zone(id)
}

// 获取性别
func Sex(id string) string {
	val, _ := strconv.Atoi(id[16:17])
	if val%2 != 0 {
		return "男"
	}
	return "女"
}

// 获取生肖
func Zodiac(id string) string {
	zodiac := []string{"鼠", "牛", "虎", "兔", "龙", "蛇", "马", "羊", "猴", "鸡", "狗", "猪"}
	year, _ := strconv.Atoi(Year(id))
	val := (year - 1900) % 12
	return zodiac[val]
}

// 获取年份
func Year(id string) string {
	return id[6:10]
}

// 获取星座
func constellation(id string) string {
	constellation := []string{"水瓶座", "双鱼座", "白羊座", "金牛座", "双子座", "巨蟹座", "狮子座", "处女座", "天秤座", "天蝎座", "射手座", "摩羯座"}
	marginal := []int{20, 19, 21, 20, 21, 22, 23, 23, 23, 24, 23, 22}
	mon, _ := strconv.Atoi(month(id))
	month := mon - 1
	day, _ := strconv.Atoi(day(id))
	if day < marginal[month] {
		month--
	}
	if month >= 0 {
		return constellation[month]
	} else {
		return constellation[11]
	}
}

// 获取月
func month(id string) string {
	return id[10:12]
}

// 获取日
func day(id string) string {
	return id[12:14]
}

func main() {
	fmt.Println("身份证所在地:" + Address("130431199402181323"))
	fmt.Println("性别:" + Sex("130431199402181323"))
	fmt.Println("所属动物:" + Zodiac("130431199402181323"))
	fmt.Println("所属星座:" + constellation("130121199409180025"))
}
