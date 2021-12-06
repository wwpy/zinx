/**
* FileName: ipacket
* Description:
* Author:   ww
* Date:     2021/12/6 6:48 下午
 */
package ziface

type Packet interface {
	Unpack(binaryData []byte) (IMessage, error)
	Pack(msg IMessage) ([]byte, error)
	GetHeadLen() uint32
}
