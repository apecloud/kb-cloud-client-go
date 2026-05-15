// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoSchemaAnalysis struct {
	SampleSize    *int64             `json:"sampleSize,omitempty"`
	DocumentCount *int64             `json:"documentCount,omitempty"`
	Fields        []MongoSchemaField `json:"fields,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoSchemaAnalysis instantiates a new MongoSchemaAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoSchemaAnalysis() *MongoSchemaAnalysis {
	this := MongoSchemaAnalysis{}
	return &this
}

// NewMongoSchemaAnalysisWithDefaults instantiates a new MongoSchemaAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoSchemaAnalysisWithDefaults() *MongoSchemaAnalysis {
	this := MongoSchemaAnalysis{}
	return &this
}

// GetSampleSize returns the SampleSize field value if set, zero value otherwise.
func (o *MongoSchemaAnalysis) GetSampleSize() int64 {
	if o == nil || o.SampleSize == nil {
		var ret int64
		return ret
	}
	return *o.SampleSize
}

// GetSampleSizeOk returns a tuple with the SampleSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoSchemaAnalysis) GetSampleSizeOk() (*int64, bool) {
	if o == nil || o.SampleSize == nil {
		return nil, false
	}
	return o.SampleSize, true
}

// HasSampleSize returns a boolean if a field has been set.
func (o *MongoSchemaAnalysis) HasSampleSize() bool {
	return o != nil && o.SampleSize != nil
}

// SetSampleSize gets a reference to the given int64 and assigns it to the SampleSize field.
func (o *MongoSchemaAnalysis) SetSampleSize(v int64) {
	o.SampleSize = &v
}

// GetDocumentCount returns the DocumentCount field value if set, zero value otherwise.
func (o *MongoSchemaAnalysis) GetDocumentCount() int64 {
	if o == nil || o.DocumentCount == nil {
		var ret int64
		return ret
	}
	return *o.DocumentCount
}

// GetDocumentCountOk returns a tuple with the DocumentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoSchemaAnalysis) GetDocumentCountOk() (*int64, bool) {
	if o == nil || o.DocumentCount == nil {
		return nil, false
	}
	return o.DocumentCount, true
}

// HasDocumentCount returns a boolean if a field has been set.
func (o *MongoSchemaAnalysis) HasDocumentCount() bool {
	return o != nil && o.DocumentCount != nil
}

// SetDocumentCount gets a reference to the given int64 and assigns it to the DocumentCount field.
func (o *MongoSchemaAnalysis) SetDocumentCount(v int64) {
	o.DocumentCount = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *MongoSchemaAnalysis) GetFields() []MongoSchemaField {
	if o == nil || o.Fields == nil {
		var ret []MongoSchemaField
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoSchemaAnalysis) GetFieldsOk() (*[]MongoSchemaField, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *MongoSchemaAnalysis) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given []MongoSchemaField and assigns it to the Fields field.
func (o *MongoSchemaAnalysis) SetFields(v []MongoSchemaField) {
	o.Fields = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoSchemaAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.SampleSize != nil {
		toSerialize["sampleSize"] = o.SampleSize
	}
	if o.DocumentCount != nil {
		toSerialize["documentCount"] = o.DocumentCount
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoSchemaAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SampleSize    *int64             `json:"sampleSize,omitempty"`
		DocumentCount *int64             `json:"documentCount,omitempty"`
		Fields        []MongoSchemaField `json:"fields,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"sampleSize", "documentCount", "fields"})
	} else {
		return err
	}
	o.SampleSize = all.SampleSize
	o.DocumentCount = all.DocumentCount
	o.Fields = all.Fields

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
