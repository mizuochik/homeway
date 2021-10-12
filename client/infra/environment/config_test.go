package environment_test

import (
	"net/url"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mizuochikeita/homeway/client/infra/environment"
)

func TestConfig(t *testing.T) {
	t.Run("ReadFromEnv()", func(t *testing.T) {
		env := map[string]string{
			"SERVER_URL": "https://homeway-server",
		}
		for k, v := range env {
			os.Setenv(k, v)
		}
		defer func() {
			for k := range env {
				os.Unsetenv(k)
			}
		}()
		got, err := environment.ReadFromEnv()
		if err != nil {
			t.Fatal(err)
		}
		want := &environment.Config{
			ServerURL: &url.URL{
				Scheme: "https",
				Host:   "homeway-server",
			},
		}
		if diff := cmp.Diff(want, got); diff != "" {
			t.Error(diff)
		}
	})
}
