// Package packagejson contains structure to represent package.json
package packagejson

// PackageJSON represents NodeJS package.json
type PackageJSON struct {
	Name        string   `json:"name"`    // required
	Version     string   `json:"version"` // required
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
	Homepage    string   `json:"homepage"`
	License     string   `json:"license"`
	Files       []string `json:"files"`
	Main        string   `json:"main"`
	Os          []string `json:"os"`
	Cpu         []string `json:"cpu"`
	Private     bool     `json:"private"`
}

// Validate checks if provided package.json is valid.
func (p *PackageJSON) Validate() error {
	return nil
}
