package route

import "net/http"

type rESTfulRoute struct {
    fixedRouteList map[string]func(http.ResponseWriter, *http.Request)
}
var FixedRoute rESTfulRoute
func init() {
    FixedRoute = &rESTfulRoute{}
    FixedRoute.fixedRouteList = make(map[string]func(http.ResponseWriter, *http.Request))
}

func (this *rESTfulRoute)SetFixedRouteDir(url string,handle func(http.ResponseWriter, *http.Request))  {
    this.fixedRouteList[url] = handle
}

func (this *rESTfulRoute)RegisteFixedRoute()  {
    for k, v := range this.fixedRouteList {
        http.HandleFunc(k, v)
    }
}
