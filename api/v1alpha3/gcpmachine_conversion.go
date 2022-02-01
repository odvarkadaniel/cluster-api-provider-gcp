/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha3

import (
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1beta1 "sigs.k8s.io/cluster-api-provider-gcp/api/v1beta1"
)

// ConvertTo converts this GCPMachine to the Hub version (v1beta1).
func (src *GCPMachine) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*v1beta1.GCPMachine)

	if err := Convert_v1alpha3_GCPMachine_To_v1beta1_GCPMachine(src, dst, nil); err != nil {
		return err
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1beta1) to this version.
func (dst *GCPMachine) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*v1beta1.GCPMachine)
	if err := Convert_v1beta1_GCPMachine_To_v1alpha3_GCPMachine(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion.
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}

	return nil
}

// ConvertTo converts this GCPMachineList to the Hub version (v1beta1).
func (src *GCPMachineList) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*v1beta1.GCPMachineList)
	return Convert_v1alpha3_GCPMachineList_To_v1beta1_GCPMachineList(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1beta1) to this version.
func (dst *GCPMachineList) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*v1beta1.GCPMachineList)
	return Convert_v1beta1_GCPMachineList_To_v1alpha3_GCPMachineList(src, dst, nil)
}