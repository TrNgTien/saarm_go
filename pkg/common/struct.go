package common

type PaginationQuery struct {
	Limit  int
	Offset int
	Page   int
}

type UploadWaterMeter struct {
	CroppedFile string `json:"croppedFile"`
	OriginalFile string `json:"originalFile"`
}
