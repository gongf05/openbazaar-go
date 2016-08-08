package corehttp

import (
	"net"
	"net/http"

	prometheus "gx/ipfs/QmdhsRK1EK2fvAz2i2SH5DEfkL6seDuyMYEsxKa9Braim3/client_golang/prometheus"

	core "github.com/ipfs/go-ipfs/core"
)

// This adds the scraping endpoint which Prometheus uses to fetch metrics.
func MetricsScrapingOption(path string) ServeOption {
	return func(n *core.IpfsNode, _ net.Listener, mux *http.ServeMux) (*http.ServeMux, error) {
		mux.Handle(path, prometheus.UninstrumentedHandler())
		return mux, nil
	}
}

// This adds collection of net/http-related metrics
func MetricsCollectionOption(handlerName string) ServeOption {
	return func(_ *core.IpfsNode, _ net.Listener, mux *http.ServeMux) (*http.ServeMux, error) {
		childMux := http.NewServeMux()
		mux.HandleFunc("/", prometheus.InstrumentHandler(handlerName, childMux))
		return childMux, nil
	}
}

var (
	peersTotalMetric = prometheus.NewDesc(
		prometheus.BuildFQName("ipfs", "p2p", "peers_total"),
		"Number of connected peers", []string{"transport"}, nil)
)

type IpfsNodeCollector struct {
	Node *core.IpfsNode
}

func (_ IpfsNodeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- peersTotalMetric
}

func (c IpfsNodeCollector) Collect(ch chan<- prometheus.Metric) {
	for tr, val := range c.PeersTotalValues() {
		ch <- prometheus.MustNewConstMetric(
			peersTotalMetric,
			prometheus.GaugeValue,
			val,
			tr,
		)
	}
}

func (c IpfsNodeCollector) PeersTotalValues() map[string]float64 {
	vals := make(map[string]float64)
	for _, conn := range c.Node.PeerHost.Network().Conns() {
		tr := ""
		for _, proto := range conn.RemoteMultiaddr().Protocols() {
			tr = tr + "/" + proto.Name
		}
		vals[tr] = vals[tr] + 1
	}
	return vals
}
