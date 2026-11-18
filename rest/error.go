package rest

type EpicGamesRestError struct {
	ErrorCode          string   `json:"errorCode"`
	ErrorMessage       string   `json:"errorMessage"`
	NumericErrorCode   int      `json:"numericErrorCode"`
	OriginatingService string   `json:"originatingService"`
	MessageVars        []string `json:"messageVars,omitempty"`
	Intent             string   `json:"intent"`
	Continuation       string   `json:"continuation,omitempty"`
	ContinuationURL    string   `json:"continuationUrl,omitempty"`
	CorrectiveAction   string   `json:"correctiveAction,omitempty"`
}
