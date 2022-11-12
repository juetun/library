package upload_operate

import (
	"reflect"
	"testing"
)

func TestParseTextEditorContent(t *testing.T) {
	type argStruct struct {
		textContent string
	}
	tests := []struct {
		name              string
		args              argStruct
		wantKeysDescImg   map[string]string
		wantKeysDescVideo map[string]string
		wantErr           bool
	}{
		{
			args: argStruct{textContent: `<p><br></p><div data-w-e-type="video" data-w-e-is-void>
<video poster="video|spu|17" controls="true" width="auto" height="auto"><source src="https://juetun-test.oss-cn-beijing.aliyuncs.com/spu/202211/12/221112210004_my9JDb5xjlzYkEQD7Bq1r0L46W8adP.mp4" type="video/mp4"/></video>
</div><p><img src="https://juetun-test.oss-cn-beijing.aliyuncs.com/media/spu_src/202211/12/221112205954_joRMeLawEv67Y394orNqB9zn1Xyxpk_1422x800.jpeg" alt="image|spu|87" data-href="image|spu|87" style=""/></p><video poster="video|spu|17" controls="true" width="auto" height="auto"><source src="https://juetun-test.oss-cn-beijing.aliyuncs.com/spu/202211/12/221112210004_my9JDb5xjlzYkEQD7Bq1r0L46W8adP.mp2224" type="video/mp4"/></video>`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKeysDescImg, gotKeysDescVideo, err := ParseTextEditorContent(tt.args.textContent)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTextEditorContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKeysDescImg, tt.wantKeysDescImg) {
				t.Errorf("ParseTextEditorContent() gotKeysDescImg = %v, want %v", gotKeysDescImg, tt.wantKeysDescImg)
			}
			if !reflect.DeepEqual(gotKeysDescVideo, tt.wantKeysDescVideo) {
				t.Errorf("ParseTextEditorContent() gotKeysDescVideo = %v, want %v", gotKeysDescVideo, tt.wantKeysDescVideo)
			}
		})
	}
}
