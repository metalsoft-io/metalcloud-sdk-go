/*
MetalSoft REST API

MetalSoft REST API documentation

API version: 2.0
Contact: support@metalsoft.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// ExtensionInputOptions - Input options.
type ExtensionInputOptions struct {
	ExtensionInputBoolean *ExtensionInputBoolean
	ExtensionInputInteger *ExtensionInputInteger
	ExtensionInputOsTemplate *ExtensionInputOsTemplate
	ExtensionInputServerType *ExtensionInputServerType
	ExtensionInputString *ExtensionInputString
}

// ExtensionInputBooleanAsExtensionInputOptions is a convenience function that returns ExtensionInputBoolean wrapped in ExtensionInputOptions
func ExtensionInputBooleanAsExtensionInputOptions(v *ExtensionInputBoolean) ExtensionInputOptions {
	return ExtensionInputOptions{
		ExtensionInputBoolean: v,
	}
}

// ExtensionInputIntegerAsExtensionInputOptions is a convenience function that returns ExtensionInputInteger wrapped in ExtensionInputOptions
func ExtensionInputIntegerAsExtensionInputOptions(v *ExtensionInputInteger) ExtensionInputOptions {
	return ExtensionInputOptions{
		ExtensionInputInteger: v,
	}
}

// ExtensionInputOsTemplateAsExtensionInputOptions is a convenience function that returns ExtensionInputOsTemplate wrapped in ExtensionInputOptions
func ExtensionInputOsTemplateAsExtensionInputOptions(v *ExtensionInputOsTemplate) ExtensionInputOptions {
	return ExtensionInputOptions{
		ExtensionInputOsTemplate: v,
	}
}

// ExtensionInputServerTypeAsExtensionInputOptions is a convenience function that returns ExtensionInputServerType wrapped in ExtensionInputOptions
func ExtensionInputServerTypeAsExtensionInputOptions(v *ExtensionInputServerType) ExtensionInputOptions {
	return ExtensionInputOptions{
		ExtensionInputServerType: v,
	}
}

// ExtensionInputStringAsExtensionInputOptions is a convenience function that returns ExtensionInputString wrapped in ExtensionInputOptions
func ExtensionInputStringAsExtensionInputOptions(v *ExtensionInputString) ExtensionInputOptions {
	return ExtensionInputOptions{
		ExtensionInputString: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ExtensionInputOptions) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ExtensionInputBoolean
	err = newStrictDecoder(data).Decode(&dst.ExtensionInputBoolean)
	if err == nil {
        if len(dst.ExtensionInputBoolean.AdditionalProperties) > 0 {
            dst.ExtensionInputBoolean = nil
        } else {
            jsonExtensionInputBoolean, _ := json.Marshal(dst.ExtensionInputBoolean)
            if string(jsonExtensionInputBoolean) == "{}" { // empty struct
                dst.ExtensionInputBoolean = nil
            } else {
                if err = validator.Validate(dst.ExtensionInputBoolean); err != nil {
                    dst.ExtensionInputBoolean = nil
                } else {
                    match++
                }
            }
        }
	} else {
		dst.ExtensionInputBoolean = nil
	}

	// try to unmarshal data into ExtensionInputInteger
	err = newStrictDecoder(data).Decode(&dst.ExtensionInputInteger)
	if err == nil {
        if len(dst.ExtensionInputInteger.AdditionalProperties) > 0 {
            dst.ExtensionInputInteger = nil
        } else {
            jsonExtensionInputInteger, _ := json.Marshal(dst.ExtensionInputInteger)
            if string(jsonExtensionInputInteger) == "{}" { // empty struct
                dst.ExtensionInputInteger = nil
            } else {
                if err = validator.Validate(dst.ExtensionInputInteger); err != nil {
                    dst.ExtensionInputInteger = nil
                } else {
                    match++
                }
            }
        }
	} else {
		dst.ExtensionInputInteger = nil
	}

	// try to unmarshal data into ExtensionInputOsTemplate
	err = newStrictDecoder(data).Decode(&dst.ExtensionInputOsTemplate)
	if err == nil {
        if len(dst.ExtensionInputOsTemplate.AdditionalProperties) > 0 {
            dst.ExtensionInputOsTemplate = nil
        } else {
            jsonExtensionInputOsTemplate, _ := json.Marshal(dst.ExtensionInputOsTemplate)
            if string(jsonExtensionInputOsTemplate) == "{}" { // empty struct
                dst.ExtensionInputOsTemplate = nil
            } else {
                if err = validator.Validate(dst.ExtensionInputOsTemplate); err != nil {
                    dst.ExtensionInputOsTemplate = nil
                } else {
                    match++
                }
            }
        }
	} else {
		dst.ExtensionInputOsTemplate = nil
	}

	// try to unmarshal data into ExtensionInputServerType
	err = newStrictDecoder(data).Decode(&dst.ExtensionInputServerType)
	if err == nil {
        if len(dst.ExtensionInputServerType.AdditionalProperties) > 0 {
            dst.ExtensionInputServerType = nil
        } else {
            jsonExtensionInputServerType, _ := json.Marshal(dst.ExtensionInputServerType)
            if string(jsonExtensionInputServerType) == "{}" { // empty struct
                dst.ExtensionInputServerType = nil
            } else {
                if err = validator.Validate(dst.ExtensionInputServerType); err != nil {
                    dst.ExtensionInputServerType = nil
                } else {
                    match++
                }
            }
        }
	} else {
		dst.ExtensionInputServerType = nil
	}

	// try to unmarshal data into ExtensionInputString
	err = newStrictDecoder(data).Decode(&dst.ExtensionInputString)
	if err == nil {
        if len(dst.ExtensionInputString.AdditionalProperties) > 0 {
            dst.ExtensionInputString = nil
        } else {
            jsonExtensionInputString, _ := json.Marshal(dst.ExtensionInputString)
            if string(jsonExtensionInputString) == "{}" { // empty struct
                dst.ExtensionInputString = nil
            } else {
                if err = validator.Validate(dst.ExtensionInputString); err != nil {
                    dst.ExtensionInputString = nil
                } else {
                    match++
                }
            }
        }
	} else {
		dst.ExtensionInputString = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ExtensionInputBoolean = nil
		dst.ExtensionInputInteger = nil
		dst.ExtensionInputOsTemplate = nil
		dst.ExtensionInputServerType = nil
		dst.ExtensionInputString = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ExtensionInputOptions)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ExtensionInputOptions)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ExtensionInputOptions) MarshalJSON() ([]byte, error) {
	if src.ExtensionInputBoolean != nil {
		return json.Marshal(&src.ExtensionInputBoolean)
	}

	if src.ExtensionInputInteger != nil {
		return json.Marshal(&src.ExtensionInputInteger)
	}

	if src.ExtensionInputOsTemplate != nil {
		return json.Marshal(&src.ExtensionInputOsTemplate)
	}

	if src.ExtensionInputServerType != nil {
		return json.Marshal(&src.ExtensionInputServerType)
	}

	if src.ExtensionInputString != nil {
		return json.Marshal(&src.ExtensionInputString)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ExtensionInputOptions) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ExtensionInputBoolean != nil {
		return obj.ExtensionInputBoolean
	}

	if obj.ExtensionInputInteger != nil {
		return obj.ExtensionInputInteger
	}

	if obj.ExtensionInputOsTemplate != nil {
		return obj.ExtensionInputOsTemplate
	}

	if obj.ExtensionInputServerType != nil {
		return obj.ExtensionInputServerType
	}

	if obj.ExtensionInputString != nil {
		return obj.ExtensionInputString
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ExtensionInputOptions) GetActualInstanceValue() (interface{}) {
	if obj.ExtensionInputBoolean != nil {
		return *obj.ExtensionInputBoolean
	}

	if obj.ExtensionInputInteger != nil {
		return *obj.ExtensionInputInteger
	}

	if obj.ExtensionInputOsTemplate != nil {
		return *obj.ExtensionInputOsTemplate
	}

	if obj.ExtensionInputServerType != nil {
		return *obj.ExtensionInputServerType
	}

	if obj.ExtensionInputString != nil {
		return *obj.ExtensionInputString
	}

	// all schemas are nil
	return nil
}

type NullableExtensionInputOptions struct {
	value *ExtensionInputOptions
	isSet bool
}

func (v NullableExtensionInputOptions) Get() *ExtensionInputOptions {
	return v.value
}

func (v *NullableExtensionInputOptions) Set(val *ExtensionInputOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionInputOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionInputOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionInputOptions(val *ExtensionInputOptions) *NullableExtensionInputOptions {
	return &NullableExtensionInputOptions{value: val, isSet: true}
}

func (v NullableExtensionInputOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionInputOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


