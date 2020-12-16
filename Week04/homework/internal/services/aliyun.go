/*
 * @Author: qiuling
 * @Date: 2019-05-08 11:26:58
 * @Last Modified by: qiuling
 * @Last Modified time: 2019-12-04 16:13:38
 */

package services

import (
	"errors"
	"io"
	"time"

	. "github.com/wlxpkg/base"
	"github.com/wlxpkg/base/log"

	"github.com/gookit/config/v2"
	"github.com/ql2005/aliyun-sts-go-sdk/sts"

	// "strconv"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"hash"
)

const (
	// region      string = "cn-hangzhou"
	endpoint   = "https://oss-cn-hangzhou.aliyuncs.com"
	expireTime = 3600
	// base64Table string = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var dirMap = map[string]string{
	"tools":     "tools/",
	"course":    "course/",
	"user":      "user/",
	"discovery": "discovery/",
	"game":      "game/",
	"shop":      "shop/",
	"grant":     "grant/",
	"common":    "common/",
	"message":   "message/",
}

var accessKeyId string
var accessKeySecret string
var host string
var bucket string
var accessSTSId string
var accessSTSSecret string
var roleArn string

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	// Callback    string `json:"callback"`
}

type StsToken struct {
	AccessKeyId     string   `json:"accessid"`
	AccessKeySecret string   `json:"access_secret"`
	SecurityToken   string   `json:"security_token"`
	Bucket          string   `json:"bucket_name"`
	Dir             []string `json:"upload_path"`
	Endpoint        string   `json:"endpoint"`
	Expire          int64    `json:"expire"`
}

type Aliyun struct {
}

func init() {
	accessKeyId = config.GetEnv("OSS_ID")
	accessKeySecret = config.GetEnv("OSS_SECRET")
	host = config.GetEnv("OSS_HOST")
	bucket = config.GetEnv("OSS_BUCKET")
	accessSTSId = config.GetEnv("OSS_STS_ID")
	accessSTSSecret = config.GetEnv("OSS_STS_SECRET")
	roleArn = config.GetEnv("OSS_STS_ARN")
}

// GetSign web 直传签名
func (a *Aliyun) GetSign(dir string) (policyToken PolicyToken, err error) {

	uploadDir := dirMap[dir]
	if uploadDir == "" {
		err = errors.New("ERR_PARAM")
		return
	}

	now := time.Now().Unix()
	expireEnd := now + expireTime
	var tokenExpire = getGmtIso8601(expireEnd)

	//create post policy json
	var config ConfigStruct
	config.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, uploadDir)
	config.Conditions = append(config.Conditions, condition)

	//calucate signature
	result, err := json.Marshal(config)
	if err != nil {
		log.Warn(err)
		err = errors.New("ERR_UNKNOW_ERROR")
		return
	}

	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(accessKeySecret))
	_, _ = io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	/* var callbackParam CallbackParam
	   callbackParam.CallbackUrl = callbackUrl
	   callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	   callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
	   callback_str,err:=json.Marshal(callbackParam)
	   if err != nil {
	       fmt.Println("callback json err:", err)
	   }
	   callbackBase64 := base64.StdEncoding.EncodeToString(callback_str) */
	policyToken = PolicyToken{
		accessKeyId,
		host,
		expireEnd,
		string(signedStr),
		string(debyte),
		uploadDir,
	}
	return
}

func getGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

/* func getUnixTime(time Time) {
	const TIME_LAYOUT = "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	start_time, _ := time.ParseInLocation(TIME_LAYOUT, AdModel.Stime, loc)
} */

// OssToken 通过 sts 认证获取客户端 SecurityToken
func (a *Aliyun) OssToken() (stsToken StsToken, err error) {
	stsClient := sts.NewClient(accessSTSId, accessSTSSecret, roleArn, "sts_client")

	resp, err := stsClient.AssumeRole(uint(expireTime))
	if err != nil {
		log.Err(err)
		err = errors.New("ERR_UNKNOW_ERROR")
		return
	}

	// R(resp.Credentials, "resp")
	var dirs []string

	for _, dir := range dirMap {
		dirs = append(dirs, dir)
	}

	stsToken = StsToken{
		AccessKeyId:     resp.Credentials.AccessKeyId,
		AccessKeySecret: resp.Credentials.AccessKeySecret,
		SecurityToken:   resp.Credentials.SecurityToken,
		Bucket:          bucket,
		Dir:             dirs,
		Endpoint:        endpoint,
		Expire:          Time2Unix(resp.Credentials.Expiration),
	}

	return
}
