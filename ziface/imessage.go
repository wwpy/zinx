/**
* FileName: imessage
* Description: 提供消息的基本方法
* Author:   ww
* Date:     2021/12/6 5:29 下午
 */
package ziface

/*
* 将请求的一个消息封装到message中，定义抽象层接口
*/
type IMessage interface {
	GetDataLen() uint32 //获取消息数据段长度
	GetMsgID() uint32   //获取消息ID
	GetData() []byte    //获取消息内容

	SetDataLen(uint32) //设置消息数据段长度
	SetMsgID(uint32)   //设置消息ID
	SetData([]byte)    //设置消息内容
}