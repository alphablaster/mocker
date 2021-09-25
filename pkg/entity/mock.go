package entity

type Mock struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        int    `json:"status"`
	ContentType   string `json:"content_type"`
	RequestMethod string `json:"request_method"`
	RoutePath     string `json:"route_path"`
	BodyType      string `json:"body_type"`
	BodyContent   string `json:"body_content"`
	ScriptType    string `json:"script_type"`
	Script        string `json:"script"`
	Active        bool   `json:"active"`
	MockOrder     int    `json:"mock_order"`
}
