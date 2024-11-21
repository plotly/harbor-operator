// Code generated by "stringer -type=Controller -linecomment"; DO NOT EDIT.

package controllers

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Core-0]
	_ = x[JobService-1]
	_ = x[Portal-2]
	_ = x[Registry-3]
	_ = x[RegistryController-4]
	_ = x[Exporter-5]
	_ = x[Trivy-6]
	_ = x[Harbor-7]
	_ = x[HarborCluster-8]
	_ = x[HarborConfigurationCm-9]
	_ = x[HarborConfiguration-10]
	_ = x[HarborProject-11]
	_ = x[HarborServerConfiguration-12]
	_ = x[PullSecretBinding-13]
	_ = x[Namespace-14]
}

const _Controller_name = "corejobserviceportalregistryregistryctlexportertrivyharborharborclusterharborconfigurationcmharborconfigurationharborprojectharborserverconfigurationpullsecretbindingnamespace"

var _Controller_index = [...]uint8{0, 4, 14, 20, 28, 39, 47, 52, 58, 71, 92, 111, 124, 149, 166}

func (i Controller) String() string {
	if i < 0 || i >= Controller(len(_Controller_index)-1) {
		return "Controller(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Controller_name[_Controller_index[i]:_Controller_index[i+1]]
}
