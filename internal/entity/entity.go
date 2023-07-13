package entity

type EventResponse struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	Country     string `json:"country"`
}

type EventRequest struct {
	TypeId      int    `json:"typeid"`
	Description string `json:"description"`
	CountryId   int    `json:"countryid"`
}

type CountryMetric struct {
	Country    string `json:"country"`
	EventCount int    `json:"event_count"`
}

// Data representa un dato compuesto por ID y descripción
type Data struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
