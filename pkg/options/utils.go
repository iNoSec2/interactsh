package options

import (
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/utils/auth/pdcp"
	updateutils "github.com/projectdiscovery/utils/update"
)

const Version = "1.2.4"

var banner = (`
    _       __                       __       __  
   (_)___  / /____  _________ ______/ /______/ /_ 
  / / __ \/ __/ _ \/ ___/ __ '/ ___/ __/ ___/ __ \
 / / / / / /_/  __/ /  / /_/ / /__/ /_(__  ) / / /
/_/_/ /_/\__/\___/_/   \__,_/\___/\__/____/_/ /_/
`)

func ShowBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\tprojectdiscovery.io\n\n")
}

// GetUpdateCallback returns a callback function that updates interactsh
func GetUpdateCallback(assetName string) func() {
	return func() {
		ShowBanner()
		updateutils.GetUpdateToolFromRepoCallback(assetName, Version, "interactsh")()
	}
}

// AuthWithPDCP is used to authenticate with PDCP
func AuthWithPDCP() {
	ShowBanner()
	pdcp.CheckNValidateCredentials("interactsh")
}
