package balancer

import "errors"

var (
	NoHostError                = errors.New("no host")
	AlgorithmNotSupportedError = errors.New("algorithm not supported")
)

// Balancer interface is the load balancer for the reverse proxy
type Balancer interface {
	Add(string)                     // add the host for proxy
	Remove(string)                  // remove the host for proxy
	Balance(string) (string, error) // select a host to receive the request
	Inc(string)                     // increase connect for proxy host
	Done(string2 string)            // decrease connect for proxy host
}

// Factory is the factory that generates Balancer,
// and the factory design pattern is used here
type Factory func([]string) Balancer

var factories = make(map[string]Factory)

func Build(algorithm string, hosts []string) (Balancer, error) {
	factory, ok := factories[algorithm]
	if !ok {
		return nil, AlgorithmNotSupportedError
	}
	return factory(hosts), nil
}
