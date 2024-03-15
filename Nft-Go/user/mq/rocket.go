/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mq

import (
	"Nft-Go/global"
	"Nft-Go/user/internal/logic"
	"Nft-Go/user/internal/model"
	"context"
	"dubbo.apache.org/dubbo-go/v3/logger/zap"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/dubbogo/gost/log/logger"
	"github.com/spf13/viper"
	"os"
)

func InitMq() {
	rlog.SetLogger(&ZapLogger{})
	// 设置推送消费者
	c, _ := rocketmq.NewPushConsumer(
		//消费组
		consumer.WithGroupName(viper.Get("rocketmq.group").(string)),
		// namesrv地址
		consumer.WithNameServer([]string{viper.Get("rocketmq.nameserver").(string)}),
		consumer.WithConsumerModel(consumer.BroadCasting),
	)
	// 必须先在 开始前
	err := c.Subscribe("Nft-Go", consumer.MessageSelector{}, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range ext {
			logger.Info("Rocketmq Received:%v \n", ext[i])
			json := global.GetFastJson()
			data, err := json.ParseBytes(ext[i].Body)
			switch ext[i].GetTags() {
			case "createUser":
				if err != nil {
					return 0, err
				}
				err = logic.SaveNotice(model.Notice{
					Title:       data.Get("title").String(),
					Description: data.Get("description").String(),
					PublishTime: data.Get("publishTime").String(),
					UserAddress: data.Get("userAddress").String(),
					Type:        data.Get("type").GetInt(),
				})
				if err != nil {
					return 0, err
				}
				break
			default:
				break
			}
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		logger.Info(err.Error())
	}
	err = c.Start()
	if err != nil {
		logger.Info(err.Error())
		os.Exit(-1)
	}
	logger.Info("mq connect success")

}

type ZapLogger struct {
	logger *zap.Logger
}

func (l *ZapLogger) Debug(msg string, fields map[string]interface{}) {
	logger.Info(msg, fields)
}

func (l *ZapLogger) Info(msg string, fields map[string]interface{}) {
	logger.Info(msg, fields)
}

func (l *ZapLogger) Warning(msg string, fields map[string]interface{}) {
	logger.Warn(msg, fields)
}

func (l *ZapLogger) Error(msg string, fields map[string]interface{}) {
	logger.Error(msg, fields)
}

func (l *ZapLogger) Fatal(msg string, fields map[string]interface{}) {
	logger.Fatal(msg, fields)
}

func (l *ZapLogger) Level(level string) {
	logger.SetLoggerLevel(level)
}

func (l *ZapLogger) OutputPath(path string) (err error) {
	logger.Info("output path is not supported")
	return nil
}