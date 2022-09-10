package edge

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestEdgeApi(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	ssml := `<speak xmlns="http://www.w3.org/2001/10/synthesis" x111mlns:mstts="http://www.w3.org/2001/mstts" xmlns:emo="http://www.w3.org/2009/10/emotionml" version="1.0" xml:lang="en-US"><voice name="zh-CN-XiaoxiaoNeural"><prosody rate="200%" pitch="+0Hz">　　半年后一天，苏浩再次尝试控制身上的血气运动，原本以为会一如既往般毫无动静，没想到意识操控的那部分血气竟然往控制方向移动了一丝。就是这一丝移动，让苏浩欣喜若狂。</prosody></voice></speak>`
	_, err := GetAudio(ssml, "webm-24khz-16bit-mono-opus")
	if err != nil {
		log.Fatal(err)
		return
	}
}