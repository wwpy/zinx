/**
* FileName:
* Description: 连接管理,包括添加、删除、通过一个连接ID获得连接对象，当前连接数量、清空全部连接等方法
* Author:   ww
* Date:     2021/12/6 5:06 下午
 */

package ziface

type IConnManager interface {
	Add(conn IConnection)                   // 添加链接
	Remove(conn IConnection)                // 删除连接
	Get(connID uint32) (IConnection, error) // 通过ConnID获取链接
	Len() int                               // 获取连接个数
	ClearConn()                             // 删除并停止所有链接
}
