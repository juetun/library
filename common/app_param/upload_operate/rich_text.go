package upload_operate

import (
	"regexp"
)

func parseVideo(textContent string) (keysDescVideo map[string]string) {
	keysDescVideo = map[string]string{}
	compileRegex := regexp.MustCompile(`<video poster="([^"]*)".*</video>`)
	matchArr := compileRegex.FindAllStringSubmatch(textContent, -1)
	keysDescVideo = make(map[string]string, len(matchArr))
	for _, item := range matchArr {
		if item[0] == "" {
			continue
		}
		keysDescVideo[item[1]] = item[0]
	}
	return
}

func parseImg(textContent string) (keysDescImg map[string]string) {
	keysDescImg = map[string]string{}
	compileRegex := regexp.MustCompile(`<img src="[^"]*" alt="[^"]*" data-href="([^"]*)"[^\/]*/>`)
	matchArr := compileRegex.FindAllStringSubmatch(textContent, -1)
	keysDescImg = make(map[string]string, len(matchArr))
	for _, item := range matchArr {
		if item[0] == "" {
			continue
		}
		keysDescImg[item[1]] = item[0]
	}
	return
}

func ParseTextEditorContent(textContent string) (keysDescImg, keysDescVideo map[string]string, err error) {

	keysDescVideo = parseVideo(textContent)
	keysDescImg = parseImg(textContent)

	return
}
