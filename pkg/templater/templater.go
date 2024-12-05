package templater

import (
	"strconv"

	"github.com/randomtoy/go-8ball/pkg/translator"
)

type Templater struct {
	Code   string
	Value  string
	String string
}

func NewTemplater(code string) *Templater {
	return &Templater{
		Code: code,
	}

}

func (t *Templater) Generate(value int) (*Templater, error) {
	valueStr := strconv.Itoa(value)
	str, err := translator.GetStringByCode(t.Code, valueStr)
	if err != nil {
		return t, err
	}
	t.Value = valueStr
	t.String = str
	return t, err
}
