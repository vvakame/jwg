// generated by jwg -type Sample,Inner -output misc/fixture/g/model_json.go misc/fixture/g; DO NOT EDIT

package g

import (
	"encoding/json"
)

// for Inner
type InnerJson struct {
	A string `json:"a,omitempty"`
}

type InnerJsonList []*InnerJson

type InnerPropertyEncoder func(src *Inner, dest *InnerJson) error

type InnerPropertyDecoder func(src *InnerJson, dest *Inner) error

type InnerPropertyInfo struct {
	name    string
	Encoder InnerPropertyEncoder
	Decoder InnerPropertyDecoder
}

type InnerJsonBuilder struct {
	_properties map[string]*InnerPropertyInfo
	A           *InnerPropertyInfo
}

func NewInnerJsonBuilder() *InnerJsonBuilder {
	return &InnerJsonBuilder{
		_properties: map[string]*InnerPropertyInfo{},
		A: &InnerPropertyInfo{
			name: "A",
			Encoder: func(src *Inner, dest *InnerJson) error {
				if src == nil {
					return nil
				}
				dest.A = src.A
				return nil
			},
			Decoder: func(src *InnerJson, dest *Inner) error {
				if src == nil {
					return nil
				}
				dest.A = src.A
				return nil
			},
		},
	}
}

func (b *InnerJsonBuilder) AddAll() *InnerJsonBuilder {
	b._properties["A"] = b.A
	return b
}

func (b *InnerJsonBuilder) Add(info *InnerPropertyInfo) *InnerJsonBuilder {
	b._properties[info.name] = info
	return b
}

func (b *InnerJsonBuilder) Remove(info *InnerPropertyInfo) *InnerJsonBuilder {
	delete(b._properties, info.name)
	return b
}

