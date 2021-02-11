package redis

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

//创建redis连接池
var pool *redis.Pool

//Init 初始化
func Init() {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle:     16, //最初的连接数量
		MaxActive:   0,  //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 60, //连接关闭时间60秒 （60秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

//GetConn 获取一个连接
func GetConn() redis.Conn {
	return pool.Get()
}

//SetStruct 保存map信息
func SetStruct(key string, u interface{}, extime int) error {
	//获取一个连接
	c := pool.Get()
	defer c.Close()
	//遍历结构体存储数据
	ut := reflect.TypeOf(u)
	num := ut.NumField()
	uv := reflect.ValueOf(u)
	kind := uv.Kind()
	if kind != reflect.Struct {
		return errors.New("目标信息的类型不是struct")
	}
	for i := 0; i < num; i++ {
		_, err := c.Do("HSet", key, ut.Field(i).Name, uv.Field(i))
		if err != nil {
			return err
		}
	}
	//设置过期时间
	_, err := c.Do("expire", key, extime)
	if err != nil {
		return err
	}
	return nil
}

//HGetInt 获取hash的int类型数据
func HGetInt(key, name string) (int, error) {
	//获取一个连接
	c := pool.Get()
	defer c.Close()
	//获取数据
	val, err := c.Do("HGet", key, name)
	if err != nil {
		return 0, err
	}
	if val == nil {
		return 0, errors.New("数据不存在")
	}
	//val是[]byte类型, 断言
	if v, ok := val.([]byte); ok {
		i, err := strconv.Atoi(string(v))
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return 0, errors.New("断言失败")
}

//HGetString 获取hash的string类型数据
func HGetString(key, name string) (string, error) {
	//获取一个连接
	c := pool.Get()
	defer c.Close()
	//获取数据
	val, err := c.Do("HGet", key, name)
	if err != nil {
		return "", err
	}
	if val == nil {
		return "", errors.New("数据不存在")
	}
	//val是[]byte类型, 断言
	if v, ok := val.([]byte); ok {
		return string(v), nil
	}
	return "", errors.New("断言失败")
}
