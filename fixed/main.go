package main

import (
	"fmt"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	fmt.Printf("%+v\n", v1.PodList{})
	fmt.Printf("%+v\n", metav1.ListOptions{})
}


