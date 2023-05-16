/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by protoc-gen-go.
// source: mesos/docker/spec.proto
// DO NOT EDIT!

/*
Package docker_spec is a generated protocol buffer package.

It is generated from these files:

	mesos/docker/spec.proto

It has these top-level messages:

	ImageReference
	Config
*/
package docker_spec

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// The docker image reference:
//
//	[REGISTRY_HOST[:REGISTRY_PORT]/]REPOSITORY[:TAG|@DIGEST]
//
// See more details in:
// https://github.com/docker/distribution/blob/master/reference/reference.go
type ImageReference struct {
	// The registry in the following format: 'HOST:PORT'.
	Registry *string `protobuf:"bytes,1,opt,name=registry" json:"registry,omitempty"`
	// The repository name (e.g., library/busybox, ubuntu).
	Repository *string `protobuf:"bytes,2,req,name=repository" json:"repository,omitempty"`
	// For example: latest, 2.0.2, etc.
	Tag *string `protobuf:"bytes,3,opt,name=tag" json:"tag,omitempty"`
	// In the format of 'ALGORITHM:HEX'.
	Digest           *string `protobuf:"bytes,4,opt,name=digest" json:"digest,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ImageReference) Reset()                    { *m = ImageReference{} }
func (m *ImageReference) String() string            { return proto.CompactTextString(m) }
func (*ImageReference) ProtoMessage()               {}
func (*ImageReference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ImageReference) GetRegistry() string {
	if m != nil && m.Registry != nil {
		return *m.Registry
	}
	return ""
}

func (m *ImageReference) GetRepository() string {
	if m != nil && m.Repository != nil {
		return *m.Repository
	}
	return ""
}

func (m *ImageReference) GetTag() string {
	if m != nil && m.Tag != nil {
		return *m.Tag
	}
	return ""
}

func (m *ImageReference) GetDigest() string {
	if m != nil && m.Digest != nil {
		return *m.Digest
	}
	return ""
}

// *
// The config structure for docker config file
// (e.g., ~/.docker/config.json or ~/.dockercfg).
type Config struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The base64 encoded 'auth' and 'email' pair in docker
// config file.
type Config_Auth struct {
	Auth             *string `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	Email            *string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Config_Auth) Reset()                    { *m = Config_Auth{} }
func (m *Config_Auth) String() string            { return proto.CompactTextString(m) }
func (*Config_Auth) ProtoMessage()               {}
func (*Config_Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Config_Auth) GetAuth() string {
	if m != nil && m.Auth != nil {
		return *m.Auth
	}
	return ""
}

func (m *Config_Auth) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageReference)(nil), "docker.spec.ImageReference")
	proto.RegisterType((*Config)(nil), "docker.spec.Config")
	proto.RegisterType((*Config_Auth)(nil), "docker.spec.Config.Auth")
}

func init() { proto.RegisterFile("mesos/docker/spec.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0xcd, 0xcf, 0x8a, 0x83, 0x30,
	0x10, 0xc7, 0x71, 0xfc, 0xb3, 0xb2, 0xce, 0xc2, 0x52, 0x86, 0xd2, 0x06, 0x0f, 0x45, 0x3c, 0x79,
	0xd2, 0x9e, 0x7b, 0x2b, 0x3d, 0xf5, 0xea, 0x1b, 0x04, 0x1d, 0x63, 0x68, 0x35, 0x92, 0xc4, 0x82,
	0x6f, 0x5f, 0x8c, 0x52, 0x7a, 0xfb, 0x7d, 0x3f, 0x30, 0x0c, 0x1c, 0x7b, 0x32, 0xca, 0x94, 0x8d,
	0xaa, 0x1f, 0xa4, 0x4b, 0x33, 0x52, 0x5d, 0x8c, 0x5a, 0x59, 0x85, 0x7f, 0x2b, 0x15, 0x0b, 0x65,
	0x2f, 0xf8, 0xbf, 0xf7, 0x5c, 0x50, 0x45, 0x2d, 0x69, 0x1a, 0x6a, 0xc2, 0x04, 0x7e, 0x35, 0x09,
	0x69, 0xac, 0x9e, 0x99, 0x97, 0x7a, 0x79, 0x5c, 0x7d, 0x1a, 0x4f, 0x00, 0x9a, 0x46, 0x65, 0xa4,
	0x55, 0x7a, 0x66, 0x7e, 0xea, 0xe7, 0x71, 0xf5, 0x25, 0xb8, 0x83, 0xc0, 0x72, 0xc1, 0x02, 0x77,
	0xb6, 0x4c, 0x3c, 0x40, 0xd4, 0x48, 0x41, 0xc6, 0xb2, 0xd0, 0xe1, 0x56, 0xd9, 0x05, 0xa2, 0x9b,
	0x1a, 0x5a, 0x29, 0x92, 0x33, 0x84, 0xd7, 0xc9, 0x76, 0x88, 0x10, 0xf2, 0xc9, 0x76, 0xdb, 0x4f,
	0xb7, 0x71, 0x0f, 0x3f, 0xd4, 0x73, 0xf9, 0x64, 0xbe, 0xc3, 0x35, 0xde, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x75, 0x93, 0x1a, 0x54, 0xda, 0x00, 0x00, 0x00,
}
