package corp

import "go-interface/tools"

func corpName(sliName string) string {
	name := sliName + tools.RandInt()
	return name
}
