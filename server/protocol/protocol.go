package protocol

// client/src/protocol.ts should be kept in sync with this

type Entity struct {
	ID     uint64  `json:"id"`
	Name   *string `json:"name"`
	Color  string  `json:"color"`
	Sprite *string `json:"sprite"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Radius float64 `json:"radius"`
}

type LogEvent struct {
	LogType    int    `json:"logType"`
	ProtagName string `json:"protagName"`
	AntagName  string `json:"antagName"`
	Resources  int    `json:"resources"`
}

type WorldState struct {
	Time     uint64   `json:"time"`
	Entities []Entity `json:"entities"`
}

type Event struct {
	Type string
	Data interface{}
}
