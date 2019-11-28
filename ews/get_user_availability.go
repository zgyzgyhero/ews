package ews

type GetUserAvailabilityRequest struct {
	XMLName struct{} `xml:"m:GetUserAvailabilityRequest"`
}

type GetUserAvailabilityResponse struct {
}

// GetUserAvailability
//https://docs.microsoft.com/en-us/exchange/client-developer/web-service-reference/getuseravailability-operation
func GetUserAvailability(c *Client, r *GetUserAvailabilityRequest) (*GetUserAvailabilityResponse, error) {

	return nil, nil
}
