/**
* FileName: server
* Description:
* Author:   ww
* Date:     2021/12/6 8:02 下午
 */
package znet

import (
	"fmt"
	"github.com/wwpy/zinx/utils"
	"github.com/wwpy/zinx/ziface"
	"net"
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
	s := &Server{
		Name: utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP: utils.GlobalObject.Host,
		Port: utils.GlobalObject.TCPPort,
		msgHandler: NewMsgHandle(),
		ConnMgr: NewConnManager(),
		packet: NewDataPack(),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

/**
* Start 开启网络服务
 */
func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s,listenner at IP: %s, Port %d is starting\n", s.Name, s.IP, s.Port)

	go func() {
		// 启动worker工作池机制
		s.msgHandler.StartWorkerPool()
		// 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		// 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			panic(err)
		}
		// 监听成功
		fmt.Println("start zinx server  ", s.Name, " success, now listening...")

		//TODO server.go 应该有一个自动生成ID的方法
		var cID uint32
		cID = 0

		// 启动server网络连接业务
		for {
			// 阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

			// 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				conn.Close()
				continue
			}

			// 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConnection(s, conn, cID, s.msgHandler)
			cID++

			// 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()
}

/**
* Stop 停止服务
 */
func (s *Server) Stop() {
	fmt.Println("[STOP] zinx server , name ", s.Name)

	// 将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()
}

/**
* Serve 运行服务
 */
func (s *Server) Serve() {
	s.Start()

	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加

	//阻塞,否则主Go退出， listenner的go将会退出
	select {}
}

/**
* 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
 */
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgID, router)
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

func init() {

}
