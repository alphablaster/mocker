package entity

type Mock struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Status        int    `json:"status"`
	ContentType   string `json:"content_type" db:"content_type"`
	RequestMethod string `json:"request_method" db:"request_method"`
	RoutePath     string `json:"route_path" db:"route_path"`
	BodyType      string `json:"body_type" db:"body_type"`
	BodyContent   string `json:"body_content" db:"body_content"`
	ScriptType    string `json:"script_type" db:"script_type"`
	Script        string `json:"script"`
	Active        bool   `json:"active"`
	MockOrder     int    `json:"mock_order" db:"mock_order"`
}
