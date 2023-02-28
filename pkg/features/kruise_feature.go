package features

import (
	"fmt"
	"github.com/ChenWoChong/ones/pkg/util/feature"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/component-base/featuregate"
	"os"
	"strings"
)

const (
	// KruiseDaemon enables the features relied on kruise-daemon, such as image pulling and container restarting.
	KruiseDaemon featuregate.Feature = "KruiseDaemon"

	// PodWebhook enables webhook for Pods creations. This is also related to SidecarSet.
	PodWebhook featuregate.Feature = "PodWebhook"

	// CloneSetShortHash enables CloneSet controller only set revision hash name to pod label.
	CloneSetShortHash featuregate.Feature = "CloneSetShortHash"

	// KruisePodReadinessGate enables Kruise webhook to inject 'KruisePodReady' readiness-gate to
	// all Pods during creation.
	// Otherwise, it will only be injected to Pods created by Kruise workloads.
	KruisePodReadinessGate featuregate.Feature = "KruisePodReadinessGate"

	// PreDownloadImageForInPlaceUpdate enables cloneset/statefulset controllers to create ImagePullJobs to
	// pre-download images for in-place update.
	PreDownloadImageForInPlaceUpdate featuregate.Feature = "PreDownloadImageForInPlaceUpdate"

	// CloneSetPartitionRollback enables CloneSet controller to rollback Pods to currentRevision
	// when number of updateRevision pods is bigger than (replicas - partition).
	CloneSetPartitionRollback featuregate.Feature = "CloneSetPartitionRollback"

	// ResourcesDeletionProtection enables protection for resources deletion, currently supports
	// Namespace, CustomResourcesDefinition, Deployment, StatefulSet, ReplicaSet, CloneSet, Advanced StatefulSet, UnitedDeployment.
	// It is only supported for Kubernetes version >= 1.16
	// Note that if it is enabled during Kruise installation or upgrade, Kruise will require more authorities:
	// 1. Webhook for deletion operation of namespace, crd, deployment, statefulset, replicaset and workloads in Kruise.
	// 2. ClusterRole for reading all resource types, because CRD validation needs to list the CRs of this CRD.
	ResourcesDeletionProtection featuregate.Feature = "ResourcesDeletionProtection"

	// PodUnavailableBudgetDeleteGate enables PUB capability to protect pod from deletion and eviction
	PodUnavailableBudgetDeleteGate featuregate.Feature = "PodUnavailableBudgetDeleteGate"

	// PodUnavailableBudgetUpdateGate enables PUB capability to protect pod from in-place update
	PodUnavailableBudgetUpdateGate featuregate.Feature = "PodUnavailableBudgetUpdateGate"

	// WorkloadSpread enable WorkloadSpread to constrain the spread of the workload.
	WorkloadSpread featuregate.Feature = "WorkloadSpread"

	// DaemonWatchingPod enables kruise-daemon to list watch pods that belong to the same node.
	DaemonWatchingPod featuregate.Feature = "DaemonWatchingPod"

	// TemplateNoDefaults to control whether webhook should inject pod default fields into pod template
	// and pvc default fields into pvc template.
	// If TemplateNoDefaults is false, webhook should inject default fields only when the template changed.
	TemplateNoDefaults featuregate.Feature = "TemplateNoDefaults"

	// InPlaceUpdateEnvFromMetadata enables Kruise to in-place update a container in Pod
	// when its env from labels/annotations changed and pod is in-place updating.
	InPlaceUpdateEnvFromMetadata featuregate.Feature = "InPlaceUpdateEnvFromMetadata"

	// Enables policies controlling deletion of PVCs created by a StatefulSet.
	StatefulSetAutoDeletePVC featuregate.Feature = "StatefulSetAutoDeletePVC"

	// SidecarSetPatchPodMetadataDefaultsAllowed whether sidecarSet patch pod metadata is allowed
	SidecarSetPatchPodMetadataDefaultsAllowed featuregate.Feature = "SidecarSetPatchPodMetadataDefaultsAllowed"

	// PodProbeMarkerGate enable Kruise provide the ability to execute custom Probes.
	// Note: custom probe execution requires kruise daemon, so currently only traditional Kubelet is supported, not virtual-kubelet.
	PodProbeMarkerGate featuregate.Feature = "PodProbeMarkerGate"

	// PreDownloadImageForDaemonSetUpdate enables daemonset-controller to create ImagePullJobs to
	// pre-download images for update.
	PreDownloadImageForDaemonSetUpdate featuregate.Feature = "PreDownloadImageForDaemonSetUpdate"
)

