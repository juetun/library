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
				EmsAddress: `{"id":1,"province_id":"510000","city_id":"510100","area_id":"510104","province":"","city":"","area":"","title":"家里","address":"士大夫阿斯蒂芬","area_address":"","contact_user":"李女士","contact_phone":"15108352617","full_address":""}
`,
				SKusFreight: `[{"num":1,"sku":{"sku_id":"441542713672007784","sku_name":"阿斯顿发射点阿斯顿发射点阿斯顿发射点阿斯顿发射点阿斯顿发射点鲜活生鲜螃蟹现货水水鞋子32码对8只 鲜活生鲜螃蟹现货水水","thumbnail":"image|spu|13","thumbnail_url":"","lock_key":"","sku_att_relate_id":0,"image":"","video":"","user_hid":1,"shop_id":1,"sku_status":1,"weight":"0.00","price":"1.00","market_cost":"0.00","price_cost":"0.00","shop_sale_code":"support_code1","provide_channel":0,"provide_sale_code":"asdfadsf0101","sale_num":0,"sale_online_time":"0001-01-01 00:00:00","sale_over_time":null,"volume":"","flag_tester":2,"have_bind_spu":0,"created_at":"0001-01-01 00:00:00","updated_at":"2023-05-08 20:58:11"},"sku_relate":{"id":262,"shop_id":1,"product_id":"440958537740459917","category_id":0,"parent_id":0,"pk":"1","sku_name":"馆长在离石区政府办事时被区长张海文辱骂，推倒致伤，现状态很不稳定，不能正常开馆","sku_id":"441542713672007784","price":"1.00","is_not_attr_name":2,"property_id":1,"spu_status":1,"sale_type":1,"down_payment":"0.00","final_payment":"0.00","freight_template":1,"sale_online_time":"2023-05-07 14:55:48","sale_over_time":null,"final_start_time":"2023-05-07 14:55:48","final_over_time":"2023-05-07 14:55:48","sales_tax_rate":"0.00","sales_tax_rate_value":"0.00","max_limit_num":500,"min_limit_num":1,"is_leaf":1,"sku_status":1,"have_gift":2,"created_at":"2023-05-07 14:55:49","updated_at":"2023-05-07 14:56:24"},"spu":{"product_id":"440958537740459917","title":"国家双减政策的实施,让","user_hid":0,"thumbnail_url":"","brand_id":1,"video":"","video_url":"","shop_id":1,"status":1,"sub_title":"","min_price":"1.00","max_price":"50.00","min_price_cost":"0.00","max_price_cost":"","min_down_payment":"0.00","max_down_payment":"0.00","service_ids":"[1]","keywords":"","sale_num":0,"freight_type":1,"freight_template":1,"total_stock":0,"category_id":2,"sale_type":1,"pull_on_time":"2023-05-07 14:56:24","pull_off_time":null,"sale_online_time":"2023-05-07 14:56:24","sale_over_time":null,"final_start_time":"2000-01-01 00:00:00","final_over_time":"2000-01-01 00:00:00","delivery_time":null,"sale_count_show":0,"relate_type":0,"relate_item_id":"","relate_buy_count":0,"relate_buy_amount":"0.00","settle_type":2,"flag_tester":1},"shop":{"shop_id":1,"name":"测试店铺测试店铺测试店铺测","logo_url":"","bg_image_url":"","bg_image":"image|shop_logo|132","shop_type":2,"shop_entry_type":4,"status":1,"flag_tester":2,"admin_user_hid":1,"need_verify_status":2,"verify_status":"0","created_at":"2023-04-21 17:15:56","updated_at":"2023-04-21 17:16:12"},"ems_address_freight":null,"TemplateFreight":{"freight_model":{"id":1,"shop_id":1,"title":"指定条件包邮","province_id":310000,"city_id":310100,"area_id":310110,"free_freight":2,"pricing_mode":2,"have_use":2,"sale_area":"[{\"a\":[\"140000\",\"330400\",\"371600\",\"120100\",\"130900\",\"320000\",\"360500\",\"131000\",\"341700\",\"361000\",\"150900\",\"341300\",\"360400\",\"130100\",\"140800\",\"330100\",\"330600\",\"340100\",\"370000\",\"a_2\",\"130000\",\"130200\",\"370900\",\"371100\",\"140600\",\"331000\",\"340700\",\"341200\",\"360300\",\"150800\",\"320200\",\"320600\",\"130500\",\"140300\",\"140500\",\"140900\",\"150500\",\"321200\",\"330800\",\"340400\",\"321000\",\"130600\",\"321300\",\"340000\",\"340200\",\"341100\",\"320300\",\"150600\",\"340500\",\"370200\",\"139000\",\"330700\",\"371300\",\"140400\",\"141000\",\"152500\",\"330300\",\"340800\",\"130800\",\"340600\",\"140100\",\"360700\",\"a_1\",\"150100\",\"320100\",\"341800\",\"360000\",\"360600\",\"361100\",\"130700\",\"141100\",\"310100\",\"330500\",\"340300\",\"110000\",\"140200\",\"371200\",\"150700\",\"360100\",\"360200\",\"370500\",\"130400\",\"360900\",\"371700\",\"131100\",\"310000\",\"320800\",\"341000\",\"371000\",\"320900\",\"330000\",\"331100\",\"360800\",\"371500\",\"150300\",\"320700\",\"370600\",\"120000\",\"150000\",\"150200\",\"152900\",\"330200\",\"150400\",\"321100\",\"341600\",\"370300\",\"130300\",\"152200\",\"370400\",\"370700\",\"371400\",\"320500\",\"370100\",\"140700\",\"320400\",\"341500\",\"370800\",\"110100\",\"330900\"],\"fg\":\"12\",\"fp\":\"1.00\",\"eg\":\"2\",\"ep\":\"1.00\"},{\"a\":[\"410200\",\"410400\",\"411200\",\"411300\",\"420000\",\"420500\",\"421000\",\"422800\",\"430500\",\"430700\",\"410900\",\"411000\",\"411700\",\"420300\",\"420600\",\"421300\",\"430300\",\"431200\",\"a_3\",\"410600\",\"430100\",\"431000\",\"410100\",\"411500\",\"411600\",\"419000\",\"420100\",\"420900\",\"430400\",\"430600\",\"421200\",\"430000\",\"430200\",\"431100\",\"410500\",\"411100\",\"410000\",\"410300\",\"410800\",\"420700\",\"420800\",\"421100\",\"431300\",\"410700\",\"411400\",\"420200\",\"429000\",\"430800\",\"430900\",\"433100\"],\"fg\":\"12\",\"fp\":\"3.00\",\"eg\":\"3\",\"ep\":\"1.00\"}]","free_condition":"[{\"a\":[\"410900\",\"420100\",\"421300\",\"419000\",\"420500\",\"420600\",\"420000\",\"410800\",\"411200\",\"411700\",\"420700\",\"410400\",\"410200\",\"410300\",\"410700\",\"411300\",\"411400\",\"420300\",\"420900\",\"410100\",\"421200\",\"411000\",\"411100\",\"411600\",\"420200\",\"410600\",\"422800\",\"429000\",\"410000\",\"411500\",\"420800\",\"421000\",\"421100\",\"410500\"],\"ft\":2,\"fp\":\"10.00\",\"fn\":2},{\"a\":[\"310000\",\"310100\"],\"ft\":2,\"fp\":\"10.00\",\"fn\":23}]","postage_condition":1,"created_at":"2023-03-29 22:34:17","updated_at":"2023-05-08 00:00:00"},"areas":null},"from_city_id":"0"},{"num":1,"sku":{"sku_id":"441542713672269928","sku_name":"尺码","thumbnail":"image|spu|21","thumbnail_url":"","lock_key":"","sku_att_relate_id":0,"image":"","video":"","user_hid":1,"shop_id":1,"sku_status":1,"weight":"0.00","price":"1.00","market_cost":"0.00","price_cost":"0.00","shop_sale_code":"","provide_channel":0,"provide_sale_code":"","sale_num":0,"sale_online_time":"2023-02-20 09:03:22","sale_over_time":null,"volume":"","flag_tester":1,"have_bind_spu":0,"created_at":"2023-02-20 09:03:22","updated_at":"2023-04-05 10:01:14"},"sku_relate":{"id":264,"shop_id":1,"product_id":"440958537740394381","category_id":0,"parent_id":0,"pk":"4","sku_name":"尺码","sku_id":"441542713672269928","price":"1.00","is_not_attr_name":2,"property_id":4,"spu_status":1,"sale_type":1,"down_payment":"0.00","final_payment":"0.00","freight_template":1,"sale_online_time":"2023-05-07 14:59:27","sale_over_time":null,"final_start_time":"2023-05-07 14:59:27","final_over_time":"2023-05-07 14:59:27","sales_tax_rate":"0.00","sales_tax_rate_value":"0.00","max_limit_num":500,"min_limit_num":1,"is_leaf":1,"sku_status":1,"have_gift":2,"created_at":"2023-05-07 14:59:27","updated_at":"2023-05-07 14:59:53"},"spu":{"product_id":"440958537740394381","title":"多项数据创新高带动中国经济回暖","user_hid":0,"thumbnail_url":"","brand_id":1,"video":"","video_url":"","shop_id":1,"status":1,"sub_title":"","min_price":"1.00","max_price":"1.00","min_price_cost":"0.00","max_price_cost":"","min_down_payment":"0.00","max_down_payment":"0.00","service_ids":"[1]","keywords":"123","sale_num":0,"freight_type":1,"freight_template":1,"total_stock":0,"category_id":11,"sale_type":2,"pull_on_time":"2023-05-07 14:59:53","pull_off_time":null,"sale_online_time":"2023-05-07 14:59:53","sale_over_time":null,"final_start_time":"2000-01-01 00:00:00","final_over_time":"2000-01-01 00:00:00","delivery_time":"2023-05-23 00:00:00","sale_count_show":0,"relate_type":0,"relate_item_id":"","relate_buy_count":0,"relate_buy_amount":"0.00","settle_type":2,"flag_tester":1},"shop":{"shop_id":1,"name":"测试店铺测试店铺测试店铺测","logo_url":"","bg_image_url":"","bg_image":"image|shop_logo|132","shop_type":2,"shop_entry_type":4,"status":1,"flag_tester":2,"admin_user_hid":1,"need_verify_status":2,"verify_status":"0","created_at":"2023-04-21 17:15:56","updated_at":"2023-04-21 17:16:12"},"ems_address_freight":null,"TemplateFreight":{"freight_model":{"id":0,"shop_id":0,"title":"","province_id":0,"city_id":0,"area_id":0,"free_freight":0,"pricing_mode":0,"have_use":0,"sale_area":"","free_condition":"","postage_condition":0,"created_at":"0001-01-01 00:00:00","updated_at":"0001-01-01 00:00:00"},"areas":null},"from_city_id":"0"}]
`,
			},
		},
	}
	var err error
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				resEmsAddress ResultGetByAddressIdsItem
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
			t.Logf("%v", tt.want)
		})
	}
}
