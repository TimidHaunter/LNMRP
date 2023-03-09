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
	var category int

	fmt.Print("你想生成多少个视频:")
	fmt.Scan(&number)

	fmt.Print("你想生成哪类视频（1.乡村，2.悬疑）:")
	fmt.Scan(&category)

	start := time.Now()
	if number > 0 && number <= 100 {
		for i := 1; i <= number; i++ {
			createVideo(i, category)
		}
	} else if number > 100 {
		fmt.Printf("生成的视频个数太多，建议100以内\n")
	}

	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}

func createVideo(i int, category int) {
	fmt.Printf("第" + fmt.Sprint(i) + "次生成视频开始\n")

	//1.图片合成视频
	img2Video(category)

	//2.分辨率
	transcodeVideo()

	//3.背景
	addBg(category)

	//4.配音乐
	addMusic(category)

	/**
	 * 5.加文案
	 *
	 * 通用文案
	 *
	 * 本故事纯属虚构 7x24=168 (720-168)/2=276
	 */

	addText(i)

	//6.转比特率
	transcodeBit(i)

	fmt.Printf("第" + fmt.Sprint(i) + "次生成视频结束\n")
}

func img2Video(category int) {
	rand.Seed(time.Now().UnixNano())

	dirName := ".\\video\\images\\" + fmt.Sprint(category) + "\\"
	var count = getFileNum(dirName)

	//初始化数组
	var expSlice []int
	for i := 1; i <= count; i++ {
		expSlice = append(expSlice, i)
	}

	// fmt.Print(expSlice)
	// os.Exit(0)

	//打乱图片顺序
	rand.Shuffle(count, func(i, j int) {
		expSlice[i], expSlice[j] = expSlice[j], expSlice[i]
	})

	/**
	 * 图片转场特效
	 */
	var gallery [12]string
	gallery[0] = "fade"
	gallery[1] = "fadeblack"
	gallery[2] = "fadewhite"
	gallery[3] = "wipeleft"
	gallery[4] = "wiperight"
	gallery[5] = "wipeup"
	gallery[6] = "wipedown"
	gallery[7] = "slideleft"
	gallery[8] = "slideright"
	gallery[9] = "slideup"
	gallery[10] = "slidedown"
	gallery[11] = "smoothleft"
	gallery_hit := gallery[rand.Intn(12)]

	// 4(3+1)+5(1+3+1)+5(1+3+1)+4(3+1)
	cmdArguments := []string{
		"-loop", "1", "-t", "4", "-i", ".\\video\\images\\" + fmt.Sprint(category) + "\\" + fmt.Sprint(expSlice[1]) + ".jpg",
		"-loop", "1", "-t", "5", "-i", ".\\video\\images\\" + fmt.Sprint(category) + "\\" + fmt.Sprint(expSlice[2]) + ".jpg",
		"-loop", "1", "-t", "5", "-i", ".\\video\\images\\" + fmt.Sprint(category) + "\\" + fmt.Sprint(expSlice[3]) + ".jpg",
		"-loop", "1", "-t", "4", "-i", ".\\video\\images\\" + fmt.Sprint(category) + "\\" + fmt.Sprint(expSlice[4]) + ".jpg",
		"-filter_complex", "[0][1]xfade=transition=" + gallery_hit + ":duration=1:offset=3[V01]; [V01][2]xfade=transition=" + gallery_hit + ":duration=1:offset=7[V02]; [V02][3]xfade=transition=" + gallery_hit + ":duration=1:offset=11,format=yuv420p[video]",
		"-map", "[video]",
		"-movflags", "+faststart",
		".\\video\\res\\video.mp4"}

	cmd := exec.Command("ffmpeg", cmdArguments...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("合视频失败,选择图片为No." + fmt.Sprint(expSlice[1]) + "-" + fmt.Sprint(expSlice[2]) + "-" + fmt.Sprint(expSlice[3]) + "-" + fmt.Sprint(expSlice[4]) + "\n")
		os.RemoveAll(".\\video\\res\\video.mp4")
	} else {
		fmt.Printf("合视频成功\n")
	}
}

