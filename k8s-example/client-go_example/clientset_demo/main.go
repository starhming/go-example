package main

import (
	"context"
	"flag"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var err error
	var config *rest.Config
	var kubeconfig *string

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(可选) kubeconfig 文件的绝对路径")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "kubeconfig 文件的绝对路径")
	}
	flag.Parse()

	// 首先使用 inCluster 模式(需要去配置对应的 RBAC 权限，默认的sa是default->是没有获取deployments的List权限)
	if config, err = rest.InClusterConfig(); err != nil {
		// 使用 KubeConfig 文件创建集群配置 Config 对象
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}

	// 创建 ClientSet 实例
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// 设置 list options
	listOptions := metav1.ListOptions{
		LabelSelector: "",
		FieldSelector: "",
	}

	// 获取 default 命名空间下的 pod 列表
	pods, err := clientset.CoreV1().Pods("etcd").List(context.TODO(), listOptions)
	if err != nil {
		log.Fatal(err)
	}
	for _, pod := range pods.Items {
		log.Println(pod.Name)
	}
}
