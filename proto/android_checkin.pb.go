// Copyright 2014 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Logging information for Android "checkin" events (automatic, periodic
// requests made by Android devices to the server).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: android_checkin.proto

package __

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

// enum values correspond to the type of device.
// Used in the AndroidCheckinProto and Device proto.
type DeviceType int32

const (
	// Android Device
	DeviceType_DEVICE_ANDROID_OS DeviceType = 1
	// Apple IOS device
	DeviceType_DEVICE_IOS_OS DeviceType = 2
	// Chrome browser - Not Chrome OS.  No hardware records.
	DeviceType_DEVICE_CHROME_BROWSER DeviceType = 3
	// Chrome OS
	DeviceType_DEVICE_CHROME_OS DeviceType = 4
)

// Enum value maps for DeviceType.
var (
	DeviceType_name = map[int32]string{
		1: "DEVICE_ANDROID_OS",
		2: "DEVICE_IOS_OS",
		3: "DEVICE_CHROME_BROWSER",
		4: "DEVICE_CHROME_OS",
	}
	DeviceType_value = map[string]int32{
		"DEVICE_ANDROID_OS":     1,
		"DEVICE_IOS_OS":         2,
		"DEVICE_CHROME_BROWSER": 3,
		"DEVICE_CHROME_OS":      4,
	}
)

func (x DeviceType) Enum() *DeviceType {
	p := new(DeviceType)
	*p = x
	return p
}

func (x DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_android_checkin_proto_enumTypes[0].Descriptor()
}

func (DeviceType) Type() protoreflect.EnumType {
	return &file_android_checkin_proto_enumTypes[0]
}

func (x DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *DeviceType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = DeviceType(num)
	return nil
}

// Deprecated: Use DeviceType.Descriptor instead.
func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_android_checkin_proto_rawDescGZIP(), []int{0}
}

type ChromeBuildProto_Platform int32

const (
	ChromeBuildProto_PLATFORM_WIN   ChromeBuildProto_Platform = 1
	ChromeBuildProto_PLATFORM_MAC   ChromeBuildProto_Platform = 2
	ChromeBuildProto_PLATFORM_LINUX ChromeBuildProto_Platform = 3
	ChromeBuildProto_PLATFORM_CROS  ChromeBuildProto_Platform = 4
	ChromeBuildProto_PLATFORM_IOS   ChromeBuildProto_Platform = 5
	// Just a placeholder. Likely don't need it due to the presence of the
	// Android GCM on phone/tablet devices.
	ChromeBuildProto_PLATFORM_ANDROID ChromeBuildProto_Platform = 6
)

// Enum value maps for ChromeBuildProto_Platform.
var (
	ChromeBuildProto_Platform_name = map[int32]string{
		1: "PLATFORM_WIN",
		2: "PLATFORM_MAC",
		3: "PLATFORM_LINUX",
		4: "PLATFORM_CROS",
		5: "PLATFORM_IOS",
		6: "PLATFORM_ANDROID",
	}
	ChromeBuildProto_Platform_value = map[string]int32{
		"PLATFORM_WIN":     1,
		"PLATFORM_MAC":     2,
		"PLATFORM_LINUX":   3,
		"PLATFORM_CROS":    4,
		"PLATFORM_IOS":     5,
		"PLATFORM_ANDROID": 6,
	}
)

func (x ChromeBuildProto_Platform) Enum() *ChromeBuildProto_Platform {
	p := new(ChromeBuildProto_Platform)
	*p = x
	return p
}

func (x ChromeBuildProto_Platform) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChromeBuildProto_Platform) Descriptor() protoreflect.EnumDescriptor {
	return file_android_checkin_proto_enumTypes[1].Descriptor()
}

func (ChromeBuildProto_Platform) Type() protoreflect.EnumType {
	return &file_android_checkin_proto_enumTypes[1]
}

func (x ChromeBuildProto_Platform) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ChromeBuildProto_Platform) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ChromeBuildProto_Platform(num)
	return nil
}

// Deprecated: Use ChromeBuildProto_Platform.Descriptor instead.
func (ChromeBuildProto_Platform) EnumDescriptor() ([]byte, []int) {
	return file_android_checkin_proto_rawDescGZIP(), []int{0, 0}
}

