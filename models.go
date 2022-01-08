package main

type Request struct {
	URL               string     `json:"url"`
	Method            string     `json:"method"`
	Headers           [][]string `json:"headers"`
	Body              string     `json:"body"`
	AllowRedirect     bool       `json:"allowRedirect"`
	Proxy             string     `json:"proxy"`
	Timeout           int64      `json:"timeout"`
	PseudoHeaderOrder []string   `json:"pseudoHeaderOrder"`
}

type Response struct {
	Headers [][]string `json:"headers"`
	Body    string     `json:"body"`
}
