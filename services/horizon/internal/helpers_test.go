package horizon

import (
	"time"

	"github.com/stellar/throttled"

	"github.com/stellar/go/network"
	"github.com/stellar/go/services/horizon/internal/test"
	supportLog "github.com/stellar/go/support/log"
)

func NewTestApp() *App {
	return NewApp(NewTestConfig())
}

func NewTestConfig() Config {
	return Config{
		DatabaseURL:            test.DatabaseURL(),
		StellarCoreDatabaseURL: test.StellarCoreDatabaseURL(),
		RateQuota: &throttled.RateQuota{
			MaxRate:  throttled.PerHour(1000),
			MaxBurst: 100,
		},
		ConnectionTimeout: 55 * time.Second, // Default
		LogLevel:          supportLog.InfoLevel,
		NetworkPassphrase: network.TestNetworkPassphrase,
	}
}

func NewRequestHelper(app *App) test.RequestHelper {
	return test.NewRequestHelper(app.web.router)
}
