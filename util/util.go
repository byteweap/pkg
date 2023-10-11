// Package util 常用工具类
package util

import (
	"fmt"
	"math/rand"
	"time"
)

// SortIntArrAsc 升序排序数组
func SortIntArrAsc(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				// tmp := array[i]
				// array[i] = array[j]
				// array[j] = tmp
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

// EqualArr 判断两个数组是否完全相同
func EqualArr(_arr1 []int, _arr2 []int) bool {

	arr1 := SortIntArrAsc(_arr1)
	arr2 := SortIntArrAsc(_arr2)

	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// RemoveElement 数组删除指定元素,相同的全部删除
func RemoveElement(_array []int, _value int) []int {
	tmpArr := make([]int, 0)
	for _, v := range _array {
		if v != _value {
			tmpArr = append(tmpArr, v)
		}
	}
	return tmpArr
}

func RemoveStrElement(_array []string, _value string) []string {
	tmpArr := make([]string, 0)
	isRemove := false
	for _, v := range _array {
		if v != _value {
			tmpArr = append(tmpArr, v)
		} else {
			if !isRemove { //一次只删除一个
				isRemove = true
			} else {
				tmpArr = append(tmpArr, v)
			}
		}
	}
	return tmpArr
}

// RandArr 乱序一个数组
func RandArr(_array []int) []int {
	tmpArr := make([]int, 0)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(_array)
	for i := 0; i < length; i++ {
		ranIndex := random.Intn(len(_array))
		val := _array[ranIndex]
		tmpArr = append(tmpArr, val)
		_array = RemoveElement(_array, val)
	}
	return tmpArr
}

func RandStrArr(_array []string) []string {
	tmpArr := make([]string, 0)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := len(_array)
	for i := 0; i < length; i++ {
		ranIndex := random.Intn(len(_array))
		val := _array[ranIndex]
		tmpArr = append(tmpArr, val)
		_array = RemoveStrElement(_array, val)
	}
	return tmpArr
}

// GetRandomNum 获取一个 count位 的随机数
func GetRandomNum(count int) int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	number := 0
	for i := 0; i < count; i++ {
		ran := random.Intn(9) //0~9
		if i == count-1 {
			ran = random.Intn(8) + 1 //第一位不为0
		}

		value := ran * Power(10, i)
		//logs.Info("value:%v", value)
		number += value
	}

	return number
}

// GetRanExceptX 获取一个 [0,count) 的随机数
func GetRanExceptX(count int) int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	ran := random.Intn(count) //0~9
	return ran
}

// GetRanContainsX 获取一个 [1,count] 的随机数
func GetRanContainsX(count int) int {
	return GetRanExceptX(count) + 1
}

// GetRanBtween 获取一个 [0,count) 的随机数
func GetRanBtween(_min int, _max int) int {
	tmp := _max - _min + 1
	num := GetRanExceptX(tmp) + _min //[0,_max-_min +1) ->[_min,_max+1)
	return num
}

// Power 求value的 ct 次冥
func Power(value int, ct int) int {
	if ct == 0 {
		return 1
	}
	_tmp := value //保存底数
	for i := 0; i < ct-1; i++ {
		value *= _tmp
	}
	//log.Info("ct:%v------>value:%v", ct, value)
	return value
}

// GetMaxElement 获取数组中的最大元素
func GetMaxElement(_array []int) int {
	_data := _array[0]
	for _, v := range _array {
		if v > _data {
			_data = v
		}
	}
	return _data
}

// HasElement 判断数组中是否含有指定元素
func HasElement(_array []int, _element int) bool {
	for _, v := range _array {
		if v == _element {
			return true
		}
	}
	return false
}

// GetZuheResult 求排列组合 组合算法(从n个中取出m个数)---------------------------------------------------------
// _m 待取数组,_n:取n个
func GetZuheResult(_m []int, _n int) [][]int {
	n := len(_m)
	indexs := zuheResult(n, _n)
	result := findNumsByIndexs(_m, indexs)
	return result
}

func zuheResult(n int, m int) [][]int {
	if m < 1 || m > n {
		fmt.Println("Illegal argument. Param m must between 1 and len(nums).")
		return [][]int{}
	}
	//保存最终结果的数组，总数直接通过数学公式计算
	result := make([][]int, 0, mathZuhe(n, m))
	//保存每一个组合的索引的数组，1表示选中，0表示未选中
	indexs := make([]int, n)
	for i := 0; i < n; i++ {
		if i < m {
			indexs[i] = 1
		} else {
			indexs[i] = 0
		}
	}
	//第一个结果
	result = addTo(result, indexs)
	for {
		find := false
		//每次循环将第一次出现的 1 0 改为 0 1，同时将左侧的1移动到最左侧
		for i := 0; i < n-1; i++ {
			if indexs[i] == 1 && indexs[i+1] == 0 {
				find = true
				indexs[i], indexs[i+1] = 0, 1
				if i > 1 {
					moveOneToLeft(indexs[:i])
				}
				result = addTo(result, indexs)
				break
			}
		}
		//本次循环没有找到 1 0 ，说明已经取到了最后一种情况
		if !find {
			break
		}
	}
	return result
}

//将ele复制后添加到arr中，返回新的数组
func addTo(arr [][]int, ele []int) [][]int {
	newEle := make([]int, len(ele))
	copy(newEle, ele)
	arr = append(arr, newEle)
	return arr
}

func moveOneToLeft(leftNums []int) {
	//计算有几个1
	sum := 0
	for i := 0; i < len(leftNums); i++ {
		if leftNums[i] == 1 {
			sum++
		}
	}
	//将前sum个改为1，之后的改为0
	for i := 0; i < len(leftNums); i++ {
		if i < sum {
			leftNums[i] = 1
		} else {
			leftNums[i] = 0
		}
	}
}

//根据索引号数组得到元素数组
func findNumsByIndexs(nums []int, indexs [][]int) [][]int {
	if len(indexs) == 0 {
		return [][]int{}
	}
	result := make([][]int, len(indexs))
	for i, v := range indexs {
		line := make([]int, 0)
		for j, v2 := range v {
			if v2 == 1 {
				line = append(line, nums[j])
			}
		}
		result[i] = line
	}
	return result
}

//数学方法计算排列数(从n中取m个数)
func mathPailie(n int, m int) int {
	return jieCheng(n) / jieCheng(n-m)
}

//数学方法计算组合数(从n中取m个数)
func mathZuhe(n int, m int) int {
	return jieCheng(n) / (jieCheng(n-m) * jieCheng(m))
}

//阶乘
func jieCheng(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// RemoveSubArr 删除子数组
func RemoveSubArr(_array []int, _subArr []int) []int {
	for i := 0; i < len(_subArr); i++ {
	label:
		for j := 0; j < len(_array); j++ {
			if _array[j] == _subArr[i] {
				_array = RemoveIntElementOne(_array, _subArr[i])
				break label //只删除一个
			}
		}
	}
	return _array
}

// RemoveIntElementOne 数组删除指定value元素，只删除一个
func RemoveIntElementOne(_array []int, _value int) []int {
	tmpArr := make([]int, 0)
	isRemove := false
	for _, v := range _array {
		if v != _value {
			tmpArr = append(tmpArr, v)
		} else {
			if !isRemove { //一次只删除一个
				isRemove = true
			} else {
				tmpArr = append(tmpArr, v)
			}
		}
	}
	return tmpArr
}
