//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 知乎：https://www.zhihu.com/people/shirdonl
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package cache

import (
	"github.com/garyburd/redigo/redis"
)

//初始化Redis连接池
func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   10000,
		MaxActive: 12000, //最大连接数
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

const LuaScript = `
        local goods_key = KEYS[1]
        local goods_total_key = ARGV[1]
        local goods_sold_key = ARGV[2]
        local goods_total_number = tonumber(redis.call('HGET', goods_key, goods_total_key))
        local goods_sold_number = tonumber(redis.call('HGET', goods_key, goods_sold_key))
		-- 查看是否还有剩余的商品,增加订单数量,返回结果值
        if(goods_total_number >= goods_sold_number) then
            return redis.call('HINCRBY', goods_key, goods_sold_key, 1)
        end
        return 0
`

//远程订单存储健值
type RemoteCacheKeys struct {
	SpikeOrderHashKey  string //Redis中秒杀订单hash结构Key
	TotalInventoryKey  string //Hash结构中总订单库存Key
	QuantityOfOrderKey string //Hash结构中已有订单数量Key
}

//远端统一扣库存
func (RemoteCacheKeys *RemoteCacheKeys) DeductStock(conn redis.Conn) bool {
	lua := redis.NewScript(1, LuaScript)
	result, err := redis.Int(lua.Do(conn,
		RemoteCacheKeys.SpikeOrderHashKey,
		RemoteCacheKeys.TotalInventoryKey,
		RemoteCacheKeys.QuantityOfOrderKey))
	if err != nil {
		return false
	}
	return result != 0
}
