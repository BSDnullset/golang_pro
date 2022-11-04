package main

import (
	"fmt"
)

func addKubernetes_Clusters(clusters []string, cluster string) []string (
	return append(clusters, cluster)
)

func main() {
	fmt.println("clusters slice segregation")
	clusters := []string{} // this is a void cluster slice
	result := addKubernetes_Clusters(clusters, "alpine")
	result := addKubernetes_Clusters(clusters, "fedora")
	resutl := addKubernetes_Clusters(clusters, "ubuntu")
	result := addKubernetes_Clusters(clusters, "debian")
	fmt.println(result)
	// will print the clusters: (alpine, fedora, ubuntu, debian)
}
