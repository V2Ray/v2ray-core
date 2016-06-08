// +build json

package transport

import "encoding/json"

func (this *Config) UnmarshalJSON(data []byte) error {
	type JsonConfig struct {
		ConnectionReuse bool `json:"connectionReuse"`
	}
	jsonConfig := new(JsonConfig)
	if err := json.Unmarshal(data, jsonConfig); err != nil {
		return err
	}
	this.ConnectionReuse = jsonConfig.ConnectionReuse
	return nil
}
