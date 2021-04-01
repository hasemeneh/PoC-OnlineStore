package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hasemeneh/PoC-OnlineStore/svc/warehouse/src/constants"
)

type MetaData map[string]json.RawMessage

func (m MetaData) Value() (driver.Value, error) {
	if len(m) == 0 {
		return "", nil
	}
	bs, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(bs), nil
}

func (m *MetaData) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("Incompatible data type %v, data must be bytes", reflect.TypeOf(src))
	}
	if len(source) == 0 {
		return nil
	}
	return json.Unmarshal(source, m)
}

func (m MetaData) IsExist(key string) bool {
	_, ok := m[key]
	return ok
}

func (m MetaData) Remove(key string) {
	delete(m, key)
}

func (m MetaData) GetString(key string) (string, error) {
	var (
		s string
	)

	err := m.GetCustom(key, &s)
	if err != nil {
		return "", err
	}
	return s, nil
}
func (m MetaData) GetCustom(key string, result interface{}) error {
	byteJson, err := m.getByteJson(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(byteJson, result)
}
func (m MetaData) getByteJson(key string) ([]byte, error) {

	value, ok := m[key]
	if !ok {
		return nil, constants.KeyNotFoundError
	}
	return value, nil

}
func (m MetaData) GetInt64(key string) (int64, error) {
	var (
		s int64
	)
	err := m.GetCustom(key, &s)
	if err != nil {
		return 0, err
	}
	return s, nil
}

func (m MetaData) GetFloat64(key string) (float64, error) {
	var (
		s float64
	)
	err := m.GetCustom(key, &s)
	if err != nil {
		return 0, err
	}
	return s, nil
}

func (m MetaData) GetBool(key string) (bool, error) {
	var (
		s bool
	)
	err := m.GetCustom(key, &s)
	if err != nil {
		return false, err
	}
	return s, nil
}
