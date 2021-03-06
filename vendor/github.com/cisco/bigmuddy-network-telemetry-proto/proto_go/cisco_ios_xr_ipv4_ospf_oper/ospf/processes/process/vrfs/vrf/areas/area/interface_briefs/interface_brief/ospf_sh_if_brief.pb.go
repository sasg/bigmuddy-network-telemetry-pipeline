// Code generated by protoc-gen-go.
// source: ospf_sh_if_brief.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_areas_area_interface_briefs_interface_brief is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_if_brief.proto

It has these top-level messages:
	OspfShIfBrief_KEYS
	OspfShIfBrief
	OspfShInterfaceMadj
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_areas_area_interface_briefs_interface_brief

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

// OSPF Interface Brief Information
type OspfShIfBrief_KEYS struct {
	ProcessName   string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName       string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AreaId        uint32 `protobuf:"varint,3,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	InterfaceName string `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *OspfShIfBrief_KEYS) Reset()                    { *m = OspfShIfBrief_KEYS{} }
func (m *OspfShIfBrief_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShIfBrief_KEYS) ProtoMessage()               {}
func (*OspfShIfBrief_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShIfBrief_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShIfBrief_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShIfBrief_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShIfBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type OspfShIfBrief struct {
	// Interface name
	InterfaceName string `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Area ID string in decimal or dotted-decimal format
	InterfaceArea string `protobuf:"bytes,51,opt,name=interface_area,json=interfaceArea" json:"interface_area,omitempty"`
	// Interface IP Address
	InterfaceAddress string `protobuf:"bytes,52,opt,name=interface_address,json=interfaceAddress" json:"interface_address,omitempty"`
	// Interface IP Mask
	InterfaceMask uint32 `protobuf:"varint,53,opt,name=interface_mask,json=interfaceMask" json:"interface_mask,omitempty"`
	// Interface link cost
	InterfaceLinkCost uint32 `protobuf:"varint,54,opt,name=interface_link_cost,json=interfaceLinkCost" json:"interface_link_cost,omitempty"`
	// Interface OSPF state
	OspfInterfaceState string `protobuf:"bytes,55,opt,name=ospf_interface_state,json=ospfInterfaceState" json:"ospf_interface_state,omitempty"`
	// Interface in fast detect hold down state
	InterfaceFastDetectHoldDown bool `protobuf:"varint,56,opt,name=interface_fast_detect_hold_down,json=interfaceFastDetectHoldDown" json:"interface_fast_detect_hold_down,omitempty"`
	// Total number of Neighbors
	InterfaceNeighborCount uint32 `protobuf:"varint,57,opt,name=interface_neighbor_count,json=interfaceNeighborCount" json:"interface_neighbor_count,omitempty"`
	// Total number of Adjacent Neighbors
	InterfaceAdjNeighborCount uint32 `protobuf:"varint,58,opt,name=interface_adj_neighbor_count,json=interfaceAdjNeighborCount" json:"interface_adj_neighbor_count,omitempty"`
	// If true, interface is multi-area
	InterfaceisMadj bool `protobuf:"varint,59,opt,name=interfaceis_madj,json=interfaceisMadj" json:"interfaceis_madj,omitempty"`
	// Total number of multi-area
	InterfaceMadjCount uint32 `protobuf:"varint,60,opt,name=interface_madj_count,json=interfaceMadjCount" json:"interface_madj_count,omitempty"`
	// Information for multi-area on the interface
	InterfaceMadjList []*OspfShInterfaceMadj `protobuf:"bytes,61,rep,name=interface_madj_list,json=interfaceMadjList" json:"interface_madj_list,omitempty"`
}

func (m *OspfShIfBrief) Reset()                    { *m = OspfShIfBrief{} }
func (m *OspfShIfBrief) String() string            { return proto.CompactTextString(m) }
func (*OspfShIfBrief) ProtoMessage()               {}
func (*OspfShIfBrief) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OspfShIfBrief) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceArea() string {
	if m != nil {
		return m.InterfaceArea
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceAddress() string {
	if m != nil {
		return m.InterfaceAddress
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceMask() uint32 {
	if m != nil {
		return m.InterfaceMask
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceLinkCost() uint32 {
	if m != nil {
		return m.InterfaceLinkCost
	}
	return 0
}

func (m *OspfShIfBrief) GetOspfInterfaceState() string {
	if m != nil {
		return m.OspfInterfaceState
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceFastDetectHoldDown() bool {
	if m != nil {
		return m.InterfaceFastDetectHoldDown
	}
	return false
}

func (m *OspfShIfBrief) GetInterfaceNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceNeighborCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceAdjNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceAdjNeighborCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceisMadj() bool {
	if m != nil {
		return m.InterfaceisMadj
	}
	return false
}

func (m *OspfShIfBrief) GetInterfaceMadjCount() uint32 {
	if m != nil {
		return m.InterfaceMadjCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceMadjList() []*OspfShInterfaceMadj {
	if m != nil {
		return m.InterfaceMadjList
	}
	return nil
}

// OSPF Interface Multi-Area Information
type OspfShInterfaceMadj struct {
	// Area ID string in decimal or dotted-decimal format
	InterfaceArea string `protobuf:"bytes,1,opt,name=interface_area,json=interfaceArea" json:"interface_area,omitempty"`
	// Area ID
	MadjAreaId uint32 `protobuf:"varint,2,opt,name=madj_area_id,json=madjAreaId" json:"madj_area_id,omitempty"`
	// Number of Neighbors
	InterfaceNeighborCount uint32 `protobuf:"varint,3,opt,name=interface_neighbor_count,json=interfaceNeighborCount" json:"interface_neighbor_count,omitempty"`
	// Total number of Adjacent Neighbors
	InterfaceAdjNeighborCount uint32 `protobuf:"varint,4,opt,name=interface_adj_neighbor_count,json=interfaceAdjNeighborCount" json:"interface_adj_neighbor_count,omitempty"`
	// Interface link cost
	InterfaceLinkCost uint32 `protobuf:"varint,5,opt,name=interface_link_cost,json=interfaceLinkCost" json:"interface_link_cost,omitempty"`
	// Interface OSPF state
	OspfInterfaceState string `protobuf:"bytes,6,opt,name=ospf_interface_state,json=ospfInterfaceState" json:"ospf_interface_state,omitempty"`
}

func (m *OspfShInterfaceMadj) Reset()                    { *m = OspfShInterfaceMadj{} }
func (m *OspfShInterfaceMadj) String() string            { return proto.CompactTextString(m) }
func (*OspfShInterfaceMadj) ProtoMessage()               {}
func (*OspfShInterfaceMadj) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OspfShInterfaceMadj) GetInterfaceArea() string {
	if m != nil {
		return m.InterfaceArea
	}
	return ""
}

func (m *OspfShInterfaceMadj) GetMadjAreaId() uint32 {
	if m != nil {
		return m.MadjAreaId
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceNeighborCount
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceAdjNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceAdjNeighborCount
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceLinkCost() uint32 {
	if m != nil {
		return m.InterfaceLinkCost
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetOspfInterfaceState() string {
	if m != nil {
		return m.OspfInterfaceState
	}
	return ""
}

func init() {
	proto.RegisterType((*OspfShIfBrief_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.areas.area.interface_briefs.interface_brief.ospf_sh_if_brief_KEYS")
	proto.RegisterType((*OspfShIfBrief)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.areas.area.interface_briefs.interface_brief.ospf_sh_if_brief")
	proto.RegisterType((*OspfShInterfaceMadj)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.areas.area.interface_briefs.interface_brief.ospf_sh_interface_madj")
}

func init() { proto.RegisterFile("ospf_sh_if_brief.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 534 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xd6, 0xad, 0x1b, 0xde, 0x06, 0xc3, 0xc0, 0xf0, 0x04, 0x12, 0xa5, 0x12, 0x52, 0x11,
	0x52, 0x84, 0xb6, 0x01, 0xe3, 0x4f, 0x68, 0x5a, 0x41, 0x4c, 0x6c, 0x5c, 0x74, 0x57, 0x5c, 0x59,
	0x6e, 0xec, 0xac, 0x4e, 0xdb, 0xb8, 0xf2, 0x31, 0x1d, 0x77, 0x3c, 0x05, 0xcf, 0xc1, 0x1b, 0xf0,
	0x38, 0x3c, 0x07, 0xf2, 0xe9, 0x4f, 0x9c, 0xa8, 0x03, 0x21, 0x24, 0x6e, 0xa2, 0xe4, 0x7c, 0xdf,
	0x77, 0xfc, 0xf9, 0xb3, 0x4f, 0xc8, 0xb6, 0x81, 0x51, 0xca, 0xa1, 0xc7, 0x75, 0xca, 0xbb, 0x56,
	0xab, 0x34, 0x1e, 0x59, 0xe3, 0x0c, 0x3d, 0x4f, 0x34, 0x24, 0x86, 0x6b, 0x03, 0xfc, 0x8b, 0xe5,
	0x7a, 0x34, 0xde, 0xe7, 0xc8, 0x34, 0x23, 0x65, 0x63, 0xff, 0xe6, 0x79, 0x89, 0x02, 0x50, 0x30,
	0x7b, 0x8b, 0xc7, 0x36, 0xc5, 0x47, 0x2c, 0xac, 0x12, 0x80, 0xcf, 0x58, 0xe7, 0x4e, 0xd9, 0x54,
	0x24, 0x6a, 0xd2, 0x1d, 0xaa, 0x85, 0xe6, 0xb7, 0x88, 0xdc, 0xaa, 0x7a, 0xe0, 0x1f, 0xde, 0x7e,
	0x3a, 0xa3, 0xf7, 0xc9, 0xc6, 0xb4, 0x33, 0xcf, 0xc5, 0x50, 0xb1, 0xa8, 0x11, 0xb5, 0xae, 0x74,
	0xd6, 0xa7, 0xb5, 0x8f, 0x62, 0xa8, 0xe8, 0x0e, 0x59, 0x1b, 0xdb, 0x74, 0x02, 0x2f, 0x21, 0xbc,
	0x3a, 0xb6, 0x29, 0x42, 0xb7, 0xc9, 0xaa, 0x37, 0xc0, 0xb5, 0x64, 0xb5, 0x46, 0xd4, 0xda, 0xec,
	0xd4, 0xfd, 0xe7, 0xb1, 0xa4, 0x0f, 0xc8, 0xd5, 0xc2, 0x03, 0x2a, 0x97, 0x51, 0xb9, 0x39, 0xaf,
	0x7a, 0x7d, 0xf3, 0xe7, 0x0a, 0xd9, 0xaa, 0xfa, 0x5a, 0xa0, 0xdd, 0x5d, 0xa0, 0x2d, 0xd3, 0xfc,
	0xb2, 0x6c, 0xaf, 0x42, 0x3b, 0xb4, 0x4a, 0xd0, 0x47, 0xe4, 0x7a, 0x40, 0x93, 0xd2, 0x2a, 0x00,
	0xb6, 0x8f, 0xcc, 0xad, 0x82, 0x39, 0xa9, 0x97, 0x7b, 0x0e, 0x05, 0xf4, 0xd9, 0x13, 0xdc, 0x56,
	0xd1, 0xf3, 0x54, 0x40, 0x9f, 0xc6, 0xe4, 0x46, 0x41, 0x1b, 0xe8, 0xbc, 0xcf, 0x13, 0x03, 0x8e,
	0x3d, 0x45, 0x6e, 0xb1, 0xdc, 0x89, 0xce, 0xfb, 0x47, 0x06, 0x1c, 0x7d, 0x4c, 0x6e, 0xe2, 0x2e,
	0x0b, 0x11, 0x38, 0xe1, 0x14, 0x7b, 0x86, 0x36, 0xa8, 0xc7, 0x8e, 0x67, 0xd0, 0x99, 0x47, 0x68,
	0x9b, 0xdc, 0x2b, 0xc8, 0xa9, 0x00, 0xc7, 0xa5, 0x72, 0x2a, 0x71, 0xbc, 0x67, 0x06, 0x92, 0x4b,
	0x73, 0x91, 0xb3, 0x83, 0x46, 0xd4, 0x5a, 0xeb, 0xdc, 0x99, 0xd3, 0xde, 0x09, 0x70, 0x6d, 0x24,
	0xbd, 0x37, 0x03, 0xd9, 0x36, 0x17, 0x39, 0x3d, 0x20, 0x2c, 0x48, 0x52, 0xe9, 0xf3, 0x5e, 0xd7,
	0x58, 0x9e, 0x98, 0xcf, 0xb9, 0x63, 0xcf, 0xd1, 0xec, 0x76, 0x91, 0xe9, 0x14, 0x3e, 0xf2, 0x28,
	0x7d, 0x43, 0xee, 0x86, 0xa9, 0x65, 0x55, 0xf5, 0x0b, 0x54, 0xef, 0x04, 0x01, 0x66, 0xe5, 0x06,
	0x0f, 0x49, 0x91, 0xae, 0x06, 0x3e, 0x14, 0x32, 0x63, 0x2f, 0xd1, 0xf1, 0xb5, 0xa0, 0x7e, 0x2a,
	0x64, 0xe6, 0xd3, 0x09, 0x43, 0x97, 0xd9, 0x74, 0x8d, 0x57, 0xb8, 0x06, 0x0d, 0xa2, 0x97, 0xd9,
	0xa4, 0xf9, 0xf7, 0x28, 0x3c, 0x00, 0x94, 0x0c, 0x34, 0x38, 0xf6, 0xba, 0x51, 0x6b, 0xad, 0xef,
	0x7e, 0x8d, 0xff, 0xd3, 0x58, 0xc5, 0xf3, 0xab, 0x5b, 0xf2, 0x12, 0xdc, 0x00, 0x6f, 0xf9, 0x44,
	0x83, 0x6b, 0xfe, 0x58, 0x0a, 0x7e, 0x02, 0x25, 0xf6, 0x82, 0x7b, 0x1c, 0x2d, 0xba, 0xc7, 0x0d,
	0xb2, 0x81, 0x1b, 0x9d, 0xcd, 0xdb, 0x12, 0xa6, 0x43, 0x7c, 0xed, 0x70, 0x32, 0x73, 0xbf, 0x3b,
	0xed, 0xda, 0x3f, 0x9d, 0xf6, 0xf2, 0x9f, 0x4e, 0xfb, 0x92, 0x81, 0x58, 0xf9, 0xdb, 0x81, 0xa8,
	0x5f, 0x36, 0x10, 0xdd, 0x3a, 0xfe, 0x31, 0xf7, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x87, 0xb4,
	0x53, 0x23, 0x4b, 0x05, 0x00, 0x00,
}
