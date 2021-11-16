package utils

import (
	"bytes"

	
)
import jsoniter "github.com/json-iterator/go"

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Marshal is exported by gin/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gin/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gin/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gin/json package.
	
)

/// <summary>
/// Json序列化对象成[]byte
/// </summary>
/// <param name="v">被序列化的对象</param>
/// <returns>序列化后的字节数组</returns>
func JSONEncode(v interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	
	e := json.NewEncoder(buffer).Encode(v)
	if e != nil {
		return nil, e
	}
	byteSlice := buffer.Bytes()

	return byteSlice, nil
}

/// <summary>
/// Json反序列化[]byte到对象,解决Marshal解码科学计数法的问题（float->科学计数法）
/// </summary>
/// <param name="b">被反序列化的字节数组</param>
/// <param name="v">被反序列化的对象,必须是对象指针</param>
/// <returns></returns>
func JSONDecode(b []byte, v interface{}) (err error) {
	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err = d.Decode(v)

	return
}
