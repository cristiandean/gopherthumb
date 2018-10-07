package context

type Server struct {
	Port string
	Ip   string
}

type MainConfig struct {
	Server Server
}


var Config = &MainConfig{
	Server: Server{
		Port: "8000",
		Ip: "0.0.0.0",
	},
}