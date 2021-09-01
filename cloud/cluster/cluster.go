// Cluster provides the means for coordinating the schedulers and workers that
// make up a Scoot system. This is achieved mainly through the Cluster type,
// individual Nodes, and Subscriptions to cluster changes.
package cluster

import (
	"sort"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/twitter/scoot/common/stats"
)

var ClusterChBufferSize = 100

type NodeReqChType chan chan []Node

// cluster implementation of Cluster
type Cluster struct {
	state *state

	stat stats.StatsReceiver

	updateFreq time.Duration

	// latestFetchedNodesCh, contains the latest node list from fetcher.
	fetchedNodesCh chan []Node

	// nodesReqCh, used (by groupcache) to request the current list of nodes
	NodesReqCh NodeReqChType

	// nodesUpdatesChan, the node updates that the scheduler needs to process.
	NodesUpdatesCh chan []NodeUpdate

	useNodesUpdatesCh bool
	bufferSize        int

	priorNodeUpdateTime  time.Time
	priorFetchUpdateTime time.Time
}

// NewCluster creates a Cluster object and starts its processing loop and returns either a nodes updates channel
// (reporting node add/remove events) or nodes list channel (reporting the current list of nodes).
// The processing loop continuously gets the latest nodes list from the fetched nodes channel and either
// put the list of added/removed nodes on the nodes updates channel or if the nodes list is requested put the
// current list of nodes on the nodes list channel
func NewCluster(stat stats.StatsReceiver, fetcher Fetcher, useNodesUpdatesCh bool, updateFreq time.Duration, chBufferSize int) (chan []NodeUpdate, NodeReqChType) {
	fetchedNodesCh := StartFetchCron(fetcher, updateFreq, chBufferSize, stat)

	s := makeState([]Node{})
	var updatesCh chan []NodeUpdate = nil
	var nodesReqCh NodeReqChType = nil
	if useNodesUpdatesCh {
		updatesCh = make(chan []NodeUpdate, chBufferSize)
	} else {
		nodesReqCh = make(NodeReqChType)
	}
	c := &Cluster{
		state:                s,
		updateFreq:           updateFreq,
		stat:                 stat,
		fetchedNodesCh:       fetchedNodesCh,
		NodesReqCh:           nodesReqCh,
		NodesUpdatesCh:       updatesCh,
		useNodesUpdatesCh:    useNodesUpdatesCh,
		priorNodeUpdateTime:  time.Now(),
		priorFetchUpdateTime: time.Now(),
		bufferSize:           chBufferSize,
	}
	if stat == nil {
		c.stat = stats.NilStatsReceiver()
	}
	// logging
	chDesc := []string{}
	if c.NodesUpdatesCh != nil {
		chDesc = append(chDesc, "NodesUpdateCh")
	}
	if c.NodesReqCh != nil {
		chDesc = append(chDesc, "NodesReqCh")
	}
	log.Infof("cluster loop starting with frequency %s, reporting on %s with channel buffer size %d", c.updateFreq, strings.Join(chDesc[:], ","), chBufferSize)

	go c.loop()

	return updatesCh, nodesReqCh
}

// loop continuously get the latest list of nodes (on channel, typically set by fetcher).
// If the cluster is using the nodes updates channel, process the nodes putting the list
// of nodeAdd, nodeRemove entries on the nodes updates channel.  Otherwise if a list of nodes
// has been requested, put the most latest list of nodes on the nodes list channel
func (c *Cluster) loop() {
	// start processing loop
	ticker := time.NewTicker(c.updateFreq)
	for range ticker.C {
		lastestNodeList := c.getLatestNodesList() // get latest fetched nodes
		if lastestNodeList != nil {
			sort.Sort(NodeSorter(lastestNodeList))
			updates := c.state.setAndDiff(lastestNodeList) // compute updates and update local state
			if c.NodesUpdatesCh != nil {
				c.addToCurrentNodeUpdates(updates)
			}
		}

		if c.NodesReqCh != nil {
			select {
			case respCh := <-c.NodesReqCh:
				respCh <- c.getNodes()
			default:
			}
		}
	}
}

// pull (up to ClusterChBufferSize) node lists off the fetchedNodesCh
// keep only the latest list of nodes from the fetcher
func (c *Cluster) getLatestNodesList() []Node {
	var currentNodes []Node
	haveFetchedNodes := false
	i := 0
LOOP:
	for ; i < c.bufferSize; i++ {
		select {
		case currentNodes = <-c.fetchedNodesCh:
			haveFetchedNodes = true
		default:
			break LOOP
		}
	}

	if !haveFetchedNodes {
		return c.getNodes()
	}
	return currentNodes
}

// addToCurrentNodeUpdates put the node updates on the nodes updates channel.
func (c *Cluster) addToCurrentNodeUpdates(updates []NodeUpdate) {
	c.NodesUpdatesCh <- updates
	if len(updates) > 0 {
		log.Infof("cluster has %d new node updates", len(updates))
		// record time since last time saw node
		c.stat.Gauge(stats.ClusterNodeUpdateFreqMs).Update(time.Since(c.priorNodeUpdateTime).Milliseconds())
		c.priorNodeUpdateTime = time.Now()
	}
}

// GetNodes get (a copy of) the list of nodes last fetched by fetcher
func (c *Cluster) getNodes() []Node {
	ret := []Node{}
	for _, node := range c.state.nodes {
		ret = append(ret, node)
	}
	return ret
}
