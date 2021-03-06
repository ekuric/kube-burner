// Copyright 2020 The Kube-burner Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

// Spec configuration root
type Spec struct {
	// GlobalConfig defines global configuration parameters
	GlobalConfig GlobalConfig `yaml:"global,omitempty"`
	// Jobs list of kube-burner jobs
	Jobs []Job `yaml:"jobs"`
}

// IndexerConfig holds the indexer configuration
type IndexerConfig struct {
	// Enabled enable indexer
	Enabled bool `yaml:"enabled"`
	// Type type of indexer
	Type string `yaml:"type"`
	// ESServers List of ElasticSearch instances
	ESServers []string `yaml:"esServers"`
	// DefaultIndex default index to send prometheus metrics
	DefaultIndex string `yaml:"defaultIndex"`
	Port         int    `yaml:"port"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	// InsecureSkipVerify disable TLS ceriticate verification
	InsecureSkipVerify bool `yaml:"insecureSkipVerify"`
}

// Measurement holds the measurement configuration
type Measurement struct {
	// Name is the name the measurement
	Name string `yaml:"name"`
	// ESIndex contains the ElasticSearch index used to
	// index the metrics
	ESIndex string `yaml:"esIndex"`
}

// GlobalConfig holds the global configuration
type GlobalConfig struct {
	// Kubeconfig path to a valid kubeconfig file
	Kubeconfig string `yaml:"kubeconfig"`
	// IndexerConfig contains a IndexerConfig definition
	IndexerConfig IndexerConfig `yaml:"indexerConfig"`
	// Write prometheus metrics to a file
	WriteToFile bool `yaml:"writeToFile"`
	// Directory to save metrics files in
	MetricsDirectory string `yaml:"metricsDirectory"`
	// Measurements describes a list of measurements kube-burner
	// will take along with job
	Measurements []Measurement `yaml:"measurements"`
}

// Object defines an object that kube-burner will create
type Object struct {
	// ObjectTemplate path to a valid YAML definition of a k8s resource
	ObjectTemplate string `yaml:"objectTemplate"`
	// Replicas number of replicas to create of the given object
	Replicas int `yaml:"replicas"`
	// InputVars contains a map of arbitrary input variables
	// that can be introduced by the users
	InputVars map[string]string `yaml:"inputVars"`
}

// Job defines a kube-burner job
type Job struct {
	// IterationCount how many times to execute the job
	JobIterations int `yaml:"jobIterations"`
	// IterationDelay how many milliseconds to wait between each job iteration
	JobIterationDelay int `yaml:"jobIterationDelay"`
	// IterationPause how many milliseconds to pause after finishing the job
	JobPause int `yaml:"jobPause"`
	// PodWait wait for all pods to be running before moving forward to the next iteration
	PodWait bool `yaml:"podWait"`
	// Cleanup clean up old namespaces
	Cleanup bool `yaml:"cleanup"`
	// Namespace namespace base name to use
	Namespace string `yaml:"namespace"`
	// whether to create namespaces or not with each iteration
	Namespaced bool `yaml:"namespaced"`
	// NamespacedIterations create a namespace per iteration
	NamespacedIterations bool `yaml:"namespacedIterations"`
	// Name name of the iteration
	Name string `yaml:"name"`
	// Objects list of objects
	Objects []Object `yaml:"objects"`
	// Max number of queries per second
	QPS int `yaml:"qps"`
	// Maximum burst for throttle
	Burst int `yaml:"burst"`
}
