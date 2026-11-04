// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import (
	"fmt"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

type DolphinDBLicenseInfo struct {
	// Native expiration calendar date in YYYY.MM.DD format.
	Expiration string `json:"expiration"`
	// Native memory quota in GB per node.
	MaxMemoryPerNode float64 `json:"maxMemoryPerNode"`
	MaxCoresPerNode  int64   `json:"maxCoresPerNode"`
	// Native license node quota; not a Kubernetes Pod count.
	MaxNodes   int64  `json:"maxNodes"`
	ClientName string `json:"clientName"`
	// Native license type, 0 ordinary file, 1 fingerprint, 2 online, 3 license server.
	LicenseType int32 `json:"licenseType"`
	// Optional native source type; absent on older versions.
	LicenseSourceType *int32 `json:"licenseSourceType,omitempty"`
	Authorization     string `json:"authorization"`
	Version           string `json:"version"`
	ModuleNames       string `json:"moduleNames"`
	Modules           int64  `json:"modules"`
	ProductKey        string `json:"productKey"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDolphinDBLicenseInfo instantiates a new DolphinDBLicenseInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDolphinDBLicenseInfo(expiration string, maxMemoryPerNode float64, maxCoresPerNode int64, maxNodes int64, clientName string, licenseType int32, authorization string, version string, moduleNames string, modules int64, productKey string) *DolphinDBLicenseInfo {
	this := DolphinDBLicenseInfo{}
	this.Expiration = expiration
	this.MaxMemoryPerNode = maxMemoryPerNode
	this.MaxCoresPerNode = maxCoresPerNode
	this.MaxNodes = maxNodes
	this.ClientName = clientName
	this.LicenseType = licenseType
	this.Authorization = authorization
	this.Version = version
	this.ModuleNames = moduleNames
	this.Modules = modules
	this.ProductKey = productKey
	return &this
}

// NewDolphinDBLicenseInfoWithDefaults instantiates a new DolphinDBLicenseInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDolphinDBLicenseInfoWithDefaults() *DolphinDBLicenseInfo {
	this := DolphinDBLicenseInfo{}
	return &this
}

// GetExpiration returns the Expiration field value.
func (o *DolphinDBLicenseInfo) GetExpiration() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetExpirationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value.
func (o *DolphinDBLicenseInfo) SetExpiration(v string) {
	o.Expiration = v
}

// GetMaxMemoryPerNode returns the MaxMemoryPerNode field value.
func (o *DolphinDBLicenseInfo) GetMaxMemoryPerNode() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.MaxMemoryPerNode
}

// GetMaxMemoryPerNodeOk returns a tuple with the MaxMemoryPerNode field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetMaxMemoryPerNodeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxMemoryPerNode, true
}

// SetMaxMemoryPerNode sets field value.
func (o *DolphinDBLicenseInfo) SetMaxMemoryPerNode(v float64) {
	o.MaxMemoryPerNode = v
}

// GetMaxCoresPerNode returns the MaxCoresPerNode field value.
func (o *DolphinDBLicenseInfo) GetMaxCoresPerNode() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxCoresPerNode
}

// GetMaxCoresPerNodeOk returns a tuple with the MaxCoresPerNode field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetMaxCoresPerNodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxCoresPerNode, true
}

// SetMaxCoresPerNode sets field value.
func (o *DolphinDBLicenseInfo) SetMaxCoresPerNode(v int64) {
	o.MaxCoresPerNode = v
}

// GetMaxNodes returns the MaxNodes field value.
func (o *DolphinDBLicenseInfo) GetMaxNodes() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.MaxNodes
}

// GetMaxNodesOk returns a tuple with the MaxNodes field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetMaxNodesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxNodes, true
}

// SetMaxNodes sets field value.
func (o *DolphinDBLicenseInfo) SetMaxNodes(v int64) {
	o.MaxNodes = v
}

// GetClientName returns the ClientName field value.
func (o *DolphinDBLicenseInfo) GetClientName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetClientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientName, true
}

// SetClientName sets field value.
func (o *DolphinDBLicenseInfo) SetClientName(v string) {
	o.ClientName = v
}

// GetLicenseType returns the LicenseType field value.
func (o *DolphinDBLicenseInfo) GetLicenseType() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetLicenseTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseType, true
}

// SetLicenseType sets field value.
func (o *DolphinDBLicenseInfo) SetLicenseType(v int32) {
	o.LicenseType = v
}

// GetLicenseSourceType returns the LicenseSourceType field value if set, zero value otherwise.
func (o *DolphinDBLicenseInfo) GetLicenseSourceType() int32 {
	if o == nil || o.LicenseSourceType == nil {
		var ret int32
		return ret
	}
	return *o.LicenseSourceType
}

// GetLicenseSourceTypeOk returns a tuple with the LicenseSourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetLicenseSourceTypeOk() (*int32, bool) {
	if o == nil || o.LicenseSourceType == nil {
		return nil, false
	}
	return o.LicenseSourceType, true
}

// HasLicenseSourceType returns a boolean if a field has been set.
func (o *DolphinDBLicenseInfo) HasLicenseSourceType() bool {
	return o != nil && o.LicenseSourceType != nil
}

// SetLicenseSourceType gets a reference to the given int32 and assigns it to the LicenseSourceType field.
func (o *DolphinDBLicenseInfo) SetLicenseSourceType(v int32) {
	o.LicenseSourceType = &v
}

// GetAuthorization returns the Authorization field value.
func (o *DolphinDBLicenseInfo) GetAuthorization() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Authorization
}

// GetAuthorizationOk returns a tuple with the Authorization field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetAuthorizationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authorization, true
}

// SetAuthorization sets field value.
func (o *DolphinDBLicenseInfo) SetAuthorization(v string) {
	o.Authorization = v
}

// GetVersion returns the Version field value.
func (o *DolphinDBLicenseInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *DolphinDBLicenseInfo) SetVersion(v string) {
	o.Version = v
}

// GetModuleNames returns the ModuleNames field value.
func (o *DolphinDBLicenseInfo) GetModuleNames() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModuleNames
}

// GetModuleNamesOk returns a tuple with the ModuleNames field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetModuleNamesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModuleNames, true
}

// SetModuleNames sets field value.
func (o *DolphinDBLicenseInfo) SetModuleNames(v string) {
	o.ModuleNames = v
}

// GetModules returns the Modules field value.
func (o *DolphinDBLicenseInfo) GetModules() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetModulesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modules, true
}

// SetModules sets field value.
func (o *DolphinDBLicenseInfo) SetModules(v int64) {
	o.Modules = v
}

// GetProductKey returns the ProductKey field value.
func (o *DolphinDBLicenseInfo) GetProductKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProductKey
}

// GetProductKeyOk returns a tuple with the ProductKey field value
// and a boolean to check if the value has been set.
func (o *DolphinDBLicenseInfo) GetProductKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductKey, true
}

// SetProductKey sets field value.
func (o *DolphinDBLicenseInfo) SetProductKey(v string) {
	o.ProductKey = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DolphinDBLicenseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	toSerialize["expiration"] = o.Expiration
	toSerialize["maxMemoryPerNode"] = o.MaxMemoryPerNode
	toSerialize["maxCoresPerNode"] = o.MaxCoresPerNode
	toSerialize["maxNodes"] = o.MaxNodes
	toSerialize["clientName"] = o.ClientName
	toSerialize["licenseType"] = o.LicenseType
	if o.LicenseSourceType != nil {
		toSerialize["licenseSourceType"] = o.LicenseSourceType
	}
	toSerialize["authorization"] = o.Authorization
	toSerialize["version"] = o.Version
	toSerialize["moduleNames"] = o.ModuleNames
	toSerialize["modules"] = o.Modules
	toSerialize["productKey"] = o.ProductKey

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DolphinDBLicenseInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Expiration        *string  `json:"expiration"`
		MaxMemoryPerNode  *float64 `json:"maxMemoryPerNode"`
		MaxCoresPerNode   *int64   `json:"maxCoresPerNode"`
		MaxNodes          *int64   `json:"maxNodes"`
		ClientName        *string  `json:"clientName"`
		LicenseType       *int32   `json:"licenseType"`
		LicenseSourceType *int32   `json:"licenseSourceType,omitempty"`
		Authorization     *string  `json:"authorization"`
		Version           *string  `json:"version"`
		ModuleNames       *string  `json:"moduleNames"`
		Modules           *int64   `json:"modules"`
		ProductKey        *string  `json:"productKey"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	if all.Expiration == nil {
		return fmt.Errorf("required field expiration missing")
	}
	if all.MaxMemoryPerNode == nil {
		return fmt.Errorf("required field maxMemoryPerNode missing")
	}
	if all.MaxCoresPerNode == nil {
		return fmt.Errorf("required field maxCoresPerNode missing")
	}
	if all.MaxNodes == nil {
		return fmt.Errorf("required field maxNodes missing")
	}
	if all.ClientName == nil {
		return fmt.Errorf("required field clientName missing")
	}
	if all.LicenseType == nil {
		return fmt.Errorf("required field licenseType missing")
	}
	if all.Authorization == nil {
		return fmt.Errorf("required field authorization missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	if all.ModuleNames == nil {
		return fmt.Errorf("required field moduleNames missing")
	}
	if all.Modules == nil {
		return fmt.Errorf("required field modules missing")
	}
	if all.ProductKey == nil {
		return fmt.Errorf("required field productKey missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"expiration", "maxMemoryPerNode", "maxCoresPerNode", "maxNodes", "clientName", "licenseType", "licenseSourceType", "authorization", "version", "moduleNames", "modules", "productKey"})
	} else {
		return err
	}
	o.Expiration = *all.Expiration
	o.MaxMemoryPerNode = *all.MaxMemoryPerNode
	o.MaxCoresPerNode = *all.MaxCoresPerNode
	o.MaxNodes = *all.MaxNodes
	o.ClientName = *all.ClientName
	o.LicenseType = *all.LicenseType
	o.LicenseSourceType = all.LicenseSourceType
	o.Authorization = *all.Authorization
	o.Version = *all.Version
	o.ModuleNames = *all.ModuleNames
	o.Modules = *all.Modules
	o.ProductKey = *all.ProductKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
