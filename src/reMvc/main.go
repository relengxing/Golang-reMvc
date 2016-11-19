package main

import (
    "reMvc/route"
    "reMvc/models"
    "fmt"
    "reMvc/models/db/mysql"
)

func main() {
    mysql.Connect("root","1234","user")
    route.StaticRoute.SetStaticRouteDir("/css/","static")
    route.StaticRoute.SetStaticRouteDir("/js/","static")
    route.FixedRoute.SetFixedRouteDir("/404",route.NotfoundHandler)
    fmt.Println(models.GetPageUser(0))
    //route.Route.Start(":9696")
}
