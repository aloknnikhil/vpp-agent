// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	"github.com/ligato/vpp-agent/examples/kvscheduler/mock_plugins/l2plugin/model"
	"github.com/ligato/vpp-agent/pkg/idxvpp"
)

////////// type-safe key-value pair with metadata //////////

type BridgeDomainKVWithMetadata struct {
	Key      string
	Value    *mock_l2.BridgeDomain
	Metadata *idxvpp.OnlyIndex
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type BridgeDomainDescriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *mock_l2.BridgeDomain) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *mock_l2.BridgeDomain) error
	Create               func(key string, value *mock_l2.BridgeDomain) (metadata *idxvpp.OnlyIndex, err error)
	Delete               func(key string, value *mock_l2.BridgeDomain, metadata *idxvpp.OnlyIndex) error
	Update               func(key string, oldValue, newValue *mock_l2.BridgeDomain, oldMetadata *idxvpp.OnlyIndex) (newMetadata *idxvpp.OnlyIndex, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *mock_l2.BridgeDomain, metadata *idxvpp.OnlyIndex) bool
	Retrieve             func(correlate []BridgeDomainKVWithMetadata) ([]BridgeDomainKVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *mock_l2.BridgeDomain) []KeyValuePair
	Dependencies         func(key string, value *mock_l2.BridgeDomain) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type BridgeDomainDescriptorAdapter struct {
	descriptor *BridgeDomainDescriptor
}

func NewBridgeDomainDescriptor(typedDescriptor *BridgeDomainDescriptor) *KVDescriptor {
	adapter := &BridgeDomainDescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *BridgeDomainDescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castBridgeDomainValue(key, oldValue)
	typedNewValue, err2 := castBridgeDomainValue(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *BridgeDomainDescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castBridgeDomainValue(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *BridgeDomainDescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castBridgeDomainValue(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *BridgeDomainDescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castBridgeDomainValue(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castBridgeDomainValue(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castBridgeDomainMetadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *BridgeDomainDescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castBridgeDomainValue(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castBridgeDomainMetadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *BridgeDomainDescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castBridgeDomainValue(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castBridgeDomainValue(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castBridgeDomainMetadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *BridgeDomainDescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []BridgeDomainKVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castBridgeDomainValue(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castBridgeDomainMetadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			BridgeDomainKVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *BridgeDomainDescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castBridgeDomainValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *BridgeDomainDescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castBridgeDomainValue(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castBridgeDomainValue(key string, value proto.Message) (*mock_l2.BridgeDomain, error) {
	typedValue, ok := value.(*mock_l2.BridgeDomain)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castBridgeDomainMetadata(key string, metadata Metadata) (*idxvpp.OnlyIndex, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(*idxvpp.OnlyIndex)
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}