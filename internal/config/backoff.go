package config

import "time"

type BackoffSettings struct {
	NormalPeriod      time.Duration `fig:"normal_period"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period"`
	MaxAbnormalPeriod time.Duration `fig:"max_abnormal_period"`
}

const defaultBackoffPeriod = time.Minute

var defaultBackoffSettings = BackoffSettings{
	NormalPeriod:      defaultBackoffPeriod,
	MinAbnormalPeriod: defaultBackoffPeriod,
	MaxAbnormalPeriod: defaultBackoffPeriod,
}
