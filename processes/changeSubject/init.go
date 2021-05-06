package changeSubject

import "tcj3-kadai-tuika-kun/types"

var (
	Config *types.ConfigYaml
)

type flexJson struct {
	Contents []flexJsonContent `json:"contents"`
}

type flexJsonContent struct {
	Type     string               `json:"type"`
	Layout   string               `json:"layout"`
	Contents []flexJsonContent2nd `json:"contents"`
}

type flexJsonContent2nd struct {
	Type    string                    `json:"type"`
	Text    string                    `json:"text,omitempty"`
	Margin  string                    `json:"margin,omitempty"`
	Gravity string                    `json:"gravity,omitempty"`
	Align   string                    `json:"align,omitempty"`
	Size    string                    `json:"size,omitempty"`
	Action  *flexJsonContent2ndAction `json:"action,omitempty"`
	Height  string                    `json:"height,omitempty"`
	Flex    int                       `json:"flex,omitempty"`
}

type flexJsonContent2ndAction struct {
	Type  string `json:"type,omitempty"`
	Label string `json:"label,omitempty"`
	Text  string `json:"text,omitempty"`
}
