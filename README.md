# Go Lang HelloWorld Based on [flamingo-golang-hello- By John](https://golangexample.com/flamingo-a-go-based-framework-for-pluggable-web-projects/)

1. Start By Running

```go
 go run main.go
```

2. Routes are Defined in `./config/routes.yaml`

```yaml
-path: /
name: index
controller: flaimngo.render(tpl="index")
```
- Note: Make Sure to Add Routes as Per the Offical Go Lang Docs


3. HTML Template Files are structured in  