package githubstats

import (
	"testing"
)

func TestParseCommitsURL(t *testing.T) {

	expected := "https://api.github.com/repos/donvito/learngo/commits"
	result := parseCommitsURL("https://api.github.com/repos/donvito/learngo/commits{/sha}")

	if expected != result {
		println("Expected %s, result is %s", expected, result)
		t.Error("Result not as expected", result)
	}
}
