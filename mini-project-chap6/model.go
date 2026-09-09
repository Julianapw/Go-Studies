package main

type UserResponse struct {
	Profile         string `json:"profile"`
	Orders          string `json:"orders"`
	Recommendations string `json:"recommendations"`
	Error           string `json:"error,omitempty"`
}
