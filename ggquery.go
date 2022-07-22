package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

type Client struct {
	ctx     context.Context
	service *admin.Service
}

func New(ctx context.Context, adminEmail string, serviceAccountKey []byte) (*Client, error) {
	config, err := google.JWTConfigFromJSON(serviceAccountKey,
		admin.AdminDirectoryGroupReadonlyScope,
		admin.AdminDirectoryGroupMemberReadonlyScope,
		admin.AdminDirectoryUserReadonlyScope)

	config.Subject = adminEmail

	if err != nil {
		return nil, err
	}

	ts := config.TokenSource(ctx)

	srv, err := admin.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		return nil, err
	}

	return &Client{
		ctx:     ctx,
		service: srv,
	}, nil
}

func (c *Client) GetGroups(query string) ([]*admin.Group, error) {
	g := make([]*admin.Group, 0)
	var err error

	if query != "" {
		err = c.service.Groups.List().Customer("my_customer").Query(query).Pages(context.TODO(), func(groups *admin.Groups) error {
			g = append(g, groups.Groups...)
			return nil
		})
	} else {
		err = c.service.Groups.List().Customer("my_customer").Pages(context.TODO(), func(groups *admin.Groups) error {
			g = append(g, groups.Groups...)
			return nil
		})

	}
	return g, err
}

func Fetch(query string) error {
	creds, err := ioutil.ReadFile("./credentials.json")
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	email := os.Getenv("ADMIN_EMAIL")

	client, err := New(ctx, email, creds)
	if err != nil {
		return err
	}

	fmt.Printf("Searching for groups with query: %s\n", query)
	fmt.Println("----------------------------------")
	groups, err := client.GetGroups(query)
	if err != nil {
		return err
	}

	if len(groups) < 1 {
		return fmt.Errorf("no results found for query: %s", query)
	}

	for _, group := range groups {
		fmt.Printf("name: %s\n", group.Name)
		fmt.Printf("email: %s\n", group.Email)
		fmt.Println("----------------------------------")
	}
	return nil
}

func main() {
	argsWithoutProg := os.Args[1:]

	query := ""
	if len(argsWithoutProg) > 0 {
		query = argsWithoutProg[0]
	}

	err := Fetch(query)
	if err != nil {
		panic(err)
	}
}
