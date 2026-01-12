package mcp

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type baseItem[T any] struct {
	TemplateID string `json:"templateId"`
	Attributes T      `json:"attributes"`
	Quantity   int    `json:"quantity"`
}

type Item baseItem[json.RawMessage]

func (m *Item) ReadInto(v any) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Pointer || val.IsNil() {
		return fmt.Errorf("v must be a non-nil pointer")
	}

	elem := val.Elem()
	if elem.Kind() != reflect.Struct {
		return fmt.Errorf("v must be a pointer to a struct")
	}

	if field := elem.FieldByName("TemplateID"); field.IsValid() && field.CanSet() {
		field.SetString(m.TemplateID)
	}

	if field := elem.FieldByName("Quantity"); field.IsValid() && field.CanSet() {
		field.SetInt(int64(m.Quantity))
	}

	if field := elem.FieldByName("Attributes"); field.IsValid() && field.CanSet() {
		if err := json.Unmarshal(m.Attributes, field.Addr().Interface()); err != nil {
			return fmt.Errorf("failed to unmarshal attributes: %w", err)
		}
	}

	return nil
}
