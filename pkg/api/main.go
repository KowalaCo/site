package api

func Run() {
	ReadConfig()
	ConnectDatabase()
	StartRouter()
}
