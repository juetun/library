package freight

import (
	"encoding/json"
	"testing"
)

//邮费计算测试
func TestNewPriceFreight(t *testing.T) {
	type args struct {
		EmsAddress  string `json:"ems_address"`
		SKusFreight string `json:"s_kus_freight"`
	}
	tests := []struct {
		name string
		args args
		want *PriceFreightResult
	}{
		{
			args: args{
				EmsAddress: `{"id":1,"title":"家里","user_hid":1,"province_id":"510000","city_id":"510100","area_id":"510104","address":"士大夫阿斯蒂芬","zip_code":"123123","contact_user":"李女士","contact_phone":"15108352617","status":1}`,
				SKusFreight: `[{"num":1,"sku":{"sku_id":"441542713672269928","sku_name":"尺码","thumbnail":"image|spu|21","thumbnail_url":"","lock_key":"","sku_att_relate_id":0,"image":"","video":"","user_hid":1,"shop_id":1,"sku_status":1,"weight":"0.00","price":"1.00","market_cost":"0.00","price_cost":"0.00","shop_sale_code":"","provide_channel":0,"provide_sale_code":"","sale_num":0,"sale_online_time":"2023-02-20 09:03:22","sale_over_time":null,"volume":"","flag_tester":1,"have_bind_spu":0,"created_at":"2023-02-20 09:03:22","updated_at":"2023-04-05 10:01:14"},"sku_relate":{"id":264,"shop_id":1,"product_id":"440958537740394381","category_id":0,"parent_id":0,"pk":"4","sku_name":"尺码","sku_id":"441542713672269928","price":"1.00","is_not_attr_name":2,"property_id":4,"spu_status":1,"sale_type":1,"down_payment":"0.00","final_payment":"0.00","freight_template":1,"sale_online_time":"2023-05-07 14:59:27","sale_over_time":null,"final_start_time":"2023-05-07 14:59:27","final_over_time":"2023-05-07 14:59:27","sales_tax_rate":"0.00","sales_tax_rate_value":"0.00","max_limit_num":500,"min_limit_num":1,"is_leaf":1,"sku_status":1,"have_gift":2,"created_at":"2023-05-07 14:59:27","updated_at":"2023-05-07 14:59:53"},"spu":{"product_id":"440958537740394381","title":"多项数据创新高带动中国经济回暖","user_hid":0,"thumbnail_url":"","brand_id":1,"video":"","video_url":"","shop_id":1,"status":1,"sub_title":"","min_price":"1.00","max_price":"1.00","min_price_cost":"0.00","max_price_cost":"","min_down_payment":"0.00","max_down_payment":"0.00","service_ids":"[1]","keywords":"123","sale_num":0,"freight_type":1,"freight_template":1,"total_stock":0,"category_id":11,"sale_type":2,"pull_on_time":"2023-05-07 14:59:53","pull_off_time":null,"sale_online_time":"2023-05-07 14:59:53","sale_over_time":null,"final_start_time":"2000-01-01 00:00:00","final_over_time":"2000-01-01 00:00:00","delivery_time":"2023-05-23 00:00:00","sale_count_show":0,"relate_type":0,"relate_item_id":"","relate_buy_count":0,"relate_buy_amount":"0.00","settle_type":2,"flag_tester":1},"shop":{"shop_id":1,"name":"测试店铺测试店铺测试店铺测","logo_url":"","bg_image_url":"","bg_image":"image|shop_logo|132","shop_type":2,"shop_entry_type":4,"status":1,"flag_tester":2,"admin_user_hid":1,"need_verify_status":2,"verify_status":"0","created_at":"2023-04-21 17:15:56","updated_at":"2023-04-21 17:16:12"},"ems_address_freight":null,"TemplateFreight":{"freight_model":{"id":0,"shop_id":0,"title":"","province_id":0,"city_id":0,"area_id":0,"free_freight":0,"pricing_mode":0,"have_use":0,"sale_area":"","free_condition":"","postage_condition":0,"created_at":"0001-01-01 00:00:00","updated_at":"0001-01-01 00:00:00"},"areas":null},"from_city_id":"0"}]
`,
			},
		},
	}
	var err error
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				resEmsAddress EmsAddressFreight
				sKusFreight   []*SkuFreightSingle
			)
			if err = json.Unmarshal([]byte(tt.args.EmsAddress), &resEmsAddress); err != nil {
				t.Errorf("UnmarshalEmsAddress() = %v", err.Error())
				return
			}
			if err = json.Unmarshal([]byte(tt.args.SKusFreight), &sKusFreight); err != nil {
				t.Errorf("Unmarshal() = %v", err.Error())
				return
			}
			//邮费计算逻辑
			if tt.want, err = NewPriceFreight(
				//freight.OptionFreightContext(r.Context),
				OptionFreightEmsAddress(&resEmsAddress),
			).
				AppendNeedCalSKus(sKusFreight...).
				Calculate(); err != nil {
				t.Errorf("NewPriceFreight() = %v", err.Error())
			}

		})
	}
}
