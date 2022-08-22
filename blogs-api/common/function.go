package common

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"git.oa00.com/go/mysql"
	"git.oa00.com/go/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"kkblogs/blogs-api/constant"
	"kkblogs/blogs-api/lib/bean"
	"kkblogs/blogs-api/model"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// @Title 随机数种子
func init() {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
}

// MD5 @Title md5加密
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// RandStr @Title 生成随机数
func RandStr(n int, str ...string) string {
	s := "abcfghijklmnopqwedasdrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	if len(str) > 0 {
		s = str[0]
	}
	res := ""
	for i := 0; i < n; i++ {
		res += string(s[rand.Intn(len(s))])
	}
	return res
}

// GetVerErr @Title 处理验证错误消息
func GetVerErr(err error) string {
	er, ok := err.(validator.ValidationErrors)
	if ok {
		field := er[0].Field()
		if field == er[0].StructField() {
			field = strings.ToLower(field[0:1]) + field[1:]
		}
		switch er[0].Tag() {
		case "required":
			return field + "不能为空"
		case "min":
			if er[0].Type().String() == "string" {
				return field + "不能小于" + er[0].Param() + "位"
			}
			return field + "不能小于" + er[0].Param()
		case "gte":
			if er[0].Type().String() == "string" {
				return field + "不能小于" + er[0].Param() + "位"
			}
			return field + "不能小于" + er[0].Param()
		}
		return field + "错误"
	} else {
		return err.Error()
		//return "参数格式错误"
	}
}

// InArray @Title 是否在数组中
func InArray(need interface{}, data interface{}) bool {
	if datas, ok := data.([]int); ok {
		for _, item := range datas {
			if item == need {
				return true
			}
		}
	}
	if datas, ok := data.([]uint); ok {
		for _, item := range datas {
			if item == need {
				return true
			}
		}
	}
	if datas, ok := data.([]string); ok {
		for _, item := range datas {
			if item == need {
				return true
			}
		}
	}
	return false
}

// CreateDateDir 根据当前日期来创建文件夹
func CreateDateDir(Path string) string {
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(Path, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0755)
	}
	return folderName
}

// DownloadFormUrl @Title 下载文件
func DownloadFormUrl(src string, filename string) error {
	res, err := http.Get(src)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, bytes.NewReader(body)); err != nil {
		return err
	}
	return nil
}

// Utf8ToGbk @Title utf8转gbk
func Utf8ToGbk(source string) string {
	result, _, _ := transform.String(simplifiedchinese.GBK.NewEncoder(), source)
	return result
}

// GbkToUtf8 @Title gbk转utf8
func GbkToUtf8(source string) string {
	result, _, _ := transform.String(simplifiedchinese.GBK.NewDecoder(), source)
	return result
}

// GetPassword @Title 密码加密
func GetPassword(password, salt string) string {
	return MD5(salt + MD5(password+salt))
}

// GetDeviceType @Title 获取设备类型
func GetDeviceType(c *gin.Context) uint {
	switch strings.ToLower(c.GetHeader("Device-Type")) {
	case "android":
		return constant.DeviceTypeAndroid
	case "ios":
		return constant.DeviceTypeIos
	default:
		return constant.DeviceTypeWeb
	}
}

// IsPhone @Title 是否手机号
func IsPhone(phone string) bool {
	return regexp.MustCompile(`^1\d{10}$`).MatchString(phone)
}

// GetAddress @Title 返回具体地址
func GetAddress(province, city, county, address string) string {
	if InArray(city, []string{"县", "市辖区", "自治区直辖县级行政区划", "省直辖县级行政区划"}) {
		city = ""
	}
	if InArray(county, []string{"市中区", "市辖区"}) {
		county = ""
	}
	return province + city + county + address
}

// GetCity @Title 返回城市
func GetCity(province, city string) string {
	if InArray(city, []string{"县", "市辖区", "自治区直辖县级行政区划", "省直辖县级行政区划"}) {
		return province
	}
	return city
}

