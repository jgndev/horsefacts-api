package types

type Breed struct {
	ID          int      `json:"ID"`
	Name        string   `json:"Name"`
	Country     string   `json:"Country"`
	Colors      []string `json:"Colors"`
	Established string   `json:"Established"`
}
