package daemon

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/noironetworks/cilium-net/common"
	"github.com/noironetworks/cilium-net/common/types"

	. "github.com/noironetworks/cilium-net/Godeps/_workspace/src/github.com/hashicorp/consul/api"
	. "github.com/noironetworks/cilium-net/Godeps/_workspace/src/gopkg.in/check.v1"
)

var (
	lbls = types.Labels{
		"foo":    "bar",
		"foo2":   "=bar2",
		"key":    "",
		"foo==":  "==",
		`foo\\=`: `\=`,
		`//=/`:   "",
		`%`:      `%ed`,
	}
	lbls2 = types.Labels{
		"foo":  "bar",
		"foo2": "=bar2",
	}
)

func (ds *DaemonSuite) SetUpTest(c *C) {
	c.Skip("To test, enable a consul environment")
	// Make client config
	conf := DefaultConfig()

	d, err := NewDaemon("", nil, nil, conf)
	c.Assert(err, Equals, nil)
	ds.d = d
	d.consul.KV().DeleteTree(common.OperationalPath, nil)
}

func (ds *DaemonSuite) TestLabels(c *C) {
	c.Skip("To test, enable a consul environment")
	//Set up last free ID with zero
	kv := ds.d.consul.KV()
	byteJSON, err := json.Marshal(0)
	c.Assert(err, Equals, nil)
	p := &KVPair{Key: common.LastFreeIDKeyPath, Value: byteJSON}
	_, err = kv.Put(p, nil)
	c.Assert(err, Equals, nil)

	kvPair, _, err := kv.Get(common.LastFreeIDKeyPath, nil)
	c.Assert(err, Equals, nil)
	var id int
	err = json.Unmarshal(kvPair.Value, &id)
	c.Assert(err, Equals, nil)
	c.Assert(id, Equals, 0)

	id, err = ds.d.GetLabelsID(lbls)
	c.Assert(err, Equals, nil)
	c.Assert(id, Equals, 0)

	id, err = ds.d.GetLabelsID(lbls)
	c.Assert(err, Equals, nil)
	c.Assert(id, Equals, 0)

	id, err = ds.d.GetLabelsID(lbls2)
	c.Assert(err, Equals, nil)
	c.Assert(id, Equals, 1)

	id, err = ds.d.GetLabelsID(lbls2)
	c.Assert(err, Equals, nil)
	c.Assert(id, Equals, 1)

	id, err = ds.d.GetLabelsID(lbls)
	c.Assert(err, Equals, nil)
	c.Assert(id, Equals, 0)

	//Get labels from ID
	gotLabels, err := ds.d.GetLabels(0)
	c.Assert(err, Equals, nil)
	c.Assert(reflect.DeepEqual(*gotLabels, lbls), Equals, true)

	gotLabels, err = ds.d.GetLabels(1)
	c.Assert(err, Equals, nil)
	c.Assert(reflect.DeepEqual(*gotLabels, lbls2), Equals, true)
}

func (ds *DaemonSuite) TestMaxSetOfLabels(c *C) {
	c.Skip("To test, enable a consul environment")
	//Set up last free ID with common.MaxSetOfLabels - 1
	kv := ds.d.consul.KV()
	byteJSON, err := json.Marshal((common.MaxSetOfLabels - 1))
	c.Assert(err, Equals, nil)
	p := &KVPair{Key: common.LastFreeIDKeyPath, Value: byteJSON}
	_, err = kv.Put(p, nil)
	c.Assert(err, Equals, nil)

	kvPair, _, err := kv.Get(common.LastFreeIDKeyPath, nil)
	c.Assert(err, Equals, nil)
	var id int
	err = json.Unmarshal(kvPair.Value, &id)
	c.Assert(err, Equals, nil)
	c.Assert(id, Equals, (common.MaxSetOfLabels - 1))

	id, err = ds.d.GetLabelsID(lbls)
	c.Assert(err, Equals, nil)
	c.Assert(id, Equals, (common.MaxSetOfLabels - 1))

	_, err = ds.d.GetLabelsID(lbls2)
	c.Assert(strings.Contains(err.Error(), "maximum"), Equals, true)

	id, err = ds.d.GetLabelsID(lbls)
	c.Assert(err, Equals, nil)
	c.Assert(id, Equals, (common.MaxSetOfLabels - 1))
}