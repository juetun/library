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
				SKusFreight: `[{"num":1,"sku":{"sku_id":"441542713672007784","sku_name":"阿斯顿发射点阿斯顿发射点阿斯顿发射点阿斯顿发射点阿斯顿发射点鲜活生鲜螃蟹现货水水鞋子32码对8只 鲜活生鲜螃蟹现货水水","thumbnail":"image|spu|13","thumbnail_url":"","lock_key":"","sku_att_relate_id":0,"image":"","video":"","user_hid":1,"shop_id":1,"sku_status":1,"weight":"1.50","price":"1.00","market_cost":"0.00","price_cost":"0.00","shop_sale_code":"support_code1","provide_channel":0,"provide_sale_code":"asdfadsf0101","sale_num":0,"sale_online_time":"0001-01-01 00:00:00","sale_over_time":null,"volume":"","flag_tester":2,"have_bind_spu":0,"created_at":"0001-01-01 00:00:00","updated_at":"2023-05-12 09:09:08"},"sku_relate":{"id":262,"shop_id":1,"product_id":"440958537740459917","category_id":0,"parent_id":0,"pk":"1","sku_name":"馆长在离石区政府办事时被区长张海文辱骂，推倒致伤，现状态很不稳定，不能正常开馆","sku_id":"441542713672007784","price":"1.00","is_not_attr_name":2,"property_id":1,"spu_status":1,"sale_type":1,"down_payment":"0.00","final_payment":"0.00","freight_template":2,"sale_online_time":"2023-05-07 14:55:48","sale_over_time":null,"final_start_time":"2023-05-07 14:55:48","final_over_time":"2023-05-07 14:55:48","sales_tax_rate":"0.00","sales_tax_rate_value":"0.00","max_limit_num":500,"min_limit_num":1,"is_leaf":1,"sku_status":1,"have_gift":2,"created_at":"2023-05-07 14:55:49","updated_at":"2023-05-13 16:26:51"},"spu":{"product_id":"440958537740459917","title":"国家双减政策的实施,让","user_hid":0,"thumbnail_url":"","brand_id":1,"video":"","video_url":"","shop_id":1,"status":1,"sub_title":"","min_price":"1.00","max_price":"50.00","min_price_cost":"0.00","max_price_cost":"","min_down_payment":"0.00","max_down_payment":"0.00","service_ids":"[1]","keywords":"","sale_num":0,"freight_type":1,"freight_template":2,"total_stock":0,"category_id":2,"sale_type":1,"pull_on_time":"2023-05-13 16:26:51","pull_off_time":null,"sale_online_time":"2023-05-13 16:26:51","sale_over_time":null,"final_start_time":"2000-01-01 00:00:00","final_over_time":"2000-01-01 00:00:00","delivery_time":null,"sale_count_show":0,"relate_type":0,"relate_item_id":"","relate_buy_count":0,"relate_buy_amount":"0.00","settle_type":2,"flag_tester":1},"shop":{"shop_id":1,"name":"测试店铺测试店铺测试店铺测测试店铺测试店铺测试店铺测测试店铺测试店铺测试店铺测测试店铺测试店铺测试店铺测","logo_url":"","bg_image_url":"","bg_image":"image|shop_logo|132","shop_type":2,"shop_entry_type":4,"status":1,"flag_tester":2,"admin_user_hid":1,"need_verify_status":2,"verify_status":"0","created_at":"2023-04-21 17:15:56","updated_at":"2023-04-21 17:16:12"},"ems_address_freight":null,"template_freight":{"freight_model":{"id":2,"shop_id":1,"title":"指定条件包邮(2)","province_id":110000,"city_id":110100,"area_id":110102,"free_freight":2,"pricing_mode":1,"have_use":2,"sale_area":"[{\"a\":[\"152200\",\"330700\",\"330800\",\"371100\",\"371300\",\"140900\",\"320000\",\"330900\",\"511300\",\"131100\",\"150500\",\"320200\",\"330200\",\"140100\",\"150200\",\"341200\",\"360700\",\"513400\",\"a_1\",\"510100\",\"130700\",\"140700\",\"150700\",\"310100\",\"340300\",\"340800\",\"371700\",\"510700\",\"511100\",\"512000\",\"130800\",\"310000\",\"320900\",\"130300\",\"321300\",\"370200\",\"511400\",\"120000\",\"150800\",\"152900\",\"360600\",\"361100\",\"370300\",\"510900\",\"131000\",\"330000\",\"360100\",\"361000\",\"370700\",\"371600\",\"130000\",\"150900\",\"340600\",\"371500\",\"140200\",\"141000\",\"320100\",\"370000\",\"513300\",\"130200\",\"360000\",\"370400\",\"370500\",\"510800\",\"511700\",\"110000\",\"130500\",\"130900\",\"140400\",\"150300\",\"330500\",\"340500\",\"120100\",\"320600\",\"321000\",\"510500\",\"511600\",\"320400\",\"320500\",\"360400\",\"511000\",\"511800\",\"130600\",\"140800\",\"341600\",\"360800\",\"320800\",\"330400\",\"370600\",\"511900\",\"130400\",\"140500\",\"330100\",\"341800\",\"130100\",\"140000\",\"341300\",\"510400\",\"150600\",\"340700\",\"370900\",\"371000\",\"510000\",\"a_2\",\"141100\",\"150100\",\"330600\",\"331000\",\"139000\",\"150000\",\"370100\",\"152500\",\"320300\",\"320700\",\"371400\",\"140300\",\"321100\",\"321200\",\"330300\",\"340000\",\"340100\",\"341700\",\"110100\",\"150400\",\"331100\",\"360900\",\"370800\",\"371200\",\"513200\",\"140600\",\"340400\",\"341000\",\"341500\",\"360300\",\"510300\",\"511500\",\"340200\",\"341100\",\"360200\",\"360500\",\"510600\"],\"fg\":\"0\",\"fp\":\"0.00\",\"eg\":\"0\",\"ep\":\"0.00\"},{\"a\":[\"220200\",\"420600\",\"421300\",\"430700\",\"431000\",\"451200\",\"532300\",\"632200\",\"211200\",\"429000\",\"630200\",\"211100\",\"532500\",\"640500\",\"210400\",\"232700\",\"421000\",\"420200\",\"540200\",\"220400\",\"230800\",\"350600\",\"533100\",\"620000\",\"620600\",\"211400\",\"220000\",\"420300\",\"530300\",\"210200\",\"350100\",\"440800\",\"445200\",\"350400\",\"440500\",\"220300\",\"410500\",\"350800\",\"420500\",\"421100\",\"610000\",\"610700\",\"411500\",\"460100\",\"640000\",\"210500\",\"441300\",\"450500\",\"460200\",\"640100\",\"653000\",\"532800\",\"542400\",\"610800\",\"420000\",\"632700\",\"410900\",\"441200\",\"211000\",\"220700\",\"411600\",\"431300\",\"530500\",\"530800\",\"621100\",\"211300\",\"440100\",\"411700\",\"431200\",\"433100\",\"451100\",\"530700\",\"231000\",\"430300\",\"610900\",\"611000\",\"530600\",\"654000\",\"230600\",\"420800\",\"422800\",\"450600\",\"451400\",\"540500\",\"450000\",\"540400\",\"650100\",\"210700\",\"210800\",\"410000\",\"445300\",\"650500\",\"450200\",\"540000\",\"620100\",\"430000\",\"542500\",\"620400\",\"640200\",\"350900\",\"420100\",\"460300\",\"652300\",\"350500\",\"410700\",\"620700\",\"621000\",\"652700\",\"222400\",\"230000\",\"653100\",\"430400\",\"430800\",\"440300\",\"441500\",\"530400\",\"610100\",\"230700\",\"430600\",\"469000\",\"430500\",\"530100\",\"632800\",\"411200\",\"450800\",\"460400\",\"533400\",\"420900\",\"621200\",\"640300\",\"650200\",\"230900\",\"350700\",\"440400\",\"442000\",\"450700\",\"532900\",\"220600\",\"440600\",\"a_3\",\"210900\",\"441400\",\"620300\",\"410600\",\"640400\",\"654200\",\"411100\",\"441900\",\"230500\",\"430100\",\"441700\",\"620500\",\"654300\",\"a_6\",\"a_7\",\"220100\",\"220500\",\"230100\",\"445100\",\"450100\",\"650000\",\"410100\",\"441600\",\"620900\",\"451300\",\"530000\",\"652900\",\"440900\",\"210000\",\"230300\",\"350000\",\"410400\",\"450900\",\"451000\",\"500100\",\"653200\",\"231100\",\"350200\",\"410200\",\"411300\",\"532600\",\"620200\",\"350300\",\"410300\",\"410800\",\"420700\",\"440700\",\"441800\",\"500000\",\"230400\",\"411400\",\"610200\",\"632500\",\"a_5\",\"430200\",\"430900\",\"610300\",\"610600\",\"630100\",\"632600\",\"411000\",\"419000\",\"450400\",\"533300\",\"620800\",\"659000\",\"450300\",\"500200\",\"540300\",\"220800\",\"610500\",\"622900\",\"230200\",\"421200\",\"460000\",\"610400\",\"540100\",\"630000\",\"632300\",\"652800\",\"623000\",\"210600\",\"431100\",\"440000\",\"530900\",\"231200\",\"a_4\",\"210100\",\"210300\",\"440200\",\"650400\"],\"fg\":\"\",\"fp\":\"\",\"eg\":\"\",\"ep\":\"\"}]","free_condition":"[{\"a\":[\"320700\",\"340400\",\"330300\",\"341700\",\"361000\",\"320900\",\"321000\",\"321100\",\"360700\",\"320000\",\"320400\",\"360200\",\"330800\",\"341500\",\"360000\",\"340100\",\"341000\",\"330900\",\"331100\",\"340000\",\"340300\",\"341300\",\"360500\",\"341800\",\"360400\",\"361100\",\"341200\",\"341600\",\"360900\",\"321300\",\"330100\",\"330200\",\"320500\",\"330700\",\"360100\",\"331000\",\"a_1\",\"320300\",\"320800\",\"321200\",\"360800\",\"320100\",\"330600\",\"340200\",\"310100\",\"320200\",\"340700\",\"360300\",\"341100\",\"330400\",\"340500\",\"340600\",\"340800\",\"310000\",\"320600\",\"330000\",\"330500\",\"360600\"],\"ft\":1,\"fp\":\"0.00\",\"fn\":0}]","postage_condition":1,"created_at":"2023-05-13 15:09:59","updated_at":"2023-05-13 00:00:00"}},"from_city_id":"0"},{"num":1,"sku":{"sku_id":"441542713672269928","sku_name":"尺码","thumbnail":"image|spu|21","thumbnail_url":"","lock_key":"","sku_att_relate_id":0,"image":"","video":"","user_hid":1,"shop_id":1,"sku_status":1,"weight":"1.25","price":"1.00","market_cost":"0.00","price_cost":"0.00","shop_sale_code":"","provide_channel":0,"provide_sale_code":"","sale_num":0,"sale_online_time":"2023-02-20 09:03:22","sale_over_time":null,"volume":"","flag_tester":1,"have_bind_spu":0,"created_at":"2023-02-20 09:03:22","updated_at":"2023-05-12 09:13:12"},"sku_relate":{"id":264,"shop_id":1,"product_id":"440958537740394381","category_id":0,"parent_id":0,"pk":"4","sku_name":"尺码","sku_id":"441542713672269928","price":"1.00","is_not_attr_name":2,"property_id":4,"spu_status":1,"sale_type":1,"down_payment":"0.00","final_payment":"0.00","freight_template":2,"sale_online_time":"2023-05-07 14:59:27","sale_over_time":null,"final_start_time":"2023-05-07 14:59:27","final_over_time":"2023-05-07 14:59:27","sales_tax_rate":"0.00","sales_tax_rate_value":"0.00","max_limit_num":500,"min_limit_num":1,"is_leaf":1,"sku_status":1,"have_gift":2,"created_at":"2023-05-07 14:59:27","updated_at":"2023-05-13 16:27:01"},"spu":{"product_id":"440958537740394381","title":"多项数据创新高带动中国经济回暖","user_hid":0,"thumbnail_url":"","brand_id":1,"video":"","video_url":"","shop_id":1,"status":1,"sub_title":"","min_price":"1.00","max_price":"1.00","min_price_cost":"0.00","max_price_cost":"","min_down_payment":"0.00","max_down_payment":"0.00","service_ids":"[1]","keywords":"123","sale_num":0,"freight_type":1,"freight_template":2,"total_stock":0,"category_id":11,"sale_type":2,"pull_on_time":"2000-01-01 00:00:00","pull_off_time":null,"sale_online_time":"2000-01-01 00:00:00","sale_over_time":null,"final_start_time":"2000-01-01 00:00:00","final_over_time":"2000-01-01 00:00:00","delivery_time":"2023-05-23 00:00:00","sale_count_show":0,"relate_type":0,"relate_item_id":"","relate_buy_count":0,"relate_buy_amount":"0.00","settle_type":2,"flag_tester":1},"shop":{"shop_id":1,"name":"测试店铺测试店铺测试店铺测测试店铺测试店铺测试店铺测测试店铺测试店铺测试店铺测测试店铺测试店铺测试店铺测","logo_url":"","bg_image_url":"","bg_image":"image|shop_logo|132","shop_type":2,"shop_entry_type":4,"status":1,"flag_tester":2,"admin_user_hid":1,"need_verify_status":2,"verify_status":"0","created_at":"2023-04-21 17:15:56","updated_at":"2023-04-21 17:16:12"},"ems_address_freight":null,"template_freight":{"freight_model":{"id":2,"shop_id":1,"title":"指定条件包邮(2)","province_id":110000,"city_id":110100,"area_id":110102,"free_freight":2,"pricing_mode":1,"have_use":2,"sale_area":"[{\"a\":[\"152200\",\"330700\",\"330800\",\"371100\",\"371300\",\"140900\",\"320000\",\"330900\",\"511300\",\"131100\",\"150500\",\"320200\",\"330200\",\"140100\",\"150200\",\"341200\",\"360700\",\"513400\",\"a_1\",\"510100\",\"130700\",\"140700\",\"150700\",\"310100\",\"340300\",\"340800\",\"371700\",\"510700\",\"511100\",\"512000\",\"130800\",\"310000\",\"320900\",\"130300\",\"321300\",\"370200\",\"511400\",\"120000\",\"150800\",\"152900\",\"360600\",\"361100\",\"370300\",\"510900\",\"131000\",\"330000\",\"360100\",\"361000\",\"370700\",\"371600\",\"130000\",\"150900\",\"340600\",\"371500\",\"140200\",\"141000\",\"320100\",\"370000\",\"513300\",\"130200\",\"360000\",\"370400\",\"370500\",\"510800\",\"511700\",\"110000\",\"130500\",\"130900\",\"140400\",\"150300\",\"330500\",\"340500\",\"120100\",\"320600\",\"321000\",\"510500\",\"511600\",\"320400\",\"320500\",\"360400\",\"511000\",\"511800\",\"130600\",\"140800\",\"341600\",\"360800\",\"320800\",\"330400\",\"370600\",\"511900\",\"130400\",\"140500\",\"330100\",\"341800\",\"130100\",\"140000\",\"341300\",\"510400\",\"150600\",\"340700\",\"370900\",\"371000\",\"510000\",\"a_2\",\"141100\",\"150100\",\"330600\",\"331000\",\"139000\",\"150000\",\"370100\",\"152500\",\"320300\",\"320700\",\"371400\",\"140300\",\"321100\",\"321200\",\"330300\",\"340000\",\"340100\",\"341700\",\"110100\",\"150400\",\"331100\",\"360900\",\"370800\",\"371200\",\"513200\",\"140600\",\"340400\",\"341000\",\"341500\",\"360300\",\"510300\",\"511500\",\"340200\",\"341100\",\"360200\",\"360500\",\"510600\"],\"fg\":\"0\",\"fp\":\"0.00\",\"eg\":\"0\",\"ep\":\"0.00\"},{\"a\":[\"220200\",\"420600\",\"421300\",\"430700\",\"431000\",\"451200\",\"532300\",\"632200\",\"211200\",\"429000\",\"630200\",\"211100\",\"532500\",\"640500\",\"210400\",\"232700\",\"421000\",\"420200\",\"540200\",\"220400\",\"230800\",\"350600\",\"533100\",\"620000\",\"620600\",\"211400\",\"220000\",\"420300\",\"530300\",\"210200\",\"350100\",\"440800\",\"445200\",\"350400\",\"440500\",\"220300\",\"410500\",\"350800\",\"420500\",\"421100\",\"610000\",\"610700\",\"411500\",\"460100\",\"640000\",\"210500\",\"441300\",\"450500\",\"460200\",\"640100\",\"653000\",\"532800\",\"542400\",\"610800\",\"420000\",\"632700\",\"410900\",\"441200\",\"211000\",\"220700\",\"411600\",\"431300\",\"530500\",\"530800\",\"621100\",\"211300\",\"440100\",\"411700\",\"431200\",\"433100\",\"451100\",\"530700\",\"231000\",\"430300\",\"610900\",\"611000\",\"530600\",\"654000\",\"230600\",\"420800\",\"422800\",\"450600\",\"451400\",\"540500\",\"450000\",\"540400\",\"650100\",\"210700\",\"210800\",\"410000\",\"445300\",\"650500\",\"450200\",\"540000\",\"620100\",\"430000\",\"542500\",\"620400\",\"640200\",\"350900\",\"420100\",\"460300\",\"652300\",\"350500\",\"410700\",\"620700\",\"621000\",\"652700\",\"222400\",\"230000\",\"653100\",\"430400\",\"430800\",\"440300\",\"441500\",\"530400\",\"610100\",\"230700\",\"430600\",\"469000\",\"430500\",\"530100\",\"632800\",\"411200\",\"450800\",\"460400\",\"533400\",\"420900\",\"621200\",\"640300\",\"650200\",\"230900\",\"350700\",\"440400\",\"442000\",\"450700\",\"532900\",\"220600\",\"440600\",\"a_3\",\"210900\",\"441400\",\"620300\",\"410600\",\"640400\",\"654200\",\"411100\",\"441900\",\"230500\",\"430100\",\"441700\",\"620500\",\"654300\",\"a_6\",\"a_7\",\"220100\",\"220500\",\"230100\",\"445100\",\"450100\",\"650000\",\"410100\",\"441600\",\"620900\",\"451300\",\"530000\",\"652900\",\"440900\",\"210000\",\"230300\",\"350000\",\"410400\",\"450900\",\"451000\",\"500100\",\"653200\",\"231100\",\"350200\",\"410200\",\"411300\",\"532600\",\"620200\",\"350300\",\"410300\",\"410800\",\"420700\",\"440700\",\"441800\",\"500000\",\"230400\",\"411400\",\"610200\",\"632500\",\"a_5\",\"430200\",\"430900\",\"610300\",\"610600\",\"630100\",\"632600\",\"411000\",\"419000\",\"450400\",\"533300\",\"620800\",\"659000\",\"450300\",\"500200\",\"540300\",\"220800\",\"610500\",\"622900\",\"230200\",\"421200\",\"460000\",\"610400\",\"540100\",\"630000\",\"632300\",\"652800\",\"623000\",\"210600\",\"431100\",\"440000\",\"530900\",\"231200\",\"a_4\",\"210100\",\"210300\",\"440200\",\"650400\"],\"fg\":\"\",\"fp\":\"\",\"eg\":\"\",\"ep\":\"\"}]","free_condition":"[{\"a\":[\"320700\",\"340400\",\"330300\",\"341700\",\"361000\",\"320900\",\"321000\",\"321100\",\"360700\",\"320000\",\"320400\",\"360200\",\"330800\",\"341500\",\"360000\",\"340100\",\"341000\",\"330900\",\"331100\",\"340000\",\"340300\",\"341300\",\"360500\",\"341800\",\"360400\",\"361100\",\"341200\",\"341600\",\"360900\",\"321300\",\"330100\",\"330200\",\"320500\",\"330700\",\"360100\",\"331000\",\"a_1\",\"320300\",\"320800\",\"321200\",\"360800\",\"320100\",\"330600\",\"340200\",\"310100\",\"320200\",\"340700\",\"360300\",\"341100\",\"330400\",\"340500\",\"340600\",\"340800\",\"310000\",\"320600\",\"330000\",\"330500\",\"360600\"],\"ft\":1,\"fp\":\"0.00\",\"fn\":0}]","postage_condition":1,"created_at":"2023-05-13 15:09:59","updated_at":"2023-05-13 00:00:00"}},"from_city_id":"0"},{"num":1,"sku":{"sku_id":"441543074433466472","sku_name":"32码","thumbnail":"image|spu|36","thumbnail_url":"","lock_key":"","sku_att_relate_id":0,"image":"","video":"","user_hid":1,"shop_id":1,"sku_status":1,"weight":"0.00","price":"52.00","market_cost":"0.00","price_cost":"30.00","shop_sale_code":"122","provide_channel":0,"provide_sale_code":"333","sale_num":0,"sale_online_time":"2023-02-20 09:10:47","sale_over_time":null,"volume":"","flag_tester":1,"have_bind_spu":0,"created_at":"2023-02-20 09:10:47","updated_at":"2023-04-05 10:04:10"},"sku_relate":{"id":273,"shop_id":1,"product_id":"440958537740263309","category_id":0,"parent_id":269,"pk":"2_4","sku_name":"32码","sku_id":"441543074433466472","price":"52.00","is_not_attr_name":2,"property_id":4,"spu_status":1,"sale_type":1,"down_payment":"0.00","final_payment":"0.00","freight_template":2,"sale_online_time":"2023-05-08 17:30:56","sale_over_time":null,"final_start_time":"2023-05-08 17:30:56","final_over_time":"2023-05-08 17:30:56","sales_tax_rate":"0.00","sales_tax_rate_value":"0.00","max_limit_num":500,"min_limit_num":1,"is_leaf":1,"sku_status":1,"have_gift":2,"created_at":"2023-05-08 17:30:56","updated_at":"2023-05-13 16:27:24"},"spu":{"product_id":"440958537740263309","title":"比亚迪回应长沙工厂“排队辞职”","user_hid":0,"thumbnail_url":"","brand_id":1,"video":"","video_url":"","shop_id":1,"status":1,"sub_title":"","min_price":"1.00","max_price":"52.00","min_price_cost":"0.00","max_price_cost":"","min_down_payment":"0.00","max_down_payment":"0.00","service_ids":"[1]","keywords":"","sale_num":0,"freight_type":1,"freight_template":2,"total_stock":0,"category_id":2,"sale_type":1,"pull_on_time":"2023-05-13 16:27:24","pull_off_time":null,"sale_online_time":"2023-05-13 16:27:24","sale_over_time":null,"final_start_time":"2000-01-01 00:00:00","final_over_time":"2000-01-01 00:00:00","delivery_time":null,"sale_count_show":0,"relate_type":0,"relate_item_id":"","relate_buy_count":0,"relate_buy_amount":"0.00","settle_type":2,"flag_tester":1},"shop":{"shop_id":1,"name":"测试店铺测试店铺测试店铺测测试店铺测试店铺测试店铺测测试店铺测试店铺测试店铺测测试店铺测试店铺测试店铺测","logo_url":"","bg_image_url":"","bg_image":"image|shop_logo|132","shop_type":2,"shop_entry_type":4,"status":1,"flag_tester":2,"admin_user_hid":1,"need_verify_status":2,"verify_status":"0","created_at":"2023-04-21 17:15:56","updated_at":"2023-04-21 17:16:12"},"ems_address_freight":null,"template_freight":{"freight_model":{"id":2,"shop_id":1,"title":"指定条件包邮(2)","province_id":110000,"city_id":110100,"area_id":110102,"free_freight":2,"pricing_mode":1,"have_use":2,"sale_area":"[{\"a\":[\"152200\",\"330700\",\"330800\",\"371100\",\"371300\",\"140900\",\"320000\",\"330900\",\"511300\",\"131100\",\"150500\",\"320200\",\"330200\",\"140100\",\"150200\",\"341200\",\"360700\",\"513400\",\"a_1\",\"510100\",\"130700\",\"140700\",\"150700\",\"310100\",\"340300\",\"340800\",\"371700\",\"510700\",\"511100\",\"512000\",\"130800\",\"310000\",\"320900\",\"130300\",\"321300\",\"370200\",\"511400\",\"120000\",\"150800\",\"152900\",\"360600\",\"361100\",\"370300\",\"510900\",\"131000\",\"330000\",\"360100\",\"361000\",\"370700\",\"371600\",\"130000\",\"150900\",\"340600\",\"371500\",\"140200\",\"141000\",\"320100\",\"370000\",\"513300\",\"130200\",\"360000\",\"370400\",\"370500\",\"510800\",\"511700\",\"110000\",\"130500\",\"130900\",\"140400\",\"150300\",\"330500\",\"340500\",\"120100\",\"320600\",\"321000\",\"510500\",\"511600\",\"320400\",\"320500\",\"360400\",\"511000\",\"511800\",\"130600\",\"140800\",\"341600\",\"360800\",\"320800\",\"330400\",\"370600\",\"511900\",\"130400\",\"140500\",\"330100\",\"341800\",\"130100\",\"140000\",\"341300\",\"510400\",\"150600\",\"340700\",\"370900\",\"371000\",\"510000\",\"a_2\",\"141100\",\"150100\",\"330600\",\"331000\",\"139000\",\"150000\",\"370100\",\"152500\",\"320300\",\"320700\",\"371400\",\"140300\",\"321100\",\"321200\",\"330300\",\"340000\",\"340100\",\"341700\",\"110100\",\"150400\",\"331100\",\"360900\",\"370800\",\"371200\",\"513200\",\"140600\",\"340400\",\"341000\",\"341500\",\"360300\",\"510300\",\"511500\",\"340200\",\"341100\",\"360200\",\"360500\",\"510600\"],\"fg\":\"0\",\"fp\":\"0.00\",\"eg\":\"0\",\"ep\":\"0.00\"},{\"a\":[\"220200\",\"420600\",\"421300\",\"430700\",\"431000\",\"451200\",\"532300\",\"632200\",\"211200\",\"429000\",\"630200\",\"211100\",\"532500\",\"640500\",\"210400\",\"232700\",\"421000\",\"420200\",\"540200\",\"220400\",\"230800\",\"350600\",\"533100\",\"620000\",\"620600\",\"211400\",\"220000\",\"420300\",\"530300\",\"210200\",\"350100\",\"440800\",\"445200\",\"350400\",\"440500\",\"220300\",\"410500\",\"350800\",\"420500\",\"421100\",\"610000\",\"610700\",\"411500\",\"460100\",\"640000\",\"210500\",\"441300\",\"450500\",\"460200\",\"640100\",\"653000\",\"532800\",\"542400\",\"610800\",\"420000\",\"632700\",\"410900\",\"441200\",\"211000\",\"220700\",\"411600\",\"431300\",\"530500\",\"530800\",\"621100\",\"211300\",\"440100\",\"411700\",\"431200\",\"433100\",\"451100\",\"530700\",\"231000\",\"430300\",\"610900\",\"611000\",\"530600\",\"654000\",\"230600\",\"420800\",\"422800\",\"450600\",\"451400\",\"540500\",\"450000\",\"540400\",\"650100\",\"210700\",\"210800\",\"410000\",\"445300\",\"650500\",\"450200\",\"540000\",\"620100\",\"430000\",\"542500\",\"620400\",\"640200\",\"350900\",\"420100\",\"460300\",\"652300\",\"350500\",\"410700\",\"620700\",\"621000\",\"652700\",\"222400\",\"230000\",\"653100\",\"430400\",\"430800\",\"440300\",\"441500\",\"530400\",\"610100\",\"230700\",\"430600\",\"469000\",\"430500\",\"530100\",\"632800\",\"411200\",\"450800\",\"460400\",\"533400\",\"420900\",\"621200\",\"640300\",\"650200\",\"230900\",\"350700\",\"440400\",\"442000\",\"450700\",\"532900\",\"220600\",\"440600\",\"a_3\",\"210900\",\"441400\",\"620300\",\"410600\",\"640400\",\"654200\",\"411100\",\"441900\",\"230500\",\"430100\",\"441700\",\"620500\",\"654300\",\"a_6\",\"a_7\",\"220100\",\"220500\",\"230100\",\"445100\",\"450100\",\"650000\",\"410100\",\"441600\",\"620900\",\"451300\",\"530000\",\"652900\",\"440900\",\"210000\",\"230300\",\"350000\",\"410400\",\"450900\",\"451000\",\"500100\",\"653200\",\"231100\",\"350200\",\"410200\",\"411300\",\"532600\",\"620200\",\"350300\",\"410300\",\"410800\",\"420700\",\"440700\",\"441800\",\"500000\",\"230400\",\"411400\",\"610200\",\"632500\",\"a_5\",\"430200\",\"430900\",\"610300\",\"610600\",\"630100\",\"632600\",\"411000\",\"419000\",\"450400\",\"533300\",\"620800\",\"659000\",\"450300\",\"500200\",\"540300\",\"220800\",\"610500\",\"622900\",\"230200\",\"421200\",\"460000\",\"610400\",\"540100\",\"630000\",\"632300\",\"652800\",\"623000\",\"210600\",\"431100\",\"440000\",\"530900\",\"231200\",\"a_4\",\"210100\",\"210300\",\"440200\",\"650400\"],\"fg\":\"\",\"fp\":\"\",\"eg\":\"\",\"ep\":\"\"}]","free_condition":"[{\"a\":[\"320700\",\"340400\",\"330300\",\"341700\",\"361000\",\"320900\",\"321000\",\"321100\",\"360700\",\"320000\",\"320400\",\"360200\",\"330800\",\"341500\",\"360000\",\"340100\",\"341000\",\"330900\",\"331100\",\"340000\",\"340300\",\"341300\",\"360500\",\"341800\",\"360400\",\"361100\",\"341200\",\"341600\",\"360900\",\"321300\",\"330100\",\"330200\",\"320500\",\"330700\",\"360100\",\"331000\",\"a_1\",\"320300\",\"320800\",\"321200\",\"360800\",\"320100\",\"330600\",\"340200\",\"310100\",\"320200\",\"340700\",\"360300\",\"341100\",\"330400\",\"340500\",\"340600\",\"340800\",\"310000\",\"320600\",\"330000\",\"330500\",\"360600\"],\"ft\":1,\"fp\":\"0.00\",\"fn\":0}]","postage_condition":1,"created_at":"2023-05-13 15:09:59","updated_at":"2023-05-13 00:00:00"}},"from_city_id":"0"}]
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
			var bt []byte
			if bt,err=json.Marshal(tt.want);err!=nil{
				t.Errorf("Marshal() = %v", err.Error())
				return
			}
			t.Logf("%v", string(bt))
		})
	}
}
