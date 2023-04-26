package bridge

import (
	"design-patterns/structural/bridge/other_2"
	"testing"
)

func TestChromeOs(t *testing.T) {
	chromeForUbuntu := other_2.NewChromeBeta(&other_2.Ubuntu{})
	chromeForUbuntu.OpenURL()

	chromeForWindows := other_2.NewChromeBeta(&other_2.Windows{})
	chromeForWindows.OpenURL()
}
