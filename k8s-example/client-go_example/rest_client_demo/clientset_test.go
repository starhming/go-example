package main

import (
	"context"
	"fmt"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestClientSetQueryNode(t *testing.T) {
	InitKubeConfig()

	list, _ := GetClientSet().CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	for _, item := range list.Items {
		str := item.String()
		fmt.Println(str)

	}

}