type ChromeBuildProto_Channel int32

const (
	ChromeBuildProto_CHANNEL_STABLE  ChromeBuildProto_Channel = 1
	ChromeBuildProto_CHANNEL_BETA    ChromeBuildProto_Channel = 2
	ChromeBuildProto_CHANNEL_DEV     ChromeBuildProto_Channel = 3
	ChromeBuildProto_CHANNEL_CANARY  ChromeBuildProto_Channel = 4
	ChromeBuildProto_CHANNEL_UNKNOWN ChromeBuildProto_Channel = 5 // for tip of tree or custom builds
)

// Enum value maps for ChromeBuildProto_Channel.
var (
	ChromeBuildProto_Channel_name = map[int32]string{
		1: "CHANNEL_STABLE",
		2: "CHANNEL_BETA",
		3: "CHANNEL_DEV",
		4: "CHANNEL_CANARY",
		5: "CHANNEL_UNKNOWN",
	}
	ChromeBuildProto_Channel_value = map[string]int32{
		"CHANNEL_STABLE":  1,
		"CHANNEL_BETA":    2,
		"CHANNEL_DEV":     3,
		"CHANNEL_CANARY":  4,
		"CHANNEL_UNKNOWN": 5,
	}
)

func (x ChromeBuildProto_Channel) Enum() *ChromeBuildProto_Channel {
	p := new(ChromeBuildProto_Channel)
	*p = x
	return p
}

func (x ChromeBuildProto_Channel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChromeBuildProto_Channel) Descriptor() protoreflect.EnumDescriptor {
	return file_android_checkin_proto_enumTypes[2].Descriptor()
}

func (ChromeBuildProto_Channel) Type() protoreflect.EnumType {
	return &file_android_checkin_proto_enumTypes[2]
}

func (x ChromeBuildProto_Channel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ChromeBuildProto_Channel) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ChromeBuildProto_Channel(num)
	return nil
}

// Deprecated: Use ChromeBuildProto_Channel.Descriptor instead.
func (ChromeBuildProto_Channel) EnumDescriptor() ([]byte, []int) {
	return file_android_checkin_proto_rawDescGZIP(), []int{0, 1}
}

// Build characteristics unique to the Chrome browser, and Chrome OS
type ChromeBuildProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The platform of the device.
	Platform *ChromeBuildProto_Platform `protobuf:"varint,1,opt,name=platform,enum=proto.ChromeBuildProto_Platform" json:"platform,omitempty"`
	// The Chrome instance's version.
	ChromeVersion *string `protobuf:"bytes,2,opt,name=chrome_version,json=chromeVersion" json:"chrome_version,omitempty"`
	// The Channel (build type) of Chrome.
	Channel *ChromeBuildProto_Channel `protobuf:"varint,3,opt,name=channel,enum=proto.ChromeBuildProto_Channel" json:"channel,omitempty"`
}

func (x *ChromeBuildProto) Reset() {
	*x = ChromeBuildProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_android_checkin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChromeBuildProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChromeBuildProto) ProtoMessage() {}

