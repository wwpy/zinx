/**
* FileName: connmanager
* Description: 连接管理
* Author:   ww
* Date:     2021/12/10 8:27 上午
 */
package znet

import (
	"errors"
	"fmt"
	"github.com/wwpy/zinx/ziface"
	"sync/atomic"
)

type ConnManager struct {
	connections atomic.Value
}

/**
* 创建一个链接管理
 */
func NewConnManager() *ConnManager {
	var cm = &ConnManager{}
	connections := make(map[uint32]ziface.IConnection)
	cm.connections.Store(connections)
	return cm
}

/**
* 添加连接
 */
func (connMgr *ConnManager) Add(conn ziface.IConnection) {
	connections := connMgr.connections.Load().(map[uint32]ziface.IConnection)

	// 将conn连接添加到ConnManager中
	connections[conn.GetConnID()] = conn
	connMgr.connections.Store(connections)

	fmt.Println("connection add to ConnManager successfully: conn num = ", connMgr.Len())
}

/**
* 删除连接
 */
func (connMgr *ConnManager) Remove(conn ziface.IConnection) {
	connections := connMgr.connections.Load().(map[uint32]ziface.IConnection)
	// 删除连接信息
	delete(connections, conn.GetConnID())
	connMgr.connections.Store(connections)

	fmt.Println("connection Remove ConnID=", conn.GetConnID(), " successfully: conn num = ", connMgr.Len())
}

/**
* 根据ConnID获取链接
 */
func (connMgr *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	connections := connMgr.connections.Load().(map[uint32]ziface.IConnection)

	if conn, ok := connections[connID]; ok {
		return conn, nil
	}

	return nil, errors.New("connection not found")
}

/**
* 获取当前连接
 */
func (connMgr *ConnManager) Len() int {
	connections := connMgr.connections.Load().(map[uint32]ziface.IConnection)
	return len(connections)
}

/**
* 清除并停止所有连接
 */
func (connMgr *ConnManager) ClearConn() {
	connections := connMgr.connections.Load().(map[uint32]ziface.IConnection)

	// 停止并删除全部的连接信息
	for connID, conn := range connections {
		// 停止
		conn.Stop()
		// 删除
		delete(connections, connID)
	}
	connMgr.connections.Store(connections)

	fmt.Println("Clear All Connections successfully: conn num = ", connMgr.Len())
}

func (connMgr *ConnManager) ClearOneConn(connID uint32) {
	connections := connMgr.connections.Load().(map[uint32]ziface.IConnection)

	if conn, ok := connections[connID]; ok {
		// 停止
		conn.Stop()
		// 删除
		delete(connections, connID)
		connMgr.connections.Store(connections)
		fmt.Println("Clear Connections ID:  ", connID, "succeed")
		return
	}

	fmt.Println("Clear Connections ID:  ", connID, "err")
	return
}
