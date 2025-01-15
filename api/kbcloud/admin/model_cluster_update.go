// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// ClusterUpdate ClusterUpdate is the payload to update a KubeBlocks cluster
type ClusterUpdate struct {
	// The termination policy of cluster.
	TerminationPolicy *ClusterTerminationPolicy `json:"terminationPolicy,omitempty"`
	// Display name of cluster.
	DisplayName *string `json:"displayName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewClusterUpdate instantiates a new ClusterUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewClusterUpdate() *ClusterUpdate {
	this := ClusterUpdate{}
	var terminationPolicy ClusterTerminationPolicy = ClusterTerminationPolicyDelete
	this.TerminationPolicy = &terminationPolicy
	return &this
}

// NewClusterUpdateWithDefaults instantiates a new ClusterUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewClusterUpdateWithDefaults() *ClusterUpdate {
	this := ClusterUpdate{}
	var terminationPolicy ClusterTerminationPolicy = ClusterTerminationPolicyDelete
	this.TerminationPolicy = &terminationPolicy
	return &this
}

// GetTerminationPolicy returns the TerminationPolicy field value if set, zero value otherwise.
func (o *ClusterUpdate) GetTerminationPolicy() ClusterTerminationPolicy {
	if o == nil || o.TerminationPolicy == nil {
		var ret ClusterTerminationPolicy
		return ret
	}
	return *o.TerminationPolicy
}

// GetTerminationPolicyOk returns a tuple with the TerminationPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetTerminationPolicyOk() (*ClusterTerminationPolicy, bool) {
	if o == nil || o.TerminationPolicy == nil {
		return nil, false
	}
	return o.TerminationPolicy, true
}

// HasTerminationPolicy returns a boolean if a field has been set.
func (o *ClusterUpdate) HasTerminationPolicy() bool {
	return o != nil && o.TerminationPolicy != nil
}

// SetTerminationPolicy gets a reference to the given ClusterTerminationPolicy and assigns it to the TerminationPolicy field.
func (o *ClusterUpdate) SetTerminationPolicy(v ClusterTerminationPolicy) {
	o.TerminationPolicy = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ClusterUpdate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClusterUpdate) HasDisplayName() bool {
	return o != nil && o.DisplayName != nil
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ClusterUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ClusterUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.TerminationPolicy != nil {
		toSerialize["terminationPolicy"] = o.TerminationPolicy
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ClusterUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TerminationPolicy *ClusterTerminationPolicy `json:"terminationPolicy,omitempty"`
		DisplayName       *string                   `json:"displayName,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"terminationPolicy", "displayName"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.TerminationPolicy != nil && !all.TerminationPolicy.IsValid() {
		hasInvalidField = true
	} else {
		o.TerminationPolicy = all.TerminationPolicy
	}
	o.DisplayName = all.DisplayName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
