package dot

import "fmt"

func (self Environment) Provision() (errs []error) {
	pm := MarshalOS(self.OS).PackageManager()
	for _, profile := range self.Profiles {
		// Profile; 3 layers of custom provisioning
		// Layer 1: Individual Package Install/Remove
		err := pm.SudoInstallPackages(profile.Packages.Install)
		if err != nil {
			errs = append(errs, err)
		}
		err = pm.SudoRemovePackages(profile.Packages.Remove)
		if err != nil {
			errs = append(errs, err)
		}
		// Layer 2: Individual Config Action From/To
		configFileErrs := profile.CopyOrLinkConfigFiles()
		if len(configFileErrs) > 0 {
			fmt.Println("[error] config file copy or link error")
			errs = append(errs, configFileErrs...)
		}
		// Layer 3: Post Install Commands
		postInstallCommandErrs := profile.ExecutePostInstallCommands()
		if len(postInstallCommandErrs) > 0 {
			fmt.Println("[error] post install command error")
			errs = append(errs, postInstallCommandErrs...)
		}
	}
	return errs
}
