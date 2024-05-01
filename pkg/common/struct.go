package common

type PaginationQuery struct {
	Limit  int
	Offset int
	Page   int
}

type UploadWaterMeter struct {
	File string `json:"file"`
}
