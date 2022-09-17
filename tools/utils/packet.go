package utils

import (
	"container/list"
)

type Packet struct {
	Des   PacketDes //包头
	bytes []byte    //包体
}
type PacketDes struct {
	Cmd         uint32 //命令号
	Len         uint32 //包长度
	UserId      uint32 //用户id
	IsBroad     bool   //是否广播
	IsClient    bool   //是否是客户端
	ServerToken uint32 //服务器token
}

func NewPacket() Packet {
	return Packet{}
}

func (packet *Packet) SetData(data []byte) {
	packet.bytes = data
}

func (packet *Packet) SetUserId(userId uint32) {
	packet.Des.UserId = userId
}

func (packet *Packet) SetCmd(cmd uint32) {
	packet.Des.Cmd = cmd
}

func (packet *Packet) GetData() []byte {
	return packet.bytes
}

/*构建来自服务器的消息包*/
func BuildArrayFromServer(bytes []byte) (list.List, error) {
	bb := ByteBuffer{}
	bb.WriteBytes(bytes)
	list := list.List{}

	for {
		if bb.bytes.Len() == 0 {
			return list, nil
		}
		cmd, err := bb.ReadUint32()
		if err != nil {
			ZapLogger.Error(err.Error())
			return list, err
		}
		len, err := bb.ReadUint32()
		if err != nil {
			ZapLogger.Error(err.Error())
			return list, err
		}
		userId, err := bb.ReadUint32()
		if err != nil {
			ZapLogger.Error(err.Error())
			return list, err
		}
		isClient, err := bb.ReadBool()
		if err != nil {
			ZapLogger.Error(err.Error())
			return list, err
		}
		isBroad, err := bb.ReadBool()
		if err != nil {
			ZapLogger.Error(err.Error())
			return list, err
		}
		serverToken, err := bb.ReadUint32()
		if err != nil {
			ZapLogger.Error(err.Error())
			return list, err
		}
		bodyLen := len - 18
		packet := Packet{}
		packet.Des.UserId = userId
		packet.Des.Cmd = cmd
		packet.Des.IsClient = isClient
		packet.Des.IsBroad = isBroad
		packet.Des.ServerToken = serverToken
		bytesLen := bb.bytes.Len()
		if bodyLen > 0 && bytesLen > 0 {
			bbs := make([]byte, bytesLen)
			_, err := bb.bytes.Read(bbs)
			if err != nil {
				return list, err
			}
			packet.SetData(bbs)
		}
		list.PushFront(packet)
	}
}

/*构建来自客户端的消息包*/
func BuildArrayFromClient(bytes []byte) (list.List, error) {
	list := list.List{}
	bb := ByteBuffer{}
	bb.WriteBytes(bytes)
	for {
		if bb.bytes.Len() == 0 {
			return list, nil
		}
		cmd, err := bb.ReadUint32()
		if err != nil {
			ZapLogger.Error(err.Error())
			return list, err
		}
		dataLen, err := bb.ReadUint32()
		if err != nil {
			ZapLogger.Error(err.Error())
			return list, err
		}
		_, err1 := bb.ReadUint32()
		if err1 != nil {
			ZapLogger.Error(err1.Error())
			return list, err1
		}
		_, err2 := bb.ReadUint32()
		if err2 != nil {
			ZapLogger.Error(err2.Error())
			return list, err2
		}
		bodySize := dataLen - 16
		packet := Packet{}
		packet.SetCmd(cmd)
		packet.Des.Len = dataLen
		packet.Des.IsClient = true
		packet.Des.IsBroad = false
		bytesLen := bb.bytes.Len()
		if bodySize > 0 && bytesLen > 0 {
			bbs := make([]byte, bytesLen)
			_, err := bb.bytes.Read(bbs)
			if err != nil {
				return list, err
			}
			packet.SetData(bbs)
		}
		list.PushFront(packet)
	}
}

func BuildFromOnlyServer(bytes []byte) (Packet, error) {
	packet := Packet{}
	bb := ByteBuffer{}
	bb.bytes.Read(bytes)
	cmd, err := bb.ReadUint32()
	if err != nil {
		ZapLogger.Error(err.Error())
		return packet, err
	}
	_, err1 := bb.ReadUint32()
	if err1 != nil {
		ZapLogger.Error(err1.Error())
		return packet, err1
	}
	userId, err := bb.ReadUint32()
	if err != nil {
		ZapLogger.Error(err.Error())
		return packet, err
	}
	isClient, err := bb.ReadBool()
	if err != nil {
		ZapLogger.Error(err.Error())
		return packet, err
	}
	isBroad, err := bb.ReadBool()
	if err != nil {
		ZapLogger.Error(err.Error())
		return packet, err
	}
	serverToken, err := bb.ReadUint32()
	if err != nil {
		ZapLogger.Error(err.Error())
		return packet, err
	}
	packet.SetCmd(cmd)
	packet.Des.UserId = userId
	packet.Des.IsClient = isClient
	packet.Des.IsBroad = isBroad
	packet.Des.ServerToken = serverToken
	bytesLen := bb.bytes.Len()
	if bytesLen > 0 {
		bytes := make([]byte, bytesLen)
		bb.bytes.Read(bytes)
		packet.SetData(bytes)
	}
	return packet, nil
}

