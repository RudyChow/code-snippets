package redis

import "encoding/json"

// Snippet 代码段
type Snippet struct {
	Language string `json:"language" form:"language" xml:"language" binding:"required"`
	Version  string `json:"version" form:"version"  xml:"version" binding:"required"`
	Code     string `json:"code" form:"code" xml:"code" binding:"required"`
}

// MarshalBinary 实现接口
func (m *Snippet) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
