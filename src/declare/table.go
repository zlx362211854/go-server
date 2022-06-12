package declare


type Table struct {
	Id string `json:"id""`
	CreateTime int64 `form:"createTime" json:"createTime" xml:"createTime"`
	ModelName string `json:"modelName" binding:"required"`
	ModelKey  string `json:"modelKey" binding:"required"`
	Entity    string `json:"entity" binding:"required"`
	Columns    []struct {
		DataIndex    string `json:"dataIndex" binding:"required"`
		Title  string `json:"title" binding:"required"`
		Rules []string `json:"rules" binding:"required"`
	} `json:"columns" binding:"required"`
}
