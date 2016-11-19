package route

import "net/http"

type rESTfulRoute struct {
    rESTfulRouteList map[string]RESTfulDefault
}

var RESTfulRoute rESTfulRoute
func init() {
    RESTfulRoute = &rESTfulRoute{}
    RESTfulRoute.rESTfulRouteList = make(map[string]func(http.ResponseWriter, *http.Request))
}

func (this *rESTfulRoute)SetRESTfulRouteDir(url string,restful RESTfulDefault)  {
    this.rESTfulRouteList[url] = restful
}

func (this *rESTfulRoute)RegisteRESTfulRoute()  {
    for k, v := range this.rESTfulRouteList {
        http.HandleFunc(k,v.Transfer)
    }
}


