package xclient

import (
	"log"
	"net/http"
	"strings"
	"time"
)

/*
	type MultiServersDiscovery struct {
		r       *rand.Rand   // generate random number
		mu      sync.RWMutex // protect following
		servers []string
		index   int // record the selected position for robin algorithm
	}
*/
type XRegistryDiscovery struct {
	*MultiServersDiscovery // 复用
	registry               string
	timeout                time.Duration // 服务列表的过期时间
	lastUpdate             time.Time     // 最后从注册中心更新服务列表的时间
}

const defaultUpdateTimeout = time.Second * 10

func NewXRegistryDiscovery(registerAddr string, timeout time.Duration) *XRegistryDiscovery {
	if timeout == 0 {
		timeout = defaultUpdateTimeout
	}
	d := &XRegistryDiscovery{
		MultiServersDiscovery: NewMultiServerDiscovery(make([]string, 0)),
		registry:              registerAddr,
		timeout:               timeout,
	}
	return d
}

func (d *XRegistryDiscovery) Update(servers []string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.servers = servers
	d.lastUpdate = time.Now()
	return nil
}

func (d *XRegistryDiscovery) Refresh() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.lastUpdate.Add(d.timeout).After(time.Now()) { // 未过期, 无需重新向注册中心拉去服务列表
		return nil
	}
	log.Println("rpc registry: refresh servers from registry", d.registry)
	resp, err := http.Get(d.registry)
	if err != nil {
		log.Println("rpc registry refresh err:", err)
		return err
	}
	servers := strings.Split(resp.Header.Get("X-rpc-Servers"), ",")
	d.servers = make([]string, 0, len(servers))
	for _, server := range servers {
		if strings.TrimSpace(server) != "" {
			d.servers = append(d.servers, strings.TrimSpace(server))
		}
	}
	d.lastUpdate = time.Now() // 更新最新拉去服务列表时间
	return nil
}

// Get服务之前要先刷新服务列表
func (d *XRegistryDiscovery) Get(mode SelectMode) (string, error) {
	if err := d.Refresh(); err != nil {
		return "", err
	}
	return d.MultiServersDiscovery.Get(mode)
}

func (d *XRegistryDiscovery) GetAll() ([]string, error) {
	if err := d.Refresh(); err != nil {
		return nil, err
	}
	return d.MultiServersDiscovery.GetAll()
}
