package upload_operate

type (
	MapUploadInfo map[string]UploadInfo
	UploadInfo    struct {
		Img   *UploadImage `json:"img"`
		Video *UploadVideo `json:"video"`
		Music *UploadMusic `json:"music"`
	}
)
