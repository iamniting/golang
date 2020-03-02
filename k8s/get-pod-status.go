package main

import (
    "fmt"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
)

func main() {
    config_file := "/home/nigoyal/.kube/config"

    // use the current context in kubeconfig
    config, err := clientcmd.BuildConfigFromFlags("", config_file)
    if err != nil { fmt.Println(err); return }
    // fmt.Println(config)

    // create the clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil { fmt.Println(err); return }
    // fmt.Println(clientset)

    // get list of pods
    pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
    if err != nil { fmt.Println(err); return }
    // fmt.Println(pods.Items)

    // iterate over list of pods
    for _, pod := range pods.Items {
        // fmt.Printf("%+v\n\n", pod)
        fmt.Print(pod.ObjectMeta.Name)
        // fmt.Printf("%+v\n\n", pod.Status)

        if pod.Status.Phase != "Running" {
            fmt.Print(" Pod is not running\n"); continue
        }

        flag := true
        // iterate over list of conditions for pod
        for _, cnd := range pod.Status.Conditions {
            if cnd.Type == "Ready" && cnd.Status == "True" {
                fmt.Print(" Pod is running\n")
                flag = false
                break
            }
        }

        if flag == true { fmt.Print(" Pod is not running\n") }
    }
}
