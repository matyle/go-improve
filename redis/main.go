package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

// 定义一组常量
const (
	REDIS_IP   = "127.0.0.1"
	REDIS_PORT = "6379"
	REDIS_PWD  = ""
	REDIS_DB   = 0
)

// 定义一个redis.client类型的变量
var client *redis.Client

// 初始化函数
func init() {
	// 实例化一个redis客户端
	client = redis.NewClient(&redis.Options{
		Addr:     REDIS_IP + ":" + REDIS_PORT, // ip:port
		Password: REDIS_PWD,                   // redis连接密码
		DB:       REDIS_DB,                    // 选择的redis库
		PoolSize: 20,                          // 设置连接数,默认是10个连接
	})
}
func main() {
	// redis 全局命令
	// 获取redis所有的键,返回包含所有键的slice
	keys := client.Keys("*").Val()
	fmt.Println(keys)
	// 获取redis中的有多少个键,返回整数
	size := client.DbSize().Val()
	fmt.Println(size)
	// 判断一个键是否存在,有一个存在返回整数1,有两个存在返回整数2...
	exist := client.Exists("age", "name").Val()
	fmt.Println(exist)
	// 删除键,删除成功返回删除的数,删除失败返回0
	del := client.Del("unknownKey").Val()
	fmt.Println(del)
	// 查看键的有效时间
	ttl := client.TTL("age").Val()
	fmt.Println(ttl)
	// 给键设置有效时间,设置成功返回true,失败返回false
	expire := client.Expire("age", time.Second*86400).Val()
	fmt.Println(expire)
	// 查看键的类型(string,hash,list,set,zset...)
	Rtype := client.Type("store:finish:bill:list").Val()
	fmt.Println(Rtype)
	// 给键重命令,成功返回true,失败false
	Rname := client.RenameNX("age", "newAge").Val()
	fmt.Println(Rname)
	// 从redis中随机返回一个键
	key := client.RandomKey().Val()
	fmt.Println(key)

}

func operateString() {

	defer client.Close()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// 设置一组键值对,并社会有效期
	set1 := client.Set("age", 10, time.Hour*24).Val()
	fmt.Println(set1) // OK
	//设置一组键值对,设置的键不存在的时候才能设置成功
	set2 := client.SetNX("age", "20", time.Hour*12).Val()
	fmt.Println(set2) // false
	//设置一组键值对,设置的键必须存在的时候才能设置成功
	set3 := client.SetXX("age", "30", time.Second*86400).Val()
	fmt.Println(set3) // true
	// 批量设置
	set4 := client.MSet("age1", "40", "age2", "50").Val()
	fmt.Println(set4) // OK
	// 获取一个键的值
	get1 := client.Get("age2").Val()
	fmt.Println(get1) // 50
	// 批量获取,获取成功返回slice类型的结果数据
	get2 := client.MGet("age", "age1", "age2").Val()
	fmt.Println(get2) // [30 40 50]
	// 对指定的键进行自增操作
	incr1 := client.Incr("age").Val()
	fmt.Println(incr1) // 31
	// 对指定键进行自减操作
	decr1 := client.Decr("age1").Val()
	fmt.Println(decr1) //39
	// 自增指定的值
	incr2 := client.IncrBy("age", 10).Val()
	fmt.Println(incr2) // 41
	// 自减指定的值
	decr2 := client.DecrBy("age1", 10).Val()
	fmt.Println(decr2) // 29
	// 在key后面追加指定的值,返回字符串的长度
	append1 := client.Append("age2", "abcd").Val()
	fmt.Println(append1) // 6
	// 获取一个键的值得长度
	strlen1 := client.StrLen("age2").Val()
	fmt.Println(strlen1) //6
	// 设置一个键的值,并返回原有的值
	getset1 := client.GetSet("age2", "hello golang").Val()
	fmt.Println(getset1) // 50abcd
	// 设置键的值,在指定的位置
	_ = client.SetRange("age2", 0, "H")
	fmt.Println(client.Get("age2").Val()) // Hello golang
	// 截取一个键的值的部分,返回截取的部分
	newStr := client.GetRange("age2", 6, 11).Val()
	fmt.Println(newStr) //golang
}
