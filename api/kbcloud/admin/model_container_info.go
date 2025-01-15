// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ContainerInfo Container information
type ContainerInfo struct {
	// Container name
	Name *string `json:"name,omitempty"`
	// Container image
	Image *string `json:"image,omitempty"`
	// Container ready status
	Ready *bool `json:"ready,omitempty"`
	// Container restart count
	RestartCount *int32 `json:"restartCount,omitempty"`
	// Liveness probe configuration
	LivenessProbe interface{} `json:"livenessProbe,omitempty"`
	// Readiness probe configuration
	ReadinessProbe interface{} `json:"readinessProbe,omitempty"`
	// Container resource usage
	Resources interface{} `json:"resources,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerInfo instantiates a new ContainerInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerInfo() *ContainerInfo {
	this := ContainerInfo{}
	return &this
}

// NewContainerInfoWithDefaults instantiates a new ContainerInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerInfoWithDefaults() *ContainerInfo {
	this := ContainerInfo{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainerInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainerInfo) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainerInfo) SetName(v string) {
	o.Name = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *ContainerInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *ContainerInfo) HasImage() bool {
	return o != nil && o.Image != nil
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *ContainerInfo) SetImage(v string) {
	o.Image = &v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *ContainerInfo) GetReady() bool {
	if o == nil || o.Ready == nil {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerInfo) GetReadyOk() (*bool, bool) {
	if o == nil || o.Ready == nil {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *ContainerInfo) HasReady() bool {
	return o != nil && o.Ready != nil
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *ContainerInfo) SetReady(v bool) {
	o.Ready = &v
}

// GetRestartCount returns the RestartCount field value if set, zero value otherwise.
func (o *ContainerInfo) GetRestartCount() int32 {
	if o == nil || o.RestartCount == nil {
		var ret int32
		return ret
	}
	return *o.RestartCount
}

// GetRestartCountOk returns a tuple with the RestartCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerInfo) GetRestartCountOk() (*int32, bool) {
	if o == nil || o.RestartCount == nil {
		return nil, false
	}
	return o.RestartCount, true
}

// HasRestartCount returns a boolean if a field has been set.
func (o *ContainerInfo) HasRestartCount() bool {
	return o != nil && o.RestartCount != nil
}

// SetRestartCount gets a reference to the given int32 and assigns it to the RestartCount field.
func (o *ContainerInfo) SetRestartCount(v int32) {
	o.RestartCount = &v
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *ContainerInfo) GetLivenessProbe() interface{} {
	if o == nil || o.LivenessProbe == nil {
		var ret interface{}
		return ret
	}
	return o.LivenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerInfo) GetLivenessProbeOk() (*interface{}, bool) {
	if o == nil || o.LivenessProbe == nil {
		return nil, false
	}
	return &o.LivenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *ContainerInfo) HasLivenessProbe() bool {
	return o != nil && o.LivenessProbe != nil
}

// SetLivenessProbe gets a reference to the given interface{} and assigns it to the LivenessProbe field.
func (o *ContainerInfo) SetLivenessProbe(v interface{}) {
	o.LivenessProbe = v
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *ContainerInfo) GetReadinessProbe() interface{} {
	if o == nil || o.ReadinessProbe == nil {
		var ret interface{}
		return ret
	}
	return o.ReadinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerInfo) GetReadinessProbeOk() (*interface{}, bool) {
	if o == nil || o.ReadinessProbe == nil {
		return nil, false
	}
	return &o.ReadinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *ContainerInfo) HasReadinessProbe() bool {
	return o != nil && o.ReadinessProbe != nil
}

// SetReadinessProbe gets a reference to the given interface{} and assigns it to the ReadinessProbe field.
func (o *ContainerInfo) SetReadinessProbe(v interface{}) {
	o.ReadinessProbe = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *ContainerInfo) GetResources() interface{} {
	if o == nil || o.Resources == nil {
		var ret interface{}
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerInfo) GetResourcesOk() (*interface{}, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return &o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *ContainerInfo) HasResources() bool {
	return o != nil && o.Resources != nil
}

// SetResources gets a reference to the given interface{} and assigns it to the Resources field.
func (o *ContainerInfo) SetResources(v interface{}) {
	o.Resources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.Ready != nil {
		toSerialize["ready"] = o.Ready
	}
	if o.RestartCount != nil {
		toSerialize["restartCount"] = o.RestartCount
	}
	if o.LivenessProbe != nil {
		toSerialize["livenessProbe"] = o.LivenessProbe
	}
	if o.ReadinessProbe != nil {
		toSerialize["readinessProbe"] = o.ReadinessProbe
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainerInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name           *string     `json:"name,omitempty"`
		Image          *string     `json:"image,omitempty"`
		Ready          *bool       `json:"ready,omitempty"`
		RestartCount   *int32      `json:"restartCount,omitempty"`
		LivenessProbe  interface{} `json:"livenessProbe,omitempty"`
		ReadinessProbe interface{} `json:"readinessProbe,omitempty"`
		Resources      interface{} `json:"resources,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "image", "ready", "restartCount", "livenessProbe", "readinessProbe", "resources"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Image = all.Image
	o.Ready = all.Ready
	o.RestartCount = all.RestartCount
	o.LivenessProbe = all.LivenessProbe
	o.ReadinessProbe = all.ReadinessProbe
	o.Resources = all.Resources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
