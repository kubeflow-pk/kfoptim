/*


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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OptimJobSpec defines the desired state of OptimJob
type OptimJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of OptimJob. Edit OptimJob_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// OptimJobStatus defines the observed state of OptimJob
type OptimJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// OptimJob is the Schema for the optimjobs API
type OptimJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OptimJobSpec   `json:"spec,omitempty"`
	Status OptimJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OptimJobList contains a list of OptimJob
type OptimJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OptimJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OptimJob{}, &OptimJobList{})
}

func (r *OptimJob) validateOptimJob() error {
	var allErrs field.ErrorList
	//if err := r.validateOptimJobName(); err != nil {
	//	allErrs = append(allErrs, err)
	//}
	//if err := r.validateCronJobSpec(); err != nil {
	//	allErrs = append(allErrs, err)
	//}
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "batch.tutorial.kubebuilder.io", Kind: "CronJob"},
		r.Name, allErrs)
}