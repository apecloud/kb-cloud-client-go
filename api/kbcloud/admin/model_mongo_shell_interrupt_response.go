// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoShellInterruptResponse struct {
	// whether the current operation was interrupted
	Interrupted *bool `json:"interrupted,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoShellInterruptResponse instantiates a new MongoShellInterruptResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoShellInterruptResponse() *MongoShellInterruptResponse {
	this := MongoShellInterruptResponse{}
	return &this
}

// NewMongoShellInterruptResponseWithDefaults instantiates a new MongoShellInterruptResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoShellInterruptResponseWithDefaults() *MongoShellInterruptResponse {
	this := MongoShellInterruptResponse{}
	return &this
}

// GetInterrupted returns the Interrupted field value if set, zero value otherwise.
func (o *MongoShellInterruptResponse) GetInterrupted() bool {
	if o == nil || o.Interrupted == nil {
		var ret bool
		return ret
	}
	return *o.Interrupted
}

// GetInterruptedOk returns a tuple with the Interrupted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellInterruptResponse) GetInterruptedOk() (*bool, bool) {
	if o == nil || o.Interrupted == nil {
		return nil, false
	}
	return o.Interrupted, true
}

// HasInterrupted returns a boolean if a field has been set.
func (o *MongoShellInterruptResponse) HasInterrupted() bool {
	return o != nil && o.Interrupted != nil
}

// SetInterrupted gets a reference to the given bool and assigns it to the Interrupted field.
func (o *MongoShellInterruptResponse) SetInterrupted(v bool) {
	o.Interrupted = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoShellInterruptResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Interrupted != nil {
		toSerialize["interrupted"] = o.Interrupted
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoShellInterruptResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Interrupted *bool `json:"interrupted,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"interrupted"})
	} else {
		return err
	}
	o.Interrupted = all.Interrupted

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
