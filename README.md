What's that? Why?
==========
Inspired from tools like puppet or ansible the idea is to give ops to model their systems with a simple API. The aim is, that you only one binary to your target system and execute it. All configuration files, templates, etc will be included in this binary(!). This should easen the pain with any runtime dependend things like python, shared c libraries 

Just "let it fall" on your target system

NOTE: I'm still learning go - I think a lot of things could be expressed better. But I prefer a clean and simple style. 

Usage
==========
First of all as a platform go is used. Go is ideal for the aim of this experiment. So you're gonna need the go toolsuite

Because you're gonna embed all your configs, templates etc into your binary, you'll need packr2

```
go get -u github.com/gobuffalo/packr/v2/pack
```

This will install the packr2 binary.

Now create following project(e.g. "my-setuper") structure:

```

```

Run `go get ` for basic infrastructure library of setuper

Create a file main.go
```
package main

import (
	github.com/ui-kreinhard/go-setuperapt"
	github.com/ui-kreinhard/go-setuperexecutor"
)

func main() {
    executor.NewExecutor().
        Plan("Copy important file",
			files.CopyDeferred("important", "/etc/important"))).
		Plan("Add user myuser",
			userGroups.CreateUserWithoutPasswordDeferred("myuser")).
		Plan("Ensure group autologin exists",
			userGroups.AddGroupDeferred("mygroup")).
		Plan("Add user pi to group autologin",
			userGroups.AddUserToGroupDeferred("myuser", "mygroup")).
		Plan("Update Repo",
			apt.CheckForUpdatesDeferred()).
        Run()
}
```
Create a file staticAssets/file/important
```
Hello world 42 1337
```


Run in your project
```
   packr2
   go build
```

Execute the resulting binary on your target system

Current Modules
====

* apt
* files
* templates
* scripts
* systemd
* userGroups

Note that all modules can be very buggy. It worked on my maschine :)

Status
======
Highly experimental - api will definetly change. Currently only tested with debian 10.2. No support/testing for other distributions like Centos or fedora currently planned. I don't like them - so you have to implement package 