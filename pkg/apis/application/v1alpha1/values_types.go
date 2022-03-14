package v1alpha1

import (
	"encoding/json"

	"sigs.k8s.io/yaml"
)

// ValuesObject is an Object
// +patchStrategy=replace
// +protobuf.options.(gogoproto.goproto_stringer)=false
type ValuesObject struct {
	Object `json:",inline" protobuf:"bytes,1,opt,name=object"`
}

// UnmarshalJSON unmarshalls the Plugin from JSON, and also validates that it is a map exactly one key
func (p *ValuesObject) UnmarshalJSON(value []byte) error {
	if err := p.Object.UnmarshalJSON(value); err != nil {
		return err
	}
	// by validating the structure in UnmarshallJSON, we prevent bad data entering the system at the point of
	// parsing, which means we do not need validate
	m := map[string]interface{}{}
	if err := json.Unmarshal(p.Object.Value, &m); err != nil {
		return err
	}

	return nil
}

// UnmarshalMap unmarshalls the Plugin from JSON, and also validates that it is a map exactly one key
func (p *ValuesObject) UnmarshalMap() (map[string]interface{}, error) {

	m := map[string]interface{}{}
	if err := json.Unmarshal(p.Object.Value, &m); err != nil {
		return map[string]interface{}{}, err
	}

	return m, nil
}

func (p ValuesObject) IsEmpty() bool {
	return len(p.Object.Value) == 0
}

// String formats the ValuesObject as a string, caching the result if not calculated.
// String is an expensive operation and caching this result significantly reduces the cost of
// normal parse / marshal operations on Quantity.
func (q *ValuesObject) String() string {
	b, err := yaml.JSONToYAML(q.Object.Value)
	if err != nil {
		return "<nil>"
	}
	return string(b)
}

// ToUnstructured implements the value.UnstructuredConverter interface.
func (q ValuesObject) ToUnstructured() interface{} {
	return q.String()
}
