package transport

type (
	GetDirectionStepRes struct {
		ServiceRes
		ID          string
		Name        string
		Description string
		Duration    string
	}
)
