package dom //import "github.com/kpmy/ypk/dom"

import (
	"bytes"
	"encoding/xml"

	"reflect"

	"github.com/kpmy/ypk/halt"
)

func Decode(b *bytes.Buffer) (ret Entity, err error) {
	dd := &domm{buf: b}
	if err = dd.Unmarshal(); err == nil {
		ret = dd
	}
	return
}

func EncodeWithHeader(el Element) *bytes.Buffer {
	dd := &domm{model: el, header: true}
	return dd.Produce()
}

func Encode(el Element) *bytes.Buffer {
	dd := &domm{model: el}
	return dd.Produce()
}

type Entity interface {
	Model() Element
	Produce() *bytes.Buffer
}

type domm struct {
	header bool
	buf    *bytes.Buffer
	model  Element
}

func (x *domm) Type() string {
	return x.model.Name()
}

func (x *domm) Model() Element {
	return x.model
}

func (x *domm) Produce() (ret *bytes.Buffer) {
	if data, err := xml.Marshal(x); err == nil {
		if x.header {
			ret = bytes.NewBufferString(xml.Header)
			ret.Write(data)
		} else {
			ret = bytes.NewBuffer(data)
		}
	} else {
		halt.As(100, ret)
	}
	return
}

func (x *domm) Unmarshal() (err error) {
	d := xml.NewDecoder(x.buf)
	var _t xml.Token
	var this Element
	for stop := false; !stop && err == nil; {
		if _t, err = d.RawToken(); err == nil {
			switch t := _t.(type) {
			case xml.StartElement:
				el := Elem(ThisName(t.Name))
				if x.model == nil {
					x.model = el
					this = el
				} else {
					this.AppendChild(el)
					this = el
				}
				for _, a := range t.Attr {
					this.Attr(ThisName(a.Name), a.Value)
				}
			case xml.CharData:
				if this != nil {
					this.AppendChild(Txt(string(t)))
				} else {
					stop = true
				}
			case xml.EndElement:
				if this != nil {
					if p := this.Parent(); p != nil {
						this = p.(Element)
					} else {
						stop = true
					}
				} else {
					stop = true
				}
			case nil:
			default:
				halt.As(100, reflect.TypeOf(t))
			}
		}
	}
	return
}

func (x *domm) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
	start.Name.Local = x.model.Name()
	for k, v := range x.model.AttrAsMap() {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: k}, Value: v})
	}
	e.EncodeToken(start)
	for _, _l := range x.model.Children() {
		switch l := _l.(type) {
		case Element:
			child := &domm{}
			child.model = l
			e.Encode(child)
		case Text:
			e.EncodeToken(xml.CharData(l.Data()))
		default:
			halt.As(100, reflect.TypeOf(l))
		}
	}
	e.EncodeToken(start.End())
	return
}
