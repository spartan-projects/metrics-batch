package main

import (
	"context"
	"google.golang.org/appengine/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientSet, err := metricsv.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	if podMetricsList, err := clientSet.MetricsV1beta1().PodMetricses("").List(context.TODO(), metav1.ListOptions{}); err != nil {
		panic(err)
	} else {
		log.Infof(context.TODO(), "PodMetricLists: %s", podMetricsList)
	}
}
