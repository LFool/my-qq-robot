package main

import (
	"context"
	"demo/dict"
	"demo/event"
	m_botgo "demo/m-botgo"
	"demo/m-botgo/message"
	m_dto "demo/m-dto"
	"demo/m-token"
	"demo/websocket"
	"fmt"
	"math/rand"
	"os"

	"gopkg.in/yaml.v2"
	"log"
	"path"
	"runtime"
	"strings"
	"time"
)

var jokeList = dict.GetJokeAll()

func GetOneJoke() string {
	rand.Seed(time.Now().UnixNano())
	return jokeList[rand.Intn(len(jokeList))][1]
}

func main() {
	configName := "robot-config.yaml"
	appId, token, err := getConfigInfo(configName)
	if err != nil {
		log.Fatal(err)
	}
	botToken := m_token.BotToken(appId, token)
	api := m_botgo.NewOpenAPI(botToken).WithTimeout(30 * time.Second)
	ctx := context.Background()
	ws, err := api.WS(ctx, nil, "") //websocket
	if err != nil {
		log.Fatalln("websocket 错误， err = ", err)
	}

	m := make(map[string]int)

	var atMessage event.ATMessageEventHandler = func(event *m_dto.WSPayload, data *m_dto.WSATMessageData) error {

		if strings.Contains(data.Content, "打卡") { // 如果@机器人并输入「打卡」
			id := data.Author.ID
			username := data.Author.Username
			if _, ok := m[id]; ok {
				_, err := api.PostMessage(ctx, data.ChannelID, &m_dto.MessageToCreate{MsgID: data.ID,
					Content: username + "，你已经打过卡了哦～",
				})
				if err != nil {
					return err
				}
			} else {
				m[id] = 1
				_, err := api.PostMessage(ctx, data.ChannelID, &m_dto.MessageToCreate{MsgID: data.ID,
					Content: "你好 " + username + "，打卡成功～～",
				})
				if err != nil {
					return err
				}
			}
		} else if strings.Contains(data.Content, "笑话") { // 如果@机器人并输入「讲个笑话」，则回复「一个笑话」
			_, err := api.PostMessage(ctx, data.ChannelID, &m_dto.MessageToCreate{MsgID: data.ID,
				Content: GetOneJoke() + message.Emoji(20),
			})
			if err != nil {
				return err
			}
		}
		return nil
	}

	intent := websocket.RegisterHandlers(atMessage) // 注册socket消息处理
	err = m_botgo.NewSessionManager().Start(ws, botToken, &intent)
	if err != nil {
		return
	} // 启动socket监听
}

// 获取配置文件中的信息
func getConfigInfo(fileName string) (uint64, string, error) {
	// 获取当前 go 程调用栈所执行的函数的文件和行号信息
	// 忽略 pc 和 line
	_, filePath, _, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("runtime.Caller(1) 读取失败")
	}
	file := fmt.Sprintf("%s/%s", path.Dir(filePath), fileName)
	var conf struct {
		AppID uint64 `yaml:"appid"`
		Token string `yaml:"token"`
	}
	content, err := os.ReadFile(file)
	if err != nil {
		log.Print("os.ReadFile() 读取失败")
		return 0, "", err
	}
	if err = yaml.Unmarshal(content, &conf); err != nil {
		log.Print("yaml.Unmarshal(content, &conf) 读取失败")
		return 0, "", err
	}
	return conf.AppID, conf.Token, nil
}
