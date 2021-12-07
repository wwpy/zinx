/**
* FileName: options
* Description:
* Author:   ww
* Date:     2021/12/7 9:25 上午
 */
package znet

import "github.com/wwpy/zinx/ziface"

type Option func(s *Server)

// 实现Packet 接口可自由实现数据包解析格式，如果没有则使用默认解析格式
func WithPacket(pack ziface.Packet) Option {
	return func(s *Server) {
		s.packet = pack
	}
}
