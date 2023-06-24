package apps

func SetRouter(routerApp, RouterPath, Appkey string) {
	routers[routerApp] = RouterPath
	routersUuid[routerApp] = Appkey
}

func GetRouter(routerApp string) string {
	return routers[routerApp]
}

func getRouerUuid(routerApp string) string {
	return routersUuid[routerApp]
}
