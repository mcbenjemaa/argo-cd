package v1alpha1

import (
	"encoding/json"
)

// +kubebuilder:validation:Type=object
type Object struct {
	Value json.RawMessage `json:"-" protobuf:"bytes,1,opt,name=value"`
}

func (i *Object) UnmarshalJSON(value []byte) error {
	return i.Value.UnmarshalJSON(value)
}

func (i Object) MarshalJSON() ([]byte, error) {
	return i.Value.MarshalJSON()
}

func (i Object) OpenAPISchemaType() []string {
	return []string{"object"}
}

func (i Object) OpenAPISchemaFormat() string { return "" }

// +protobuf=true
// +protobuf.embed=string
// +protobuf.options.marshal=false
// +protobuf.options.(gogoproto.goproto_stringer)=false
// +k8s:deepcopy-gen=true
// +k8s:openapi-gen=true
// type Quantity struct {
// 	// i is the quantity in int64 scaled form, if d.Dec == nil
// 	i int64Amount
// 	// d is the quantity in inf.Dec form if d.Dec != nil
// 	d infDecAmount
// 	// s is the generated value of this quantity to avoid recalculation
// 	s string

// 	// Change Format at will. See the comment for Canonicalize for
// 	// more details.
// 	Format
// }

// // CanonicalValue allows a quantity amount to be converted to a string.
// type CanonicalValue interface {
// 	// AsCanonicalBytes returns a byte array representing the string representation
// 	// of the value mantissa and an int32 representing its exponent in base-10. Callers may
// 	// pass a byte slice to the method to avoid allocations.
// 	AsCanonicalBytes(out []byte) ([]byte, int32)
// 	// AsCanonicalBase1024Bytes returns a byte array representing the string representation
// 	// of the value mantissa and an int32 representing its exponent in base-1024. Callers
// 	// may pass a byte slice to the method to avoid allocations.
// 	AsCanonicalBase1024Bytes(out []byte) ([]byte, int32)
// }

// // Format lists the three possible formattings of a quantity.
// type Format string

// const (
// 	DecimalExponent = Format("DecimalExponent") // e.g., 12e6
// 	BinarySI        = Format("BinarySI")        // e.g., 12Mi (12 * 2^20)
// 	DecimalSI       = Format("DecimalSI")       // e.g., 12M  (12 * 10^6)
// )
