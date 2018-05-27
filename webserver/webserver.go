package main

import (
  "fmt"
  "log"
  "github.com/gorilla/mux"
  "github.com/gorilla/securecookie"
  "net/http"
  "time"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "github.com/olguncengiz/AppointmentApp/microservice/appointment"
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

func getRole(request *http.Request) (role string) {
  if cookie, err := request.Cookie("session"); err == nil {
    cookieValue := make(map[string]string)
    if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
      role = cookieValue["role"]
    }
  }
  return role
}

func setSession(userName string, role string, response http.ResponseWriter) {
  value := map[string]string{
    "name": userName,
    "role": role,
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

const (
  address     = "localhost:50888"
  defaultName = "world"
  defaultDate = "2018-05-26"
  defaultTime = "13:00"
  defaultStatus = "r"
)

// Appointment Delete Handler
func appointmentDeleteHandler(response http.ResponseWriter, request *http.Request) {
  rName := request.FormValue("username")
  
  redirectTarget := "/adminPanel"
  
  if rName != ""{
    // Set up a connection to the server.
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewAppointmentClient(conn)

    // Contact the server and print out its response.
    //clientDeadline := time.Now().Add(3000)
    //ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    r, err := c.DeleteAppointment(ctx, &pb.ClientInfo{Name: rName})
    if err != nil {
      log.Fatalf("could not delete appointment: %v", err)
    }
    log.Printf("Reply: %s", r.Message)
  }
  http.Redirect(response, request, redirectTarget, 302)
}

// Appointment Approve Handler
func appointmentApproveHandler(response http.ResponseWriter, request *http.Request) {
  rName := request.FormValue("username")
  
  redirectTarget := "/adminPanel"
  
  if rName != ""{
    // Set up a connection to the server.
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewAppointmentClient(conn)

    // Contact the server and print out its response.
    //clientDeadline := time.Now().Add(3000)
    //ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()

    r, err := c.ApproveAppointment(ctx, &pb.ClientInfo{Name: rName})
    if err != nil {
      log.Fatalf("could not approve appointment: %v", err)
    }
    log.Printf("Reply: %s", r.Message)
  }
  http.Redirect(response, request, redirectTarget, 302)
}

// Appointment Request Handler
func appointmentRequestHandler(response http.ResponseWriter, request *http.Request) {
  rName := request.FormValue("username")
  rDate := request.FormValue("date")
  rTime := request.FormValue("time")
  rStatus := "r"
  redirectTarget := "/userPanel"
  
  if rName != "" && rDate != "" && rTime != "" {
    // Set up a connection to the server.
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewAppointmentClient(conn)

    // Contact the server and print out its response.
    //clientDeadline := time.Now().Add(time.Duration(10) * time.Second)
    //ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()
    clientInfo := &pb.ClientInfo{Name: rName}
    appInfo := &pb.AppointmentInfo{Client: clientInfo, Date: rDate, Time: rTime, Status: rStatus}
    r, err := c.RequestAppointment(ctx, &pb.AppointmentReq{AppInfo: appInfo})
    if err != nil {
      log.Fatalf("could not request appointment: %v", err)
    }
    log.Printf("Reply: %s", r.Message)
  }
  http.Redirect(response, request, redirectTarget, 302)
}

// Appointment Move Handler
func appointmentMoveHandler(response http.ResponseWriter, request *http.Request) {
  rName := request.FormValue("username")
  rDate := request.FormValue("date")
  rTime := request.FormValue("time")
  rStatus := "a"
  redirectTarget := "/adminPanel"
  
  if rName != "" && rDate != "" && rTime != "" {
    // Set up a connection to the server.
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
      log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewAppointmentClient(conn)

    // Contact the server and print out its response.
    //clientDeadline := time.Now().Add(time.Duration(10) * time.Second)
    //ctx, cancel := context.WithDeadline(context.Background(), clientDeadline)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
    defer cancel()
    clientInfo := &pb.ClientInfo{Name: rName}
    appInfo := &pb.AppointmentInfo{Client: clientInfo, Date: rDate, Time: rTime, Status: rStatus}
    r, err := c.MoveAppointment(ctx, &pb.AppointmentReq{AppInfo: appInfo})
    if err != nil {
      log.Fatalf("could not move appointment: %v", err)
    }
    log.Printf("Reply: %s", r.Message)
  }
  http.Redirect(response, request, redirectTarget, 302)
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
      setSession(name, userRole, response)
      redirectTarget = "/userPanel"
    } else if userRole == "admin" {
      setSession(name, userRole, response)
      redirectTarget = "/adminPanel"
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
<div>
<b>Request Appointment</b>
<form method="post" action="/requestAppointment">
    <label for="username">User name</label>
    <input type="text" id="username" name="username">
    <label for="date">Date</label>
    <input type="date" id="date" name="date">
    <label for="time">Time</label>
    <input type="time" id="time" name="time">
    <button type="submit">Request</button>
</form>
</div>
<br>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`

// Admin Panel
const adminPanel = `
<h1>Admin Panel</h1>
<hr>
<small>User: %s</small>
<div>
<b>Delete Appointment</b>
<form method="post" action="/deleteAppointment">
    <label for="username">User name</label>
    <input type="text" id="username" name="username">
    <button type="submit">Delete</button>
</form>
</div>
<br>
<div>
<b>Approve Appointment</b>
<form method="post" action="/approveAppointment">
    <label for="username">User name</label>
    <input type="text" id="username" name="username">
    <button type="submit">Approve</button>
</form>
</div>
<br>
<div>
<div>
<b>Move Appointment</b>
<form method="post" action="/moveAppointment">
    <label for="username">User name</label>
    <input type="text" id="username" name="username">
    <label for="date">New Date</label>
    <input type="date" id="date" name="date">
    <label for="time">New Time</label>
    <input type="time" id="time" name="time">
    <button type="submit">Move</button>
</form>
</div>
<br>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`

func userPanelHandler(response http.ResponseWriter, request *http.Request) {
  userName := getUserName(request)
  role := getRole(request)

  if userName != "" && role == "user" {
    fmt.Fprintf(response, userPanel, userName)
  } else {
    http.Redirect(response, request, "/", 302)
  }
}

func adminPanelHandler(response http.ResponseWriter, request *http.Request) {
  userName := getUserName(request)
  role := getRole(request)

  if userName != "" && role == "admin" {
    fmt.Fprintf(response, adminPanel, userName)
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
  router.HandleFunc("/adminPanel", adminPanelHandler)

  router.HandleFunc("/login", loginHandler).Methods("POST")
  router.HandleFunc("/logout", logoutHandler).Methods("POST")
  router.HandleFunc("/requestAppointment", appointmentRequestHandler).Methods("POST")
  router.HandleFunc("/deleteAppointment", appointmentDeleteHandler).Methods("POST")
  router.HandleFunc("/approveAppointment", appointmentApproveHandler).Methods("POST")
  router.HandleFunc("/moveAppointment", appointmentMoveHandler).Methods("POST")

  http.Handle("/", router)
  http.ListenAndServe(":8080", nil)
}