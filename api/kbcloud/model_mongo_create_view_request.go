// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoCreateViewRequest struct {
	// view name
	Name *string `json:"name,omitempty"`
	// source collection name
	ViewOn *string `json:"viewOn,omitempty"`
	// read-only aggregation pipeline backing the view
	Pipeline []interface{} `json:"pipeline,omitempty"`
	// optional view collation
	Collation interface{} `json:"collation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoCreateViewRequest instantiates a new MongoCreateViewRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoCreateViewRequest() *MongoCreateViewRequest {
	this := MongoCreateViewRequest{}
	return &this
}

// NewMongoCreateViewRequestWithDefaults instantiates a new MongoCreateViewRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoCreateViewRequestWithDefaults() *MongoCreateViewRequest {
	this := MongoCreateViewRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MongoCreateViewRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCreateViewRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MongoCreateViewRequest) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MongoCreateViewRequest) SetName(v string) {
	o.Name = &v
}

// GetViewOn returns the ViewOn field value if set, zero value otherwise.
func (o *MongoCreateViewRequest) GetViewOn() string {
	if o == nil || o.ViewOn == nil {
		var ret string
		return ret
	}
	return *o.ViewOn
}

// GetViewOnOk returns a tuple with the ViewOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCreateViewRequest) GetViewOnOk() (*string, bool) {
	if o == nil || o.ViewOn == nil {
		return nil, false
	}
	return o.ViewOn, true
}

// HasViewOn returns a boolean if a field has been set.
func (o *MongoCreateViewRequest) HasViewOn() bool {
	return o != nil && o.ViewOn != nil
}

// SetViewOn gets a reference to the given string and assigns it to the ViewOn field.
func (o *MongoCreateViewRequest) SetViewOn(v string) {
	o.ViewOn = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *MongoCreateViewRequest) GetPipeline() []interface{} {
	if o == nil || o.Pipeline == nil {
		var ret []interface{}
		return ret
	}
	return o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCreateViewRequest) GetPipelineOk() (*[]interface{}, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return &o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *MongoCreateViewRequest) HasPipeline() bool {
	return o != nil && o.Pipeline != nil
}

// SetPipeline gets a reference to the given []interface{} and assigns it to the Pipeline field.
func (o *MongoCreateViewRequest) SetPipeline(v []interface{}) {
	o.Pipeline = v
}

// GetCollation returns the Collation field value if set, zero value otherwise.
func (o *MongoCreateViewRequest) GetCollation() interface{} {
	if o == nil || o.Collation == nil {
		var ret interface{}
		return ret
	}
	return o.Collation
}

// GetCollationOk returns a tuple with the Collation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoCreateViewRequest) GetCollationOk() (*interface{}, bool) {
	if o == nil || o.Collation == nil {
		return nil, false
	}
	return &o.Collation, true
}

// HasCollation returns a boolean if a field has been set.
func (o *MongoCreateViewRequest) HasCollation() bool {
	return o != nil && o.Collation != nil
}

// SetCollation gets a reference to the given interface{} and assigns it to the Collation field.
func (o *MongoCreateViewRequest) SetCollation(v interface{}) {
	o.Collation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoCreateViewRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ViewOn != nil {
		toSerialize["viewOn"] = o.ViewOn
	}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	if o.Collation != nil {
		toSerialize["collation"] = o.Collation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoCreateViewRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name      *string       `json:"name,omitempty"`
		ViewOn    *string       `json:"viewOn,omitempty"`
		Pipeline  []interface{} `json:"pipeline,omitempty"`
		Collation interface{}   `json:"collation,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "viewOn", "pipeline", "collation"})
	} else {
		return err
	}
	o.Name = all.Name
	o.ViewOn = all.ViewOn
	o.Pipeline = all.Pipeline
	o.Collation = all.Collation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
