# AppointmentApp
A lawyer appointment management application written in Go.

# General Architecture
The overall architecture of this application is defined as below:

JS Web Client -> Golang webserver -> Golang gRPC Microservice

# Assumptions
- User registration will be ignored.
- Initially, there will be 2 users and 1 admin account in the system:
  + Username: user1, Password: user, Role: user
  + Username: user2, Password: user, Role: user
  + Username: admin, Password: admin, Role: admin

# Learning Curves
Below are some of the obstacles I had to overcome for this task:
- Learning about Golang syntax
- Learning about Golang project folder structure (Best Practices)
- Understanding Mutex Concepts
- Understanding Protocol Buffers
- Finding the right command to compile protocol buffer file: 
  + "protoc --go_out=plugins=grpc:. ./*.proto"

# Problems
Below are the problems I faced during this project:
- Lost too much time on "protoc" tool error (Couldn't find the .exe file to compile the protocol buffer files)
- Working with pointers and pointer syntax (* and &) after a long time (since university) is a little confusing

# Used libraries
- Gorilla MUX
- Gorilla SecureCookie
- gRPC

# Resources
Below are online resources I used during this project:
-Login/Logout: https://gist.github.com/mschoebel/9398202
- Proto File Compile: https://medium.com/namely-labs/go-protocol-buffers-57b49e28bc4a
- Golang Syntax and Basics: https://gobyexample.com/
- Golang Mux Templates: https://github.com/meshhq/golang-html-template-tutorial
- Golang Mux Templates: https://meshstudio.io/blog/2017-11-06-serving-html-with-golang/

# Lessons Learnt
- You need to be in the same directory with "webserver.go" when executing "go run webserver.go" command if you want Gorilla MUX's server static files option
