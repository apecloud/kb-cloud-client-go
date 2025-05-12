// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import (
	"time"

	"github.com/apecloud/kb-cloud-client-go/api/common"
)

// Key mutiple kv pair
type Key struct {
	// The unique identifier of the key.
	Id *string `json:"id,omitempty"`
	// The name of the key.
	Name *string `json:"name,omitempty"`
	// the encryption algorithm
	Algorithm *EncryptionAlgorithm `json:"algorithm,omitempty"`
	// The intended usage of the key.
	KeyUsage *EncryptionKeyUsage `json:"keyUsage,omitempty"`
	// the clusters which use the key
	Clusters []string `json:"clusters,omitempty"`
	// the backups which use the key
	Backups []string `json:"backups,omitempty"`
	// The creation timestamp of the key.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The last update timestamp of the key.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// the name of data key
	DataKeyName *string `json:"dataKeyName,omitempty"`
	// the sysmetric key
	SysmetricKey *string `json:"sysmetricKey,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewKey instantiates a new Key object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewKey() *Key {
	this := Key{}
	return &this
}

// NewKeyWithDefaults instantiates a new Key object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewKeyWithDefaults() *Key {
	this := Key{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Key) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Key) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Key) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Key) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Key) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Key) SetName(v string) {
	o.Name = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *Key) GetAlgorithm() EncryptionAlgorithm {
	if o == nil || o.Algorithm == nil {
		var ret EncryptionAlgorithm
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetAlgorithmOk() (*EncryptionAlgorithm, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *Key) HasAlgorithm() bool {
	return o != nil && o.Algorithm != nil
}

// SetAlgorithm gets a reference to the given EncryptionAlgorithm and assigns it to the Algorithm field.
func (o *Key) SetAlgorithm(v EncryptionAlgorithm) {
	o.Algorithm = &v
}

// GetKeyUsage returns the KeyUsage field value if set, zero value otherwise.
func (o *Key) GetKeyUsage() EncryptionKeyUsage {
	if o == nil || o.KeyUsage == nil {
		var ret EncryptionKeyUsage
		return ret
	}
	return *o.KeyUsage
}

// GetKeyUsageOk returns a tuple with the KeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetKeyUsageOk() (*EncryptionKeyUsage, bool) {
	if o == nil || o.KeyUsage == nil {
		return nil, false
	}
	return o.KeyUsage, true
}

// HasKeyUsage returns a boolean if a field has been set.
func (o *Key) HasKeyUsage() bool {
	return o != nil && o.KeyUsage != nil
}

// SetKeyUsage gets a reference to the given EncryptionKeyUsage and assigns it to the KeyUsage field.
func (o *Key) SetKeyUsage(v EncryptionKeyUsage) {
	o.KeyUsage = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *Key) GetClusters() []string {
	if o == nil || o.Clusters == nil {
		var ret []string
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetClustersOk() (*[]string, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return &o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *Key) HasClusters() bool {
	return o != nil && o.Clusters != nil
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *Key) SetClusters(v []string) {
	o.Clusters = v
}

// GetBackups returns the Backups field value if set, zero value otherwise.
func (o *Key) GetBackups() []string {
	if o == nil || o.Backups == nil {
		var ret []string
		return ret
	}
	return o.Backups
}

// GetBackupsOk returns a tuple with the Backups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetBackupsOk() (*[]string, bool) {
	if o == nil || o.Backups == nil {
		return nil, false
	}
	return &o.Backups, true
}

// HasBackups returns a boolean if a field has been set.
func (o *Key) HasBackups() bool {
	return o != nil && o.Backups != nil
}

// SetBackups gets a reference to the given []string and assigns it to the Backups field.
func (o *Key) SetBackups(v []string) {
	o.Backups = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Key) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Key) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Key) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Key) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Key) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Key) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetDataKeyName returns the DataKeyName field value if set, zero value otherwise.
func (o *Key) GetDataKeyName() string {
	if o == nil || o.DataKeyName == nil {
		var ret string
		return ret
	}
	return *o.DataKeyName
}

// GetDataKeyNameOk returns a tuple with the DataKeyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetDataKeyNameOk() (*string, bool) {
	if o == nil || o.DataKeyName == nil {
		return nil, false
	}
	return o.DataKeyName, true
}

// HasDataKeyName returns a boolean if a field has been set.
func (o *Key) HasDataKeyName() bool {
	return o != nil && o.DataKeyName != nil
}

// SetDataKeyName gets a reference to the given string and assigns it to the DataKeyName field.
func (o *Key) SetDataKeyName(v string) {
	o.DataKeyName = &v
}

// GetSysmetricKey returns the SysmetricKey field value if set, zero value otherwise.
func (o *Key) GetSysmetricKey() string {
	if o == nil || o.SysmetricKey == nil {
		var ret string
		return ret
	}
	return *o.SysmetricKey
}

// GetSysmetricKeyOk returns a tuple with the SysmetricKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetSysmetricKeyOk() (*string, bool) {
	if o == nil || o.SysmetricKey == nil {
		return nil, false
	}
	return o.SysmetricKey, true
}

// HasSysmetricKey returns a boolean if a field has been set.
func (o *Key) HasSysmetricKey() bool {
	return o != nil && o.SysmetricKey != nil
}

// SetSysmetricKey gets a reference to the given string and assigns it to the SysmetricKey field.
func (o *Key) SetSysmetricKey(v string) {
	o.SysmetricKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Key) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.KeyUsage != nil {
		toSerialize["keyUsage"] = o.KeyUsage
	}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	if o.Backups != nil {
		toSerialize["backups"] = o.Backups
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DataKeyName != nil {
		toSerialize["dataKeyName"] = o.DataKeyName
	}
	if o.SysmetricKey != nil {
		toSerialize["sysmetricKey"] = o.SysmetricKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Key) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id           *string              `json:"id,omitempty"`
		Name         *string              `json:"name,omitempty"`
		Algorithm    *EncryptionAlgorithm `json:"algorithm,omitempty"`
		KeyUsage     *EncryptionKeyUsage  `json:"keyUsage,omitempty"`
		Clusters     []string             `json:"clusters,omitempty"`
		Backups      []string             `json:"backups,omitempty"`
		CreatedAt    *time.Time           `json:"createdAt,omitempty"`
		UpdatedAt    *time.Time           `json:"updatedAt,omitempty"`
		DataKeyName  *string              `json:"dataKeyName,omitempty"`
		SysmetricKey *string              `json:"sysmetricKey,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"id", "name", "algorithm", "keyUsage", "clusters", "backups", "createdAt", "updatedAt", "dataKeyName", "sysmetricKey"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = all.Name
	if all.Algorithm != nil && !all.Algorithm.IsValid() {
		hasInvalidField = true
	} else {
		o.Algorithm = all.Algorithm
	}
	if all.KeyUsage != nil && !all.KeyUsage.IsValid() {
		hasInvalidField = true
	} else {
		o.KeyUsage = all.KeyUsage
	}
	o.Clusters = all.Clusters
	o.Backups = all.Backups
	o.CreatedAt = all.CreatedAt
	o.UpdatedAt = all.UpdatedAt
	o.DataKeyName = all.DataKeyName
	o.SysmetricKey = all.SysmetricKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
