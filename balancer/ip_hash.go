package balancer

import (
	"hash/crc32"
	"sync"
)

func init() {
	factories[IPHashBalancer] = NewIPHash
}

type IPHash struct {
	sync.RWMutex
	hosts []string
}

func NewIPHash(hosts []string) Balancer {
	return &IPHash{
		hosts: hosts,
	}
}

func (r *IPHash) Add(host string) {
	r.Lock()
	defer r.Unlock()
	for _, h := range r.hosts {
		if h == host {
			return
		}
	}
	r.hosts = append(r.hosts, host)
}

func (r *IPHash) Remove(host string) {
	r.Lock()
	defer r.Unlock()
	for i, h := range r.hosts {
		if h == host {
			r.hosts = append(r.hosts[:i], r.hosts[i+1:]...)
			return
		}
	}
}

func (r *IPHash) Balance(key string) (string, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.hosts) == 0 {
		return "", NoHostError
	}
	value := crc32.ChecksumIEEE([]byte(key)) % uint32(len(r.hosts))
	return r.hosts[value], nil
}

// Inc .
func (r *IPHash) Inc(_ string) {

}

// Done .
func (r *IPHash) Done(_ string) {

}
