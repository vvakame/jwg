package b

// test for struct with json tags definition

type Sample struct {
	A string `json:"foo!"`
	B string `json:",omitempty"`
	C int    `json:",string"`
	D int    `json:",omitempty,string"`
	E string `json:"-"`
}
