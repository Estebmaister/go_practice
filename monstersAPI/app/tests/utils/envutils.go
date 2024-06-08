package utilstests

import (
	"fmt"
	"os"
	"regexp"

	"battle-of-monsters/app/config"
)

// this should match the project root folder name.
const rootFolder = "assessment-cc-go-sr-01"

// LoadEnv load the env files relative to the root folder from any nested test directory.
func LoadEnv() {
	re := regexp.MustCompile(fmt.Sprintf("^(.*%s)", rootFolder))
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	errMsg := "update the `testutils.rootFolder` variable to match the name of the project root folder"

	if rootPath == nil {
		panic(errMsg)
	}

	if err := os.Chdir(string(rootPath)); err != nil {
		panic(errMsg)
	}

	os.Setenv("GO_ENVIRONMENT", "TEST")
	config.Load()
}
