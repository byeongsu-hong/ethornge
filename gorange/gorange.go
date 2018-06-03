package gorange

import "../utils"

func Launch(gnOpt *GenesisOption, gethOpt *GethOption) (err error) {
	if err = utils.CreateDir(gethOpt.DataDir); err != nil {
		return
	}

	if err = gnOpt.Accounts.Export(gethOpt.AccountDir); err != nil {
		return
	}

	return gnOpt.init(gethOpt)
}
