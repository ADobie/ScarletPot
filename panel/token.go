package panel

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math/big"
	"scarletpot/utils/md5"
	"time"
)

// TODO： 生成api token相关

// 获取用户jwt-token 通过decode jwt-token 使用其中的 username 和 userToken 来生成apiID
// 开发测试环境先直接获取用户发送的username 和 userToken 来生成,但是会有被抓包修改的安全风险，所以后期采用jwtToken传输信息
// 在生成api秘钥之前 首先需要验证用户传输的username 和 userToken 是否匹配
type GenMsg struct {
	Timestamp int64  `json:"timestamp"`
	Username  string `json:"username"`
	UserToken string `json:"userToken"`
}

type GenRetMsg struct {
	APIId  string `json:"apiId"`
	APISec string `json:"apiSec"`
	Msg    string `json:"msg"`
}

func (s *Service) genApiToken(c *gin.Context) (int, interface{}) {
	// apiId 用户名+获取的时间戳 16位md5  apiSecret 随机生成的字符串
	var dataGen GenMsg
	err := c.BindJSON(&dataGen)
	if err != nil {
		return s.errJSON(500, 10000, "JSON解析失败")
	}
	id := genApiId(dataGen.Username, dataGen.UserToken)
	sec := createRandomString(32)

	ret, _ := json.Marshal(GenRetMsg{
		APIId:  id,
		APISec: sec,
		Msg:    "api秘钥生成成功",
	})

	return s.successJSON(ret)
}

func genApiId(username string, userToken string) string {
	apiId := md5.Md5(username + userToken + time.Now().String())
	return apiId
}

func createRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$_"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
