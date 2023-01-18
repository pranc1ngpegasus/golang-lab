package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v48/github"
	loghttp "github.com/motemen/go-loghttp"
	"github.com/spf13/viper"
)

type Config struct {
	AppID            int64  `mapstructure:"GH_APP_ID"`
	InstallationID   int64  `mapstructure:"GH_APP_INSTALLATION_ID"`
	PrivateKeyName   string `mapstructure:"GH_APP_PRIVATE_KEY_NAME"`
	OrganizationName string `mapstructure:"GH_ORG_NAME"`
}

func main() {
	viper.SetConfigFile("sample.env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	ctx := context.Background()

	smclient, err := secretmanager.NewClient(ctx)
	if err != nil {
		panic(err)
	}

	defer func() {
		smclient.Close()
	}()

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: config.PrivateKeyName,
	}
	result, err := smclient.AccessSecretVersion(ctx, req)
	if err != nil {
		panic(err)
	}

	transport, err := ghinstallation.New(
		&loghttp.Transport{},
		config.AppID,
		config.InstallationID,
		result.Payload.Data,
	)
	if err != nil {
		panic(err)
	}
	client := github.NewClient(&http.Client{Transport: transport})

	members, _, err := client.Organizations.ListMembers(ctx, config.OrganizationName, &github.ListMembersOptions{})
	if err != nil {
		panic(err)
	}

	val, err := json.MarshalIndent(members, "", "    ")
	if err != nil {
		panic(err)
	}

	log.Default().Printf("%s", val)
}
