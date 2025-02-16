package redis

import "encoding/json"

type User struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
}

// 序列化
func (m *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// 反序列化
func (m *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

type UserArray []User

// 序列化
func (m *UserArray) MarshalBinary() (data []byte, err error) {
	return json.Marshal(m)
}

// 反序列化
func (m *UserArray) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}
