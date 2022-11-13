package upload_operate

import "github.com/juetun/base-wrapper/lib/base"

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

func (r *ArgUploadGetInfo) Default(c *base.Context) (err error) {

	return
}

func NewArgUploadGetInfo() (res *ArgUploadGetInfo) {
	res = &ArgUploadGetInfo{
		ImgKeys:   make([]string, 0, 50),
		VideoKeys: make([]string, 0, 10),
		MusicKey:  make([]string, 0, 10),
		Material:  make([]string, 0, 50),
		File:      make([]string, 0, 50),
	}
	return
}

func NewResultMapUploadInfo() (res *ResultMapUploadInfo) {
	res = &ResultMapUploadInfo{
		Img:      make(map[string]*UploadImage, 50),
		Video:    make(map[string]*UploadVideo, 10),
		Music:    make(map[string]*UploadMusic, 10),
		Material: make(map[string]*UploadMaterial, 50),
		File:     make(map[string]*UploadFile, 50),
	}
	return
}
