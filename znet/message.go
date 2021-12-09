/**
* FileName: message
* Description:
* Author:   ww
* Date:     2021/12/9 9:15 上午
 */
package znet

type Message struct {
	DataLen uint32 // 消息长度
	ID      uint32 // 消息ID
	Data    []byte // 消息内容
}

/**
* 创建一个Message消息包
 */
func NewMsgPackage(ID uint32, data []byte) *Message {
	return &Message{
		DataLen: uint32(len(data)),
		ID:      ID,
		Data:    data,
	}
}

/**
* 获取消息数据段长度
 */
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

/**
* 获取消息ID
 */
func (msg *Message) GetMsgID() uint32 {
	return msg.ID
}

/**
* 获取消息内容
 */
func (msg *Message) GetData() []byte {
	return msg.Data
}

/**
* 设置消息数据段长度
 */
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

/*
* 设置消息ID
 */
func (msg *Message) SetMsgID(msgID uint32) {
	msg.ID = msgID
}

/**
* 设置消息内容
 */
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
