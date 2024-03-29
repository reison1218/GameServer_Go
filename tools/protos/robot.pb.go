// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.7.1
// source: proto_files/robot.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

///机器人proto结构体
type RobotPt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RobotId  uint32             `protobuf:"varint,1,opt,name=robot_id,json=robotId,proto3" json:"robot_id,omitempty"`    //机器人id
	NickName string             `protobuf:"bytes,2,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`  //机器人名称
	Cter     *BattleCharacterPt `protobuf:"bytes,3,opt,name=cter,proto3" json:"cter,omitempty"`                          //机器人选择的角色数据
	State    uint32             `protobuf:"varint,4,opt,name=state,proto3" json:"state,omitempty"`                       //机器人状态，是否已经准备 0:未准备，1：准备
	TeamId   uint32             `protobuf:"varint,5,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`       //机器人所属队伍id
	JoinTime uint64             `protobuf:"varint,6,opt,name=join_time,json=joinTime,proto3" json:"join_time,omitempty"` //玩家进入房间的时间
}

func (x *RobotPt) Reset() {
	*x = RobotPt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_robot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotPt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotPt) ProtoMessage() {}

func (x *RobotPt) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_robot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotPt.ProtoReflect.Descriptor instead.
func (*RobotPt) Descriptor() ([]byte, []int) {
	return file_proto_files_robot_proto_rawDescGZIP(), []int{0}
}

func (x *RobotPt) GetRobotId() uint32 {
	if x != nil {
		return x.RobotId
	}
	return 0
}

func (x *RobotPt) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *RobotPt) GetCter() *BattleCharacterPt {
	if x != nil {
		return x.Cter
	}
	return nil
}

func (x *RobotPt) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *RobotPt) GetTeamId() uint32 {
	if x != nil {
		return x.TeamId
	}
	return 0
}

func (x *RobotPt) GetJoinTime() uint64 {
	if x != nil {
		return x.JoinTime
	}
	return 0
}

//地图proto封装结构体
type MapCellPt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                        //快的配置id
	Index       uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`                                  //块的下标
	Element     uint32 `protobuf:"varint,3,opt,name=element,proto3" json:"element,omitempty"`                              //元素
	IsWorldCell bool   `protobuf:"varint,4,opt,name=is_world_cell,json=isWorldCell,proto3" json:"is_world_cell,omitempty"` //是否是世界块
	X           int32  `protobuf:"varint,5,opt,name=x,proto3" json:"x,omitempty"`                                          //x坐标
	Y           int32  `protobuf:"varint,6,opt,name=y,proto3" json:"y,omitempty"`                                          //y坐标
}

func (x *MapCellPt) Reset() {
	*x = MapCellPt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_robot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapCellPt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapCellPt) ProtoMessage() {}

func (x *MapCellPt) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_robot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapCellPt.ProtoReflect.Descriptor instead.
func (*MapCellPt) Descriptor() ([]byte, []int) {
	return file_proto_files_robot_proto_rawDescGZIP(), []int{1}
}

func (x *MapCellPt) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MapCellPt) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *MapCellPt) GetElement() uint32 {
	if x != nil {
		return x.Element
	}
	return 0
}

func (x *MapCellPt) GetIsWorldCell() bool {
	if x != nil {
		return x.IsWorldCell
	}
	return false
}

func (x *MapCellPt) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *MapCellPt) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

//地图proto封装结构体
type TileMapPt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                            //地图配置id
	CellMaps []*MapCellPt `protobuf:"bytes,2,rep,name=cell_maps,json=cellMaps,proto3" json:"cell_maps,omitempty"` //地图块
}

func (x *TileMapPt) Reset() {
	*x = TileMapPt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_robot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TileMapPt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TileMapPt) ProtoMessage() {}

func (x *TileMapPt) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_robot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TileMapPt.ProtoReflect.Descriptor instead.
func (*TileMapPt) Descriptor() ([]byte, []int) {
	return file_proto_files_robot_proto_rawDescGZIP(), []int{2}
}

func (x *TileMapPt) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TileMapPt) GetCellMaps() []*MapCellPt {
	if x != nil {
		return x.CellMaps
	}
	return nil
}

///机器人房间proto结构体
type RobotRoomPt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId uint64 `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"` //房间id
}

