package client

type Client interface {
	Connect(hst string, prt int)
}
