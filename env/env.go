// Package env provides a library for determining information about the processes' environment.
package env

import (
	"os"
)

const (
	// BRANCH is the code's current git branch.
	BRANCH string = "BRANCH"

	// DEV is the dev branch name
	DEV string = "dev"

	// MAIN is the prod branch name
	MAIN string = "main"

	// PODNAME is the name of the pod currently executing the process.
	PODNAME string = "POD_NAME"

	// KubernetesServiceHostConst is the kubernetes service host of the current
	// process.
	KubernetesServiceHostConst string = "KUBERNETES_SERVICE_HOST"
)

func init() {
	currentBranch = os.Getenv(BRANCH)
	podName = os.Getenv(PODNAME)
	kubernetesServiceHost = os.Getenv(KubernetesServiceHostConst)
}

// currentBranch variable holds pod's name in memory to avoid OS interaction.
var currentBranch string

// IsDev checks if the current code is part of the dev branch.
func IsDev() bool {
	return IsBranch(DEV)
}

// IsMain check is the current code is part of the master branch.
func IsMain() bool {
	return IsBranch(MAIN)
}

// IsBranch checks if the current code is part of the specified branch (name).
func IsBranch(name string) bool {
	return name == currentBranch
}

// Branch retrieves the current code branch for the process.
// if the branch is not set "" is returned.
func Branch() string {
	return currentBranch
}

// podName variable holds pod's name in memory to avoid OS interaction.
var podName string

// PodName retrieves the current pod's name from the environment.
// If the pod name is not set "" is returned.
func PodName() string {
	return podName
}

// kubernetesServiceHost specifies the kubernetes service host of this
// process' pod. This value is set in a kubernetes cluster by default.
var kubernetesServiceHost string

// KubernetesServiceHost retrieves the process' pod's kubernetes service
// host.
func KubernetesServiceHost() string {
	return kubernetesServiceHost
}

// InKubeCluster specifies the whether the current process is being executed
// inside a kubernetes cluster.
func InKubeCluster() bool {
	return kubernetesServiceHost != ""
}