func (x *RobotRoomPt) Reset() {
	*x = RobotRoomPt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_robot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RobotRoomPt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotRoomPt) ProtoMessage() {}

func (x *RobotRoomPt) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_robot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotRoomPt.ProtoReflect.Descriptor instead.
func (*RobotRoomPt) Descriptor() ([]byte, []int) {
	return file_proto_files_robot_proto_rawDescGZIP(), []int{3}
}

func (x *RobotRoomPt) GetRoomId() uint64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

//请求机器人
//cmd:25001
type C_REQUEST_ROBOT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId      uint64     `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`                       //房间类型
	NeedNum     uint32     `protobuf:"varint,2,opt,name=need_num,json=needNum,proto3" json:"need_num,omitempty"`                    //需要的数量
	TileMap     *TileMapPt `protobuf:"bytes,3,opt,name=tile_map,json=tileMap,proto3" json:"tile_map,omitempty"`                     //地图数据
	AlreadyCter []uint32   `protobuf:"varint,4,rep,packed,name=already_cter,json=alreadyCter,proto3" json:"already_cter,omitempty"` //已经选择了的角色id
}

func (x *C_REQUEST_ROBOT) Reset() {
	*x = C_REQUEST_ROBOT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_robot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C_REQUEST_ROBOT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C_REQUEST_ROBOT) ProtoMessage() {}

func (x *C_REQUEST_ROBOT) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_robot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C_REQUEST_ROBOT.ProtoReflect.Descriptor instead.
func (*C_REQUEST_ROBOT) Descriptor() ([]byte, []int) {
	return file_proto_files_robot_proto_rawDescGZIP(), []int{4}
}

func (x *C_REQUEST_ROBOT) GetRoomId() uint64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *C_REQUEST_ROBOT) GetNeedNum() uint32 {
	if x != nil {
		return x.NeedNum
	}
	return 0
}

func (x *C_REQUEST_ROBOT) GetTileMap() *TileMapPt {
	if x != nil {
		return x.TileMap
	}
	return nil
}

func (x *C_REQUEST_ROBOT) GetAlreadyCter() []uint32 {
	if x != nil {
		return x.AlreadyCter
	}
	return nil
}

//请求机器人
//cmd:30002
type S_REQUEST_ROBOT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Robots []*RobotPt `protobuf:"bytes,2,rep,name=robots,proto3" json:"robots,omitempty"` //机器人
}

func (x *S_REQUEST_ROBOT) Reset() {
	*x = S_REQUEST_ROBOT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_files_robot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S_REQUEST_ROBOT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S_REQUEST_ROBOT) ProtoMessage() {}

