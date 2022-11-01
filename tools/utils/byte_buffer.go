package utils

import (
	"bytes"
	"encoding/binary"
)

type ByteBuffer struct {
	bytes bytes.Buffer
}

/*初始化ByteBuffer*/
func NewByteBuffer() ByteBuffer {
	return ByteBuffer{}
}

func (bb *ByteBuffer) WriteBytes(bytes []byte) {
	bb.bytes.Write(bytes)
}

func (bb *ByteBuffer) WriteString(str string) {
	bb.bytes.WriteString(str)
}
func (bb *ByteBuffer) WriteByte(b byte) {
	bb.bytes.WriteByte(b)
}

func (bb *ByteBuffer) WriteBool(value bool) {
	var res byte = 1
	if value {
		res = 1
	}
	bb.bytes.WriteByte(res)
}

func (bb *ByteBuffer) WriteUint32(value uint32) {
	res := Uint32ToBytes(value)
	bb.bytes.Write(res)
}

func (bb *ByteBuffer) WriteInt32(value int32) {
	res := Int32ToBytes(value)
	bb.bytes.Write(res)
}

func (bb *ByteBuffer) WriteUint16(value uint16) {
	res := Uint16ToBytes(value)
	bb.bytes.Write(res)
}

func (bb *ByteBuffer) WriteInt16(value int16) {
	res := Int16ToBytes(value)
	bb.bytes.Write(res)
}

func (bb *ByteBuffer) ReadUint32() (uint32, error) {
	bytes := make([]byte, 4)
	_, err := bb.bytes.Read(bytes)
	if err != nil {
		return 0, err
	}
	value := BytesToUint32(bytes)
	return value, nil
}

func (bb *ByteBuffer) ReadUint16() (uint16, error) {
	bytes := make([]byte, 2)
	_, err := bb.bytes.Read(bytes)
	if err != nil {
		return 0, err
	}
	value := BytesToUint16(bytes)
	return value, nil
}

func (bb *ByteBuffer) ReadInt32() (int32, error) {
	bytes := make([]byte, 4)
	_, err := bb.bytes.Read(bytes)
	if err != nil {
		return 0, err
	}
	value := BytesToInt32(bytes)
	return value, nil
}

func (bb *ByteBuffer) ReadBool() (bool, error) {
	bytes := make([]byte, 1)
	_, err := bb.bytes.Read(bytes)
	if err != nil {
		return false, err
	}
	return bytes[0] == 1, nil
}

func (bb *ByteBuffer) ReadBytes(size int) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := bb.bytes.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (bb *ByteBuffer) GetAllBytes() []byte {
	return bb.bytes.Bytes()
}

/*默认使用大端
*int32转换byte数组
 */
func Int32ToBytes(i int32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	return buf
}

/*默认使用大端
*int32转换byte数组
 */
func Uint32ToBytes(i uint32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(i))
	return buf
}

/*int32转换byte数组
 */
func Int16ToBytes(i int16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(i))
	return buf
}

/*默认使用大端
*int32转换byte数组
 */
func Uint16ToBytes(i uint16) []byte {
	buf := make([]byte, 2)
	binary.LittleEndian.PutUint16(buf, uint16(i))
	return buf
}

/*默认使用大端
*byte数组转换int32
 */
func BytesToInt32(buf []byte) int32 {
	return int32(binary.LittleEndian.Uint32(buf))
}

/*默认使用大端
*byte数组转换int32
 */
func BytesToUint32(buf []byte) uint32 {
	return uint32(binary.LittleEndian.Uint32(buf))
}

/*默认使用大端
*byte数组转换int32
 */
func BytesToInt16(buf []byte) int16 {
	return int16(binary.LittleEndian.Uint16(buf))
}

/*默认使用大端
*byte数组转换int32
 */
func BytesToUint16(buf []byte) uint16 {
	return uint16(binary.LittleEndian.Uint16(buf))
}
