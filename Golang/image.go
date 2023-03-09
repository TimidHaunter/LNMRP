package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	var number int
	var pro int

	fmt.Print("你想生成多少张图片:")
	fmt.Scan(&number)

	//素材数量
	baseDirPath := ".\\image\\base\\"
	var baseCount = getFileNum(baseDirPath)

	//文案数量
	contentDirPath := ".\\image\\content\\"
	var contentCount = getFileNum(contentDirPath)

	//角标数量
	subscriptDirPath := ".\\image\\subscript\\"
	var subscriptCount = getFileNum(subscriptDirPath)

	maxNumber := baseCount * contentCount * subscriptCount
	if number > maxNumber {
		fmt.Printf("生成的图片张数太多，建议" + fmt.Sprint(maxNumber) + "以内\n")
		os.Exit(0)
	}

	fmt.Print("你想有多少概率生成带警示语的图片（1-100）:")
	fmt.Scan(&pro)

	//开始时间
	start := time.Now()
	//初始化生成的图片编号
	var iamge []string
	rand.Seed(time.Now().UnixNano())

	//最大可以合成图片次数
	for i := 1; i <= number; i++ {
		//随机图片编号
		baseI := rand.Intn(baseCount) + 1
		contentI := rand.Intn(contentCount) + 1
		subscriptI := rand.Intn(subscriptCount) + 1
		imageNum := fmt.Sprint(baseI) + "_" + fmt.Sprint(contentI) + "_" + fmt.Sprint(subscriptI)
		// fmt.Printf(imageNum + "\n")

		if !IsContain(iamge, imageNum) {
			iamge = append(iamge, imageNum)
			createImage(i, baseI, contentI, subscriptI, pro)

		}
	}

	//结束时间
	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

/**
 * 生成图片
 */
func createImage(i, baseI, contentI, subscriptI, pro int) {

	//概率生成带有警示语的图片
	rand.Seed(time.Now().UnixNano())
	randWarning := rand.Intn(99)

	//随机出来的数字大于概率(99%)
	cmdArguments := []string{}
	if (randWarning + 1) <= pro {
		cmdArguments = []string{
			"-i", ".\\image\\base\\" + fmt.Sprint(baseI) + ".jpg",
			"-i", ".\\image\\content\\" + fmt.Sprint(contentI) + ".png",
			"-i", ".\\image\\subscript\\" + fmt.Sprint(subscriptI) + ".png",
			"-i", ".\\image\\warning\\1.png",
			"-filter_complex", "overlay=0:0,overlay=1066,overlay=0:0",
			".\\image\\res\\image_" + fmt.Sprint(i) + ".jpg"}
	} else {
		cmdArguments = []string{
			"-i", ".\\image\\base\\" + fmt.Sprint(baseI) + ".jpg",
			"-i", ".\\image\\content\\" + fmt.Sprint(contentI) + ".png",
			"-i", ".\\image\\subscript\\" + fmt.Sprint(subscriptI) + ".png",
			"-filter_complex", "overlay=0:0,overlay=1066",
			".\\image\\res\\image_" + fmt.Sprint(i) + ".jpg"}
	}

	// fmt.Println(cmdArguments)
	cmd := exec.Command("ffmpeg", cmdArguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("生成图片失败\n")
	} else {
		fmt.Printf("生成图片成功\n")
	}
}

/**
 * 获取路径下文件数量
 */
func getFileNum(dirName string) int {
	files, _ := ioutil.ReadDir(dirName)
	count := len(files)
	return count
}

/**
 * 判断元素是否在数组中
 */
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
