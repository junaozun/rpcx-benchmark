package common

import "github.com/smallnest/rpcx/protocol"

type Counter struct {
	Succ        int64 // 成功量
	Fail        int64 // 失败量
	Total       int64 // 总量
	Concurrency int64 // 并发量
	Cost        int64 // 总耗时 ms
}

func GetExecTableNum(serializeType protocol.SerializeType) string {
	switch serializeType {
	case protocol.SerializeType(8):
		return "G"
	case protocol.SerializeType(7):
		return "F"
	case protocol.SerializeType(6):
		return "E"
	case protocol.SerializeType(3):
		return "D"
	case protocol.SerializeType(2):
		return "C"
	case protocol.SerializeType(1):
		return "B"
	}
	return ""
}

func GetSValue(serializeType protocol.SerializeType) string {
	switch serializeType {
	case protocol.SerializeType(8):
		return "gogoProtobuf_GenCode"
	case protocol.SerializeType(7):
		return "gogoProtobuf"
	case protocol.SerializeType(6):
		return "jsoniterCodec"
	case protocol.SerializeType(3):
		return "MsgPack"
	case protocol.SerializeType(2):
		return "ProtoBuf"
	case protocol.SerializeType(1):
		return "JSON"
	}
	return ""
}
