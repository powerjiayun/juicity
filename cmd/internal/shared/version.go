package shared

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/juicity/juicity/config"
)

func multiline(parts ...string) string {
	return strings.Join(parts, "\n")
}

func PrintVersion(cgoEnabled int) {
	fmt.Print(multiline(
		fmt.Sprintf("juicity-client version %v", config.Version),
		fmt.Sprintf("go version %v %v/%v", runtime.Version(), runtime.GOOS, runtime.GOARCH),
		fmt.Sprintf("CGO_ENABLED: %v\n", cgoEnabled),
		"Copyright (c) 2023 juicity",
		"License GNU AGPLv3 <https://github.com/juicity/juicity/blob/main/LICENSE>",
	))
}

func GetVersion(cgoEnabled int) string {
	return multiline(
		fmt.Sprintf("juicity-client version %v", config.Version),
		fmt.Sprintf("go version %v %v/%v", runtime.Version(), runtime.GOOS, runtime.GOARCH),
		fmt.Sprintf("CGO_ENABLED: %v", cgoEnabled),
		"Copyright (c) 2023 juicity",
		"License GNU AGPLv3 <https://github.com/juicity/juicity/blob/main/LICENSE>",
	)
}