//转分辨率
func transcodeVideo() {
	cmdArguments := []string{
		"-i", ".\\video\\res\\video.mp4",
		"-vf", "scale=720:406",
		// "-b:v", "8M",
		".\\video\\res\\video_720p.mp4"}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("分辨率失败\n")
	} else {
		fmt.Printf("分辨率成功\n")
	}
	err_video := os.RemoveAll(".\\video\\res\\video.mp4")
	if err_video != nil {
		fmt.Printf("删除合成视频失败\n")
	}
}

//加背景图片
func addBg(category int) {
	rand.Seed(time.Now().UnixNano())

	dirName := ".\\video\\bg\\" + fmt.Sprint(category) + "\\"
	var count = getFileNum(dirName)

	// fmt.Print(count)
	// os.Exit(0)

	cmdArguments := []string{
		"-loop", "1", "-i", ".\\video\\bg\\" + fmt.Sprint(category) + "\\" + fmt.Sprint(rand.Intn(count-1)+1) + ".png", "-i", ".\\video\\res\\video_720p.mp4",
		"-filter_complex", "overlay=(W-w)/2:(H-h)/2:shortest=1,format=yuv420p",
		".\\video\\res\\video_bg.mp4"}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("加背景失败，选择背景No." + fmt.Sprint(rand.Intn(count-1)+1) + "\n")
	} else {
		fmt.Printf("加背景成功\n")
	}
	err_video_720p := os.RemoveAll(".\\video\\res\\video_720p.mp4")
	if err_video_720p != nil {
		fmt.Printf("删除分辨率视频失败\n")
	}
}

//加音乐
func addMusic(category int) {
	rand.Seed(time.Now().UnixNano())

	dirName := ".\\video\\audio\\" + fmt.Sprint(category) + "\\"
	var count = getFileNum(dirName)

	cmdArguments := []string{
		"-i", ".\\video\\res\\video_bg.mp4", "-i", ".\\video\\audio\\" + fmt.Sprint(category) + "\\t" + fmt.Sprint(rand.Intn(count-1)+1) + ".mp3",
		".\\video\\res\\video_audio.mp4"}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("配音乐失败\n")
	} else {
		fmt.Printf("配音乐成功\n")
	}
	err_video_bg := os.RemoveAll(".\\video\\res\\video_bg.mp4")
	if err_video_bg != nil {
		fmt.Printf("删除背景视频失败\n")
	}
}

