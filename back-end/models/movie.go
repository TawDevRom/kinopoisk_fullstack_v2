package models

type Episode struct {
	Title    string `json:"title"`
	VideoURL string `json:"video-url"`
}

type Films struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	IsSerial    bool   `json:"is_serial"`
	Description string `json:"description"`
}
