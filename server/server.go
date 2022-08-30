package server

func ServerInit() {
	router := routerInit()
	router.Run(":8080")
}
