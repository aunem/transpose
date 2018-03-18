package config

import (
	"context"

	"io/ioutil"

	"github.com/ericchiang/k8s"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// LoadConfig loads the main config from k8s or local
func LoadConfig(name, namespace string) (o *Transpose, err error) {
	log.Debugf("loading config with name: %s and namespace: %s", name, namespace)
	if namespace == "local" {
		o, err = loadLocal()
	} else {
		log.Debug("connecting to k8s")
		client, err := k8s.NewInClusterClient()
		if err != nil {
			log.Infof("could not connect to k8s: %v, using local config...", err)
			o, err = loadLocal()
		} else {
			o, err = loadK8s(client, name, namespace)
			if err != nil {
				log.Infof("could not get k8s config: %v, using local config...", err)
				o, err = loadLocal()
			}
		}
	}
	return
}

func loadK8s(cli *k8s.Client, name, namespace string) (*Transpose, error) {
	log.Debug("loading k8s config")
	var o Transpose
	err := cli.Get(context.Background(), namespace, name, &o)
	return &o, err
}

func loadLocal() (*Transpose, error) {
	log.Debug("loading local config")
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}
	var o Transpose
	err = yaml.Unmarshal(data, &o)
	return &o, err
}
