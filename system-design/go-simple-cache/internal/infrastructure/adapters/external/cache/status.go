package cache

type Status struct {
	Keys []string `json:"keys"`
	Size int      `json:"size"`
}
