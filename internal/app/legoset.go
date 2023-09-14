package app

type LegoSet struct {
	Code         string  `json:"code,omitempty"`
	Name         string  `json:"name,omitempty"`
	PieceCount   int     `json:"piece,omitempty"`
	ImageURL     string  `json:"url,omitempty"`
	Price        float64 `json:"price,omitempty"`
	CostPerPiece float64 `json:"cost,omitempty"`
}
