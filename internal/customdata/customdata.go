package customdata

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	
)

type CustomData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Color string `json:"color"`
}

// ToBinary сериализует данные в бинарный формат
func (cd *CustomData) ToBinary() ([]byte, error) {
	var buf bytes.Buffer

	data, err := json.Marshal(cd)
	if err != nil {
		return nil, err
	}

	length := uint32(len(data))
	if err := binary.Write(&buf, binary.LittleEndian, length); err != nil {
		return nil, err
	}
	if _, err := buf.Write(data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// FromBinary десериализует данные из бинарного формата
func FromBinary(data []byte) (*CustomData, error) {
	buf := bytes.NewReader(data)

	var length uint32
	if err := binary.Read(buf, binary.LittleEndian, &length); err != nil {
		return nil, err
	}

	jsonData := make([]byte, length)
	if _, err := buf.Read(jsonData); err != nil {
		return nil, err
	}

	var cd CustomData
	if err := json.Unmarshal(jsonData, &cd); err != nil {
		return nil, err
	}

	return &cd, nil
}



