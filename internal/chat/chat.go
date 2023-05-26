package chat

import (
	"fmt"

	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

func ConsoleQrCode(uuid string) {
	q, _ := qrcode.New("https://login.weixin.qq.com/l/"+uuid, qrcode.Low)
	fmt.Println(q.ToString(true))
}

func main() {
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式
	//bot := openwechat.DefaultBot()

	// 注册登陆二维码回调
	bot.UUIDCallback = ConsoleQrCode

	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	defer reloadStorage.Close()
	if err := bot.PushLogin(reloadStorage, openwechat.NewRetryLoginOption()); err != nil {
		fmt.Println("ERR push login", err)
	}

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			msg.ReplyText("pong")
		}
	}
	//bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	//// 登陆
	//if err := bot.Login(); err != nil {
	//fmt.Println(err)
	//return
	//}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的好友
	friends, err := self.Friends()
	if err != nil {
		fmt.Println("Get Friends err", err)
		return
	}
	//fmt.Printf("%+v\n", friends[0])
	for _, fr := range friends {
		//fmt.Printf("%+v\n", fr)
		name := fr.NickName
		if fr.RemarkName != "" {
			name = fr.RemarkName
		}
		fmt.Println("----", name)
	}

	// 获取所有的群组
	//groups, err := self.Groups()
	//fmt.Println(groups, err)

	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
