package main

import (
	"context"
	"fmt"

	appsV1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	fmt.Println("========[AIM : is to create new project 'gogo' and http deployment.]============")

	// PART1 : Basic things needed to access cluster

	kubeconfigPath := "/Users/anisal/Documents/go_with_kubernetes/cluster/jansun3/auth/kubeconfig"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(fmt.Errorf("failed to load kubeconfig from %s: %v", kubeconfigPath, err))
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Errorf("Fail to create the k8s client set. Errorf - %s", err))
	}

	// PART2: Create new project, if already exist continue creating next there<yet to add code for this>

	// https://pkg.go.dev/k8s.io/client-go@v0.35.0/kubernetes/typed/core/v1#NamespacesGetter
	// https://github.com/kubernetes/client-go/blob/v0.35.0/kubernetes/typed/core/v1/namespace.go#L41

	// namespace, err := clientset.CoreV1().Namespaces("gogo").Create(context.TODO(), , "k8s.io/apimachinery/pkg/apis/meta/v1".CreateOptions)
	ns1name := "gogo4"
	ns1 := v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: ns1name}}
	namespace, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &ns1, metav1.CreateOptions{})
	if err != nil {
		panic(fmt.Errorf("not able to create project %s: %v", namespace, err))
	}

	fmt.Println(namespace.Name)

	// PART3: Create new deployment in project gogo, if already exist continue creating next there

	ds1name := "http"
	image := "image-registry.openshift-image-registry.svc:5000/default/httpd-example@sha256:ac805b9a9fca8417fe61c55a043d804521f7df531dd1a01ffa6c1732d0c9358e"
	ds1 := appsV1.Deployment{
		ObjectMeta: metav1.ObjectMeta{Name: ds1name},
		Spec: appsV1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"aaa": ds1name},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:   "http",
					Labels: map[string]string{"aaa": ds1name},
				},

				Spec: v1.PodSpec{

					Containers: []v1.Container{v1.Container{
						Image: image,
						Name:  "http",
					},
					},
				},
			},
		},
	}
	deployment, err := clientset.AppsV1().Deployments(ns1name).Create(context.TODO(), &ds1, metav1.CreateOptions{})
	if err != nil {
		panic(fmt.Errorf("not able to create deployment %s: %v", deployment, err))
	}

}

/*
Yipiiii!!!!!

========[AIM : is to create new project 'gogo' and http deployment.]============
gogo4

% oc get projects | grep gogo4
gogo4                                                             Active

% oc get all -n gogo4
Warning: apps.openshift.io/v1 DeploymentConfig is deprecated in v4.14+, unavailable in v4.10000+
NAME                        READY   STATUS    RESTARTS   AGE
pod/http-8448897f7b-gv4xv   1/1     Running   0          81s

NAME                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/http   1/1     1            1           81s

NAME                              DESIRED   CURRENT   READY   AGE
replicaset.apps/http-8448897f7b   1         1         1       81s

*/
