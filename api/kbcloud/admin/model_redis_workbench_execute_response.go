// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RedisWorkbenchExecuteResponse struct {
	Status     *RedisOperationStatus  `json:"status,omitempty"`
	DurationMs *int64                 `json:"durationMs,omitempty"`
	Result     map[string]interface{} `json:"result,omitempty"`
	Guard      *RedisOperationGuard   `json:"guard,omitempty"`
	NodeId     *string                `json:"nodeId,omitempty"`
	Slot       *int64                 `json:"slot,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRedisWorkbenchExecuteResponse instantiates a new RedisWorkbenchExecuteResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRedisWorkbenchExecuteResponse() *RedisWorkbenchExecuteResponse {
	this := RedisWorkbenchExecuteResponse{}
	return &this
}

// NewRedisWorkbenchExecuteResponseWithDefaults instantiates a new RedisWorkbenchExecuteResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRedisWorkbenchExecuteResponseWithDefaults() *RedisWorkbenchExecuteResponse {
	this := RedisWorkbenchExecuteResponse{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RedisWorkbenchExecuteResponse) GetStatus() RedisOperationStatus {
	if o == nil || o.Status == nil {
		var ret RedisOperationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisWorkbenchExecuteResponse) GetStatusOk() (*RedisOperationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RedisWorkbenchExecuteResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given RedisOperationStatus and assigns it to the Status field.
func (o *RedisWorkbenchExecuteResponse) SetStatus(v RedisOperationStatus) {
	o.Status = &v
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise.
func (o *RedisWorkbenchExecuteResponse) GetDurationMs() int64 {
	if o == nil || o.DurationMs == nil {
		var ret int64
		return ret
	}
	return *o.DurationMs
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisWorkbenchExecuteResponse) GetDurationMsOk() (*int64, bool) {
	if o == nil || o.DurationMs == nil {
		return nil, false
	}
	return o.DurationMs, true
}

// HasDurationMs returns a boolean if a field has been set.
func (o *RedisWorkbenchExecuteResponse) HasDurationMs() bool {
	return o != nil && o.DurationMs != nil
}

// SetDurationMs gets a reference to the given int64 and assigns it to the DurationMs field.
func (o *RedisWorkbenchExecuteResponse) SetDurationMs(v int64) {
	o.DurationMs = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *RedisWorkbenchExecuteResponse) GetResult() map[string]interface{} {
	if o == nil || o.Result == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisWorkbenchExecuteResponse) GetResultOk() (*map[string]interface{}, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return &o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *RedisWorkbenchExecuteResponse) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given map[string]interface{} and assigns it to the Result field.
func (o *RedisWorkbenchExecuteResponse) SetResult(v map[string]interface{}) {
	o.Result = v
}

// GetGuard returns the Guard field value if set, zero value otherwise.
func (o *RedisWorkbenchExecuteResponse) GetGuard() RedisOperationGuard {
	if o == nil || o.Guard == nil {
		var ret RedisOperationGuard
		return ret
	}
	return *o.Guard
}

// GetGuardOk returns a tuple with the Guard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisWorkbenchExecuteResponse) GetGuardOk() (*RedisOperationGuard, bool) {
	if o == nil || o.Guard == nil {
		return nil, false
	}
	return o.Guard, true
}

// HasGuard returns a boolean if a field has been set.
func (o *RedisWorkbenchExecuteResponse) HasGuard() bool {
	return o != nil && o.Guard != nil
}

// SetGuard gets a reference to the given RedisOperationGuard and assigns it to the Guard field.
func (o *RedisWorkbenchExecuteResponse) SetGuard(v RedisOperationGuard) {
	o.Guard = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *RedisWorkbenchExecuteResponse) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisWorkbenchExecuteResponse) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *RedisWorkbenchExecuteResponse) HasNodeId() bool {
	return o != nil && o.NodeId != nil
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *RedisWorkbenchExecuteResponse) SetNodeId(v string) {
	o.NodeId = &v
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *RedisWorkbenchExecuteResponse) GetSlot() int64 {
	if o == nil || o.Slot == nil {
		var ret int64
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedisWorkbenchExecuteResponse) GetSlotOk() (*int64, bool) {
	if o == nil || o.Slot == nil {
		return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *RedisWorkbenchExecuteResponse) HasSlot() bool {
	return o != nil && o.Slot != nil
}

// SetSlot gets a reference to the given int64 and assigns it to the Slot field.
func (o *RedisWorkbenchExecuteResponse) SetSlot(v int64) {
	o.Slot = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RedisWorkbenchExecuteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.DurationMs != nil {
		toSerialize["durationMs"] = o.DurationMs
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Guard != nil {
		toSerialize["guard"] = o.Guard
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Slot != nil {
		toSerialize["slot"] = o.Slot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RedisWorkbenchExecuteResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status     *RedisOperationStatus  `json:"status,omitempty"`
		DurationMs *int64                 `json:"durationMs,omitempty"`
		Result     map[string]interface{} `json:"result,omitempty"`
		Guard      *RedisOperationGuard   `json:"guard,omitempty"`
		NodeId     *string                `json:"nodeId,omitempty"`
		Slot       *int64                 `json:"slot,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "durationMs", "result", "guard", "nodeId", "slot"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.DurationMs = all.DurationMs
	o.Result = all.Result
	if all.Guard != nil && all.Guard.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Guard = all.Guard
	o.NodeId = all.NodeId
	o.Slot = all.Slot

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
