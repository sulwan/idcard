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

func main() {
	fmt.Println("身份证所在地:" + Address("130105198410040316"))
	fmt.Println("性别:" + Sex("130105198410040316"))
	fmt.Println("所属动物:" + Zodiac("130124198705293638"))
}