//加文案
func addText(i int) {
	var currency_copy [9]string
	var currency_copy_num [9]int

	currency_copy[0] = "热门小说"
	currency_copy_num[0] = 4
	currency_copy[1] = "这个小说太精彩了"
	currency_copy_num[1] = 8
	currency_copy[2] = "超过瘾的热门小说"
	currency_copy_num[2] = 8
	currency_copy[3] = "真正的免费小说"
	currency_copy_num[3] = 7
	currency_copy[4] = "过瘾又好看的小说"
	currency_copy_num[4] = 8
	currency_copy[5] = "短篇爽文小说"
	currency_copy_num[5] = 6
	currency_copy[6] = "此书有毒看完上瘾"
	currency_copy_num[6] = 8
	currency_copy[7] = "劲爆小说停不下来"
	currency_copy_num[7] = 8
	currency_copy[8] = "第一章到最后一章"
	currency_copy_num[8] = 8

	// 随机一个文案
	rand.Seed(time.Now().Unix())
	randNum1 := rand.Intn(8)
	s1 := currency_copy[randNum1]
	num1 := currency_copy_num[randNum1]

	// 计算文案的x值，字体64，(720 - num*64) / 2
	x1 := (720 - num1*64) / 2

	/**
	 * 关键信息 距离头部273 字号 64
	 */
	var key_words [3]string
	key_words[0] = "全部免费"
	key_words[1] = "全部0元"
	key_words[2] = "免费阅读"
	s2 := key_words[rand.Intn(3)]
	x2 := (720 - 4*64) / 2

	/**
	 * 辅助信息 距离头部 941 字号 64
	 */
	var assist_copy [7]string
	var assist_copy_num [7]int

	assist_copy[0] = "点击视频链接"
	assist_copy_num[0] = 6
	assist_copy[1] = "快来试试吧"
	assist_copy_num[1] = 5
	assist_copy[2] = "点击下方"
	assist_copy_num[2] = 4
	assist_copy[3] = "大家都在用"
	assist_copy_num[3] = 5
	assist_copy[4] = "怕你停不下来"
	assist_copy_num[4] = 6
	assist_copy[5] = "开始阅读"
	assist_copy_num[5] = 4
	assist_copy[6] = "越看越过瘾"
	assist_copy_num[6] = 5

	// 随机一个文案
	rand.Seed(time.Now().Unix())
	randNum3 := rand.Intn(6)
	s3 := assist_copy[randNum3]
	num3 := assist_copy_num[randNum3]

	// 计算文案的x值，字体64，(720 - num*64) / 2
	x3 := (720 - num3*64) / 2

	/**
	 * 颜色
	 */
	var color [4]string
	color[0] = "#f71414"
	color[1] = "#2d59ff"
	color[2] = "#fce915"
	color[3] = "#24cf13"
	fontcolor1 := color[rand.Intn(3)]

	/**
	 * 字体包
	 */
	rand.Seed(time.Now().Unix())
	var font [5]string
	font[0] = "1.ttc"
	font[1] = "2.ttf"
	font[2] = "3.ttf"
	font[3] = "4.ttf"
	font[4] = "5.ttf"
	fontfile1 := font[rand.Intn(4)]
	fontfile2 := font[rand.Intn(4)]

	cmdArguments := []string{
		"-i", ".\\video\\res\\video_audio.mp4",
		"-vf", "drawtext=fontfile=" + fontfile1 + " :text='" + s1 + "':x=" + fmt.Sprint(x1) + ":y=177:fontsize=64:fontcolor=white:shadowy=2,drawtext=fontfile=" + fontfile2 + " :text='" + s2 + "':x=" + fmt.Sprint(x2) + ":y=273:fontsize=64:fontcolor=" + fontcolor1 + ":shadowy=2,drawtext=fontfile=" + fontfile1 + " :text='" + s3 + "':x=" + fmt.Sprint(x3) + ":y=941:fontsize=64:fontcolor=white:shadowy=2,drawtext=fontfile=SourceHanSansK-Regular.TTF :text='本故事纯属虚构':x=276:y=1232:fontsize=24:fontcolor=white:shadowy=2",
		".\\video\\res\\video_text_" + fmt.Sprint(i) + ".mp4"}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("加文案失败\n")
	} else {
		fmt.Printf("加文案成功\n")
	}
	err_video_audio := os.RemoveAll(".\\video\\res\\video_audio.mp4")
	if err_video_audio != nil {
		fmt.Printf("删除音乐视频失败\n")
	}
}

/**
 * 转码率
 */
func transcodeBit(i int) {
	cmdArguments := []string{
		"-i", ".\\video\\res\\video_text_" + fmt.Sprint(i) + ".mp4",
		"-b:v", "24M",
		".\\video\\res\\video_bit_" + fmt.Sprint(i) + ".mp4"}
	cmd := exec.Command("ffmpeg", cmdArguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("比特率失败\n")
	} else {
		fmt.Printf("比特率成功\n")
	}
	err_video := os.RemoveAll(".\\video\\res\\video_text_" + fmt.Sprint(i) + ".mp4")
	if err_video != nil {
		fmt.Printf("删除文案视频失败\n")
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
