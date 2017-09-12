package cwl

type Event struct {
	Group     string `json:"group"`
	Message   string `json:"msg"`
	Region    string `json:"region"`
	Stream    string `json:"stream"`
	Timestamp string `json:"time"`
}
