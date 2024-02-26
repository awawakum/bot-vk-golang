package models

type Keyboard struct {
	OneTime bool       `json:"one_time"`
	Inline  bool       `json:"inline"`
	Buttons [][]Button `json:"buttons"`
}

type Button struct {
	Action Action `json:"action"`
	Color  string `json:"color"`
}

type Action struct {
	Type    string `json:"type"`
	Label   string `json:"label"`
	Payload string `json:"payload"`
}

type LinkKeyboard struct {
	OneTime bool           `json:"one_time"`
	Inline  bool           `json:"inline"`
	Buttons [][]LinkButton `json:"buttons"`
}

type LinkButton struct {
	Action LinkAction `json:"action"`
}

type LinkAction struct {
	Type    string `json:"type"`
	Label   string `json:"label"`
	Payload string `json:"payload"`
	Link    string `json:"link"`
}
