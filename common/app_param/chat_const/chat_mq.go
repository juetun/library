package chat_const

import "fmt"

const (
	redisMqTopicTmp = "chat:mq:%v" //
)

//获取聊天系统监听的队列参数,分布式系统使用
func GetRedisMqTopicTmp(topicNames ...string) (res string) {
	var topicName string

	//TODO 当前共用一个，后续可在此实现分桶逻辑
	//if len(topicNames) > 0 {
	//	topicName = topicNames[0]
	//}
	res = fmt.Sprintf(redisMqTopicTmp, topicName)
	return
}
