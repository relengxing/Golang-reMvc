package route

import "net/http"

type Route struct {
    
}

func register()  {
    StaticRoute.RegisteStaticRoute()
    FixedRoute.RegisteFixedRoute()
    RESTfulRoute.RegisteRESTfulRoute()
}

func Start(port string)  {
    register()
    http.ListenAndServe(port,nil)
}