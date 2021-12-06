/**
* FileName: irouter
* Description: 消息路由接口声明
* Author:   ww
* Date:     2021/12/6 6:43 下午
 */
package ziface

/**
* 处理业务方法
* IRequest 则包含用链接信息和链接请求数据信息
 */
type IRouter interface {
	PreHandle(request IRequest)  // 在处理conn业务之前的钩子方法
	Handle(request IRequest)     // 处理conn业务的方法
	PostHandle(request IRequest) // 处理conn业务之后的钩子方法
}