func (b *InnerJsonBuilder) Convert(orig *Inner) (*InnerJson, error) {
	if orig == nil {
		return nil, nil
	}
	ret := &InnerJson{}

	for _, info := range b._properties {
		if err := info.Encoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (b *InnerJsonBuilder) ConvertList(orig []*Inner) (InnerJsonList, error) {
	if orig == nil {
		return nil, nil
	}

	list := make(InnerJsonList, len(orig))
	for idx, or := range orig {
		json, err := b.Convert(or)
		if err != nil {
			return nil, err
		}
		list[idx] = json
	}

	return list, nil
}

func (orig *InnerJson) Convert() (*Inner, error) {
	ret := &Inner{}

	b := NewInnerJsonBuilder().AddAll()
	for _, info := range b._properties {
		if err := info.Decoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (jsonList InnerJsonList) Convert() ([]*Inner, error) {
	orig := ([]*InnerJson)(jsonList)

	list := make([]*Inner, len(orig))
	for idx, or := range orig {
		obj, err := or.Convert()
		if err != nil {
			return nil, err
		}
		list[idx] = obj
	}

	return list, nil
}

func (b *InnerJsonBuilder) Marshal(orig *Inner) ([]byte, error) {
	ret, err := b.Convert(orig)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ret)
}

// for Sample
type SampleJson struct {
	A1  InnerJson     `json:"a1,omitempty"`
	A2  *InnerJson    `json:"a2,omitempty"`
	As1 []InnerJson   `json:"as1,omitempty"`
	As2 []*InnerJson  `json:"as2,omitempty"`
	As3 *[]InnerJson  `json:"as3,omitempty"`
	As4 *[]*InnerJson `json:"as4,omitempty"`
}

type SampleJsonList []*SampleJson

type SamplePropertyEncoder func(src *Sample, dest *SampleJson) error

type SamplePropertyDecoder func(src *SampleJson, dest *Sample) error

type SamplePropertyInfo struct {
	name    string
	Encoder SamplePropertyEncoder
	Decoder SamplePropertyDecoder
}

type SampleJsonBuilder struct {
	_properties map[string]*SamplePropertyInfo
	A1          *SamplePropertyInfo
	A2          *SamplePropertyInfo
	As1         *SamplePropertyInfo
	As2         *SamplePropertyInfo
	As3         *SamplePropertyInfo
	As4         *SamplePropertyInfo
}

func NewSampleJsonBuilder() *SampleJsonBuilder {
	return &SampleJsonBuilder{
		_properties: map[string]*SamplePropertyInfo{},
		A1: &SamplePropertyInfo{
			name: "A1",
			Encoder: func(src *Sample, dest *SampleJson) error {
				if src == nil {
					return nil
				}
				d, err := NewInnerJsonBuilder().AddAll().Convert(&src.A1)
				if err != nil {
					return err
				}
				dest.A1 = *d
				return nil
			},
			Decoder: func(src *SampleJson, dest *Sample) error {
				if src == nil {
					return nil
				}
				d, err := src.A1.Convert()
				if err != nil {
					return err
				}
				dest.A1 = *d
				return nil
			},
		},
		A2: &SamplePropertyInfo{
			name: "A2",
			Encoder: func(src *Sample, dest *SampleJson) error {
				if src == nil {
					return nil
				} else if src.A2 == nil {
					return nil
				}
				d, err := NewInnerJsonBuilder().AddAll().Convert(src.A2)
				if err != nil {
					return err
				}
				dest.A2 = d
				return nil
			},
			Decoder: func(src *SampleJson, dest *Sample) error {
				if src == nil {
					return nil
				} else if src.A2 == nil {
					return nil
				}
				d, err := src.A2.Convert()
				if err != nil {
					return err
				}
				dest.A2 = d
				return nil
			},
		},
		As1: &SamplePropertyInfo{
			name: "As1",
			Encoder: func(src *Sample, dest *SampleJson) error {
				if src == nil {
					return nil
				}
				b := NewInnerJsonBuilder().AddAll()
				list := make([]InnerJson, len(src.As1))
				for idx, obj := range src.As1 {
					d, err := b.Convert(&obj)
					if err != nil {
						return err
					}
					list[idx] = *d
				}
				dest.As1 = list
				return nil
			},
			Decoder: func(src *SampleJson, dest *Sample) error {
				if src == nil {
					return nil
				}
				list := make([]Inner, len(src.As1))
				for idx, obj := range src.As1 {
					d, err := obj.Convert()
					if err != nil {
						return err
					}
					list[idx] = *d
				}
				dest.As1 = list
				return nil
			},
		},
		As2: &SamplePropertyInfo{
			name: "As2",
			Encoder: func(src *Sample, dest *SampleJson) error {
				if src == nil {
					return nil
				}
				list, err := NewInnerJsonBuilder().AddAll().ConvertList(src.As2)
				if err != nil {
					return err
				}
				dest.As2 = ([]*InnerJson)(list)
				return nil
			},
			Decoder: func(src *SampleJson, dest *Sample) error {
				if src == nil {
					return nil
				}
				list := make([]*Inner, len(src.As2))
				for idx, obj := range src.As2 {
					if obj == nil {
						continue
					}
					d, err := obj.Convert()
					if err != nil {
						return err
					}
					list[idx] = d
				}
				dest.As2 = list
				return nil
			},
		},
		As3: &SamplePropertyInfo{
			name: "As3",
			Encoder: func(src *Sample, dest *SampleJson) error {
				if src == nil {
					return nil
				} else if src.As3 == nil {
					return nil
				}
				b := NewInnerJsonBuilder().AddAll()
				list := make([]InnerJson, len(*src.As3))
				for idx, obj := range *src.As3 {
					d, err := b.Convert(&obj)
					if err != nil {
						return err
					}
					list[idx] = *d
				}
				dest.As3 = &list
				return nil
			},
			Decoder: func(src *SampleJson, dest *Sample) error {
				if src == nil {
					return nil
				} else if src.As3 == nil {
					return nil
				}
				list := make([]Inner, len(*src.As3))
				for idx, obj := range *src.As3 {
					d, err := obj.Convert()
					if err != nil {
						return err
					}
					list[idx] = *d
				}
				dest.As3 = &list
				return nil
			},
		},
		As4: &SamplePropertyInfo{
			name: "As4",
			Encoder: func(src *Sample, dest *SampleJson) error {
				if src == nil {
					return nil
				} else if src.As4 == nil {
					return nil
				}
				list, err := NewInnerJsonBuilder().AddAll().ConvertList(*src.As4)
				if err != nil {
					return err
				}
				dest.As4 = (*[]*InnerJson)(&list)
				return nil
			},
			Decoder: func(src *SampleJson, dest *Sample) error {
				if src == nil {
					return nil
				} else if src.As4 == nil {
					return nil
				}
				list := make([]*Inner, len(*src.As4))
				for idx, obj := range *src.As4 {
					if obj == nil {
						continue
					}
					d, err := obj.Convert()
					if err != nil {
						return err
					}
					list[idx] = d
				}
				dest.As4 = &list
				return nil
			},
		},
	}
}

func (b *SampleJsonBuilder) AddAll() *SampleJsonBuilder {
	b._properties["A1"] = b.A1
	b._properties["A2"] = b.A2
	b._properties["As1"] = b.As1
	b._properties["As2"] = b.As2
	b._properties["As3"] = b.As3
	b._properties["As4"] = b.As4
	return b
}

func (b *SampleJsonBuilder) Add(info *SamplePropertyInfo) *SampleJsonBuilder {
	b._properties[info.name] = info
	return b
}

func (b *SampleJsonBuilder) Remove(info *SamplePropertyInfo) *SampleJsonBuilder {
	delete(b._properties, info.name)
	return b
}

func (b *SampleJsonBuilder) Convert(orig *Sample) (*SampleJson, error) {
	if orig == nil {
		return nil, nil
	}
	ret := &SampleJson{}

	for _, info := range b._properties {
		if err := info.Encoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (b *SampleJsonBuilder) ConvertList(orig []*Sample) (SampleJsonList, error) {
	if orig == nil {
		return nil, nil
	}

	list := make(SampleJsonList, len(orig))
	for idx, or := range orig {
		json, err := b.Convert(or)
		if err != nil {
			return nil, err
		}
		list[idx] = json
	}

	return list, nil
}

func (orig *SampleJson) Convert() (*Sample, error) {
	ret := &Sample{}

	b := NewSampleJsonBuilder().AddAll()
	for _, info := range b._properties {
		if err := info.Decoder(orig, ret); err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (jsonList SampleJsonList) Convert() ([]*Sample, error) {
	orig := ([]*SampleJson)(jsonList)

	list := make([]*Sample, len(orig))
	for idx, or := range orig {
		obj, err := or.Convert()
		if err != nil {
			return nil, err
		}
		list[idx] = obj
	}

	return list, nil
}

func (b *SampleJsonBuilder) Marshal(orig *Sample) ([]byte, error) {
	ret, err := b.Convert(orig)
	if err != nil {
		return nil, err
	}
	return json.Marshal(ret)
}
