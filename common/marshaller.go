package common

import "encoding/json"

// Marshaller 数据序列化程序
type Marshaller interface {
	Marshal(v any) ([]byte, error)
}

// JsonMarshaller json 数据序列化程序
type JsonMarshaller struct{}

func (jm *JsonMarshaller) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}
