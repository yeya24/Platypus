package update

import (
	"fmt"
	"os"

	"github.com/WangYihang/Platypus/lib/util/log"
	"github.com/WangYihang/Platypus/lib/util/ui"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const version = "1.2.9"

func ConfirmAndSelfUpdate() {
	log.Info("Detecting the latest version...")
	latest, found, err := selfupdate.DetectLatest("wangyihang/Platypus")
	if err != nil {
		log.Error("Error occurred while detecting version:", err)
		return
	}

	v := semver.MustParse(version)
	if !found || latest.Version.LTE(v) {
		log.Success("Current version is the latest")
		return
	}

	if !ui.PromptYesNo(fmt.Sprintf("Do you want to update to %s?", latest.Version)) {
		return
	}

	exe, err := os.Executable()
	if err != nil {
		log.Error("Could not locate executable path")
		return
	}
	log.Info("Downloading from %s", latest.AssetURL)
	if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
		log.Error("Error occurred while updating binary:", err)
		return
	}
	log.Success("Successfully updated to version", latest.Version)
}
