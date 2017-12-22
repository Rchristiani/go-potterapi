package potterapi

//BaseURL for the potterapi
const BaseURL = "https://www.potterapi.com/v1/"

//Client used to store information for potterapi
type Client struct {
	apiKey string
}

//ClientInitialize used to create a new Client
func ClientInitialize(key string) *Client {
	c := &Client{
		apiKey: key,
	}
	return c
}
