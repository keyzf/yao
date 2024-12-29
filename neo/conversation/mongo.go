package conversation

// Mongo conversation
type Mongo struct{}

// NewMongo create a new conversation
func NewMongo() *Mongo {
	return &Mongo{}
}

// UpdateChatTitle update the chat title
func (conv *Mongo) UpdateChatTitle(sid string, cid string, title string) error {
	return nil
}

// GetChats get the chat list
func (conv *Mongo) GetChats(sid string, filter ChatFilter) (*ChatGroupResponse, error) {
	return &ChatGroupResponse{
		Groups:   []ChatGroup{},
		Page:     filter.Page,
		PageSize: filter.PageSize,
		Total:    0,
		LastPage: 1,
	}, nil
}

// GetHistory get the history
func (conv *Mongo) GetHistory(sid string, cid string) ([]map[string]interface{}, error) {
	return []map[string]interface{}{}, nil
}

// SaveHistory save the history
func (conv *Mongo) SaveHistory(sid string, messages []map[string]interface{}, cid string, context map[string]interface{}) error {
	return nil
}

// GetRequest get the request
func (conv *Mongo) GetRequest(sid string, rid string) ([]map[string]interface{}, error) {
	return nil, nil
}

// SaveRequest save the request
func (conv *Mongo) SaveRequest(sid string, rid string, cid string, messages []map[string]interface{}) error {
	return nil
}

// GetChat get the chat info and its history
func (conv *Mongo) GetChat(sid string, cid string) (*ChatInfo, error) {
	return nil, nil
}

// DeleteChat deletes a specific chat and its history
func (conv *Mongo) DeleteChat(sid string, cid string) error {
	return nil
}

// DeleteAllChats deletes all chats and their histories for a user
func (conv *Mongo) DeleteAllChats(sid string) error {
	return nil
}

// SaveAssistant creates or updates an assistant
func (conv *Mongo) SaveAssistant(assistant map[string]interface{}) error {
	return nil
}

// DeleteAssistant deletes an assistant by assistant_id
func (conv *Mongo) DeleteAssistant(assistantID string) error {
	return nil
}

// GetAssistants retrieves assistants with pagination and tag filtering
func (conv *Mongo) GetAssistants(filter AssistantFilter) (*AssistantResponse, error) {
	return &AssistantResponse{
		Data:     []map[string]interface{}{},
		Page:     filter.Page,
		PageSize: filter.PageSize,
		PageCnt:  0,
		Next:     0,
		Prev:     0,
		Total:    0,
	}, nil
}
