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

	// clientset creation
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
	fmt.Println("listing cms in openshift-monitoring")
	// https://pkg.go.dev/k8s.io/api/core/v1#ConfigMapList

	for _, myconfigmap := range configmap.Items {
		fmt.Println(myconfigmap.Name)
	}

}

/*
hello
========undertsood============
https://api.xxxx:6443
listing pods in openshift-monitoring
alertmanager-main-0
alertmanager-main-1
cluster-monitoring-operator-767bd5899-rmvdl
kube-state-metrics-5df56774d9-sx92p
metrics-server-7b4d979c6b-9txls
metrics-server-7b4d979c6b-j8z7k
monitoring-plugin-68456548c-288hr
monitoring-plugin-68456548c-xbjfm
node-exporter-2twdv
node-exporter-84pll
node-exporter-n7lrv
node-exporter-qrgxc
node-exporter-sxk97
node-exporter-wdhxd
openshift-state-metrics-5f6d67f6df-nkr4k
prometheus-k8s-0
prometheus-k8s-1
prometheus-operator-5bf84fb8cf-6bvz6
prometheus-operator-admission-webhook-5fc75b4846-9vjlm
prometheus-operator-admission-webhook-5fc75b4846-dr9xj
telemeter-client-6cdfcdfd-hzzwd
thanos-querier-54c545c6dd-jxld6
thanos-querier-54c545c6dd-kzb7q
========tried============
listing cms in openshift-monitoring
alertmanager-trusted-ca-bundle
kube-root-ca.crt
kube-state-metrics-custom-resource-state-configmap
kubelet-serving-ca-bundle
metrics-client-ca
metrics-server-audit-profiles
node-exporter-accelerators-collector-config
openshift-service-ca.crt
prometheus-k8s-rulefiles-0
prometheus-trusted-ca-bundle
serving-certs-ca-bundle
telemeter-client-serving-certs-ca-bundle
telemeter-trusted-ca-bundle
telemeter-trusted-ca-bundle-56c9b9fa8d9gs
telemetry-config
*/
