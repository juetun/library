package upload_operate

import (
	"encoding/json"
	"github.com/juetun/base-wrapper/lib/base"
)

type (
	UploadInfo struct {
		//Img      *UploadImage    `json:"img"`
		Video *UploadVideo `json:"video,omitempty"`
		Music *UploadMusic `json:"music,omitempty"`
		//Material *UploadMaterial `json:"material"`
		File *UploadFile `json:"file,omitempty"`
	}
	ResultMapUploadInfo struct {
		//Img      map[string]*UploadImage    `json:"img"`
		Video    map[string]*UploadVideo `json:"video,omitempty"`
		Music    map[string]*UploadMusic `json:"music,omitempty"`
		Download map[string]*UploadFile  `json:"download,omitempty"`
		//Material map[string]*UploadMaterial `json:"material"`
		File map[string]*UploadFile `json:"file,omitempty"`
	}

	//复制文件返回参数
	ResultMapCopyUploadInfo struct {
		Video    map[string]string `json:"video,omitempty"`
		Music    map[string]string `json:"music,omitempty"`
		Download map[string]string `json:"download,omitempty"`
		File     map[string]string `json:"file,omitempty"`
	}

	ArgUploadGetInfo struct {
		//ImgKeys   []string `json:"img_keys"`
		VideoKeys []string `json:"video_keys,omitempty"`
		MusicKey  []string `json:"music_key,omitempty"`
		//Material  []string `json:"material"`
		File     []string `json:"file,omitempty"`
		Download []string `json:"download"`
		base.GetDataTypeCommon
	}

	//删除上传文件接口
	ArgUploadRemove struct {
		//ImgKeys   []string `json:"img_keys"`
		ExceptVideoKeys []string `json:"except_video_keys"` //忽略视频文件
		ExceptMusicKey  []string `json:"except_music_key"`  //忽略音频文件
		//Material  []string `json:"material"`
		ExceptFile     []string `json:"except_file"`      //忽略文件
		FileType       []string `json:"file_type"`        //文件类型
		UploadDataType string   `json:"upload_data_type"` //数据类型
		UploadDataId   string   `json:"upload_data_id"`   //数据ID
		Channel        string   `json:"channel"`          //渠道号
	}
	ResultUploadRemove struct {
		Result bool `json:"result"`
	}
)

func NewArgUploadRemove() (res *ArgUploadRemove) {
	res = &ArgUploadRemove{
		//ImgKeys:   make([]string, 0, 50),
		ExceptVideoKeys: make([]string, 0, 50),
		ExceptMusicKey:  make([]string, 0, 50),
		//Material:  make([]string, 0, 50),
		ExceptFile: make([]string, 0, 50),
	}
	return
}

//判断数据是否为空
func (r *ArgUploadRemove) IsNull() (isNull bool) {
	if r.UploadDataType == "" && r.UploadDataId == "" {
		isNull = true
		return
	}
	return
}

func (r *ArgUploadRemove) Default(c *base.Context) (err error) {
	return
}

func (r *UploadInfo) UnmarshalBinary(data []byte) (err error) {
	if data == nil {
		return
	}
	err = json.Unmarshal(data, r)
	return
}

//实现 序列化方法 encoding.BinaryMarshaler
func (r *UploadInfo) MarshalBinary() (data []byte, err error) {
	if r == nil {
		return
	}
	data, err = json.Marshal(r)
	return
}

func (r *ArgUploadGetInfo) Default(c *base.Context) (err error) {

	return
}

func NewArgUploadGetInfo() (res *ArgUploadGetInfo) {
	res = &ArgUploadGetInfo{
		//ImgKeys:   make([]string, 0, 50),
		VideoKeys: make([]string, 0, 50),
		MusicKey:  make([]string, 0, 50),
		//Material:  make([]string, 0, 50),
		File: make([]string, 0, 50),
	}
	return
}

//判断数据是否为空
func (r *ArgUploadGetInfo) IsNull() (isNull bool) {

	if len(r.VideoKeys) == 0 && len(r.MusicKey) == 0 &&
		//len(r.Material) == 0 &&
		len(r.File) == 0 && len(r.Download) == 0 {
		isNull = true
		return
	}
	return
}

func NewResultMapUploadInfo(arg *ArgUploadGetInfo) (res *ResultMapUploadInfo) {
	res = &ResultMapUploadInfo{}
	res.Video = make(map[string]*UploadVideo, len(arg.VideoKeys))
	res.Music = make(map[string]*UploadMusic, len(arg.MusicKey))
	res.File = make(map[string]*UploadFile, len(arg.File))
	res.Download = make(map[string]*UploadFile, len(arg.Download))
	return
}

func NewResultMapCopyUploadInfo(arg *ArgUploadGetInfo) (res *ResultMapCopyUploadInfo) {
	res = &ResultMapCopyUploadInfo{}
	res.Video = make(map[string]string, len(arg.VideoKeys))
	res.Music = make(map[string]string, len(arg.MusicKey))
	res.File = make(map[string]string, len(arg.File))
	res.Download = make(map[string]string, len(arg.Download))
	return
}
