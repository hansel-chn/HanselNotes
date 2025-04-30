## Link

* https://github.com/overnote/over-golang/blob/master/06-%E5%BE%AE%E6%9C%8D%E5%8A%A1/04-%E6%9C%8D%E5%8A%A1%E5%8F%91%E7%8E%B0.md

##    

* A Pod can contain one or more tightly coupled containers that need to share resources.
  Containers in the same Pod share the same IP address, port space, and storage volumes.
* 网络可以互相通信
* 共享存储卷
* Single-container Pods: Most common, where a Pod runs a single container.
  Multi-container Pods: Used when containers need to work closely together, such as a main application container and a
  sidecar container for logging or monitoring.

## daemonSet

但是，Deployment并不关心这些Pod会在集群的哪些节点上运行，在它看来，Pod的运行环境与功能是无关的，只要Pod的数量足够，应用程序应该会正常工作。

## 服务注册
Kubernetes Service
Purpose:  
Provides basic service discovery and load balancing for Pods in a Kubernetes cluster.
Exposes a group of Pods as a single network endpoint.
Features:  
Service Discovery: Uses DNS to allow Pods to find and communicate with each other.
Load Balancing: Distributes traffic across multiple Pods.
Types of Services:
ClusterIP: Internal communication within the cluster.
NodePort: Exposes the service on a static port on each node.
LoadBalancer: Exposes the service externally using a cloud provider's load balancer.
ExternalName: Maps a service to an external DNS name.