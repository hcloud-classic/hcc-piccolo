package init

import "hcc/piccolo/lib/config"

// MainInit : Main initialization function
func MainInit() error {
	err := syscheckInit()
	if err != nil {
		return err
	}

	err = loggerInit()
	if err != nil {
		return err
	}

	config.Parser()

	return nil
}
