// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoCollectionStats struct {
	// collection name
	Name *string `json:"name,omitempty"`
	// collection type (collection or view)
	Type          *string `json:"type,omitempty"`
	DocumentCount *int64  `json:"documentCount,omitempty"`
	TotalSize     *int64  `json:"totalSize,omitempty"`
	IndexCount    *int64  `json:"indexCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoCollectionStats instantiates a new MongoCollectionStats object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoCollectionStats() *MongoCollectionStats {
	this := MongoCollectionStats{}
	return &this
}

// NewMongoCollectionStatsWithDefaults instantiates a new MongoCollectionStats object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoCollectionStatsWithDefaults() *MongoCollectionStats {
	this := MongoCollectionStats{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MongoCollectionStats) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionStats) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MongoCollectionStats) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MongoCollectionStats) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MongoCollectionStats) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionStats) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MongoCollectionStats) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MongoCollectionStats) SetType(v string) {
	o.Type = &v
}

// GetDocumentCount returns the DocumentCount field value if set, zero value otherwise.
func (o *MongoCollectionStats) GetDocumentCount() int64 {
	if o == nil || o.DocumentCount == nil {
		var ret int64
		return ret
	}
	return *o.DocumentCount
}

// GetDocumentCountOk returns a tuple with the DocumentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionStats) GetDocumentCountOk() (*int64, bool) {
	if o == nil || o.DocumentCount == nil {
		return nil, false
	}
	return o.DocumentCount, true
}

// HasDocumentCount returns a boolean if a field has been set.
func (o *MongoCollectionStats) HasDocumentCount() bool {
	return o != nil && o.DocumentCount != nil
}

// SetDocumentCount gets a reference to the given int64 and assigns it to the DocumentCount field.
func (o *MongoCollectionStats) SetDocumentCount(v int64) {
	o.DocumentCount = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *MongoCollectionStats) GetTotalSize() int64 {
	if o == nil || o.TotalSize == nil {
		var ret int64
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionStats) GetTotalSizeOk() (*int64, bool) {
	if o == nil || o.TotalSize == nil {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *MongoCollectionStats) HasTotalSize() bool {
	return o != nil && o.TotalSize != nil
}

// SetTotalSize gets a reference to the given int64 and assigns it to the TotalSize field.
func (o *MongoCollectionStats) SetTotalSize(v int64) {
	o.TotalSize = &v
}

// GetIndexCount returns the IndexCount field value if set, zero value otherwise.
func (o *MongoCollectionStats) GetIndexCount() int64 {
	if o == nil || o.IndexCount == nil {
		var ret int64
		return ret
	}
	return *o.IndexCount
}

// GetIndexCountOk returns a tuple with the IndexCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCollectionStats) GetIndexCountOk() (*int64, bool) {
	if o == nil || o.IndexCount == nil {
		return nil, false
	}
	return o.IndexCount, true
}

// HasIndexCount returns a boolean if a field has been set.
func (o *MongoCollectionStats) HasIndexCount() bool {
	return o != nil && o.IndexCount != nil
}

// SetIndexCount gets a reference to the given int64 and assigns it to the IndexCount field.
func (o *MongoCollectionStats) SetIndexCount(v int64) {
	o.IndexCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoCollectionStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DocumentCount != nil {
		toSerialize["documentCount"] = o.DocumentCount
	}
	if o.TotalSize != nil {
		toSerialize["totalSize"] = o.TotalSize
	}
	if o.IndexCount != nil {
		toSerialize["indexCount"] = o.IndexCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoCollectionStats) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name          *string `json:"name,omitempty"`
		Type          *string `json:"type,omitempty"`
		DocumentCount *int64  `json:"documentCount,omitempty"`
		TotalSize     *int64  `json:"totalSize,omitempty"`
		IndexCount    *int64  `json:"indexCount,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "type", "documentCount", "totalSize", "indexCount"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Type = all.Type
	o.DocumentCount = all.DocumentCount
	o.TotalSize = all.TotalSize
	o.IndexCount = all.IndexCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
