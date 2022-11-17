package config

import "time"

type Runner struct {
	NormalPeriod      time.Duration `fig:"normal_period"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period"`
	MaxAbnormalPeriod time.Duration `fig:"normal_period"`
}

const defaultRunnerPeriod = time.Minute

var defaultRunner = Runner{
	NormalPeriod:      defaultRunnerPeriod,
	MinAbnormalPeriod: defaultRunnerPeriod,
	MaxAbnormalPeriod: defaultRunnerPeriod,
}
