// @File:  jd_check_cookies.go
// @Time:  2022/1/6 6:25 PM
// @Author: ClassmateLin
// @Email: classmatelin.site@gmail.com
// @Site: https://www.classmatelin.top
// @Description:
// @Cron: 0 */4 * * *

package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"scripts/config"
	"scripts/config/jd"
	"scripts/global"
	"scripts/structs"
)

// isValidCookie
// @description: 判断cookie是否过期
// @param:  cookie
// @return: bool
func isValidCookie(cookie string) bool {
	client := resty.New()
	url := "https://api.m.jd.com/client.action?functionId=newUserInfo&clientVersion=10.0.9&client=android&openudid=a27b83d3d1dba1cc&uuid=a27b83d3d1dba1cc&aid=a27b83d3d1dba1cc&area=19_1601_36953_50397&st=1626848394828&sign=447ffd52c08f0c8cca47ebce71579283&sv=101&body=%7B%22flag%22%3A%22nickname%22%2C%22fromSource%22%3A1%2C%22sourceLevel%22%3A1%7D&"

	resp, err := client.R().
		SetHeader("user-agent", global.GetJdUserAgent()).
		SetHeader("cookie", cookie).
		Get(url)

	if err != nil {
		return false
	}
	code := gjson.Get(resp.String(), "code").Int()
	if code != 0 {
		return false
	}
	return true
}

func main() {

	hasInvalid := false // 是否有过期的cookies

	title := "Cookies过期通知\n"

	message := "# 账号列表:\n"

	for _, user := range jd.UserList {
		isValid := isValidCookie(user.CookieStr)
		if !isValid {
			message += " - " + user.Username + "\n"
			hasInvalid = true
		}
	}

	if hasInvalid {
		fmt.Println(message)
		structs.Notify(config.VP, title, message)
	} else {
		fmt.Println("暂无过期账号...")
	}

}
