// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// Affinity Rules for pod scheduling based on node and pod relationships.
type Affinity struct {
	// Rules for scheduling pods based on node labels.
	NodeAffinity *NodeAffinity `json:"nodeAffinity,omitempty"`
	// Rules for scheduling pods to co-locate with other pods.
	PodAffinity *PodAffinity `json:"podAffinity,omitempty"`
	// Rules for scheduling pods to avoid co-locating with other pods.
	PodAntiAffinity *PodAntiAffinity `json:"podAntiAffinity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAffinity instantiates a new Affinity object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAffinity() *Affinity {
	this := Affinity{}
	return &this
}

// NewAffinityWithDefaults instantiates a new Affinity object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAffinityWithDefaults() *Affinity {
	this := Affinity{}
	return &this
}

// GetNodeAffinity returns the NodeAffinity field value if set, zero value otherwise.
func (o *Affinity) GetNodeAffinity() NodeAffinity {
	if o == nil || o.NodeAffinity == nil {
		var ret NodeAffinity
		return ret
	}
	return *o.NodeAffinity
}

// GetNodeAffinityOk returns a tuple with the NodeAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affinity) GetNodeAffinityOk() (*NodeAffinity, bool) {
	if o == nil || o.NodeAffinity == nil {
		return nil, false
	}
	return o.NodeAffinity, true
}

// HasNodeAffinity returns a boolean if a field has been set.
func (o *Affinity) HasNodeAffinity() bool {
	return o != nil && o.NodeAffinity != nil
}

// SetNodeAffinity gets a reference to the given NodeAffinity and assigns it to the NodeAffinity field.
func (o *Affinity) SetNodeAffinity(v NodeAffinity) {
	o.NodeAffinity = &v
}

// GetPodAffinity returns the PodAffinity field value if set, zero value otherwise.
func (o *Affinity) GetPodAffinity() PodAffinity {
	if o == nil || o.PodAffinity == nil {
		var ret PodAffinity
		return ret
	}
	return *o.PodAffinity
}

// GetPodAffinityOk returns a tuple with the PodAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affinity) GetPodAffinityOk() (*PodAffinity, bool) {
	if o == nil || o.PodAffinity == nil {
		return nil, false
	}
	return o.PodAffinity, true
}

// HasPodAffinity returns a boolean if a field has been set.
func (o *Affinity) HasPodAffinity() bool {
	return o != nil && o.PodAffinity != nil
}

// SetPodAffinity gets a reference to the given PodAffinity and assigns it to the PodAffinity field.
func (o *Affinity) SetPodAffinity(v PodAffinity) {
	o.PodAffinity = &v
}

// GetPodAntiAffinity returns the PodAntiAffinity field value if set, zero value otherwise.
func (o *Affinity) GetPodAntiAffinity() PodAntiAffinity {
	if o == nil || o.PodAntiAffinity == nil {
		var ret PodAntiAffinity
		return ret
	}
	return *o.PodAntiAffinity
}

// GetPodAntiAffinityOk returns a tuple with the PodAntiAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Affinity) GetPodAntiAffinityOk() (*PodAntiAffinity, bool) {
	if o == nil || o.PodAntiAffinity == nil {
		return nil, false
	}
	return o.PodAntiAffinity, true
}

// HasPodAntiAffinity returns a boolean if a field has been set.
func (o *Affinity) HasPodAntiAffinity() bool {
	return o != nil && o.PodAntiAffinity != nil
}

// SetPodAntiAffinity gets a reference to the given PodAntiAffinity and assigns it to the PodAntiAffinity field.
func (o *Affinity) SetPodAntiAffinity(v PodAntiAffinity) {
	o.PodAntiAffinity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Affinity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.NodeAffinity != nil {
		toSerialize["nodeAffinity"] = o.NodeAffinity
	}
	if o.PodAffinity != nil {
		toSerialize["podAffinity"] = o.PodAffinity
	}
	if o.PodAntiAffinity != nil {
		toSerialize["podAntiAffinity"] = o.PodAntiAffinity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Affinity) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NodeAffinity    *NodeAffinity    `json:"nodeAffinity,omitempty"`
		PodAffinity     *PodAffinity     `json:"podAffinity,omitempty"`
		PodAntiAffinity *PodAntiAffinity `json:"podAntiAffinity,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"nodeAffinity", "podAffinity", "podAntiAffinity"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.NodeAffinity != nil && all.NodeAffinity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NodeAffinity = all.NodeAffinity
	if all.PodAffinity != nil && all.PodAffinity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PodAffinity = all.PodAffinity
	if all.PodAntiAffinity != nil && all.PodAntiAffinity.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.PodAntiAffinity = all.PodAntiAffinity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
