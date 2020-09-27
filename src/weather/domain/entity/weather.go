package entity

type Weather struct {
	ID     string `json:"id"`
	Value  string `json:"value"`
	Metric string `json:"metric"`
	City   string `json:"city"`
}

func (w *Weather) Weather() interface{} {
	return &Weather{
		ID:     w.ID,
		Value:  w.Value,
		Metric: w.Metric,
		City:   w.City,
	}
}
