/*
Copyright © 2021 Dell Inc. or its subsidiaries. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package components

// Scheduler encapsulates information to deploy CSI scheduler
type Scheduler struct {
	Enable             bool     `json:"enable"`
	ServiceAccount     string   `json:"serviceAccount"`
	Image              *Image   `json:"image,omitempty"`
	Log                *Log     `json:"log,omitempty"`
	Metrics            *Metrics `json:"metrics,omitempty"`
	Patcher            *Patcher `json:"patcher,omitempty"`
	ExtenderPort       string   `json:"extenderPort,omitempty"`
	StorageProvisioner string   `json:"storageProvisioner"`
	// +nullable
	// +optional
	Resources         *ResourceRequirements `json:"resources,omitempty"`
	SecurityContext   *SecurityContext      `json:"securityContext,omitempty"`
	PodSecurityPolicy *PodSecurityPolicy    `json:"podSecurityPolicy,omitempty"`
}
