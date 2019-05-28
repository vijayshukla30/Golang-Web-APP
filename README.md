# Golang-Web-APP
* Introduction
* Installation
* Configuration
* Start Application
* API
    * /
    * /register
    * /edit/{id}
    * /delete/{id}
* Packages Used
    * database/sql
    * github.com/go-sql-driver/mysql
    * fmt
    * html/template
    * net/http
    * log
    * github.com/gorilla/mux

# Introduction
This application is basic Sample of CRUD (user) using **Gorilla Mux** for _routing_ and **HTML Templating** using _html/template_.

# Installation
[Download the archive](https://golang.org/dl/) and extract it into /usr/local, creating a Go tree in /usr/local/go. For example: 
```
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```

# Configuration
By Default you Go Project Environment is
```
cd $HOME/go/src/[project]
```
Add below lines for Ubuntu in your **.bashrc** or **.profile**
```
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

# Start Application
Navigate to 
```
cd $HOME/go/src/github.com/{username}/
```

Clone the Application
```
cd [project-name]
```
If you are running the application first time then please install all the packages from **Packages** section.

Example: Install the MUX package
```
go get github.com/gorilla/mux
```

To start the Application:

* Please create a database in MySQL
```
CREATE DATABASE golang;
```

* Create a Table User
```
DROP TABLE IF EXISTS `user`;
CREATE TABLE user(id int(0) unsigned NOT NULL AUTO_INCREMENT, first_name varchar(30) NOT NULL, last_name varchar(30) NOT NULL, email varchar(50) NOT  NULL, PRIMARY KEY (id));
```

Run the application 
```
go run main.go
```

Application will be running on PORT 8080. If you want to change the PORT number then open **main.go** file and search for 
```
http.ListenAndServe(":8080", r)
``` 
and change the port

# API
``/`` will return list of users from DB.

``/register`` can be used to create a user

``/edit{id}`` will be used to edit an user whose id is given the URL

``/delete/{id}`` it will delete the user with the specified id

# Packages

**database/sql, fmt, html/template, net/http, log** are the default packages, installed during go installation.

Since we have used some more packages in our project so there details are below:

* **github.com/go-sql-driver/mysql** For database connectivity. [Read More](https://github.com/go-sql-driver/mysql)
* **github.com/gorilla/mux**  To handle our routes. [Read More](https://github.com/gorilla/mux)
* **github.com/pilu/fresh** It is used for hot-reloading. [Read More](https://github.com/gravityblast/fresh)



