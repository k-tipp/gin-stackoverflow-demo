package main

type Reference struct {
	Ref        string      `json:"reference,omitempty"`
	Type       string      `json:"type,omitempty"`
	Identifier *Identifier `json:"identifier,omitempty"`
	Display    string      `json:"display,omitempty"`
}


