package mall

import (
	"fmt"
	"github.com/juetun/base-wrapper/lib/base"
	"strconv"
	"strings"
)

type (
	ArgGetByAddressIds struct {
		IdString string `json:"ids" form:"ids"`
		base.GetDataTypeCommon
		Ids []int64 `json:"-" form:"-"`
	}

	ResultGetByAddressIds map[int64]*ResultGetByAddressIdsItem

	ResultGetByAddressIdsItem struct {
		Id           int64  `json:"id"`
		ProvinceId   string `json:"province_id"`
		CityId       string `json:"city_id"`
		AreaId       string `json:"area_id"`
		Province     string `json:"province"`
		City         string `json:"city"`
		Area         string `json:"area"`
		Title        string `json:"title"`
		Address      string `json:"address"`
		AreaAddress  string `json:"area_address"`
		ContactUser  string `json:"contact_user"`
		ContactPhone string `json:"contact_phone"`
		FullAddress  string `json:"full_address"`
	}
)

func (r *ResultGetByAddressIdsItem) GetToCityId() (res string) {
	res = r.CityId
	return
}

func (r *ArgGetByAddressIds) Default(ctx *base.Context) (err error) {
	idString := strings.Split(r.IdString, ",")
	r.Ids = make([]int64, 0, len(idString))
	var id int64
	for _, idStr := range idString {
		if idStr == "" {
			continue
		}
		if id, err = strconv.ParseInt(idStr, 10, 64); err != nil {
			err = fmt.Errorf("地址ID格式不正确")
			return
		}
		r.Ids = append(r.Ids, id)
	}

	return
}
