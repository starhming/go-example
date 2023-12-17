package main

import (
	"flag"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	kubeconfig *string
	clientset  *kubernetes.Clientset
)

// 加载配置文件，初始化 clientset
func InitKubeConfig() {
	if home := homedir.HomeDir(); home != "" {
		// 优先从home目录下寻找 config 文件
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// 首先使用 inCluster 模式（需要去配置对应的 RBAC 权限，默认的 sa 是 default-> 是没有获取 nodes 的 List 权限）
	config, err := rest.InClusterConfig()
	if err != nil {
		// 使用 KubeConfig 文件创建集群配置 Config 对象
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}

	// 创建 clientset 实例
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

// 导出私有变量 clientset
func GetClientSet() *kubernetes.Clientset {
	return clientset
}
