package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "github.com/gorilla/securecookie"
  "net/http"
)

// Cookie Handling
var cookieHandler = securecookie.New(
  securecookie.GenerateRandomKey(64),
  securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
  if cookie, err := request.Cookie("session"); err == nil {
    cookieValue := make(map[string]string)
    if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
      userName = cookieValue["name"]
    }
  }
  return userName
}

func setSession(userName string, response http.ResponseWriter) {
  value := map[string]string{
    "name": userName,
  }
  if encoded, err := cookieHandler.Encode("session", value); err == nil {
    cookie := &http.Cookie{
      Name:  "session",
      Value: encoded,
      Path:  "/",
    }
    http.SetCookie(response, cookie)
  }
}

func clearSession(response http.ResponseWriter) {
  cookie := &http.Cookie{
    Name:   "session",
    Value:  "",
    Path:   "/",
    MaxAge: -1,
  }
  http.SetCookie(response, cookie)
}

func authenticateUser(username string, password string) string {
  role := ""
  if (username == "user1" && password == "user") ||
    (username == "user2" && password == "user") {
      role = "user"
  } else if username == "admin" && password == "admin" {
    role = "admin"
  }
  return role
}

// Login Handler
func loginHandler(response http.ResponseWriter, request *http.Request) {
  name := request.FormValue("username")
  pass := request.FormValue("password")
  redirectTarget := "/"
  if name != "" && pass != "" {
    // Authentication
    // TO-DO: This can be improved, like a database of users and credentials...
    userRole := authenticateUser(name, pass)
    if userRole == "user" {
      setSession(name, response)
      redirectTarget = "/userPanel"
    } else if userRole == "admin" {
      setSession(name, response)
      redirectTarget = "/internal"
    }
  }
  http.Redirect(response, request, redirectTarget, 302)
}

// Logout Handler
func logoutHandler(response http.ResponseWriter, request *http.Request) {
  clearSession(response)
  http.Redirect(response, request, "/", 302)
}

// Index Page
const indexPage = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="username">User name</label>
    <input type="text" id="username" name="username">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
  fmt.Fprintf(response, indexPage)
}

// Internal Page
const internalPage = `
<h1>Internal</h1>
<hr>
<small>User: %s</small>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`

func internalPageHandler(response http.ResponseWriter, request *http.Request) {
  userName := getUserName(request)
  if userName != "" {
    fmt.Fprintf(response, internalPage, userName)
  } else {
    http.Redirect(response, request, "/", 302)
  }
}

// User Panel
const userPanel = `
<h1>User Panel</h1>
<hr>
<small>User: %s</small>
<div>Appointment Request Will Come Here...</div>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`

func userPanelHandler(response http.ResponseWriter, request *http.Request) {
  userName := getUserName(request)
  if userName != "" {
    fmt.Fprintf(response, userPanel, userName)
  } else {
    http.Redirect(response, request, "/", 302)
  }
}

// Main Method
var router = mux.NewRouter()

func main() {

  router.HandleFunc("/", indexPageHandler)
  router.HandleFunc("/internal", internalPageHandler)
  router.HandleFunc("/userPanel", userPanelHandler)

  router.HandleFunc("/login", loginHandler).Methods("POST")
  router.HandleFunc("/logout", logoutHandler).Methods("POST")

  http.Handle("/", router)
  http.ListenAndServe(":8080", nil)
}