package auth_proxy



// server.go

r := mux.NewRouter()
r.HandleFunc("/login", login.LoginHandler)

r.HandleFunc("/logout", logout.LogoutHandler)