package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	NodeConditionTypeReady            = "Ready"
	NodeConditionTypeMountPropagation = "MountPropagation"
	NodeConditionTypeSchedulable      = "Schedulable"
)

const (
	NodeConditionReasonManagerPodDown            = "ManagerPodDown"
	NodeConditionReasonManagerPodMissing         = "ManagerPodMissing"
	NodeConditionReasonKubernetesNodeGone        = "KubernetesNodeGone"
	NodeConditionReasonKubernetesNodeNotReady    = "KubernetesNodeNotReady"
	NodeConditionReasonKubernetesNodePressure    = "KubernetesNodePressure"
	NodeConditionReasonUnknownNodeConditionTrue  = "UnknownNodeConditionTrue"
	NodeConditionReasonNoMountPropagationSupport = "NoMountPropagationSupport"
	NodeConditionReasonKubernetesNodeCordoned    = "KubernetesNodeCordoned"
)

const (
	DiskConditionTypeSchedulable = "Schedulable"
	DiskConditionTypeReady       = "Ready"
)

const (
	DiskConditionReasonDiskPressure          = "DiskPressure"
	DiskConditionReasonDiskFilesystemChanged = "DiskFilesystemChanged"
	DiskConditionReasonNoDiskInfo            = "NoDiskInfo"
	DiskConditionReasonDiskNotReady          = "DiskNotReady"
)

type DiskSpec struct {
	// +optional
	Path string `json:"path"`
	// +optional
	AllowScheduling bool `json:"allowScheduling"`
	// +optional
	EvictionRequested bool `json:"evictionRequested"`
	// +optional
	StorageReserved int64 `json:"storageReserved"`
	// +optional
	Tags []string `json:"tags"`
}

type DiskStatus struct {
	// +optional
	Conditions map[string]Condition `json:"conditions"`
	// +optional
	StorageAvailable int64 `json:"storageAvailable"`
	// +optional
	StorageScheduled int64 `json:"storageScheduled"`
	// +optional
	StorageMaximum int64 `json:"storageMaximum"`
	// +optional
	ScheduledReplica map[string]int64 `json:"scheduledReplica"`
	// +optional
	DiskUUID string `json:"diskUUID"`
}

type NodeSpec struct {
	// +optional
	Name string `json:"name"`
	// +optional
	Disks map[string]DiskSpec `json:"disks"`
	// +optional
	AllowScheduling bool `json:"allowScheduling"`
	// +optional
	EvictionRequested bool `json:"evictionRequested"`
	// +optional
	Tags []string `json:"tags"`
	// +optional
	EngineManagerCPURequest int `json:"engineManagerCPURequest"`
	// +optional
	ReplicaManagerCPURequest int `json:"replicaManagerCPURequest"`
}

type NodeStatus struct {
	// +optional
	Conditions map[string]Condition `json:"conditions"`
	// +optional
	DiskStatus map[string]*DiskStatus `json:"diskStatus"`
	// +optional
	Region string `json:"region"`
	// +optional
	Zone string `json:"zone"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Node struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeSpec   `json:"spec,omitempty"`
	Status NodeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Node `json:"items"`
}
