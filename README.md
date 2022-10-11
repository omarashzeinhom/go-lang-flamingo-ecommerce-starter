# Go Lang HelloWorld Based on [flamingo-golang-hello- By John](https://golangexample.com/flamingo-a-go-based-framework-for-pluggable-web-projects/)

1. Start By Running

```go
 go run main.go
```

## Steps to Recreate

- Install Pacakges in goLang By using ```go get pacakge```

```go
go get flamingo.me/dingo
go get flamingo.me/flamingo/v3
go get flamingo.me/flamingo/v3/core/gotemplate
go get flamingo.me/flamingo/v3/core/requestlogger
go get flamingo.me/flamingo/v3/core/zap
```

-Remove Pacakges in Golang By using```go get package@none```

### Project Strcuture based on [Flaming-Hello-World-Folder Strcutre](https://docs.flamingo.me/1.%20Introduction/2.%20Tutorial%20Hello%20World.html)

`./config` Directory For Confgiuration and routing files.

`./src` Directory For Source Code of The Project.

`./templates` Directory For Go Templates.

`go.mod/` The Projects Dependencies .

`main.go` Entry point for our project.

#### Routing

Note: Routes are Defined in `./config/routes.yaml`

```yaml
-path: /
name: index
controller: flaimngo.render(tpl="index")
```

- Note: Make Sure to Add Routes as Per the Offical Go Lang Docs

##### Steps

1. In `./config` Create a New File call it `config.yaml`

2. In `./config` Create a New File For Routes call it `routes.yaml` - Note : There are some built in controllers for `routes.yaml` [Built In Controllers for routes.yaml](https://docs.flamingo.me/1.%20Introduction/2.%20Tutorial%20Hello%20World.html#builtin-controllers)

3. And in `./templates` Create a New `index.html` to display Basic Content

```html
<!DOCTYPE html>
<html lang="en">
  <head> </head>

  <body>
    <h1>ANDGOEDU ECOMMERCE IN GO LANG</h1>
  </body>
</html>
```

4. Create a Custom Controller for `routes.yaml`
Add to  ```import ("flamingo.me/go-commerce-andogedu/src/helloworld")``` in `./main.go`
Add ```new(helloworld.Module),``` to ```flamingo.App([]dingo.Module{})```in `./main.go` - MUST ADD THE .dingo.Module in the main.go to load custom modules.

```go
//1. Add to import 
 import ( "flamingo.me/go-commerce-andogedu/src/helloworld")
//2. Add to func main in `./main.go`
 func main (
 flamingo.App([]dingo.Module){
  new(helloworld.Module),
 })
```

5. A controller in flamingo is  a ```struct{}``` with a couple of methods , use such as ,,Actions.

- Controllers can use ```web.Responder``` , also the ```http.Handler``` from GoLang `stdlib`.
  
#### References

1. [go-dev](https://go.dev/doc/install)
2. [flamingo-commerce](https://www.flamingo.me/flamingo-commerce.html#Start)
3. [golang- manage dependencies](https://go.dev/doc/modules/managing-dependencies#naming_module)
4. [What does undefined : fmt.Println(build) means in GOlang? while using "fmt.Println(len("hello world")"](https://stackoverflow.com/questions/46216071/what-does-undefined-fmt-printlnbuild-means-in-golang-while-using-fmt-print)
5. [Removing packages installed with go get](https://stackoverflow.com/questions/13792254/removing-packages-installed-with-go-get)
