// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// CustomOpsTask customOpsTask is the information of custom ops task
type CustomOpsTask struct {
	// the pod name
	ObjectKey *string `json:"objectKey,omitempty"`
	// namespace of the task
	Namespace *string `json:"namespace,omitempty"`
	// status of the task
	Status *string `json:"status,omitempty"`
	// target pod name of the task
	TargetPodName *string `json:"targetPodName,omitempty"`
	// retries of the task
	Retries *int32 `json:"retries,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomOpsTask instantiates a new CustomOpsTask object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomOpsTask() *CustomOpsTask {
	this := CustomOpsTask{}
	return &this
}

// NewCustomOpsTaskWithDefaults instantiates a new CustomOpsTask object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomOpsTaskWithDefaults() *CustomOpsTask {
	this := CustomOpsTask{}
	return &this
}

// GetObjectKey returns the ObjectKey field value if set, zero value otherwise.
func (o *CustomOpsTask) GetObjectKey() string {
	if o == nil || o.ObjectKey == nil {
		var ret string
		return ret
	}
	return *o.ObjectKey
}

// GetObjectKeyOk returns a tuple with the ObjectKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOpsTask) GetObjectKeyOk() (*string, bool) {
	if o == nil || o.ObjectKey == nil {
		return nil, false
	}
	return o.ObjectKey, true
}

// HasObjectKey returns a boolean if a field has been set.
func (o *CustomOpsTask) HasObjectKey() bool {
	return o != nil && o.ObjectKey != nil
}

// SetObjectKey gets a reference to the given string and assigns it to the ObjectKey field.
func (o *CustomOpsTask) SetObjectKey(v string) {
	o.ObjectKey = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CustomOpsTask) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOpsTask) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CustomOpsTask) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *CustomOpsTask) SetNamespace(v string) {
	o.Namespace = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomOpsTask) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOpsTask) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomOpsTask) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CustomOpsTask) SetStatus(v string) {
	o.Status = &v
}

// GetTargetPodName returns the TargetPodName field value if set, zero value otherwise.
func (o *CustomOpsTask) GetTargetPodName() string {
	if o == nil || o.TargetPodName == nil {
		var ret string
		return ret
	}
	return *o.TargetPodName
}

// GetTargetPodNameOk returns a tuple with the TargetPodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOpsTask) GetTargetPodNameOk() (*string, bool) {
	if o == nil || o.TargetPodName == nil {
		return nil, false
	}
	return o.TargetPodName, true
}

// HasTargetPodName returns a boolean if a field has been set.
func (o *CustomOpsTask) HasTargetPodName() bool {
	return o != nil && o.TargetPodName != nil
}

// SetTargetPodName gets a reference to the given string and assigns it to the TargetPodName field.
func (o *CustomOpsTask) SetTargetPodName(v string) {
	o.TargetPodName = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *CustomOpsTask) GetRetries() int32 {
	if o == nil || o.Retries == nil {
		var ret int32
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomOpsTask) GetRetriesOk() (*int32, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *CustomOpsTask) HasRetries() bool {
	return o != nil && o.Retries != nil
}

// SetRetries gets a reference to the given int32 and assigns it to the Retries field.
func (o *CustomOpsTask) SetRetries(v int32) {
	o.Retries = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomOpsTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ObjectKey != nil {
		toSerialize["objectKey"] = o.ObjectKey
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TargetPodName != nil {
		toSerialize["targetPodName"] = o.TargetPodName
	}
	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomOpsTask) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ObjectKey     *string `json:"objectKey,omitempty"`
		Namespace     *string `json:"namespace,omitempty"`
		Status        *string `json:"status,omitempty"`
		TargetPodName *string `json:"targetPodName,omitempty"`
		Retries       *int32  `json:"retries,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"objectKey", "namespace", "status", "targetPodName", "retries"})
	} else {
		return err
	}
	o.ObjectKey = all.ObjectKey
	o.Namespace = all.Namespace
	o.Status = all.Status
	o.TargetPodName = all.TargetPodName
	o.Retries = all.Retries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
