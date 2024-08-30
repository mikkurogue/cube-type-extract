package gen

type Generator struct {
	CubeCount int
	Metadata  CubeMetadata
}

type CubeMetadata struct {
	Cubes []Cube `json:"cubes"`
}

type Cube struct {
	Name       string     `json:"name"`
	Dimensions []FieldSet `json:"dimensions"`
	Measures   []FieldSet `json:"measures"`
}

// NOTE: Just doing this to test if we can filter dimensions and measures based on a meta prop
type Meta struct {
	Extractable bool `json:"Extractable"`
}

type FieldSet struct {
	Name string `json:"name"`
	Meta Meta   `json:"meta"`
}
