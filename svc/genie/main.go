package main

import (
	"context"
	"fmt"
	"github.com/statuzproj/statuz/utils/healthz"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"
	"os"
	"os/user"
)

func main() {

	usr, err := user.Current()
	if err != nil {
		log.Printf("Error getting user's home directory: %v\n", err)
		os.Exit(1)
	}

	kubeconfigPath := usr.HomeDir + "/.kube/config"

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Printf("Error loading kubeconfig: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating clientset: %v\n", err)
		os.Exit(1)
	}

	configMaps, err := clientset.CoreV1().ConfigMaps("kube-system").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Error listing ConfigMaps: %v\n", err)
		os.Exit(1)
	}

	for _, cm := range configMaps.Items {
		fmt.Printf("ConfigMap Name: %s\n", cm.Name)
	}

	log.Println("Work in Progress!!")
	http.HandleFunc("/healthz", healthz.HealthCheck)
	http.ListenAndServe(":8081", nil)
}
