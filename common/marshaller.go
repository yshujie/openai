package common

import "encoding/json"

// marshaller 数据序列化程序
type marshaller interface {
	marshal(v any) ([]byte, error)
}

// jsonMarshaller json 数据序列化程序
type jsonMarshaller struct{}

func (jm *jsonMarshaller) marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
