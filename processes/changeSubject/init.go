package changeSubject

import "tcj3-kadai-tuika-kun/types"

var (
	Config *types.ConfigYaml
)

type flexJson struct {
	Contents []struct {
		Type     string            `json:"type"`
		Layout   string            `json:"layout"`
		Contents []flexJsonContent `json:"contents"`
	} `json:"contents"`
}

type flexJsonContent struct {
	Type    string                `json:"type"`
	Text    string                `json:"text,omitempty"`
	Margin  string                `json:"margin,omitempty"`
	Gravity string                `json:"gravity,omitempty"`
	Align   string                `json:"align,omitempty"`
	Size    string                `json:"size,omitempty"`
	Action  flexJsonContentAction `json:"action,omitempty"`
	Height  string                `json:"height,omitempty"`
	Flex    int                   `json:"flex,omitempty"`
}

type flexJsonContentAction struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	Text  string `json:"text"`
}
