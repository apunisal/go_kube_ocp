package main

import (
	"context"

	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	fmt.Println("hello")

	fmt.Println("========undertsood============")
	kubeconfigPath := "/Users/anisal/Documents/go_with_kubernetes/cluster/jansun3/auth/kubeconfig"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(fmt.Errorf("failed to load kubeconfig from %s: %v", kubeconfigPath, err))
	}
	fmt.Println(config.Host)

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Errorf("Fail to create the k8s client set. Errorf - %s", err))
	}

	// list pods in monitoring project
	pods, err := clientset.CoreV1().Pods("openshift-monitoring").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(fmt.Errorf("Fail to get list of pods - %s", err))
	}

	fmt.Println("listing pods in openshift-monitoring")

	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	fmt.Println("========tried============")
	// https://pkg.go.dev/k8s.io/client-go@v0.35.0/kubernetes/typed/core/v1#ConfigMapInterface
	// https://pkg.go.dev/context#Context
	configmap, err := clientset.CoreV1().ConfigMaps("openshift-monitoring").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(fmt.Errorf("Fail to get list of configmap - %s", err))
	}
	fmt.Println("listing pods in openshift-monitoring")
	// https://pkg.go.dev/k8s.io/api/core/v1#ConfigMapList

	for _, myconfigmap := range configmap.Items {
		fmt.Println(myconfigmap.Name)
	}

}
