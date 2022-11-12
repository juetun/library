package upload_operate

type (
	UploadInfo struct {
		Img   *UploadImage `json:"img"`
		Video *UploadVideo `json:"video"`
		Music *UploadMusic `json:"music"`
	}
	ResultMapUploadInfo struct {
		Img   map[string]*UploadImage `json:"img"`
		Video map[string]*UploadVideo `json:"video"`
		Music map[string]*UploadMusic `json:"music"`
	}
	ArgUploadGetInfo struct {
		ImgKeys   []string `json:"img_keys"`
		VideoKeys []string `json:"video_keys"`
		MusicKey  []string `json:"music_key"`
	}
)
