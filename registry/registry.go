package registry

import (
	"errors"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/contrib/registry/kubernetes/v2"
	kreg "github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
	etcdv3 "go.etcd.io/etcd/client/v3"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func getClientSet() (*kubernetes.Clientset, error) {
	restConfig, err := rest.InClusterConfig()
	home := homedir.HomeDir()

	if err != nil {
		kubeconfig := filepath.Join(home, ".kube", "config")
		restConfig, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	}
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return clientSet, nil
}

func New(c *Conf) (kreg.Registrar, kreg.Discovery) {
	if c.Etcd.Enable {
		cli, err := etcdv3.New(etcdv3.Config{
			Endpoints: c.Etcd.Endpoints,
		})
		if err != nil {
			panic(err)
		}
		r := etcd.New(cli)
		return r, r
	} else if c.Consul.Enable {
		// consul
		consulConfig := api.DefaultConfig()
		consulConfig.Address = c.Consul.Address
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			panic(err)
		}
		r := consul.New(consulClient)
		return r, r
	} else if c.Kube.Enable {
		cli, err := getClientSet()
		if err != nil {
			panic(err)
		}
		r := kuberegistry.NewRegistry(cli)
		return r, r
	} else {
		panic(errors.New("not support"))
	}
}