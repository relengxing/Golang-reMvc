package models

import (
    "fmt"
    "reMvc/models/db/mysql"
)

type User struct {
    Uid int
    Username string
    Password string
}

func GetAllUser() []User {
    list := make([]User,0)
    rows, err := mysql.Db.Query("SELECT * FROM user;")
    if err != nil {
        fmt.Println("fetech data failed:", err.Error())
        return nil
    }
    defer rows.Close()
    for rows.Next() {
        var uid int
        var name, password string
        rows.Scan(&uid, &name, &password)
        list = append(list,User{Uid:uid,Username:name,Password:password})
    }
    return list
}

func GetPageUser(page int) []User {
    list := make([]User,0)
    rows, err := mysql.Db.Query("SELECT * FROM user limit ?,?;",5*page,5*page+5)
    if err != nil {
        fmt.Println("fetech data failed:", err.Error())
        return nil
    }
    defer rows.Close()
    for rows.Next() {
        var uid int
        var name, password string
        rows.Scan(&uid, &name, &password)
        list = append(list,User{Uid:uid,Username:name,Password:password})
    }
    return list
}
