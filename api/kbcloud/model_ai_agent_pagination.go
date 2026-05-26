// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentPagination struct {
	NextCursor *string `json:"nextCursor,omitempty"`
	HasMore    bool    `json:"hasMore"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentPagination instantiates a new AiAgentPagination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentPagination(hasMore bool) *AiAgentPagination {
	this := AiAgentPagination{}
	this.HasMore = hasMore
	return &this
}

// NewAiAgentPaginationWithDefaults instantiates a new AiAgentPagination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentPaginationWithDefaults() *AiAgentPagination {
	this := AiAgentPagination{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *AiAgentPagination) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentPagination) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *AiAgentPagination) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *AiAgentPagination) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetHasMore returns the HasMore field value.
func (o *AiAgentPagination) GetHasMore() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value
// and a boolean to check if the value has been set.
func (o *AiAgentPagination) GetHasMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMore, true
}

// SetHasMore sets field value.
func (o *AiAgentPagination) SetHasMore(v bool) {
	o.HasMore = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentPagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NextCursor != nil {
		toSerialize["nextCursor"] = o.NextCursor
	}
	toSerialize["hasMore"] = o.HasMore

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentPagination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextCursor *string `json:"nextCursor,omitempty"`
		HasMore    *bool   `json:"hasMore"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.HasMore == nil {
		return fmt.Errorf("required field hasMore missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nextCursor", "hasMore"})
	} else {
		return err
	}
	o.NextCursor = all.NextCursor
	o.HasMore = *all.HasMore

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
