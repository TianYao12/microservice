package main

/*
type Server struct {
	accountClinet		*account.Client
	catalogClient		*catalog.Client
	orderClient			*order.Client
}
*/

func NewGraphQLServer(accountUrl, catalogUrl, orderUrl string) (*Server, error) {
	accountClient, err := account.NewClient(accountUrl)
	if err != nil {
		
	}
}