var defaultFeatureGates = map[featuregate.Feature]featuregate.FeatureSpec{
	PodWebhook:        {Default: true, PreRelease: featuregate.Beta},
	KruiseDaemon:      {Default: true, PreRelease: featuregate.Beta},
	DaemonWatchingPod: {Default: true, PreRelease: featuregate.Beta},

	CloneSetShortHash:                         {Default: false, PreRelease: featuregate.Alpha},
	KruisePodReadinessGate:                    {Default: false, PreRelease: featuregate.Alpha},
	PreDownloadImageForInPlaceUpdate:          {Default: false, PreRelease: featuregate.Alpha},
	CloneSetPartitionRollback:                 {Default: false, PreRelease: featuregate.Alpha},
	ResourcesDeletionProtection:               {Default: false, PreRelease: featuregate.Alpha},
	WorkloadSpread:                            {Default: false, PreRelease: featuregate.Alpha},
	PodUnavailableBudgetDeleteGate:            {Default: false, PreRelease: featuregate.Alpha},
	PodUnavailableBudgetUpdateGate:            {Default: false, PreRelease: featuregate.Alpha},
	TemplateNoDefaults:                        {Default: false, PreRelease: featuregate.Alpha},
	InPlaceUpdateEnvFromMetadata:              {Default: false, PreRelease: featuregate.Alpha},
	StatefulSetAutoDeletePVC:                  {Default: false, PreRelease: featuregate.Alpha},
	SidecarSetPatchPodMetadataDefaultsAllowed: {Default: false, PreRelease: featuregate.Alpha},
	PodProbeMarkerGate:                        {Default: false, PreRelease: featuregate.Alpha},
	PreDownloadImageForDaemonSetUpdate:        {Default: false, PreRelease: featuregate.Alpha},
}

func init() {
	compatibleEnv()
	runtime.Must(feature.DefaultMutableFeatureGate.Add(defaultFeatureGates))
}

func compatibleEnv() {
	str := strings.TrimSpace(os.Getenv("CUSTOM_RESOURCE_ENABLE"))
	if len(str) == 0 {
		return
	}

	limits := sets.NewString(strings.Split(str, ",")...)
	if !limits.Has("SidecarSet") {
		defaultFeatureGates[PodWebhook] = featuregate.FeatureSpec{Default: false, PreRelease: featuregate.Beta}
	}
}

func SetDefaultFeatureGates() {
	if !feature.DefaultFeatureGate.Enabled(PodWebhook) {
		_ = feature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%s=false", KruisePodReadinessGate))
		_ = feature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%s=false", ResourcesDeletionProtection))
		_ = feature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%s=false", PodUnavailableBudgetDeleteGate))
		_ = feature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%s=false", PodUnavailableBudgetUpdateGate))
		_ = feature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%s=false", WorkloadSpread))
	}

	if !feature.DefaultFeatureGate.Enabled(KruiseDaemon) {
		_ = feature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%s=false", PreDownloadImageForDaemonSetUpdate))
		_ = feature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%s=false", DaemonWatchingPod))
		_ = feature.DefaultMutableFeatureGate.Set(fmt.Sprintf("%s=false", InPlaceUpdateEnvFromMetadata))
	}
}
