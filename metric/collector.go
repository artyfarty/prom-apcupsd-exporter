package metric

import (
	"local/apcupsd_exporter/apcupsd"
	"local/apcupsd_exporter/model"
	"os/exec"
	"time"

	promLog "github.com/prometheus/common/log"
)

// Collector ..
type Collector struct {
	ApcupsdAddr         string
	ApcaccessPath       string
	ApcaccessFloodLimit time.Duration
	CollectInterval     time.Duration

	initialized  bool
	collectCh    chan CollectOpts
	currModel    *model.Model
	lastOutput   *apcupsd.Output
	lastOutputTs int64
}

// Init method
func (c *Collector) Init() {
	if !c.initialized {
		c.initialized = true
		c.collectCh = make(chan CollectOpts)
		c.currModel = model.NewModel()
		c.lastOutput = apcupsd.NewOutput("")
		c.lastOutput.Parse()

		go c.loopCollect()
		go c.listenCollect()
	}
}

// GetModel method
func (c *Collector) GetModel() *model.Model {
	return c.currModel
}

// GetLastOutput method
func (c *Collector) GetLastOutput() *apcupsd.Output {
	return c.lastOutput
}

// Collect method
func (c *Collector) Collect(opts CollectOpts) {
	promLog.Infoln("Collect")

	c.collectCh <- opts
}

func (c *Collector) loopCollect() {
	for {
		promLog.Infoln("loopCollect")

		c.Collect(CollectOpts{
			PreventFlood: true,
		})
		time.Sleep(c.CollectInterval)
	}
}

func (c *Collector) listenCollect() {
	promLog.Infoln("listenCollect")
	for {
		if opts, ok := <-c.collectCh; ok {
			promLog.Infoln("listenCollect OK")
			c.collect(opts)
		} else {
			promLog.Infoln("listenCollect FAIL")
			return
		}
	}
}

func (c *Collector) collect(opts CollectOpts) {
	promLog.Infoln("collect()")

	c.updateOutput(opts)
	c.updateModel(opts)
	c.updateMetrics(opts)

	if opts.OnComplete != nil {
		opts.OnComplete <- true
	}
}

func (c *Collector) updateOutput(opts CollectOpts) {
	promLog.Infoln("updating apcupsd output..")

	ts := time.Now().UnixNano()
	if opts.PreventFlood && ts-c.lastOutputTs < int64(c.ApcaccessFloodLimit) {
		return
	}
	c.lastOutputTs = ts

	cmdResult, err := exec.Command(c.ApcaccessPath, "status", c.ApcupsdAddr).Output()
	if err != nil {
		promLog.Errorln("apcaccess exited with error")
		promLog.Errorln("  Error:", err.Error())
		promLog.Errorln("  Result:", string(cmdResult))
		cmdResult = []byte{}
	}

	c.lastOutput = apcupsd.NewOutput(string(cmdResult))
	c.lastOutput.Parse()
}

func (c *Collector) updateModel(opts CollectOpts) {
	promLog.Infoln("updating model..")

	c.currModel.Update(model.NewStateFromOutput(c.lastOutput))

	for field, diff := range c.currModel.ChangedFields {
		promLog.Infof("field changed '%s'\n  OLD: %#v\n  NEW: %#v\n", field, diff[0], diff[1])
	}
}

func (c *Collector) updateMetrics(opts CollectOpts) {
	promLog.Infoln("updating metrics..")
	for _, metric := range Metrics {
		if metric.HandlerFunc != nil {
			metric.HandlerFunc(metric, c.currModel)
		} else if metric.ValFunc != nil {
			metric.UpdateCollector(metric.ValFunc(metric, c.currModel))
		}
	}
	promLog.Infoln("metrics updated")
}

// CollectOpts ..
type CollectOpts struct {
	PreventFlood bool
	OnComplete   chan bool
}