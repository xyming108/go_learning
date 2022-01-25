package consumer

import (
	"GO_Learning/13_project/01_go_excel_batch/service"
	"GO_Learning/13_project/01_go_excel_batch/types"
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/tal-tech/go-queue/dq"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

func Consumer(conf *viper.Viper) {
	consumer := dq.NewConsumer(dq.DqConf{
		Beanstalks: []dq.Beanstalk{
			{
				Endpoint: conf.GetString("Beanstalks.Endpoint1"),
				Tube:     "tube",
			},
			{
				Endpoint: conf.GetString("Beanstalks.Endpoint2"),
				Tube:     "tube",
			},
		},
		Redis: redis.RedisConf{
			Host: conf.GetString("Beanstalks.RedisHost"),
			Type: redis.NodeType,
		},
	})

	var result types.TransferData

	consumer.Consume(func(body []byte) {
		err := json.Unmarshal(body, &result)
		if err != nil {
			return
		}

		service.ExcelOperation(result)
	})
}
