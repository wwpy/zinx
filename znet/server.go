/**
* FileName: server
* Description:
* Author:   ww
* Date:     2021/12/6 8:02 下午
 */
package znet

import (
	"fmt"
	"github.com/wwpy/zinx/ziface"
)

/**
* 定义一个Server服务类
 */
type Server struct {
	// 服务名称
	Name string
	// tcp4
	IPVersion string
	// 服务绑定的IP地址
	IP string
	// 服务绑定的端口
	Port int
	// 当前Server的消息管理模块，用来绑定MsgID和对应的处理方法
	msgHandler ziface.IMsgHandler
	// 当前Server的链接管理器
	ConnMgr ziface.IConnManager
	// 当前Server连接时创建Hook函数
	OnConnStart func(conn ziface.IConnection)
	// 当前Server连接断开时Hook函数
	OnConnStop func(conn ziface.IConnection)

	packet ziface.Packet
}

/**
* NewServer 创建一个服务句柄
 */
func NewServer(opts ...Option) ziface.IServer {
	s := &Server{}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

/**
* Start 开启网络服务
 */
func (s *Server) Start() {

}

/**
* Stop 停止服务
 */
func (s *Server) Stop() {

}

/**
* Serve 运行服务
 */
func (s *Server) Serve() {

}

/**
* 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
 */
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {

}

/**
* 获取链接管理
 */
func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

/**
* 设置Server的连接时创建Hook函数
 */
func (s *Server) SetOnConnStart(hookFunc func(ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

/**
* 设置Server的连接断开时Hook函数
 */
func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

/**
* 调用连接OnConnStart Hook函数
 */
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

/**
* 调用连接OnConnStop Hook函数
 */
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}

func (s *Server) Packet() ziface.Packet {
	return s.packet
}
