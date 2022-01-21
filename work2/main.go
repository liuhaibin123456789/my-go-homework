package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//获取linux的cpu利用率，内存使用量随时间的变化
func main() {
	var cpuArray= [100]float64{}
	var memArray=[100]float64{}
	i:=0
	j:=0
	file, err := os.OpenFile("D:\\Go Code\\demo\\fourth\\data.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("未找到,请更换绝对路径")
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	readLine:=make([]byte,1024)
	for true {
		readLine, _, err = reader.ReadLine()
		if err != nil {
			break
		}
		s := string(readLine)
		//匹配cpu和出现的mem位置
		if checkCPu(s) {
			get, err := regexGetCpu(s)
			if err != nil {
				fmt.Println(fmt.Sprintf("%s",err))
				return
			}
			//剪切字符串
			s2 := get[0][0]
			n := strings.Split(s2, " id")
			//获取空闲cpu百分比
			s3, err1 :=strconv.ParseFloat(n[0],64)
			if err1!=nil {
				fmt.Println(fmt.Sprintf("%s",err))
			}
			//转化空闲cpu百分比为cpu利用率百分比并储存至数组cpuArray
			cpuArray[i]=1-s3/100
			i++
		}
		if checkMem(s) {
			get, err := regexGetMem(s)
			if err != nil {
				fmt.Println(fmt.Sprintf("%s",err))
				return
			}
			//剪切字符串
			s2 := get[0][0]
			n := strings.Split(s2, " used")
			//获取mem 内存使用量
			s3, err1 :=strconv.ParseFloat(n[0],64)
			if err1!=nil {
				fmt.Println(fmt.Sprintf("%s",err))
			}
			//储存至数组memArray
			memArray[j]=s3
			j++
		}
	}
	fmt.Println("cpuArray->",cpuArray)
	fmt.Println("memArray->",memArray)
}

//checkCPu 定位CPu
func checkCPu(line string) bool {
	str:="%Cpu(s):"
	if strings.Contains(line,str) {
		return true
	}
	return false
}

//checkMem 定位Mem
func checkMem(line string) bool {
	str:="KiB Mem :"
	if strings.Contains(line,str) {
		return true
	}
	return false
}

//regexGetCpu 匹配cpu正则，获取数据
func regexGetCpu(line string) ([][]string, error) {
	//正则表达式
	regex:="[0-9]{1,2}[.]{1}[0-9][ ]{1}[i][d][,]"
	compile, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	stringSubMatch := compile.FindAllStringSubmatch(line,1)
	return stringSubMatch, nil
}

//regexGetMem 匹配mem正则，获取mem数据
func regexGetMem(line string) ([][]string, error) {
	//正则表达式
	regex:="[0-9]+[ ][u][s][e][d][,]"
	compile, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}
	stringSubMatch := compile.FindAllStringSubmatch(line,1)
	return stringSubMatch, nil
}