package discovery

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/common/model"
	promDiscovery "github.com/prometheus/prometheus/discovery"
	"github.com/prometheus/prometheus/discovery/targetgroup"
	"github.com/sapcc/ipmi_sd/pkg/adapter"
	"github.com/sapcc/ipmi_sd/pkg/config"
	"github.com/sapcc/ipmi_sd/pkg/netbox"
	"github.com/sapcc/ipmi_sd/pkg/writer"
)

type (
	SwitchDiscovery struct {
		manager         *promDiscovery.Manager
		adapter         adapter.Adapter
		netbox          *netbox.Netbox
		region          string
		refreshInterval int
		logger          log.Logger
		status          *Status
		outputFile      string
		cfg             switchConfig
	}
	switchConfig struct {
		RefreshInterval int      `yaml:"refresh_interval"`
		NetboxHost      string   `yaml:"netbox_host"`
		NetboxAPIToken  string   `yaml:"netbox_api_token"`
		TargetsFileName string   `yaml:"targets_file_name"`
		Switches        []device `yaml:"switches"`
	}
	device struct {
		Name         string `yaml:"name"`
		Manufacturer string `yaml:"manufacturer"`
	}
)

func init() {
	Register("switch", NewSwitchDiscovery)
}

func NewSwitchDiscovery(disc interface{}, ctx context.Context, m *promDiscovery.Manager, opts config.Options, w writer.Writer, l log.Logger) (d Discovery, err error) {
	var cfg switchConfig
	if err := UnmarshalHandler(disc, &cfg); err != nil {
		return nil, err
	}
	fmt.Println(cfg)
	nClient, err := netbox.New(cfg.NetboxHost, cfg.NetboxAPIToken)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return d, err
	}

	a := adapter.NewPrometheus(ctx, m, cfg.TargetsFileName, w, l)

	return &SwitchDiscovery{
		adapter:         a,
		manager:         m,
		netbox:          nClient,
		region:          opts.Region,
		refreshInterval: cfg.RefreshInterval,
		logger:          l,
		status:          &Status{Up: false},
		outputFile:      cfg.TargetsFileName,
		cfg:             cfg,
	}, err

}

func (sd *SwitchDiscovery) Run(ctx context.Context, ch chan<- []*targetgroup.Group) {
	for c := time.Tick(time.Duration(sd.refreshInterval) * time.Second); ; {
		tgs, err := sd.getSwitches()
		if err == nil {
			sd.status.Lock()
			sd.status.Up = true
			sd.status.Unlock()
			ch <- tgs
		} else {
			level.Error(log.With(sd.logger, "component", "SwitchDiscovery")).Log("err", err)
			sd.status.Lock()
			sd.status.Up = false
			sd.status.Unlock()
			continue
		}
		// Wait for ticker or exit when ctx is closed.
		select {
		case <-c:
			continue
		case <-ctx.Done():
			return
		}
	}
}

func (sd *SwitchDiscovery) getSwitches() (tgroups []*targetgroup.Group, err error) {

	for _, device := range sd.cfg.Switches {
		tg, err := sd.loadSwitches(device)
		if err != nil {
			return tgroups, err
		} else {
			tgroups = append(tgroups, tg...)
		}
	}
	return tgroups, err
}

func (sd *SwitchDiscovery) loadSwitches(d device) (tgroups []*targetgroup.Group, err error) {
	devices, err := sd.netbox.DevicesByRegion(d.Name, d.Manufacturer, sd.region)
	if err != nil {
		return tgroups, err
	}

	for _, device := range devices {
		if device.PrimaryIP != nil {
			if err != nil {
				level.Error(log.With(sd.logger, "component", "SwitchDiscovery")).Log("cannot find ip address of switch %d. Error: %s", device.ID, err)
				continue
			}
		}

		if strings.ToUpper(*device.Status.Label) != "ACTIVE" {
			level.Debug(log.With(sd.logger, "component", "SwitchDiscovery")).Log("Ignoring device: %s. Status: %s", device.ID, *device.Status.Label)
			continue
		}

		tgroup := targetgroup.Group{
			Source:  *device.PrimaryIP.Address,
			Labels:  make(model.LabelSet),
			Targets: make([]model.LabelSet, 0, 1),
		}

		target := model.LabelSet{model.AddressLabel: model.LabelValue(*device.PrimaryIP.Address)}
		labels := model.LabelSet{
			model.LabelName("job"):          model.LabelValue(fmt.Sprintf("switch_%s/netbox", d.Name)),
			model.LabelName("server_name"):  model.LabelValue(device.Name),
			model.LabelName("state"):        model.LabelValue(*device.Status.Label),
			model.LabelName("manufacturer"): model.LabelValue(*device.DeviceType.Manufacturer.Name),
			model.LabelName("model"):        model.LabelValue(*device.DeviceType.Model),
			model.LabelName("server_id"):    model.LabelValue(strconv.Itoa(int(device.ID))),
		}

		if device.Serial != "" {
			labels[model.LabelName("serial")] = model.LabelValue(device.Serial)
		}

		tgroup.Labels = labels
		tgroup.Targets = append(tgroup.Targets, target)
		tgroups = append(tgroups, &tgroup)
	}
	return tgroups, err
}

func (sd *SwitchDiscovery) Up() bool {
	return sd.status.Up

}
func (sd *SwitchDiscovery) Lock() {
	sd.status.Lock()

}
func (sd *SwitchDiscovery) Unlock() {
	sd.status.Unlock()
}

func (sd *SwitchDiscovery) GetOutputFile() string {
	return sd.outputFile
}

func (sd *SwitchDiscovery) StartAdapter() {
	sd.adapter.Run()
}

func (sd *SwitchDiscovery) GetAdapter() adapter.Adapter {
	return sd.adapter
}

func (sd *SwitchDiscovery) GetManager() *promDiscovery.Manager {
	return sd.manager
}
