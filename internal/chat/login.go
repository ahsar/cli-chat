package chat

import (
	"fmt"
	"log"

	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

var (
	bot   *openwechat.Bot
	self  *openwechat.Self
	msgch chan *openwechat.Message
	err   error
)

// 登录wechat
func Init(ch chan *openwechat.Message) {
	log.Println("now login")
	if bot != nil && bot.Alive() {
		return
	}
	msgch = ch

	bot = openwechat.DefaultBot(openwechat.Desktop) // 桌面模式

	// 注册登陆二维码回调
	bot.UUIDCallback = consoleQrCode

	// 免重复扫描
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	if err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		log.Println("ERR push login", err)
	}

	// 注册消息处理函数
	bot.MessageHandler = messageHandler

	// 获取登陆的用户
	self, err = bot.GetCurrentUser()
	if err != nil {
		log.Println(err)
		return
	}

	bot.MessageHandler = messageHandler
	go func() {
		if r := bot.Block(); r != nil {
			fmt.Println("wechat 已退出", r)
		}
	}()
}

func Logout() {
	if bot == nil || !bot.Alive() {
		return
	}
	bot.Logout()
}

func consoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}
