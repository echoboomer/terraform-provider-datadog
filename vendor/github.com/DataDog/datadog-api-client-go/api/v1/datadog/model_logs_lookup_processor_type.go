/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// LogsLookupProcessorType Type of logs lookup processor.
type LogsLookupProcessorType string

// List of LogsLookupProcessorType
const (
	LOGSLOOKUPPROCESSORTYPE_LOOKUP_PROCESSOR LogsLookupProcessorType = "lookup-processor"
)

func (v *LogsLookupProcessorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogsLookupProcessorType(value)
	for _, existing := range []LogsLookupProcessorType{"lookup-processor"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogsLookupProcessorType", *v)
}

// Ptr returns reference to LogsLookupProcessorType value
func (v LogsLookupProcessorType) Ptr() *LogsLookupProcessorType {
	return &v
}

type NullableLogsLookupProcessorType struct {
	value *LogsLookupProcessorType
	isSet bool
}

func (v NullableLogsLookupProcessorType) Get() *LogsLookupProcessorType {
	return v.value
}

func (v *NullableLogsLookupProcessorType) Set(val *LogsLookupProcessorType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsLookupProcessorType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsLookupProcessorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsLookupProcessorType(val *LogsLookupProcessorType) *NullableLogsLookupProcessorType {
	return &NullableLogsLookupProcessorType{value: val, isSet: true}
}

func (v NullableLogsLookupProcessorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsLookupProcessorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
