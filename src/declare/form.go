package declare

type Form struct {
	Id string `json:"id""`
	CreateTime int64 `form:"createTime" json:"createTime" xml:"createTime"`
	ModelName string `json:"modelName" binding:"required"`
	ModelKey  string `json:"modelKey" binding:"required"`
	Entity    string `json:"entity" binding:"required"`
	Fields    []struct {
		Key    string `json:"key" binding:"required"`
		Label  string `json:"label" binding:"required"`
		UIType string `json:"uiType" binding:"required"`
		Rules []struct {
			Required bool `json:"required" binding:"required"`
			Message string `json:"message" binding:"required"`
		} `json:"rules" binding:"required"`
	} `json:"fields" binding:"required"`
}

