package consumer

type ConsumerRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=1,max=100"` // Name of the Consumer.
	Desc     string `json:"desc" validate:"required,alphanum,min=1,max=100"`     // Description of usage scenarios.
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
