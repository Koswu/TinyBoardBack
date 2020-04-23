package setting

var Server struct{
	HttpPort int
	ReadTimeout int
	WriteTimeout int
}

func loadServer(){
	serverSection := getSection("server")
	Server.HttpPort = serverSection.Key("HTTP_PORT").MustInt(8080)
	Server.ReadTimeout = serverSection.Key("READ_TIMEOUT").MustInt(60)
	Server.WriteTimeout = serverSection.Key("WRITE_TIMEOUT").MustInt(60)
}
