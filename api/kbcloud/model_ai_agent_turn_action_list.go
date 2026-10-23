// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type AiAgentTurnActionList struct {
	Items []AiAgentTurnAction `json:"items"`
	// Raw action cursor for the next page; independent of coalesced item IDs and timestamps.
	NextCursor *string `json:"nextCursor,omitempty"`
	// Whether more actions are available. Pages finish the last turn before coalescing, so item counts must not be used to detect the final page.
	HasMore *bool `json:"hasMore,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentTurnActionList instantiates a new AiAgentTurnActionList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentTurnActionList(items []AiAgentTurnAction) *AiAgentTurnActionList {
	this := AiAgentTurnActionList{}
	this.Items = items
	return &this
}

// NewAiAgentTurnActionListWithDefaults instantiates a new AiAgentTurnActionList object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiAgentTurnActionListWithDefaults() *AiAgentTurnActionList {
	this := AiAgentTurnActionList{}
	return &this
}

// GetItems returns the Items field value.
func (o *AiAgentTurnActionList) GetItems() []AiAgentTurnAction {
	if o == nil {
		var ret []AiAgentTurnAction
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AiAgentTurnActionList) GetItemsOk() (*[]AiAgentTurnAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value.
func (o *AiAgentTurnActionList) SetItems(v []AiAgentTurnAction) {
	o.Items = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *AiAgentTurnActionList) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTurnActionList) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *AiAgentTurnActionList) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *AiAgentTurnActionList) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *AiAgentTurnActionList) GetHasMore() bool {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentTurnActionList) GetHasMoreOk() (*bool, bool) {
	if o == nil || o.HasMore == nil {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *AiAgentTurnActionList) HasHasMore() bool {
	return o != nil && o.HasMore != nil
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *AiAgentTurnActionList) SetHasMore(v bool) {
	o.HasMore = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentTurnActionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["items"] = o.Items
	if o.NextCursor != nil {
		toSerialize["nextCursor"] = o.NextCursor
	}
	if o.HasMore != nil {
		toSerialize["hasMore"] = o.HasMore
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentTurnActionList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items      *[]AiAgentTurnAction `json:"items"`
		NextCursor *string              `json:"nextCursor,omitempty"`
		HasMore    *bool                `json:"hasMore,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Items == nil {
		return fmt.Errorf("required field items missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "nextCursor", "hasMore"})
	} else {
		return err
	}
	o.Items = *all.Items
	o.NextCursor = all.NextCursor
	o.HasMore = all.HasMore

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