func (x *S_REQUEST_ROBOT) ProtoReflect() protoreflect.Message {
	mi := &file_proto_files_robot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S_REQUEST_ROBOT.ProtoReflect.Descriptor instead.
func (*S_REQUEST_ROBOT) Descriptor() ([]byte, []int) {
	return file_proto_files_robot_proto_rawDescGZIP(), []int{5}
}

func (x *S_REQUEST_ROBOT) GetRobots() []*RobotPt {
	if x != nil {
		return x.Robots
	}
	return nil
}

var File_proto_files_robot_proto protoreflect.FileDescriptor

var file_proto_files_robot_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x07, 0x52, 0x6f,
	0x62, 0x6f, 0x74, 0x50, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a,
	0x04, 0x63, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x50, 0x74, 0x52, 0x04, 0x63, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6a,
	0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x6a, 0x6f, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x09, 0x4d, 0x61, 0x70,
	0x43, 0x65, 0x6c, 0x6c, 0x50, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x69, 0x73, 0x5f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69,
	0x73, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x22, 0x4b, 0x0a, 0x09, 0x54, 0x69, 0x6c, 0x65, 0x4d, 0x61,
	0x70, 0x50, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x09, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x6d, 0x61, 0x70, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x4d, 0x61, 0x70, 0x43, 0x65, 0x6c, 0x6c, 0x50, 0x74, 0x52, 0x08, 0x63, 0x65, 0x6c, 0x6c, 0x4d,
	0x61, 0x70, 0x73, 0x22, 0x26, 0x0a, 0x0b, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x52, 0x6f, 0x6f, 0x6d,
	0x50, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x96, 0x01, 0x0a, 0x0f,
	0x43, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x52, 0x4f, 0x42, 0x4f, 0x54, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x65, 0x64,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x65, 0x65, 0x64,
	0x4e, 0x75, 0x6d, 0x12, 0x2c, 0x0a, 0x08, 0x74, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x54,
	0x69, 0x6c, 0x65, 0x4d, 0x61, 0x70, 0x50, 0x74, 0x52, 0x07, 0x74, 0x69, 0x6c, 0x65, 0x4d, 0x61,
	0x70, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x63, 0x74, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79,
	0x43, 0x74, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x0f, 0x53, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x5f, 0x52, 0x4f, 0x42, 0x4f, 0x54, 0x12, 0x27, 0x0a, 0x06, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x50, 0x74, 0x52, 0x06, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73,
	0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_files_robot_proto_rawDescOnce sync.Once
	file_proto_files_robot_proto_rawDescData = file_proto_files_robot_proto_rawDesc
)

func file_proto_files_robot_proto_rawDescGZIP() []byte {
	file_proto_files_robot_proto_rawDescOnce.Do(func() {
		file_proto_files_robot_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_files_robot_proto_rawDescData)
	})
	return file_proto_files_robot_proto_rawDescData
}

var file_proto_files_robot_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_files_robot_proto_goTypes = []interface{}{
	(*RobotPt)(nil),           // 0: protos.RobotPt
	(*MapCellPt)(nil),         // 1: protos.MapCellPt
	(*TileMapPt)(nil),         // 2: protos.TileMapPt
	(*RobotRoomPt)(nil),       // 3: protos.RobotRoomPt
	(*C_REQUEST_ROBOT)(nil),   // 4: protos.C_REQUEST_ROBOT
	(*S_REQUEST_ROBOT)(nil),   // 5: protos.S_REQUEST_ROBOT
	(*BattleCharacterPt)(nil), // 6: protos.BattleCharacterPt
}
var file_proto_files_robot_proto_depIdxs = []int32{
	6, // 0: protos.RobotPt.cter:type_name -> protos.BattleCharacterPt
	1, // 1: protos.TileMapPt.cell_maps:type_name -> protos.MapCellPt
	2, // 2: protos.C_REQUEST_ROBOT.tile_map:type_name -> protos.TileMapPt
	0, // 3: protos.S_REQUEST_ROBOT.robots:type_name -> protos.RobotPt
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_files_robot_proto_init() }
func file_proto_files_robot_proto_init() {
	if File_proto_files_robot_proto != nil {
		return
	}
	file_proto_files_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_files_robot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotPt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_files_robot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapCellPt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_files_robot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TileMapPt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_files_robot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RobotRoomPt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_files_robot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C_REQUEST_ROBOT); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_files_robot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S_REQUEST_ROBOT); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_files_robot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_files_robot_proto_goTypes,
		DependencyIndexes: file_proto_files_robot_proto_depIdxs,
		MessageInfos:      file_proto_files_robot_proto_msgTypes,
	}.Build()
	File_proto_files_robot_proto = out.File
	file_proto_files_robot_proto_rawDesc = nil
	file_proto_files_robot_proto_goTypes = nil
	file_proto_files_robot_proto_depIdxs = nil
}
