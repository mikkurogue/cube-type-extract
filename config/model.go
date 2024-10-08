package config

type Configuration struct {
	CubeUrl  string   `json:"cube_url"`
	Output   string   `json:"output"`
	FileName string   `json:"file_name"`
	Prefixes []Prefix `json:"prefixes"` // experimental try and see if we can pre-define our prefixes for the cube(s)
	Ignore   []string `json:"ignore"`
}

// the name prop is the current name, the prefix is the new name
type Prefix struct {
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
}
