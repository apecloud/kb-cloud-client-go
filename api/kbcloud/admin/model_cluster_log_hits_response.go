// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterLogHitsResponse Log hits histogram response with time-bucketed counts
type ClusterLogHitsResponse struct {
	Items []ClusterLogHitsItem `json:"items,omitempty"`
	Total *int64               `json:"total,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterLogHitsResponse instantiates a new ClusterLogHitsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterLogHitsResponse() *ClusterLogHitsResponse {
	this := ClusterLogHitsResponse{}
	return &this
}

// NewClusterLogHitsResponseWithDefaults instantiates a new ClusterLogHitsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterLogHitsResponseWithDefaults() *ClusterLogHitsResponse {
	this := ClusterLogHitsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ClusterLogHitsResponse) GetItems() []ClusterLogHitsItem {
	if o == nil || o.Items == nil {
		var ret []ClusterLogHitsItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogHitsResponse) GetItemsOk() (*[]ClusterLogHitsItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ClusterLogHitsResponse) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []ClusterLogHitsItem and assigns it to the Items field.
func (o *ClusterLogHitsResponse) SetItems(v []ClusterLogHitsItem) {
	o.Items = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ClusterLogHitsResponse) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterLogHitsResponse) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ClusterLogHitsResponse) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *ClusterLogHitsResponse) SetTotal(v int64) {
	o.Total = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterLogHitsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterLogHitsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Items []ClusterLogHitsItem `json:"items,omitempty"`
		Total *int64               `json:"total,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"items", "total"})
	} else {
		return err
	}
	o.Items = all.Items
	o.Total = all.Total

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
