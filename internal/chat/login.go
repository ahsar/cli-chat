package chat

import (
	"log"

	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

var (
	bot  *openwechat.Bot
	self *openwechat.Self
	err  error
)

// 登录wechat
func Login() {
	log.Println("now login")
	if bot != nil && bot.Alive() {
		return
	}
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
	//bot.MessageHandler = func(msg *openwechat.Message) {
	//if msg.IsText() && msg.Content == "ping" {
	//msg.ReplyText("pong")
	//}
	//}

	// 获取登陆的用户
	self, err = bot.GetCurrentUser()
	if err != nil {
		log.Println(err)
		return
	}

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	//go bot.Block()
}

func Logout() {
	bot.Logout()
}

func consoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	log.Println(q.ToString(true))
}
