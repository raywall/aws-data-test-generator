package main

type Request struct {
	FileType   string   `json:"type"`
	Size       int      `json:"size"`
	Headers    bool     `json:"headers"`
	DataStruct []Column `json:"data-structure"`
}

type Column struct {
	ID     int    `json:"column-id"`
	Header string `json:"header"`
	Type   string `json:"type"`

	Size     int `json:"size"`
	MaxValue int `json:"max-value"`
	MinValue int `json:"min-value"`
}
