package kaeufer

type leadConfig struct {
	PerPage       int
	FromTimestamp int
}

// PerPage limits the number of Lead results returned from Leads.
func PerPage(count int) func(*leadConfig) {
	return func(lc *leadConfig) {
		lc.PerPage = count
	}
}

// FromTimestamp limits the results to those created after the given Timestamp.
func FromTimestamp(ts int) func(*leadConfig) {
	return func(lc *leadConfig) {
		lc.FromTimestamp = ts
	}
}

func loadConfig(options []func(*leadConfig)) leadConfig {
	config := leadConfig{
		PerPage:       10,
		FromTimestamp: -1,
	}

	for _, cfg := range options {
		cfg(&config)
	}

	return config
}
