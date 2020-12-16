/*
 * @Author: qiuling
 * @Date: 2019-08-21 16:04:56
 * @Last Modified by: qiuling
 * @Last Modified time: 2019-12-04 16:13:38
 */
package bizs

import (
	"github.com/wlxpkg/base/amqp"
	"github.com/wlxpkg/base/beanstalk"
	"github.com/wlxpkg/base/log"
)

// 消息
var amqpProducer *amqp.Producer
var btProducer *beanstalk.Producer

func init() {
	var err error
	amqpProducer = amqp.NewProducer()
	btProducer, err = beanstalk.NewProducer()

	if err != nil {
		log.Err(err)
	}
}

// 发送 amqo 消息
func publish(topic string, data map[string]interface{}) error {
	return amqpProducer.Publish(topic, data)
}

// 发送 beanstalk 延迟消息
func delay(topic string, data map[string]interface{}, delay int64) (uint64, error) {
	return btProducer.Publish(topic, data, delay)
}
