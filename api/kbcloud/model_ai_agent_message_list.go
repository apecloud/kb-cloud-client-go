// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentMessageList struct {
	NextCursor *string               `json:"nextCursor,omitempty"`
	HasMore    *bool                 `json:"hasMore,omitempty"`
	Items      []AiAgentMessage      `json:"items,omitempty"`
	Included   *AiAgentArtifactBatch `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentMessageList instantiates a new AiAgentMessageList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentMessageList() *AiAgentMessageList {
	this := AiAgentMessageList{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *AiAgentMessageList) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessageList) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *AiAgentMessageList) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *AiAgentMessageList) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *AiAgentMessageList) GetHasMore() bool {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessageList) GetHasMoreOk() (*bool, bool) {
	if o == nil || o.HasMore == nil {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *AiAgentMessageList) HasHasMore() bool {
	return o != nil && o.HasMore != nil
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *AiAgentMessageList) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AiAgentMessageList) GetItems() []AiAgentMessage {
	if o == nil || o.Items == nil {
		var ret []AiAgentMessage
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessageList) GetItemsOk() (*[]AiAgentMessage, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AiAgentMessageList) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []AiAgentMessage and assigns it to the Items field.
func (o *AiAgentMessageList) SetItems(v []AiAgentMessage) {
	o.Items = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AiAgentMessageList) GetIncluded() AiAgentArtifactBatch {
	if o == nil || o.Included == nil {
		var ret AiAgentArtifactBatch
		return ret
	}
	return *o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentMessageList) GetIncludedOk() (*AiAgentArtifactBatch, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AiAgentMessageList) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given AiAgentArtifactBatch and assigns it to the Included field.
func (o *AiAgentMessageList) SetIncluded(v AiAgentArtifactBatch) {
	o.Included = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentMessageList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NextCursor != nil {
		toSerialize["nextCursor"] = o.NextCursor
	}
	if o.HasMore != nil {
		toSerialize["hasMore"] = o.HasMore
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentMessageList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextCursor *string               `json:"nextCursor,omitempty"`
		HasMore    *bool                 `json:"hasMore,omitempty"`
		Items      []AiAgentMessage      `json:"items,omitempty"`
		Included   *AiAgentArtifactBatch `json:"included,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nextCursor", "hasMore", "items", "included"})
	} else {
		return err
	}
	hasInvalidField := false
	o.NextCursor = all.NextCursor
	o.HasMore = all.HasMore
	o.Items = all.Items
	if all.Included != nil && all.Included.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Included = all.Included
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}
	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
