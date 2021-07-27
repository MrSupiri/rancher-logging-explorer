package controllers

import (
	"context"
	"fmt"

	flowv1beta1 "github.com/banzaicloud/logging-operator/pkg/sdk/api/v1beta1"
	loggingplumberv1alpha1 "github.com/mrsupiri/logging-pipeline-plumber/pkg/sdk/api/v1alpha1"
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func (r *FlowTestReconciler) cleanUpResources(ctx context.Context, flowTestName string) error {
	logger := log.FromContext(ctx)

	matchingLabels := &client.MatchingLabels{"loggingplumber.isala.me/flowtest": flowTestName}

	var podList v1.PodList
	if err := r.List(ctx, &podList, matchingLabels); client.IgnoreNotFound(err) != nil {
		logger.Error(err, fmt.Sprintf("failed to get provisioned %s", podList.Kind))
		return err
	}

	for _, resource := range podList.Items {
		if err := r.Delete(ctx, &resource); client.IgnoreNotFound(err) != nil {
			logger.Error(err, fmt.Sprintf("failed to delete a provisioned %s", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
			return err
		}
		logger.V(1).Info(fmt.Sprintf("%s deleted", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
	}

	var configMapList v1.ConfigMapList
	if err := r.List(ctx, &configMapList, matchingLabels); client.IgnoreNotFound(err) != nil {
		logger.Error(err, fmt.Sprintf("failed to get provisioned %s", configMapList.Kind))
		return err
	}

	for _, resource := range configMapList.Items {
		if err := r.Delete(ctx, &resource); client.IgnoreNotFound(err) != nil {
			logger.Error(err, fmt.Sprintf("failed to delete a provisioned %s", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
			return err
		}
		logger.V(1).Info(fmt.Sprintf("%s deleted", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
	}

	var flows flowv1beta1.FlowList
	if err := r.List(ctx, &flows, &client.MatchingLabels{"loggingplumber.isala.me/flowtest": flowTestName}); client.IgnoreNotFound(err) != nil {
		logger.Error(err, fmt.Sprintf("failed to get provisioned %s", flows.Kind))
		//return err
	}

	for _, resource := range flows.Items {
		if err := r.Delete(ctx, &resource); client.IgnoreNotFound(err) != nil {
			logger.Error(err, fmt.Sprintf("failed to delete a provisioned %s", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
			return err
		}
		logger.V(1).Info(fmt.Sprintf("%s deleted", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
	}

	var outputs flowv1beta1.OutputList
	if err := r.List(ctx, &outputs, &client.MatchingLabels{"loggingplumber.isala.me/flowtest": flowTestName}); client.IgnoreNotFound(err) != nil {
		logger.Error(err, fmt.Sprintf("failed to get provisioned %s", outputs.Kind))
		//return err
	}

	for _, resource := range outputs.Items {
		if err := r.Delete(ctx, &resource); client.IgnoreNotFound(err) != nil {
			logger.Error(err, fmt.Sprintf("failed to delete a provisioned %s", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
			return err
		}
		logger.V(1).Info(fmt.Sprintf("%s deleted", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
	}

	var clusterFlows flowv1beta1.ClusterFlowList
	if err := r.List(ctx, &clusterFlows, &client.MatchingLabels{"loggingplumber.isala.me/flowtest": flowTestName}); client.IgnoreNotFound(err) != nil {
		logger.Error(err, fmt.Sprintf("failed to get provisioned %s", clusterFlows.Kind))
		//return err
	}

	for _, resource := range clusterFlows.Items {
		if err := r.Delete(ctx, &resource); client.IgnoreNotFound(err) != nil {
			logger.Error(err, fmt.Sprintf("failed to delete a provisioned %s", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
			return err
		}
		logger.V(1).Info(fmt.Sprintf("%s deleted", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
	}

	var clusterOutputs flowv1beta1.ClusterOutputList
	if err := r.List(ctx, &clusterOutputs, &client.MatchingLabels{"loggingplumber.isala.me/flowtest": flowTestName}); client.IgnoreNotFound(err) != nil {
		logger.Error(err, fmt.Sprintf("failed to get provisioned %s", clusterOutputs.Kind))
		//return err
	}

	for _, resource := range clusterOutputs.Items {
		if err := r.Delete(ctx, &resource); client.IgnoreNotFound(err) != nil {
			logger.Error(err, fmt.Sprintf("failed to delete a provisioned %s", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
			return err
		}
		logger.V(1).Info(fmt.Sprintf("%s deleted", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
	}

	return nil
}

func (r *FlowTestReconciler) cleanUpOutputResources(ctx context.Context) error {
	logger := log.FromContext(ctx)

	var flowTests loggingplumberv1alpha1.FlowTestList
	if err := r.List(ctx, &flowTests, &client.MatchingLabels{"app.kubernetes.io/created-by": "logging-plumber"}); err != nil {
		logger.Error(err, fmt.Sprintf("failed to get provisioned %s", flowTests.Kind))
		return err
	}
	for _, flowTest := range flowTests.Items {
		if flowTest.Status.Status != loggingplumberv1alpha1.Completed {
			logger.V(1).Info("unfinished flowtest found, skipping cleanUpOutputResources")
			return nil
		}
	}

	matchingLabels := &client.MatchingLabels{"loggingplumber.isala.me/component": "log-aggregator"}
	var podList v1.PodList
	if err := r.List(ctx, &podList, matchingLabels); client.IgnoreNotFound(err) != nil {
		logger.Error(err, fmt.Sprintf("failed to get provisioned %s", podList.Kind))
		return err
	}

	for _, resource := range podList.Items {
		if err := r.Delete(ctx, &resource); client.IgnoreNotFound(err) != nil {
			logger.Error(err, fmt.Sprintf("failed to delete a provisioned %s", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
			return err
		}
		logger.V(1).Info(fmt.Sprintf("%s deleted", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
	}

	var serviceList v1.ServiceList
	if err := r.List(ctx, &serviceList, matchingLabels); client.IgnoreNotFound(err) != nil {
		logger.Error(err, fmt.Sprintf("failed to get provisioned %s", podList.Kind))
		return err
	}

	for _, resource := range serviceList.Items {
		if err := r.Delete(ctx, &resource); client.IgnoreNotFound(err) != nil {
			logger.Error(err, fmt.Sprintf("failed to delete a provisioned %s", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
			return err
		}
		logger.V(1).Info(fmt.Sprintf("%s deleted", resource.Kind), "uuid", resource.GetUID(), "name", resource.GetName())
	}

	return nil
}
