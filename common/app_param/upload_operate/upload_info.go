package upload_operate

type (
	UploadInfo struct {
		Img      *UploadImage    `json:"img"`
		Video    *UploadVideo    `json:"video"`
		Music    *UploadMusic    `json:"music"`
		Material *UploadMaterial `json:"material"`
		File     *UploadFile     `json:"file"`
	}
	ResultMapUploadInfo struct {
		Img      map[string]*UploadImage    `json:"img"`
		Video    map[string]*UploadVideo    `json:"video"`
		Music    map[string]*UploadMusic    `json:"music"`
		Material map[string]*UploadMaterial `json:"material"`
		File     map[string]*UploadFile     `json:"file"`
	}
	ArgUploadGetInfo struct {
		ImgKeys   []string `json:"img_keys"`
		VideoKeys []string `json:"video_keys"`
		MusicKey  []string `json:"music_key"`
		Material  []string `json:"material"`
		File      []string `json:"file"`
	}
)
