package common

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

/*
获取程序运行路径
*/
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	//return strings.Replace(dir, "\\", "/", -1)
	return dir
}

/*
*
返回程序文件名
*/
func GetExeName() string {
	fullFilename := os.Args[0]
	var filenameWithSuffix string
	filenameWithSuffix = filepath.Base(fullFilename) // 获取文件名带后缀
	var fileSuffix string
	fileSuffix = filepath.Ext(filenameWithSuffix) // 获取文件后缀
	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix) // 获取文件名
	return filenameOnly
}

/*
*
如果返回的错误为nil,说明文件或文件夹存在
如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
如果返回的错误为其它类型,则不确定是否在存在
*/
func IsPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/*
*
ip转byte
*/
func StringIPToByte(ipstring string) []byte {
	ipaddbyte := []byte{0, 0, 0, 0}
	ipSegs := strings.Split(ipstring, ".")
	for i, ipSeg := range ipSegs {
		//fmt.Printf(" %d ipSegs: %s\n",i,ipSeg);
		ipInt, _ := strconv.Atoi(ipSeg)
		ipaddbyte[i] = byte(ipInt)
	}
	return ipaddbyte
}

/*
*
转码到UTF-8
*/
func ConvertToUTF8(src []byte) (dst []byte, err error) {
	dst, err = ioutil.ReadAll(transform.NewReader(bytes.NewReader(src), simplifiedchinese.GB18030.NewDecoder()))
	if err == nil {
		return dst, nil
	}
	return nil, err
}

/*
*
MD5
*/
func EncryptionMd5(password string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	cipherStr := md5Ctx.Sum(nil)
	passString := hex.EncodeToString(cipherStr)
	return passString
}

/*
*
字符串是否为空
*/
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

/*
*
字符串是否纯数字
*/
func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

/*
*
复制文件
src 源
dst 目的
buffer 缓冲区大小
*/
func CopyFile(src, dst string, buffer int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return errors.New(fmt.Sprintf("%s is not a regular file.", src))
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() { _ = source.Close() }()

	_, err = os.Stat(dst)
	if err == nil {
		err = os.Remove(dst)
		if err != nil {
			return errors.New(fmt.Sprintf("File %s already exists.", dst))
		}
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() { _ = destination.Close() }()

	if buffer <= 0 {
		buffer = 1000000
	}
	buf := make([]byte, buffer)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

// 删除下载文件及临时文件
func DeleteFile(parent, path string) error {
	if IsEmpty(path) {
		logx.Info("delete file path empty")
		return nil
	}
	filePath := filepath.Join(parent, path)
	err := os.Remove(filePath) // 删除文件

	tempPath := fmt.Sprintf("%s.temp", filePath)
	errs := os.Remove(tempPath) // 删除临时文件

	if err != nil && errs != nil {
		return err
	}
	return nil
}

// 批量删除文件
func DeleteFiles(filePaths ...string) error {
	var err error
	for _, filePath := range filePaths {
		if IsEmpty(filePath) {
			logx.Info("delete file path empty")
			continue
		}
		err = os.Remove(filePath) // 删除文件
		if err != nil {
			logx.Error(err.Error())
		}
	}
	return err
}

/*
获取文件大小
src 文件路径
*/
func GetFileSize(src string) (int64, error) {
	fi, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	// get the size
	return fi.Size(), nil
}

/*
*
检查IP是否在一个范围内
start 开始IP 192.168.0.1
end 结束IP 192.168.0.254
ip 需检测IP 192.168.200
*/
func IsIPInRange(start, end, ip string) bool {
	ip1 := net.ParseIP(start)
	ip2 := net.ParseIP(end)
	trial := net.ParseIP(ip)
	if trial.To4() == nil {
		return false
	}
	if bytes.Compare(trial, ip1) >= 0 && bytes.Compare(trial, ip2) <= 0 {
		return true
	}
	return false
}

/*
*
查询数组中是否已包含字符串
*/
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// 判断系统位数，64位返回64，32位返回32
func GetBits() (bit int) {
	bit = 32 << (^uint(0) >> 63)
	return
}

// 批量删除map的key
func DeleteMap(value *map[string]interface{}, attrs ...string) {
	for _, v := range attrs {
		delete(*value, v)
	}
}

func HttpRequest(method, url string, body io.Reader, headers map[string]string) ([]byte, error) {
	if headers == nil {
		headers = map[string]string{
			"Content-Type": "application/json",
			"Visitor":      "docker",
		}
	}
	//生成client 参数为默认
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}, Timeout: 20 * time.Second,
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		logx.Errorf("request Error: %v", err)
		return nil, err
	}
	// 设置请求头
	for key, info := range headers {
		req.Header.Set(key, info)
	}
	//处理返回结果
	resp, err := client.Do(req)
	if err != nil {
		logx.Errorf("request Error: %v", err)
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)

	// 关闭
	if resp.StatusCode != http.StatusOK {
		err := errors.New(fmt.Sprintf("request error:%s", data))
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
		_ = resp.Close
		_ = req.Close
		client.CloseIdleConnections()
	}()
	return data, err
}

