package main

import (
	"encoding/json"
	gocon "github.com/ahmetask/go-con"
)

type StaticDB struct {
	Configs []gocon.Config
}

type StaticDBConfig struct {
	Content   interface{}
	AppName   string
	NameSpace string
	Port      string
}

func NewStaticDB(configs []StaticDBConfig) *StaticDB {
	cs := make([]gocon.Config, 0)
	for _, c := range configs {
		m, err := json.Marshal(c.Content)
		cfg := gocon.Config{
			Content: m,
			Spec: gocon.Spec{
				AppName:   c.AppName,
				Namespace: c.NameSpace,
				Port:      c.Port,
			},
		}
		if err == nil {
			cs = append(cs, cfg)
		}
	}

	return &StaticDB{
		Configs: cs,
	}
}

func (m *StaticDB) Save(config gocon.Config) error {
	updated := false
	for i := 0; i < len(m.Configs); i++ {
		if m.Configs[i].Spec.AppName == config.Spec.AppName {
			m.Configs[i] = config
			updated = true
		}
	}
	if !updated {
		m.Configs = append(m.Configs, config)
	}
	return nil
}

func (m *StaticDB) Read(appName string) (*gocon.Config, error) {
	var config gocon.Config

	for _, c := range m.Configs {
		if c.Spec.AppName == appName {
			config = c
		}
	}

	return &config, nil
}

func (m *StaticDB) ReadSpecs() ([]gocon.Spec, error) {
	specs := make([]gocon.Spec, 0)
	for _, c := range m.Configs {
		specs = append(specs, c.Spec)
	}
	return specs, nil
}
