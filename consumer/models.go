package consumer

type ConsumerRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=1,max=100"` // Name of the Consumer.
	GroupID  string `json:"group_id" validate:"required,alphanum,min=1,max=100"` // Group of the Consumer.
	Plugins  string `json:"plugins" validate:"required,alphanum,min=1,max=100"`  // Plugins that are executed during the request/response cycle.
	Desc     string `json:"desc" validate:"required,alphanum,min=1,max=100"`     // Description of usage scenarios.
	Labels   string `json:"labels" validate:"required,alphanum,min=1,max=100"`   // Attributes of the Consumer specified as key-value pairs.
}

type GetResponse struct {
	List []struct {
		Key   string `json:"key"`
		Value struct {
			CreateTime int    `json:"create_time"`
			UpdateTime int    `json:"update_time"`
			Username   string `json:"username"`
		} `json:"value"`
		CreatedIndex  int `json:"createdIndex"`
		ModifiedIndex int `json:"modifiedIndex"`
	} `json:"list"`
	Total int `json:"total"`
}

type GetByUsernameResponse struct {
	Key   string `json:"key"`
	Value struct {
		CreateTime int    `json:"create_time"`
		UpdateTime int    `json:"update_time"`
		Username   string `json:"username"`
	} `json:"value"`
	CreatedIndex  int `json:"createdIndex"`
	ModifiedIndex int `json:"modifiedIndex"`
}

type PutResponse struct {
	Key   string `json:"key"`
	Value struct {
		CreateTime int    `json:"create_time"`
		UpdateTime int    `json:"update_time"`
		Username   string `json:"username"`
	} `json:"value"`
}

type DeleteResponse struct {
	Key     string `json:"key"`
	Deleted string `json:"deleted"`
}
