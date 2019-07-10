package main

type Organization struct {
	PartOf   *Reference   `json:"partOf,omitempty"`
	Endpoint []*Reference `json:"endpoint,omitempty"`
}
