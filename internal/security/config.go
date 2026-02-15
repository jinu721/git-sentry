package security

import (
	"fmt"
	"reflect"
)

type ConfigValidator struct {
	allowedFields map[string]bool
	rules         map[string]ValidationRule
}

type ValidationRule struct {
	Required bool
	MinValue *int
	MaxValue *int
	AllowedValues []string
}

func NewConfigValidator() *ConfigValidator {
	return &ConfigValidator{
		allowedFields: map[string]bool{
			"rules":                  true,
			"auto_suggest_commits":   true,
			"auto_suggest_pushes":    true,
			"commit_message_format":  true,
			"max_files_changed":      true,
			"max_lines_changed":      true,
			"max_minutes_since_commit": true,
			"max_unpushed_commits":   true,
		},
		rules: map[string]ValidationRule{
			"max_files_changed": {
				Required: true,
				MinValue: intPtr(1),
				MaxValue: intPtr(1000),
			},
			"max_lines_changed": {
				Required: true,
				MinValue: intPtr(1),
				MaxValue: intPtr(10000),
			},
			"max_minutes_since_commit": {
				Required: true,
				MinValue: intPtr(1),
				MaxValue: intPtr(1440),
			},
			"max_unpushed_commits": {
				Required: true,
				MinValue: intPtr(1),
				MaxValue: intPtr(100),
			},
			"commit_message_format": {
				Required: false,
				AllowedValues: []string{"conventional", "simple"},
			},
		},
	}
}

func (cv *ConfigValidator) ValidateConfig(config map[string]interface{}) error {
	for field := range config {
		if !cv.allowedFields[field] {
			return fmt.Errorf("unknown config field: %s", field)
		}
	}
	
	for field, rule := range cv.rules {
		if err := cv.validateField(config, field, rule); err != nil {
			return err
		}
	}
	
	return nil
}

func (cv *ConfigValidator) validateField(config map[string]interface{}, field string, rule ValidationRule) error {
	value, exists := config[field]
	
	if rule.Required && !exists {
		return fmt.Errorf("required field missing: %s", field)
	}
	
	if !exists {
		return nil
	}
	
	if rule.MinValue != nil || rule.MaxValue != nil {
		intVal, ok := value.(int)
		if !ok {
			return fmt.Errorf("field %s must be integer", field)
		}
		
		if rule.MinValue != nil && intVal < *rule.MinValue {
			return fmt.Errorf("field %s value %d below minimum %d", field, intVal, *rule.MinValue)
		}
		
		if rule.MaxValue != nil && intVal > *rule.MaxValue {
			return fmt.Errorf("field %s value %d above maximum %d", field, intVal, *rule.MaxValue)
		}
	}
	
	if len(rule.AllowedValues) > 0 {
		strVal, ok := value.(string)
		if !ok {
			return fmt.Errorf("field %s must be string", field)
		}
		
		for _, allowed := range rule.AllowedValues {
			if strVal == allowed {
				return nil
			}
		}
		
		return fmt.Errorf("field %s has invalid value: %s", field, strVal)
	}
	
	return nil
}

func intPtr(i int) *int {
	return &i
}

func ValidateConfigStruct(config interface{}) error {
	validator := NewConfigValidator()
	
	configMap := structToMap(config)
	return validator.ValidateConfig(configMap)
}

func structToMap(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		
		if !value.CanInterface() {
			continue
		}
		
		tag := field.Tag.Get("yaml")
		if tag == "" {
			tag = field.Name
		}
		
		result[tag] = value.Interface()
	}
	
	return result
}