package mall_comment

type (
	UInfo struct {
		Content  string `json:"content"`
		NickName string `json:"nickname"`
		Score    int    `json:"score"`
		Avatar   string `json:"avatar"`
		Time     string `json:"time"`
		Size     string `json:"size"`
		Replay   int64  `json:"replay"`
		Like     int64  `json:"like"`
	}
	CommentVideoItem struct {
		Src      string `json:"src"`
		MainUrl  string `json:"mainUrl"`  //图片地址
		VideoUrl string `json:"videoUrl"` //视频地址
	}
	CommentImageItem struct {
		Src         string `json:"src"`
		SmallImgUrl string `json:"smallImgUrl"`
		BigImgUrl   string `json:"bigImgUrl"`
		ImgUrl      string `json:"imgUrl"`
	}
	CommentFollow struct {
		Days    int      `json:"days"`
		Content string   `json:"content"`
		Images  []string `json:"images"`
	}
	CommentItem struct {
		Videos   []*CommentVideoItem `json:"videos"` //视频
		Images   []*CommentImageItem `json:"images"` //图片
		Follow   CommentFollow       `json:"follow"`
		UserInfo UInfo               `json:"info"` //用户信息
	}
	OrderComment struct {
		ShopGoodBit bool           `json:"shop_good_bit"` //是否展示好评度
		GoodBit     string         `json:"good_bit"`      //好评度
		Comment     []*CommentItem `json:"Comment"`       //评论列表
		Number      int64          `json:"number"`        //评论数量
	}
)
