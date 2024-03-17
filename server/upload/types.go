package upload

import (
	"encoding/json"
)

type uploadPostResponse struct {
	ID string `json:"id"`
}

func newUploadPostResponse(id string) *uploadPostResponse {
	return &uploadPostResponse{
		ID: id,
	}
}

func (r *uploadPostResponse) GetID() string {
	return r.ID
}

// MarshalToJson marshall the response to json
func (r *uploadPostResponse) MarshalToJson() ([]byte, error) {
	return json.Marshal(r)
}
