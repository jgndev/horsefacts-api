package types

type BreedResponse struct {
	Name        string   `json:"Name"`
	Country     string   `json:"Country"`
	Colors      []string `json:"Colors"`
	Established string   `json:"Established"`
}
