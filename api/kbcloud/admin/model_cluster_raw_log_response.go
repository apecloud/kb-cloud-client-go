// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterRawLogResponse Cluster raw log is the raw log of the cluster
type ClusterRawLogResponse struct {
	Filenames []string            `json:"filenames,omitempty"`
	Items     []ClusterRawLogItem `json:"items,omitempty"`
	// Cluster log pagination is the pagination of the cluster log
	Pagination *ClusterLogPagination `json:"pagination,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterRawLogResponse instantiates a new ClusterRawLogResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterRawLogResponse() *ClusterRawLogResponse {
	this := ClusterRawLogResponse{}
	return &this
}

// NewClusterRawLogResponseWithDefaults instantiates a new ClusterRawLogResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterRawLogResponseWithDefaults() *ClusterRawLogResponse {
	this := ClusterRawLogResponse{}
	return &this
}

// GetFilenames returns the Filenames field value if set, zero value otherwise.
func (o *ClusterRawLogResponse) GetFilenames() []string {
	if o == nil || o.Filenames == nil {
		var ret []string
		return ret
	}
	return o.Filenames
}

// GetFilenamesOk returns a tuple with the Filenames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRawLogResponse) GetFilenamesOk() (*[]string, bool) {
	if o == nil || o.Filenames == nil {
		return nil, false
	}
	return &o.Filenames, true
}

// HasFilenames returns a boolean if a field has been set.
func (o *ClusterRawLogResponse) HasFilenames() bool {
	return o != nil && o.Filenames != nil
}

// SetFilenames gets a reference to the given []string and assigns it to the Filenames field.
func (o *ClusterRawLogResponse) SetFilenames(v []string) {
	o.Filenames = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ClusterRawLogResponse) GetItems() []ClusterRawLogItem {
	if o == nil || o.Items == nil {
		var ret []ClusterRawLogItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRawLogResponse) GetItemsOk() (*[]ClusterRawLogItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return &o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ClusterRawLogResponse) HasItems() bool {
	return o != nil && o.Items != nil
}

// SetItems gets a reference to the given []ClusterRawLogItem and assigns it to the Items field.
func (o *ClusterRawLogResponse) SetItems(v []ClusterRawLogItem) {
	o.Items = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ClusterRawLogResponse) GetPagination() ClusterLogPagination {
	if o == nil || o.Pagination == nil {
		var ret ClusterLogPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRawLogResponse) GetPaginationOk() (*ClusterLogPagination, bool) {
	if o == nil || o.Pagination == nil {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ClusterRawLogResponse) HasPagination() bool {
	return o != nil && o.Pagination != nil
}

// SetPagination gets a reference to the given ClusterLogPagination and assigns it to the Pagination field.
func (o *ClusterRawLogResponse) SetPagination(v ClusterLogPagination) {
	o.Pagination = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterRawLogResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Filenames != nil {
		toSerialize["filenames"] = o.Filenames
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Pagination != nil {
		toSerialize["pagination"] = o.Pagination
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterRawLogResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Filenames  []string              `json:"filenames,omitempty"`
		Items      []ClusterRawLogItem   `json:"items,omitempty"`
		Pagination *ClusterLogPagination `json:"pagination,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"filenames", "items", "pagination"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Filenames = all.Filenames
	o.Items = all.Items
	if all.Pagination != nil && all.Pagination.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagination = all.Pagination

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
