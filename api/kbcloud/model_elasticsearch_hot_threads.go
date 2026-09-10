// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type ElasticsearchHotThreads struct {
	Node      *string `json:"node,omitempty"`
	Type      string  `json:"type"`
	Content   string  `json:"content"`
	Truncated bool    `json:"truncated"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewElasticsearchHotThreads instantiates a new ElasticsearchHotThreads object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewElasticsearchHotThreads(typeVar string, content string, truncated bool) *ElasticsearchHotThreads {
	this := ElasticsearchHotThreads{}
	this.Type = typeVar
	this.Content = content
	this.Truncated = truncated
	return &this
}

// NewElasticsearchHotThreadsWithDefaults instantiates a new ElasticsearchHotThreads object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewElasticsearchHotThreadsWithDefaults() *ElasticsearchHotThreads {
	this := ElasticsearchHotThreads{}
	return &this
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *ElasticsearchHotThreads) GetNode() string {
	if o == nil || o.Node == nil {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ElasticsearchHotThreads) GetNodeOk() (*string, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *ElasticsearchHotThreads) HasNode() bool {
	return o != nil && o.Node != nil
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *ElasticsearchHotThreads) SetNode(v string) {
	o.Node = &v
}

// GetType returns the Type field value.
func (o *ElasticsearchHotThreads) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchHotThreads) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ElasticsearchHotThreads) SetType(v string) {
	o.Type = v
}

// GetContent returns the Content field value.
func (o *ElasticsearchHotThreads) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchHotThreads) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *ElasticsearchHotThreads) SetContent(v string) {
	o.Content = v
}

// GetTruncated returns the Truncated field value.
func (o *ElasticsearchHotThreads) GetTruncated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value
// and a boolean to check if the value has been set.
func (o *ElasticsearchHotThreads) GetTruncatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Truncated, true
}

// SetTruncated sets field value.
func (o *ElasticsearchHotThreads) SetTruncated(v bool) {
	o.Truncated = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ElasticsearchHotThreads) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Node != nil {
		toSerialize["node"] = o.Node
	}
	toSerialize["type"] = o.Type
	toSerialize["content"] = o.Content
	toSerialize["truncated"] = o.Truncated

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ElasticsearchHotThreads) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Node      *string `json:"node,omitempty"`
		Type      *string `json:"type"`
		Content   *string `json:"content"`
		Truncated *bool   `json:"truncated"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.Truncated == nil {
		return fmt.Errorf("required field truncated missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"node", "type", "content", "truncated"})
	} else {
		return err
	}
	o.Node = all.Node
	o.Type = *all.Type
	o.Content = *all.Content
	o.Truncated = *all.Truncated

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