// FindParents @Title 查询所有上级
func FindParents(tableName, id, parentId string, findId uint, result interface{}) {
	mysql.Db.Raw(fmt.Sprintf(`SELECT T2.* 
FROM ( 
    SELECT 
        @r AS _id, 
        (SELECT @r := %s FROM %s WHERE %s = _id)
    FROM 
        (SELECT @r := ?) vars, 
        %s h 
    WHERE @r <> 0) T1 
JOIN %s T2 
ON T1._id = T2.%s`, parentId, tableName, id, tableName, tableName, id), findId).Scan(result)
	return
}

// GenNo @Title 生成编号
func GenNo(NoPrefix string) (string, error) {
	date := time.Now().Format("20060102")
	redisKey := fmt.Sprintf("gen:%s:%s", NoPrefix, date)
	result, err := redis.Redis.Incr(redisKey).Result()
	if err != nil {
		return "", errors.New("生成失败")
	}
	redis.Redis.Expire(redisKey, time.Hour*24)
	if result > 9999 {
		return "", errors.New("生成失败")
	}
	return fmt.Sprintf("%s%s%.4d", NoPrefix, date, result), nil
}

// GenNoWaybill @Title 生成编号
func GenNoWaybill(NoPrefix string) (string, error) {
	//date := time.Now().Format("2006")
	redisKey := fmt.Sprintf("gen:%s", NoPrefix)
	result, err := redis.Redis.Incr(redisKey).Result()
	if err != nil {
		return "", errors.New("生成失败")
	}
	if result > 999999999 {
		return "", errors.New("生成失败")
	}
	return fmt.Sprintf("%s%.9d", NoPrefix, result), nil
}

// ExcelRows @Title 获取excel数据
//func ExcelRows(filename string) (rows [][]string, err error) {
//	file, err := excelize.OpenFile(filepath.Join(config.Config.Upload.Root, filename))
//	if err != nil {
//		return rows, err
//	}
//	getRows, err := file.GetRows(file.GetSheetName(0))
//	if err != nil {
//		return nil, err
//	}
//	return getRows, nil
//}

// GetRequestBody @Title 获取请求数据
func GetRequestBody(request *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("read request body err: %v", err)
	}

	_ = request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return body, nil
}

// Ternary @Title 三元运算
func Ternary(condition bool, resultA, resultB interface{}) interface{} {
	if condition {
		return resultA
	}
	return resultB
}

var timeTemplates = []string{
	"2006-01-02 15:04:05", //常规类型
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
	"15:04:05",
}

/* 时间格式字符串转换 */
func TimeStringToGoTime(tm string) time.Time {
	for i := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
		if nil == err && !t.IsZero() {
			return t
		}
	}
	return time.Time{}
}

/* 定时同步数据库 (阅读量) */
func SyncDataBaseView() {
	key := fmt.Sprintf("%s%s", bean.ViewCommentCount, "*")
	keys := redis.Redis.Keys(key)
	var article model.Article
	for _, key := range keys.Val() {
		id := strings.Split(key, ":")[1]
		articleId, _ := strconv.ParseInt(id, 10, 64)
		val := redis.Redis.Get(key).Val()
		count, _ := strconv.Atoi(val)
		article = model.Article{Id: articleId, ViewCounts: uint(count)}
		if mysql.Db.Model(&article).Updates(&article).RowsAffected <= 0 {
			log.Println("更新失败")
		}
	}
}

/* 定时同步数据库 (评论数量) */
func SyncDataBaseComment() {
	key := fmt.Sprintf("%s%s", bean.CommentCount, "*")
	keys := redis.Redis.Keys(key)
	var article model.Article
	for _, key := range keys.Val() {
		id := strings.Split(key, ":")[1]
		articleId, _ := strconv.ParseInt(id, 10, 64)
		val := redis.Redis.Get(key).Val()
		count, _ := strconv.Atoi(val)
		article = model.Article{Id: articleId, CommentCounts: uint(count)}
		if mysql.Db.Model(&article).Updates(&article).RowsAffected <= 0 {
			log.Println("更新失败")
		}
	}
}
