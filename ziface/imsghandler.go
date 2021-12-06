/**
* FileName: imsghandler
* Description: 提供worker启动、处理消息业务调用等接口
* Author:   ww
* Date:     2021/12/6 6:41 下午
 */
package ziface

/*
* 消息管理抽象层
 */
type IMsgHandler interface {
	DoMsgHandler(request IRequest)          // 以非阻塞方式处理消息
	AddRouter(msgID uint32, router IRouter) // 为消息添加具体的处理逻辑
	StartWorkerPool()                       // 启动worker工作池
	SendMsgToTaskQueue(request IRequest)    // 将消息交给TaskQueue,由worker进行处理
}
