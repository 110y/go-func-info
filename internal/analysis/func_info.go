package analysis

type FuncInfo struct {
	Name     string        `json:"name"`
	Receiver *ReceiverInfo `json:"receiver,omitempty"`

	StartPos int `json:"start_pos"`
	EndPos   int `json:"end_pos"`
}

type ReceiverInfo struct {
	Name     string `json:"name"`
	TypeName string `json:"type_name"`
}
