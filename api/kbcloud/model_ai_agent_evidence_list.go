// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type AiAgentEvidenceList struct {
	NextCursor *string                  `json:"nextCursor,omitempty"`
	HasMore    *bool                    `json:"hasMore,omitempty"`
	Items      []AiAgentEvidenceSummary `json:"items,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiAgentEvidenceList instantiates a new AiAgentEvidenceList object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiAgentEvidenceList() *AiAgentEvidenceList {
	this := AiAgentEvidenceList{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *AiAgentEvidenceList) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceList) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *AiAgentEvidenceList) HasNextCursor() bool {
	return o != nil && o.NextCursor != nil
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *AiAgentEvidenceList) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *AiAgentEvidenceList) GetHasMore() bool {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceList) GetHasMoreOk() (*bool, bool) {
	if o == nil || o.HasMore == nil {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *AiAgentEvidenceList) HasHasMore() bool {
	return o != nil && o.HasMore != nil
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *AiAgentEvidenceList) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *AiAgentEvidenceList) GetItems() []AiAgentEvidenceSummary {
	if o == nil || o.Items == nil {
		var ret []AiAgentEvidenceSummary
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiAgentEvidenceList) GetItemsOk() (*[]AiAgentEvidenceSummary, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *AiAgentEvidenceList) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []AiAgentEvidenceSummary and assigns it to the Items field.
func (o *AiAgentEvidenceList) SetItems(v []AiAgentEvidenceSummary) {
	o.Items = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiAgentEvidenceList) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiAgentEvidenceList) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextCursor *string                  `json:"nextCursor,omitempty"`
		HasMore    *bool                    `json:"hasMore,omitempty"`
		Items      []AiAgentEvidenceSummary `json:"items,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nextCursor", "hasMore", "items"})
	} else {
		return err
	}
	o.NextCursor = all.NextCursor
	o.HasMore = all.HasMore
	o.Items = all.Items
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
