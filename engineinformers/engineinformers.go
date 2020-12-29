package engineinformers

import (
	"fmt"
	"time"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

func TestInformer() {

	var namespace string = "default"
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)
	informer := informerFactory.Apps().V1().Deployments()
	informer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc:    onAdd,
			UpdateFunc: onUpdate,
			DeleteFunc: onDelete,
		})
	lister := informer.Lister()

	stopCh := make(chan struct{})
	defer close(stopCh)
	informerFactory.Start(stopCh)
	if !cache.WaitForCacheSync(stopCh, informer.Informer().HasSynced) {
		return
	}

	deployments, err := lister.Deployments(namespace).List(labels.Everything())
	if err != nil {
		panic(err)
	}
	for _, deployment := range deployments {
		fmt.Printf("%s\r\n", deployment.Name)
	}
	<-stopCh
}

func onAdd(obj interface{}) {
	deployment := obj.(*v1.Deployment)
	fmt.Printf("onAdd:%s\r\n", deployment.Name)
}

func onUpdate(old, new interface{}) {
	oldDeployment := old.(*v1.Deployment)
	newDeployment := new.(*v1.Deployment)
	if newDeployment.ResourceVersion != oldDeployment.ResourceVersion {
		fmt.Printf("onUpdate [%s]:%s.%s.%s to %s.%s.%s---->\r\n", time.Now(), oldDeployment.Name, oldDeployment.ResourceVersion, oldDeployment.Spec.Strategy.RollingUpdate.MaxSurge, newDeployment.Name, newDeployment.ResourceVersion, newDeployment.Spec.Strategy.RollingUpdate.MaxSurge)
	}
}

func onDelete(obj interface{}) {
	deployment := obj.(*v1.Deployment)
	fmt.Printf("onDelete:%s\r\n", deployment.Name)
}
