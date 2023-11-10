package screen

import (
	"encoding/json"
	"fmt"
)

type Attributes struct {
	Italic    bool // 3m
	Underline bool // 4m
	Strikeout bool // 9m
}

type Style struct {
	Emphasis Attributes // Text attributes
	FColor   Color      // foreground color
	BColor   Color      // background color
}

func (o Attributes) String() string {
	return italic(o.Italic) + underline(o.Underline) + strikeout(o.Strikeout)
}

func italic(b bool) string {
	if b {
		return "i"
	}
	return ""
}

func underline(b bool) string {
	if b {
		return "u"
	}
	return ""
}

func strikeout(b bool) string {
	if b {
		return "s"
	}
	return ""
}

func (o Style) String() string {
	return fmt.Sprintf("%v,%v,%v", o.Emphasis, o.FColor, o.BColor)
}

func (o Style) MarshalJSON() ([]byte, error) {
	m := make(map[string]string)
	m["style"] = o.Emphasis.String() + "," + o.FColor.StringSmart() + "," + o.BColor.StringSmart()
	return json.Marshal(m)
}
