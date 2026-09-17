// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

// RedisKeyScanProgress Redis key browser scan progress for the current page.
type RedisKeyScanProgress struct {
	// Scan scope, either node or cluster.
	Scope *string `json:"scope,omitempty"`
	// Total keys in the scan scope when available.
	Total *int64 `json:"total,omitempty"`
	// Number of key names returned by SCAN in this page before type filtering.
	Scanned *int64 `json:"scanned,omitempty"`
	// Number of key summaries loaded in this page after filtering and metadata lookup.
	Loaded *int64 `json:"loaded,omitempty"`
	// Whether the scan cursor is complete after this page.
	Complete *bool `json:"complete,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisKeyScanProgress instantiates a new RedisKeyScanProgress object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisKeyScanProgress() *RedisKeyScanProgress {
	this := RedisKeyScanProgress{}
	return &this
}

// NewRedisKeyScanProgressWithDefaults instantiates a new RedisKeyScanProgress object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisKeyScanProgressWithDefaults() *RedisKeyScanProgress {
	this := RedisKeyScanProgress{}
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *RedisKeyScanProgress) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyScanProgress) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *RedisKeyScanProgress) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *RedisKeyScanProgress) SetScope(v string) {
	o.Scope = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *RedisKeyScanProgress) GetTotal() int64 {
	if o == nil || o.Total == nil {
		var ret int64
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyScanProgress) GetTotalOk() (*int64, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *RedisKeyScanProgress) HasTotal() bool {
	return o != nil && o.Total != nil
}

// SetTotal gets a reference to the given int64 and assigns it to the Total field.
func (o *RedisKeyScanProgress) SetTotal(v int64) {
	o.Total = &v
}

// GetScanned returns the Scanned field value if set, zero value otherwise.
func (o *RedisKeyScanProgress) GetScanned() int64 {
	if o == nil || o.Scanned == nil {
		var ret int64
		return ret
	}
	return *o.Scanned
}

// GetScannedOk returns a tuple with the Scanned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyScanProgress) GetScannedOk() (*int64, bool) {
	if o == nil || o.Scanned == nil {
		return nil, false
	}
	return o.Scanned, true
}

// HasScanned returns a boolean if a field has been set.
func (o *RedisKeyScanProgress) HasScanned() bool {
	return o != nil && o.Scanned != nil
}

// SetScanned gets a reference to the given int64 and assigns it to the Scanned field.
func (o *RedisKeyScanProgress) SetScanned(v int64) {
	o.Scanned = &v
}

// GetLoaded returns the Loaded field value if set, zero value otherwise.
func (o *RedisKeyScanProgress) GetLoaded() int64 {
	if o == nil || o.Loaded == nil {
		var ret int64
		return ret
	}
	return *o.Loaded
}

// GetLoadedOk returns a tuple with the Loaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyScanProgress) GetLoadedOk() (*int64, bool) {
	if o == nil || o.Loaded == nil {
		return nil, false
	}
	return o.Loaded, true
}

// HasLoaded returns a boolean if a field has been set.
func (o *RedisKeyScanProgress) HasLoaded() bool {
	return o != nil && o.Loaded != nil
}

// SetLoaded gets a reference to the given int64 and assigns it to the Loaded field.
func (o *RedisKeyScanProgress) SetLoaded(v int64) {
	o.Loaded = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *RedisKeyScanProgress) GetComplete() bool {
	if o == nil || o.Complete == nil {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisKeyScanProgress) GetCompleteOk() (*bool, bool) {
	if o == nil || o.Complete == nil {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *RedisKeyScanProgress) HasComplete() bool {
	return o != nil && o.Complete != nil
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *RedisKeyScanProgress) SetComplete(v bool) {
	o.Complete = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisKeyScanProgress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if o.Scanned != nil {
		toSerialize["scanned"] = o.Scanned
	}
	if o.Loaded != nil {
		toSerialize["loaded"] = o.Loaded
	}
	if o.Complete != nil {
		toSerialize["complete"] = o.Complete
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisKeyScanProgress) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Scope    *string `json:"scope,omitempty"`
		Total    *int64  `json:"total,omitempty"`
		Scanned  *int64  `json:"scanned,omitempty"`
		Loaded   *int64  `json:"loaded,omitempty"`
		Complete *bool   `json:"complete,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"scope", "total", "scanned", "loaded", "complete"})
	} else {
		return err
	}
	o.Scope = all.Scope
	o.Total = all.Total
	o.Scanned = all.Scanned
	o.Loaded = all.Loaded
	o.Complete = all.Complete

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
