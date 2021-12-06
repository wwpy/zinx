/**
* FileName: irequest
* Description: 连接请求接口声明
* Author:   ww
* Date:     2021/12/6 6:42 下午
 */
package ziface

/**
* 客户端请求的链接信息和请求的数据封装到Request
 */
type IRequest interface {
	GetConnection() IConnection // 获取请求连接信息
	GetData() []byte            // 获取请求消息的数据
	GetMsgID() uint32           // 获取请求的消息ID
}
