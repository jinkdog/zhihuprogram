package cookie

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type (
	Cookie struct {
		Secret string
		Opt    Option //引入option结构体并赋予变量名
	}

	Option struct {
		Config http.Cookie  //引入http包里的Cookie结构体，赋予了变量名
		Ctx    *gin.Context //赋予了变量名
	}

	Config struct {
		Secret      string
		http.Cookie //直接表示引入http包里的Cookie结构体，没有赋予变量名
	}
)

func NewCookieWriter(secret string, opt ...Option) *Cookie {
	//opt为Option结构体的数组
	if len(opt) == 0 {
		return &Cookie{Secret: secret}
		//如果长度为0就返回secret
	} else {
		return &Cookie{
			Secret: secret,
			Opt:    opt[0],
		} //如果长度不为0就返回序号为0的第一个option
	}
}

// Set 写入数据的方法
func (c *Cookie) Set(key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	setSecureCookie(c, key, string(bytes)) //设置安全的cookie
}

// Get 获取数据的方法
func (c *Cookie) Get(key string, obj interface{}) bool {
	tempData, ok := getSecureCookie(c, key) //获得安全的cookie
	if !ok {
		return false
	}
	_ = json.Unmarshal([]byte(tempData), obj)
	return true
}

// Remove 删除数据的方法
func (c *Cookie) Remove(key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	setSecureCookie(c, key, string(bytes))
}

func setSecureCookie(c *Cookie, name, value string) { //设置安全的Cookie
	vs := base64.URLEncoding.EncodeToString([]byte(value)) //设置编码
	//将value的字符串切入到字符数组中，再由EncodeToString（[]byte）重新加密为字符串
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	//将现在的时间转化为 Unix time类型为int
	//FormatInt函数则将Unix time按照base的进制来转化为string类型数据base10就是十进制
	h := hmac.New(sha256.New, []byte(c.Secret))
	//返回一个hash类型数据，传入编码方式和字符组数据
	_, _ = fmt.Fprintf(h, "%s%s", vs, timestamp)
	//func Fprintf(w io.Writer, format string, a ...any) (n int, err error)
	//按照format格式向h写入vs和timestamp数据
	sig := fmt.Sprintf("%02x", h.Sum(nil))
	//按照指定的format格式将h输出为string
	//其中format的格式“%02x”的意思为
	//X 表示以十六进制形式输出
	//02 表示不足两位，前面补0输出；
	cookie := strings.Join([]string{vs, timestamp, sig}, "|")
	//Join的作用：传入string类型切片，最后输出用sep符号（在这里是“|”）隔开的字符串

	http.SetCookie(c.Opt.Ctx.Writer, //传入c结构体中包含的Opt结构体中的Ctx（实际上是gin.context）中的Writer（实质上是 ResponseWriter类型数据）
		&http.Cookie{
			Name:     name,
			Value:    cookie,
			MaxAge:   c.Opt.Config.MaxAge, //配置的最大存在时间
			Path:     "/",
			Domain:   c.Opt.Config.Domain, //配置的存在的域名
			SameSite: http.SameSite(1),
			Secure:   c.Opt.Config.Secure,
			HttpOnly: c.Opt.Config.HttpOnly,
		})
}

func getSecureCookie(c *Cookie, key string) (string, bool) { //获取安全Cookie
	cookie, err := c.Opt.Ctx.Request.Cookie(key)
	//定义为Cookie中的opt结构体中的ctx（gin.context）
	//gin.context中调用 Request      *http.Request的数据
	//Cookie又是Request的方法
	//通过传入key值返回http包中的Cookie类型结构体
	if err != nil {
		return "", false
	} //如果没有找到返回空值和false
	val, err := url.QueryUnescape(cookie.Value)
	//将http.Cookie的Value反向转化为
	//将每个3字节的"%AB "形式的编码子串转换为16进制的0xAB字节。如果任何%后面没有两个十六进制数字，它将返回一个错误。
	if val == "" || err != nil {
		//如果val值获取到了但是为空，即使err=nil也返回false
		//没有获取到则直接返回nil
		return "", false
	}

	parts := strings.SplitN(val, "|", 3)
	//将val分割成3个子串，最后一个字串包含为分割的全部
	//以出现一次“|”为分割点，最后返回一个string类型的切片
	if len(parts) != 3 { //如果最后的切片长度不为3
		return "", false
	}

	vs := parts[0]        //为字符型切片的第一个部分
	timestamp := parts[1] //为字符型切片的第二个部分
	sig := parts[2]       //为字符型切片的第三个部分

	h := hmac.New(sha256.New, []byte(c.Secret))
	//将c.Secret强制转换为byte字符切片
	//最后传入到函数中生成Hash类型
	_, _ = fmt.Fprintf(h, "%s%s", vs, timestamp)
	//以format格式将vs和timestamp输入到h中

	if fmt.Sprintf("%02x", h.Sum(nil)) != sig {
		return "", false
	}
	//和setSecureCookie的部分对应看看h.Sum返回的切片按照format的格式输出获得最终的结果为是否为sig
	res, _ := base64.URLEncoding.DecodeString(vs) //解开编码为字符型数组
	return string(res), true                      //强制转换res为string
}
