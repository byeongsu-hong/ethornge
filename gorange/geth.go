package gorange

type WSConfig struct {
	Addr    string
	Port    string
	API     string
	Origins string
}

type RPCConfig struct {
	Addr string
	Port string
	API  string
}

type GethOption struct {
	DataDir  string
	Keystore string
	WS       *WSConfig
	RPC      *RPCConfig
}
