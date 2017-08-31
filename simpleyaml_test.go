package go_simpleyaml

import (
	"fmt"
	"testing"
)

const compose  = `
version: '3'
services:
  test:
    image: hub.docker.com/public/tomcat:lastest
    volumes:
      - /opt/docker_data/release/ROOT.war:/opt/tomcat/webapps/ROOT.war
      - /opt/docker_data/release/server-conf/tiger:/opt/server-conf/tiger
      - /opt/docker_data/upload:/opt/upload
      - /opt/logs/:/opt/tomcat/logs
    ports:
      - 8080:8080
    networks:
      - test_network
    environment:
      BUILD_VERSION: 1.1

    healthcheck:
      test: ["CMD", "curl", "-f", "http://127.0.0.1:8080"]
      interval: 30s
      timeout: 10s
      retries: 3

    deploy:
      replicas: 5
      update_config:
        parallelism: 2
        delay: 2s
      restart_policy:
        condition: on-failure

networks:
  test_network:
    driver: overlay
    ipam:
      driver: default
      config:
        - subnet: 10.11.0.0/24
`

type Network struct {
	Driver string `yaml:"driver"`
	Ipam struct{
		Driver string `yaml:"driver"`
		Config []struct{
			Subnet string
		} `yaml:"config"`
	} `yaml:"ipam"`
}


func TestSimpleYaml(t *testing.T) {
	y := New()
	//y.Load("docker-compose.yml")
	err := y.Loads([]byte(compose))
	if err != nil {
		fmt.Println(err)
		return
	}

	y.Get("networks").
		Get("test_network").
		Get("ipam").
		Set("config", []string{"10.10.0.1/24"})

	d, _ := y.Dumps()
	fmt.Println(string(d))
	fmt.Println(y.Get("services").Keys()[0])



	fmt.Println("Test Network struct .....")
	network := new(Network)

	// init network data struct
	network.Driver = "overlay"
	network.Ipam.Config = make([]struct{Subnet string}, 1)
	network.Ipam.Config[0] = struct{Subnet string}{Subnet: "10.0.0.0/24"}  // init nested struct
	network.Ipam.Driver = "default"

	yml := NewYaml(network)
	d, _ = yml.Dumps()
	fmt.Println(string(d))
}