func BuildFromOnlyClient(bytes []byte) (Packet, error) {
	packet := Packet{}
	bb := ByteBuffer{}
	bb.bytes.Read(bytes)
	cmd, err := bb.ReadUint32()
	if err != nil {
		ZapLogger.Error(err.Error())
		return packet, err
	}
	_, err1 := bb.ReadUint32()
	if err1 != nil {
		ZapLogger.Error(err1.Error())
		return packet, err1
	}
	_, err2 := bb.ReadUint32()
	if err2 != nil {
		ZapLogger.Error(err2.Error())
		return packet, err2
	}
	_, err3 := bb.ReadUint32()
	if err3 != nil {
		ZapLogger.Error(err3.Error())
		return packet, err3
	}

	packet.SetCmd(cmd)
	packet.Des.IsClient = true
	packet.Des.IsBroad = false
	if bb.bytes.Len() > 0 {
		bytes := make([]byte, bb.bytes.Len())
		_, _ = bb.bytes.Read(bytes)
		packet.SetData(bytes)
	}
	return packet, nil
}

func (packet *Packet) AllToClientBytes() []byte {
	bb := ByteBuffer{}
	bb.WriteUint32(packet.Des.Cmd)
	bb.WriteUint32(packet.Des.Len)
	bb.WriteUint32(0)
	bb.WriteUint32(0)
	bb.WriteBytes(packet.bytes)
	bytes := make([]byte, bb.bytes.Cap())
	bb.bytes.Read(bytes)
	return bytes
}

func (packet *Packet) calLen() {
	lenRes := 16 + uint32(len(packet.bytes))
	packet.Des.Len = lenRes
}

func (packet *Packet) ToClientByteBuf() ByteBuffer {
	packet.calLen()
	bb := ByteBuffer{}
	bb.WriteUint32(packet.Des.Cmd)
	bb.WriteUint32(packet.Des.UserId)
	bb.WriteUint32(0)
	bb.WriteUint32(0)
	bb.WriteBytes(packet.bytes)
	return bb
}

func (packet *Packet) ToServerByteBuf() ByteBuffer {
	bb := ByteBuffer{}
	bb.WriteUint32(packet.Des.Cmd)
	bb.WriteUint32(18 + uint32(len(packet.bytes)))
	bb.WriteUint32(packet.Des.UserId)
	bb.WriteBool(packet.Des.IsClient)
	bb.WriteBool(packet.Des.IsBroad)
	bb.WriteUint32(packet.Des.ServerToken)
	bb.WriteBytes(packet.bytes)
	return bb
}

func (packet *Packet) BuildClientBytes() []byte {
	bb := packet.ToClientByteBuf()
	bytes := make([]byte, bb.bytes.Cap())
	bb.bytes.Read(bytes)
	return bytes
}

func (packet *Packet) BuildServerBytes() []byte {
	bb := packet.ToServerByteBuf()
	bytes := make([]byte, bb.bytes.Cap())
	bb.bytes.Read(bytes)
	return bytes
}

/*构建一个发送的消息包*/
func BuildPacketBytes(cmd uint32, userId uint32, bytes []byte, isServer bool, is2Client bool) []byte {
	packet := Packet{}
	packet.SetCmd(cmd)
	packet.Des.UserId = userId
	packet.Des.IsClient = is2Client
	packet.SetData(bytes)
	if isServer {
		return packet.BuildServerBytes()
	} else {
		return packet.BuildClientBytes()
	}
}

/*构建发送的消息包，主要用于服务器内部通信，直接发送到具体的某一个服务器*/
func BuildPacketBytesDirection(cmd uint32, userId uint32, bytes []byte, isServer bool, is2Client bool, serverToken uint32) []byte {
	packet := Packet{}
	packet.SetCmd(cmd)
	packet.Des.UserId = userId
	packet.Des.IsClient = is2Client
	packet.Des.ServerToken = serverToken
	packet.SetData(bytes)
	if isServer {
		return packet.BuildServerBytes()
	} else {
		return packet.BuildClientBytes()
	}
}

func BuildPushPacketBytes(cmd uint32, userId uint32, data []byte, isServer bool, is2Client bool) []byte {
	packet := Packet{}
	packet.SetCmd(cmd)
	packet.Des.UserId = userId
	packet.Des.IsClient = is2Client
	packet.Des.IsBroad = true
	packet.SetData(data)
	if isServer {
		return packet.BuildServerBytes()
	} else {
		return packet.BuildClientBytes()
	}
}
