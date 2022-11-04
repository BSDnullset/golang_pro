package main

import "fmt"

func addKubernetes_Clusters(clusters []string, cluster string) []string {
	return append(clusters, cluster)
}

func greet() string {
	return "clusters slice segregation"
}

func main() {
	fmt.Println(greet())
	clusters := []string{} // this is a void cluster slice
	result := addKubernetes_Clusters(clusters, "alpine")
	result = addKubernetes_Clusters(result, "fedora")
	result = addKubernetes_Clusters(result, "ubuntu")
	result = addKubernetes_Clusters(result, "debian")
	fmt.Println(result)
	// will print the clusters: (alpine, fedora, ubuntu, debian)
}
