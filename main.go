package main

import "fmt"

func addKubernetes_Clusters(clusters []string, cluster string) []string {
	return append(clusters, cluster)
}

func greet() string { return "clusters slice segregation" }

func addClusterID(clusters []int, cluster int) []int {
	return append(clusters, cluster)
}

func displayID_number() string { return "clusters have unique identifiers" }

func main() {
	fmt.Println(greet())
	clusters := []string{} // this is a void cluster slice
	result := addKubernetes_Clusters(clusters, "alpine")
	result = addKubernetes_Clusters(result, "fedora")
	result = addKubernetes_Clusters(result, "ubuntu")
	result = addKubernetes_Clusters(result, "debian")
	fmt.Println(result)
	// will print the clusters: (alpine, fedora, ubuntu, debian)
	// will add more here

	fmt.Println(displayID_number())
	clusterID := []int{}
	ID_ofCluster := addClusterID(clusterID, 293)
	ID_ofCluster = addClusterID(ID_ofCluster, 734)
	ID_ofCluster = addClusterID(ID_ofCluster, 835)
	ID_ofCluster = addClusterID(ID_ofCluster, 452)
	fmt.Println(ID_ofCluster)
}
