package metadata

import (
	"github.com/codeskyblue/go-sh"
)

func GetRevision(repoPath string) []byte {
	output, _ := sh.Command("git", "rev-parse", "HEAD", sh.Dir(repoPath)).Output()
	return output
}