func (x *ChromeBuildProto) ProtoReflect() protoreflect.Message {
	mi := &file_android_checkin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChromeBuildProto.ProtoReflect.Descriptor instead.
func (*ChromeBuildProto) Descriptor() ([]byte, []int) {
	return file_android_checkin_proto_rawDescGZIP(), []int{0}
}

func (x *ChromeBuildProto) GetPlatform() ChromeBuildProto_Platform {
	if x != nil && x.Platform != nil {
		return *x.Platform
	}
	return ChromeBuildProto_PLATFORM_WIN
}

func (x *ChromeBuildProto) GetChromeVersion() string {
	if x != nil && x.ChromeVersion != nil {
		return *x.ChromeVersion
	}
	return ""
}

func (x *ChromeBuildProto) GetChannel() ChromeBuildProto_Channel {
	if x != nil && x.Channel != nil {
		return *x.Channel
	}
	return ChromeBuildProto_CHANNEL_STABLE
}

// Information sent by the device in a "checkin" request.
type AndroidCheckinProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Miliseconds since the Unix epoch of the device's last successful checkin.
	LastCheckinMsec *int64 `protobuf:"varint,2,opt,name=last_checkin_msec,json=lastCheckinMsec" json:"last_checkin_msec,omitempty"`
	// The current MCC+MNC of the mobile device's current cell.
	CellOperator *string `protobuf:"bytes,6,opt,name=cell_operator,json=cellOperator" json:"cell_operator,omitempty"`
	// The MCC+MNC of the SIM card (different from operator if the
	// device is roaming, for instance).
	SimOperator *string `protobuf:"bytes,7,opt,name=sim_operator,json=simOperator" json:"sim_operator,omitempty"`
	// The device's current roaming state (reported starting in eclair builds).
	// Currently one of "{,not}mobile-{,not}roaming", if it is present at all.
	Roaming *string `protobuf:"bytes,8,opt,name=roaming" json:"roaming,omitempty"`
	// For devices supporting multiple user profiles (which may be
	// supported starting in jellybean), the ordinal number of the
	// profile that is checking in.  This is 0 for the primary profile
	// (which can't be changed without wiping the device), and 1,2,3,...
	// for additional profiles (which can be added and deleted freely).
	UserNumber *int32 `protobuf:"varint,9,opt,name=user_number,json=userNumber" json:"user_number,omitempty"`
	// Class of device.  Indicates the type of build proto
	// (IosBuildProto/ChromeBuildProto/AndroidBuildProto)
	// That is included in this proto
	Type *DeviceType `protobuf:"varint,12,opt,name=type,enum=proto.DeviceType,def=1" json:"type,omitempty"`
	// For devices running MCS on Chrome, build-specific characteristics
	// of the browser.  There are no hardware aspects (except for ChromeOS).
	// This will only be populated for Chrome builds/ChromeOS devices
	ChromeBuild *ChromeBuildProto `protobuf:"bytes,13,opt,name=chrome_build,json=chromeBuild" json:"chrome_build,omitempty"`
}

// Default values for AndroidCheckinProto fields.
const (
	Default_AndroidCheckinProto_Type = DeviceType_DEVICE_ANDROID_OS
)

func (x *AndroidCheckinProto) Reset() {
	*x = AndroidCheckinProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_android_checkin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AndroidCheckinProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndroidCheckinProto) ProtoMessage() {}

func (x *AndroidCheckinProto) ProtoReflect() protoreflect.Message {
	mi := &file_android_checkin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndroidCheckinProto.ProtoReflect.Descriptor instead.
func (*AndroidCheckinProto) Descriptor() ([]byte, []int) {
	return file_android_checkin_proto_rawDescGZIP(), []int{1}
}

func (x *AndroidCheckinProto) GetLastCheckinMsec() int64 {
	if x != nil && x.LastCheckinMsec != nil {
		return *x.LastCheckinMsec
	}
	return 0
}

func (x *AndroidCheckinProto) GetCellOperator() string {
	if x != nil && x.CellOperator != nil {
		return *x.CellOperator
	}
	return ""
}

func (x *AndroidCheckinProto) GetSimOperator() string {
	if x != nil && x.SimOperator != nil {
		return *x.SimOperator
	}
	return ""
}

func (x *AndroidCheckinProto) GetRoaming() string {
	if x != nil && x.Roaming != nil {
		return *x.Roaming
	}
	return ""
}

func (x *AndroidCheckinProto) GetUserNumber() int32 {
	if x != nil && x.UserNumber != nil {
		return *x.UserNumber
	}
	return 0
}

func (x *AndroidCheckinProto) GetType() DeviceType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_AndroidCheckinProto_Type
}

func (x *AndroidCheckinProto) GetChromeBuild() *ChromeBuildProto {
	if x != nil {
		return x.ChromeBuild
	}
	return nil
}

var File_android_checkin_proto protoreflect.FileDescriptor

