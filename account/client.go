package account

import (
	"context"

	"github.com/TianYao12/microservice/account/pb"
	"google.golang.org/grpc"
)

// client for interacting with account service via grpc
type Client struct {
	connection    *grpc.ClientConn
	service pb.AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	connection, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	service_client := pb.NewAccountServiceClient(connection)
	return &Client{connection, service_client}, nil
}

func (client *Client) Close() {
	client.connection.Close()
}

func (client *Client) PostAccount(context context.Context, name string) (*Account, error) {
	response, err := client.service.PostAccount(
		context,
		&pb.PostAccountRequest{Name: name},
	)
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:   response.Account.Id,
		Name: response.Account.Name,
	}, nil
}

func (client *Client) GetAccount(ctx context.Context, id string) (*Account, error) {
	r, err := client.service.GetAccount(
		ctx,
		&pb.GetAccountRequest{Id: id},
	)
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:   r.Account.Id,
		Name: r.Account.Name,
	}, nil
}

func (client *Client) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	r, err := client.service.GetAccounts(
		ctx,
		&pb.GetAccountsRequest{
			Skip: skip,
			Take: take,
		},
	)
	if err != nil {
		return nil, err
	}
	accounts := []Account{}
	for _, a := range r.Accounts {
		accounts = append(accounts, Account{
			ID:   a.Id,
			Name: a.Name,
		})
	}
	return accounts, nil
}
