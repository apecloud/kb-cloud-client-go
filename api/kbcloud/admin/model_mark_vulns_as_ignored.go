// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// MarkVulnsAsIgnored the data of the mark request
type MarkVulnsAsIgnored struct {
	// the CVE ID of the vulnerability
	CveId string `json:"cveId"`
	// the engine name of the vulnerability
	EngineName string `json:"engineName"`
	// the product name of the vulnerability
	ProductName string `json:"productName"`
	// the version of the product which is affected by the vulnerability
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMarkVulnsAsIgnored instantiates a new MarkVulnsAsIgnored object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMarkVulnsAsIgnored(cveId string, engineName string, productName string, version string) *MarkVulnsAsIgnored {
	this := MarkVulnsAsIgnored{}
	this.CveId = cveId
	this.EngineName = engineName
	this.ProductName = productName
	this.Version = version
	return &this
}

// NewMarkVulnsAsIgnoredWithDefaults instantiates a new MarkVulnsAsIgnored object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMarkVulnsAsIgnoredWithDefaults() *MarkVulnsAsIgnored {
	this := MarkVulnsAsIgnored{}
	return &this
}

// GetCveId returns the CveId field value.
func (o *MarkVulnsAsIgnored) GetCveId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CveId
}

// GetCveIdOk returns a tuple with the CveId field value
// and a boolean to check if the value has been set.
func (o *MarkVulnsAsIgnored) GetCveIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CveId, true
}

// SetCveId sets field value.
func (o *MarkVulnsAsIgnored) SetCveId(v string) {
	o.CveId = v
}

// GetEngineName returns the EngineName field value.
func (o *MarkVulnsAsIgnored) GetEngineName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value
// and a boolean to check if the value has been set.
func (o *MarkVulnsAsIgnored) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EngineName, true
}

// SetEngineName sets field value.
func (o *MarkVulnsAsIgnored) SetEngineName(v string) {
	o.EngineName = v
}

// GetProductName returns the ProductName field value.
func (o *MarkVulnsAsIgnored) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *MarkVulnsAsIgnored) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value.
func (o *MarkVulnsAsIgnored) SetProductName(v string) {
	o.ProductName = v
}

// GetVersion returns the Version field value.
func (o *MarkVulnsAsIgnored) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *MarkVulnsAsIgnored) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *MarkVulnsAsIgnored) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MarkVulnsAsIgnored) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["cveId"] = o.CveId
	toSerialize["engineName"] = o.EngineName
	toSerialize["productName"] = o.ProductName
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MarkVulnsAsIgnored) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CveId       *string `json:"cveId"`
		EngineName  *string `json:"engineName"`
		ProductName *string `json:"productName"`
		Version     *string `json:"version"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.CveId == nil {
		return fmt.Errorf("required field cveId missing")
	}
	if all.EngineName == nil {
		return fmt.Errorf("required field engineName missing")
	}
	if all.ProductName == nil {
		return fmt.Errorf("required field productName missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cveId", "engineName", "productName", "version"})
	} else {
		return err
	}
	o.CveId = *all.CveId
	o.EngineName = *all.EngineName
	o.ProductName = *all.ProductName
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
