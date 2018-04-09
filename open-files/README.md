# Open To Many Files

### How to use

```
cf push
curl -sS <app-url> | cf logs reader
```

Eventually you will see this line:

```
2018-04-09T14:04:14.91-0700 [APP/0] OUT open ./file.txt: too many open files
```

### To fix you need to modify `main.go`

**Change**
```
func handler(w http.ResponseWriter, r *http.Request) {
	for {
		go readFile()
	}
}
```

**To this**
```
func handler(w http.ResponseWriter, r *http.Request) {
  go readFile()
}
```
