package rest

type CreateKey struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CacheStatusResponse struct {
	Keys []string `json:"keys"`
	Size int      `json:"size"`
}
