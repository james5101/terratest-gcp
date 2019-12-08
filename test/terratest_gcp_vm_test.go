package test

import (
	"fmt"
	"os"
	"testing"
	
	"github.com/gruntwork-io/terratest/modules/gcp"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformGcpExample(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// expectedIp := terraform.Output(t, terraformOptions,"ip")
	expectedName := terraform.Output(t, terraformOptions, "name")

	assert.Equal(t, expectedName, "jamesterratest")

	SetGcpAuthEnvVars("/home/jnoonan/go/src/terratest-gcp/james-app-dev-b2cd3e47b132.json")

	projectid := "james-app-dev"

	actaulInstance := gcp.FetchInstance(t, projectid, expectedName)
	fmt.Println(actaulInstance.MachineType)
	

}

func SetGcpAuthEnvVars(path string)  {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS",path)
}

