package docker

import "fmt"

// CreateDockerExecCommand creates "docker exec" command with environment variables.
// This method returns below command string.
// "docker exec <containerName> sh -c 'export <key1>=<value1> && export <key2>=<value2> && ... && <mainProcessCommand>'"
func CreateDockerExecCommand(containerName string, envKeyValueMap map[string]string, cmdExecMainProcess string) string {
	var cmdSetEnv string
	for key, value := range envKeyValueMap {
		cmdSetEnv += fmt.Sprintf("export %s=%s && ", key, value)
	}
	return fmt.Sprintf("docker exec %s sh -c '%s %s'", containerName, cmdSetEnv, cmdExecMainProcess)
}
