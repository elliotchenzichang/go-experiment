package temporal

import "go.temporal.io/sdk/client"

func InitATemporalClient() (client.Client, error) {
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	return c, err
}
