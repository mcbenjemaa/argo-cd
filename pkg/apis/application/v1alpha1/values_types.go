package v1alpha1

import "encoding/json"

// ValuesObject is an Object
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

func (p *ValuesObject) isEmpty() bool {
	return len(p.Value) == 0
}
