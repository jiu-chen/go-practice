package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

func getWordFrequency(readFilePath string, writeFilePath string) {
	var fileText string
	var wordFrequencyMap = make(map[string]int)
	fileData, readFileErr := ioutil.ReadFile(readFilePath)
	if readFileErr != nil {
		log.Fatal(readFileErr)
	}
	fileText = string(fileData)
	f := func(c rune) bool {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			return true
		}
		return false
	}
	//根据传入的函数分割字符串，如果当前参数c不是数字或者字母，返回true作为分割符号。
	arr := strings.FieldsFunc(fileText, f)
	for _, v := range arr {
		if _, ok := wordFrequencyMap[v]; ok {
			wordFrequencyMap[v] = wordFrequencyMap[v] + 1
		} else {
			wordFrequencyMap[v] = 1
		}
	}
	fmt.Println(wordFrequencyMap)
	type wordFrequencyNum struct {
		Word string
		Num  int
	}
	var lstWordFrequencyNum []wordFrequencyNum
	for k, v := range wordFrequencyMap {
		lstWordFrequencyNum = append(lstWordFrequencyNum, wordFrequencyNum{k, v})
	}
	sort.Slice(lstWordFrequencyNum, func(i, j int) bool {
		return lstWordFrequencyNum[i].Num > lstWordFrequencyNum[j].Num
	})
	fmt.Println("按照单词出现频率由高到低排序", lstWordFrequencyNum)

	//写入文件
	var jsonBytes []byte
	var arrJsonByte string
	var err error
	for _, v := range lstWordFrequencyNum {
		jsonBytes, err = json.Marshal(v)
		if err != nil {
			log.Fatal(err)
		}
		arrJsonByte = arrJsonByte + string(jsonBytes)
	}
	err = ioutil.WriteFile(writeFilePath, []byte(arrJsonByte), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	readFilePath := "/Users/juzichen/Documents/study/go/input.txt"    //读取的文件路径
	writeFilePath := "/Users/juzichen/Documents/study/go/out_put.txt" //写入的文件路径
	getWordFrequency(readFilePath, writeFilePath)
}
