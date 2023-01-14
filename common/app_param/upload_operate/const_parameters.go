package upload_operate

import (
	"github.com/juetun/base-wrapper/lib/common/redis_pkg"
	"time"
)

const (
	FileTypeVideo    = "video"    // 视频文件
	FileTypeMusic    = "music"    // 音频文件
	FileTypeFile     = "file"     // 普通文件
	FileTypePicture  = "image"    // 普通图片
	FileTypeMaterial = "material" // 重要资料，如身份证，证书之类的数据
)

var (
	CacheUploadCache = redis_pkg.CacheProperty{ //店铺变跟资质缓存的缓存Key (此缓存为二级缓存)
		Key:    "m:upload:%s:%s",
		Expire: 60 * time.Second, //只缓存1分钟，此设置时间一定要谨慎(最好不调整)
	}
)
