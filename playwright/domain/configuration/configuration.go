//go:generate go run github.com/golang/mock/mockgen -source=$GOFILE -destination=mock/$GOFILE -package=mock

package configuration

type Configuration interface {
	Common() *Common
}

type (
	Config struct {
		Common Common
	}

	Common struct {
		Debug        bool   `env:"DEBUG"`
		GCPProjectID string `env:"GCP_PROJECT_ID"`
	}
)
