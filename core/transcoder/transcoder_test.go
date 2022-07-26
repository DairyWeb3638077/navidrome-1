package transcoder

import (
	"testing"

	"github.com/navidrome/navidrome/log"
	"github.com/navidrome/navidrome/tests"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTranscoder(t *testing.T) {
	tests.Init(t, false)
	log.SetLevel(log.LevelCritical)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Transcoder Suite")
}

var _ = Describe("createTranscodeCommand", func() {
	It("creates a valid command line", func() {
		args := createTranscodeCommand("ffmpeg -i %s -b:a %bk mp3 -", "/music library/file.mp3", 123)
		Expect(args).To(Equal([]string{"ffmpeg", "-i", "/music library/file.mp3", "-b:a", "123k", "mp3", "-"}))
	})
})
