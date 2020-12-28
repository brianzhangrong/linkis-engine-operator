package goclient

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func Test() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// for {
	pods, err := clientset.AppsV1().Deployments("istio-system").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, p := range pods.Items {
		fmt.Printf("There are %s pods in the cluster\n", p.Name)
	}

	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	// namespace := "default"
	// // pod := "nginx"
	// // pod, err := clientset.CoreV1()
	// saList, err := clientset.CoreV1().ServiceAccounts(namespace).List(metav1.ListOptions{})
	// // _, err = clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
	// // if errors.IsNotFound(err) {
	// // 	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	// // } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// // 	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	// // 		pod, namespace, statusError.ErrStatus.Message)
	// // } else if err != nil {
	// // 	panic(err.Error())
	// // } else {
	// // 	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	// // }
	// if errors.IsNotFound(err) {
	// 	fmt.Println("list sa error %s", err)
	// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 	fmt.Println("----", statusError.ErrStatus.Message)
	// } else if err != nil {

	// } else {
	// 	fmt.Println("=====", len(saList.Items))
	// }
	time.Sleep(10 * time.Second)
	// }
}
