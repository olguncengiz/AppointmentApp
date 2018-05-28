# AppointmentApp
A lawyer appointment management application written in Go.

# General Architecture
The overall architecture of this application is defined as below:

JS Web Client -> Golang webserver -> Golang gRPC Microservice

![Overall Architecture](https://raw.githubusercontent.com/olguncengiz/AppointmentApp/master/Architecture.png)

# Usage
- Start Golang gRPC server by executing these commands under "/microservice" folder: (Windows commands)
$ go build
$ server.exe

or 

$ go run server.go

- Start Golang web server by executing these commands under "/webserver" folder: (Windows commands)
$ go build
$ webserver.exe

or

$ go run webserver.go

- Open a browser and go to "http://localhost:8080"
- Login using username/password pairs below:
	* user1/user
	* user2/user
- Create an appointment request by filling the format
- Logout
- Login using username/password pair below:
	* admin/admin
- Decline, move or approve the appointment request

# Assumptions
- User registration will be ignored.
- Initially, there will be 2 users and 1 admin account in the system:
  + Username: user1, Password: user, Role: user
  + Username: user2, Password: user, Role: user
  + Username: admin, Password: admin, Role: admin
- Date and Time information is used as string. So, date format should be YYYY-MM-DD, like 2018-05-29 and time format should be HH:MM, like 15:00

# Learning Curves
Below are some of the obstacles I had to overcome for this task:
- Learning about Golang syntax
- Learning about Golang project folder structure (Best Practices)
- Understanding Mutex Concepts
- Understanding Protocol Buffers
- Learning about Proto3 syntax
- Finding the right command to compile protocol buffer file: 
  + "protoc --go_out=plugins=grpc:. ./*.proto"

# Used libraries
- Gorilla MUX
- Gorilla SecureCookie
- gRPC  

# Resources
Below are online resources I used during this project:
- Login/Logout: https://gist.github.com/mschoebel/9398202
- Proto File Compile: https://medium.com/namely-labs/go-protocol-buffers-57b49e28bc4a
- Golang Syntax and Basics: https://gobyexample.com/
- Golang Mux Templates: https://github.com/meshhq/golang-html-template-tutorial
- Golang Mux Templates: https://meshstudio.io/blog/2017-11-06-serving-html-with-golang/

# Problems
Below are the problems I faced during this project:
- Lost too much time on "protoc" tool error (Couldn't find the .exe file to compile the protocol buffer files)
- Working with pointers and pointer syntax (* and &) after a long time (since university) is a little confusing
- Cross-browser usage can mess the date and time since they are stored as strings
- Struggling with static files and HTML templates took too much time, so the web server generates the HTML files from code for now

# Future Improvements
Due to time limitations, below improvement points are left untouched. But these can be considered in future releases:
- User panel and Admin panel HTMLs are hard-coded in .go files and HTML files are generated from the code. Instead, the templates can be used for flexibility
- User registration can be added
- Appointments table is generated in .go file. Instead, AngularJS or similar tool can be used to generate the table
- CSS files can be used for better user interface and flexibility

# Lessons Learnt
- You need to be in the same directory with "webserver.go" when executing "go run webserver.go" command if you want Gorilla MUX's server static files option