// 压缩文件 filesPath  文件的路径  outPath 压缩文件存放地址
func CompressFile(outPath string, filesPath ...string) (err error) {
	if err = os.MkdirAll(filepath.Dir(outPath), os.ModePerm); err != nil {
		logx.Error(err.Error())
		return err
	}
	outFile, err := os.Create(outPath)
	if err != nil {
		logx.Error(err.Error())
		return err
	}
	defer outFile.Close()
	zipFile := zip.NewWriter(outFile) //转换为 zip 的 流
	defer zipFile.Close()

	for _, path := range filesPath {
		fileInfo, err := os.Stat(path)
		if err != nil {
			logx.Error(err.Error())
			return err
		}
		if fileInfo.IsDir() { //如果是文件夹
			//获取除了文件名称以外的全部路径
			index := strings.LastIndex(path, fileInfo.Name())
			//获取文件夹下的东西
			err = filepath.Walk(path, func(absolutePath string, value os.FileInfo, err error) error {
				innerFileInfo, err := os.Stat(absolutePath)
				if err != nil {
					logx.Error(err.Error())
					return err
				}
				header, err := zip.FileInfoHeader(innerFileInfo) //copy文件信息转化为zip的专属对象信息
				if err != nil {
					return err
				}
				header.Name = absolutePath[index:] //设置文件夹或文件在zip中的路径
				if value.IsDir() {
					header.Name = header.Name + "/"
				}
				innerWriter, err := zipFile.CreateHeader(header) //在压缩文件中创建文件
				if !value.IsDir() {
					fileOs, err := os.Open(absolutePath)
					if err != nil {
						return err
					}
					if _, err = io.Copy(innerWriter, fileOs); err != nil {
						fileOs.Close()
						return err
					}
					fileOs.Close()
				}
				return nil
			})
			if err != nil {
				logx.Error(err.Error())
				return err
			}
		} else {
			header, err := zip.FileInfoHeader(fileInfo) //copy文件信息转化为zip的专属对象信息
			if err != nil {
				return err
			}
			writer, err := zipFile.CreateHeader(header) //在压缩文件中创建文件
			if err != nil {
				return err
			}
			fileOs, err := os.Open(path)
			if err != nil {
				return err
			}
			if _, err = io.Copy(writer, fileOs); err != nil {
				fileOs.Close()
				return err
			}
			fileOs.Close()
		}
	}
	return err
}

// 判断Strings数组内是否包含 target 内容
func IsAvailableString(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

// 判断Strings数组内是否包含 target 内容
func IsAvailableInt64(target int64, str_array []int64) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

// 返回Basic鉴权

func NewBasicAuth(username, password string) string {
	return fmt.Sprintf("Basic %s", basicAuth(username, password))
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

/*
获取年月 有多少天
*/
func GetYearMonthToday(year int, month int) int {
	//有31天的月份
	day31 := map[int]bool{
		1:  true,
		3:  true,
		5:  true,
		7:  true,
		8:  true,
		10: true,
		12: true,
	}
	if day31[month] == true {
		return 31
	}
	// 有30天的月份

	day30 := map[int]bool{
		4:  true,
		6:  true,
		9:  true,
		11: true,
	}
	if day30[month] == true {
		return 30
	}
	//计算平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出二月天数
		return 29
	}
	// 得出平年二月天数
	return 28
}

/*
将数据填充为0 满足位数要求
*/
func TimeSupplement(data interface{}, length int) string {
	// 将data 转换为string

	dataString := fmt.Sprintf("%v", data)

	// 获取dataString长度
	dataLen := len(dataString)

	// 获取dataString长度  和目标长度差值
	differencelength := length - dataLen

	// 判断如果有差值并且大于1 对数据进行填充
	if differencelength > 0 {
		// 填充string
		var fillString string
		for i := 0; i < differencelength; i++ {
			fillString = fmt.Sprintf("0%v", fillString)
		}
		//将填充数据添加到原来DataString种
		dataString = fillString + dataString
	}

	return dataString
}
