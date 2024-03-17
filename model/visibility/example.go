package visibility

type OperatorRuleConfig struct {
	externConfig
}

type externConfig struct {
	windowSize int64 `json:"window_size"`
	count      int64 `json:"count"`
	continuous bool  `json:"continuous"`
}

func (e externConfig) GetWindowSize() int64 {
	return e.windowSize
}

func (e externConfig) GetCount() int64 {
	return e.count
}

func (e externConfig) GetContinuous() bool {
	return e.continuous
}
