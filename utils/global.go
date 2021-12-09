/**
* FileName: global
* Description:
* Author:   ww
* Date:     2021/12/9 8:49 上午
 */
package utils

import (
	"encoding/json"
	"github.com/wwpy/zinx/ziface"
	"github.com/wwpy/zinx/zlog"
	"io/ioutil"
	"os"
)

/**
* 存储全局参数
* 用户也可以根据 zinx.json来配置
 */
type Global struct {
	TCPServer ziface.IServer // 当前Zinx的全局Server对象
	Host      string         // 当前服务器主机IP
	TCPPort   int            // 当前服务器主机监听端口号
	Name      string         // 当前服务器名称

	Version          string // 当前Zinx版本号
	MaxPacketSize    uint32 // 数据包的最大值
	MaxConn          int    // 当前服务器主机允许的最大链接个数
	WorkerPoolSize   uint32 // 业务工作Worker池的数量
	MaxWorkerTaskLen uint32 // 业务工作Worker对应负责的任务队列最大任务存储数量
	MaxMsgChanLen    uint32 // SendBuffMsg发送消息的缓冲最大长度

	ConfFilePath string

	LogDir        string // 日志所在文件夹 默认"./log"
	LogFile       string // 日志文件名称   默认""  --如果没有设置日志文件，打印信息将打印至stderr
	LogDebugClose bool   // 是否关闭Debug日志级别调试信息 默认false  -- 默认打开debug信息
}

/*
*	定义一个全局的对象
 */
var GlobalObject *Global

/**
* 判断一个文件是否存在
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return false, err
}

/**
* 读取用户的配置文件
 */
func (g *Global) Reload() {
	if confFileExists, _ := PathExists(g.ConfFilePath); confFileExists != true {
		return
	}

	data, err := ioutil.ReadFile(g.ConfFilePath)
	if err != nil {
		panic(err)
	}

	// 将json数据解析到struct中
	err = json.Unmarshal(data, g)
	if err != nil {
		panic(err)
	}

	// Logger 设置
	if g.LogFile != "" {
		zlog.SetLogFile(g.LogDir, g.LogFile)
	}
	if g.LogDebugClose == true {
		zlog.CloseDebug()
	}
}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		pwd = "."
	}

	// 初始化GlobalObject变量，设置一些默认值
	GlobalObject = &Global{
		Name:             "ServerApp",
		Version:          "V1.0.0",
		TCPPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          12000,
		MaxPacketSize:    4096,
		ConfFilePath:     pwd + "/conf/zinx.json",
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    1024,
		LogDir:           pwd + "/log",
		LogFile:          "",
		LogDebugClose:    false,
	}

	// NOTE: 从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
}
