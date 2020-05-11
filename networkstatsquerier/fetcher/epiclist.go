package fetcher

// EpicList implements the Endpoints interface for the Epic endpoint lists
type EpicList struct{}

// GetEndpoints takes the list of epic endpoints and formats it into a ListOfEndpoints
func (el EpicList) GetEndpoints(epicList []map[string]interface{}) ListOfEndpoints {
	return getDefaultEndpoints(epicList, string(Epic))
}