var file_android_checkin_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x03, 0x0a, 0x10, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x3c, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68,
	0x72, 0x6f, 0x6d, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x22, 0x7d, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12,
	0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x57, 0x49, 0x4e, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x4d, 0x41,
	0x43, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f,
	0x4c, 0x49, 0x4e, 0x55, 0x58, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x4c, 0x41, 0x54, 0x46,
	0x4f, 0x52, 0x4d, 0x5f, 0x43, 0x52, 0x4f, 0x53, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c,
	0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x49, 0x4f, 0x53, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10,
	0x50, 0x4c, 0x41, 0x54, 0x46, 0x4f, 0x52, 0x4d, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44,
	0x10, 0x06, 0x22, 0x69, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a,
	0x0e, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10,
	0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x42, 0x45, 0x54,
	0x41, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f, 0x44,
	0x45, 0x56, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x48, 0x41, 0x4e, 0x4e, 0x45, 0x4c, 0x5f,
	0x43, 0x41, 0x4e, 0x41, 0x52, 0x59, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x48, 0x41, 0x4e,
	0x4e, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x05, 0x22, 0xba, 0x02,
	0x0a, 0x13, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x0a, 0x11, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x69, 0x6e, 0x5f, 0x6d, 0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x4d, 0x73, 0x65,
	0x63, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x65, 0x6c, 0x6c, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6d, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x69,
	0x6d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x61,
	0x6d, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x3a, 0x11, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x4e,
	0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x4f, 0x53, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3a,
	0x0a, 0x0c, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x65, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x72,
	0x6f, 0x6d, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0b, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x2a, 0x67, 0x0a, 0x0a, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x5f, 0x4f, 0x53, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x49, 0x4f, 0x53, 0x5f, 0x4f, 0x53,
	0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x48, 0x52,
	0x4f, 0x4d, 0x45, 0x5f, 0x42, 0x52, 0x4f, 0x57, 0x53, 0x45, 0x52, 0x10, 0x03, 0x12, 0x14, 0x0a,
	0x10, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x43, 0x48, 0x52, 0x4f, 0x4d, 0x45, 0x5f, 0x4f,
	0x53, 0x10, 0x04, 0x42, 0x06, 0x48, 0x03, 0x5a, 0x02, 0x2e, 0x2f,
}

var (
	file_android_checkin_proto_rawDescOnce sync.Once
	file_android_checkin_proto_rawDescData = file_android_checkin_proto_rawDesc
)

func file_android_checkin_proto_rawDescGZIP() []byte {
	file_android_checkin_proto_rawDescOnce.Do(func() {
		file_android_checkin_proto_rawDescData = protoimpl.X.CompressGZIP(file_android_checkin_proto_rawDescData)
	})
	return file_android_checkin_proto_rawDescData
}

var file_android_checkin_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_android_checkin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_android_checkin_proto_goTypes = []interface{}{
	(DeviceType)(0),                // 0: proto.DeviceType
	(ChromeBuildProto_Platform)(0), // 1: proto.ChromeBuildProto.Platform
	(ChromeBuildProto_Channel)(0),  // 2: proto.ChromeBuildProto.Channel
	(*ChromeBuildProto)(nil),       // 3: proto.ChromeBuildProto
	(*AndroidCheckinProto)(nil),    // 4: proto.AndroidCheckinProto
}
var file_android_checkin_proto_depIdxs = []int32{
	1, // 0: proto.ChromeBuildProto.platform:type_name -> proto.ChromeBuildProto.Platform
	2, // 1: proto.ChromeBuildProto.channel:type_name -> proto.ChromeBuildProto.Channel
	0, // 2: proto.AndroidCheckinProto.type:type_name -> proto.DeviceType
	3, // 3: proto.AndroidCheckinProto.chrome_build:type_name -> proto.ChromeBuildProto
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_android_checkin_proto_init() }
func file_android_checkin_proto_init() {
	if File_android_checkin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_android_checkin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChromeBuildProto); i {
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
		file_android_checkin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AndroidCheckinProto); i {
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
			RawDescriptor: file_android_checkin_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_android_checkin_proto_goTypes,
		DependencyIndexes: file_android_checkin_proto_depIdxs,
		EnumInfos:         file_android_checkin_proto_enumTypes,
		MessageInfos:      file_android_checkin_proto_msgTypes,
	}.Build()
	File_android_checkin_proto = out.File
	file_android_checkin_proto_rawDesc = nil
	file_android_checkin_proto_goTypes = nil
	file_android_checkin_proto_depIdxs = nil
}
