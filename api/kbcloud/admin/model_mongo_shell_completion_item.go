// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type MongoShellCompletionItem struct {
	// completion text
	Completion *string `json:"completion,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoShellCompletionItem instantiates a new MongoShellCompletionItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoShellCompletionItem() *MongoShellCompletionItem {
	this := MongoShellCompletionItem{}
	return &this
}

// NewMongoShellCompletionItemWithDefaults instantiates a new MongoShellCompletionItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoShellCompletionItemWithDefaults() *MongoShellCompletionItem {
	this := MongoShellCompletionItem{}
	return &this
}

// GetCompletion returns the Completion field value if set, zero value otherwise.
func (o *MongoShellCompletionItem) GetCompletion() string {
	if o == nil || o.Completion == nil {
		var ret string
		return ret
	}
	return *o.Completion
}

// GetCompletionOk returns a tuple with the Completion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoShellCompletionItem) GetCompletionOk() (*string, bool) {
	if o == nil || o.Completion == nil {
		return nil, false
	}
	return o.Completion, true
}

// HasCompletion returns a boolean if a field has been set.
func (o *MongoShellCompletionItem) HasCompletion() bool {
	return o != nil && o.Completion != nil
}

// SetCompletion gets a reference to the given string and assigns it to the Completion field.
func (o *MongoShellCompletionItem) SetCompletion(v string) {
	o.Completion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoShellCompletionItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Completion != nil {
		toSerialize["completion"] = o.Completion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoShellCompletionItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Completion *string `json:"completion,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"completion"})
	} else {
		return err
	}
	o.Completion = all.Completion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
