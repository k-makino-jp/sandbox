module gosample

go 1.14

require (
	github.com/Azure/azure-storage-queue-go v0.0.0-20191125232315-636801874cdd
	github.com/golang/mock v1.5.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	k8s.io/api v0.20.4
	k8s.io/apimachinery v0.20.4
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0 // indirect
	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009 // indirect
)

replace k8s.io/client-go => k8s.io/client-go v0.20.4