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

package controllers

import (
	"context"
	//"k8s.io/client-go/informers/batch"
	//"sort"
	//"time"

	kbatch "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	//_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	//_ "k8s.io/client-go/tools/reference"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	optimization "github.com/kubeflow-pk/kfoptim/pkg/apis/optimization/v1alpha1"
)

// OptimJobReconciler reconciles a OptimJob object
type OptimJobReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

}

// +kubebuilder:rbac:groups=optimization.kubeflow.io,resources=optimjobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=optimization.kubeflow.io,resources=optimjobs/status,verbs=get;update;patch

func (r *OptimJobReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("optimjob", req.NamespacedName)

	var job optimization.OptimJob
	if err := r.Get(ctx, req.NamespacedName, &job); err != nil {
		log.Error(err, "unable to fetch OptimJob")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var childJobs kbatch.JobList
	if err := r.List(ctx, &childJobs, client.InNamespace(req.Namespace), client.MatchingFields{jobOwnerKey: req.Name}); err != nil {
		log.Error(err, "unable to list child Jobs")
		return ctrl.Result{}, err
	}

	// find the active list of jobs
	var activeJobs []*kbatch.Job
	var successfulJobs []*kbatch.Job
	var failedJobs []*kbatch.Job
	//var mostRecentTime *time.Time // find the last run so we can update the status

	isJobFinished := func(job *kbatch.Job) (bool, kbatch.JobConditionType) {
		for _, c := range job.Status.Conditions {
			if (c.Type == kbatch.JobComplete || c.Type == kbatch.JobFailed) && c.Status == corev1.ConditionTrue {
				return true, c.Type
			}
		}

		return false, ""
	}
	// +kubebuilder:docs-gen:collapse=isJobFinished

	for i, job := range childJobs.Items {
		_, finishedType := isJobFinished(&job)
		switch finishedType {
		case "": // ongoing
			activeJobs = append(activeJobs, &childJobs.Items[i])
		case kbatch.JobFailed:
			failedJobs = append(failedJobs, &childJobs.Items[i])
		case kbatch.JobComplete:
			successfulJobs = append(successfulJobs, &childJobs.Items[i])
		}

		// We'll store the launch time in an annotation, so we'll reconstitute that from
		// the active jobs themselves.
		//scheduledTimeForJob, err := getScheduledTimeForJob(&job)
		//if err != nil {
		//	log.Error(err, "unable to parse schedule time for child job", "job", &job)
		//	continue
		//}
		//if scheduledTimeForJob != nil {
		//	if mostRecentTime == nil {
		//		mostRecentTime = scheduledTimeForJob
		//	} else if mostRecentTime.Before(*scheduledTimeForJob) {
		//		mostRecentTime = scheduledTimeForJob
		//	}
		//}

		//job.Status.Active = nil
		//for _, activeJob := range activeJobs {
		//	jobRef, err := ref.GetReference(r.Scheme, activeJob)
		//	if err != nil {
		//		log.Error(err, "unable to make reference to active job", "job", activeJob)
		//		continue
		//	}
		//	//job.Status.Active = append(job.Status.Active, *jobRef)
		//}

		log.V(1).Info("job count", "active jobs", len(activeJobs), "successful jobs", len(successfulJobs), "failed jobs", len(failedJobs))

		if err := r.Status().Update(ctx, &job); err != nil {
			log.Error(err, "unable to update CronJob status")
			return ctrl.Result{}, err
		}
		//
		//if job.Spec.Suspend != nil && *job.Spec.Suspend {
		//	log.V(1).Info("cronjob suspended, skipping")
		//	return ctrl.Result{}, nil
		//}

	}
	return ctrl.Result{}, nil
}

var (
	jobOwnerKey = ".metadata.controller"
	//apiGVStr    = batch.GroupVersion.String()
)

func (r *OptimJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	//if err := mgr.GetFieldIndexer().IndexField(context.Background(), &kbatch.Job{}, jobOwnerKey, func(rawObj client.Object) []string {
	//	// grab the job object, extract the owner...
	//	job := rawObj.(*kbatch.Job)
	//	owner := metav1.GetControllerOf(job)
	//	if owner == nil {
	//		return nil
	//	}
	//	// ...make sure it's a CronJob...
	//	if owner.APIVersion != apiGVStr || owner.Kind != "CronJob" {
	//		return nil
	//	}
	//
	//	// ...and if so, return it
	//	return []string{owner.Name}
	//}); err != nil {
	//	return err
	//}
	return ctrl.NewControllerManagedBy(mgr).
		For(&optimization.OptimJob{}).
		Complete(r)
}
