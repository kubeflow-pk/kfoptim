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
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var optimjoblog = logf.Log.WithName("optimjob-resource")

func (r *OptimJob) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-optimization-kubeflow-io-v1alpha1-optimjob,mutating=true,failurePolicy=fail,groups=optimization.kubeflow.io,resources=optimjobs,verbs=create;update,versions=v1alpha1,name=moptimjob.kb.io

var _ webhook.Defaulter = &OptimJob{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *OptimJob) Default() {
	optimjoblog.Info("default", "name", r.Name)

	//if r.Spec.ConcurrencyPolicy == "" {
	//	r.Spec.ConcurrencyPolicy = AllowConcurrent
	//}
	//if r.Spec.Suspend == nil {
	//	r.Spec.Suspend = new(bool)
	//}
	//if r.Spec.SuccessfulJobsHistoryLimit == nil {
	//	r.Spec.SuccessfulJobsHistoryLimit = new(int32)
	//	*r.Spec.SuccessfulJobsHistoryLimit = 3
	//}
	//if r.Spec.FailedJobsHistoryLimit == nil {
	//	r.Spec.FailedJobsHistoryLimit = new(int32)
	//	*r.Spec.FailedJobsHistoryLimit = 1
	//}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-optimization-kubeflow-io-v1alpha1-optimjob,mutating=false,failurePolicy=fail,groups=optimization.kubeflow.io,resources=optimjobs,versions=v1alpha1,name=voptimjob.kb.io

var _ webhook.Validator = &OptimJob{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *OptimJob) ValidateCreate() error {
	optimjoblog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return r.validateOptimJob()

}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *OptimJob) ValidateUpdate(old runtime.Object) error {
	optimjoblog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return  r.validateOptimJob()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *OptimJob) ValidateDelete() error {
	optimjoblog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}