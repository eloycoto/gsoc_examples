package api

type Foo struct {
	Version string `json:"version,omitempty"`
	Bar     *BAR   `json:"bar,omitempty"`
}

type BAR struct {
	Version string `json:"version,omitempty"`
}
