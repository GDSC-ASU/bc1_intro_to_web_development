# Day 12:

### React Router
Install React's Router by running

```bash
npm install react-router-dom 
```

inside the client's directory


### Middleware definition example:

```go

// handler is the rest api handler that we want to do a process before the request reaches it
func middlewareName(handler http.HandlerFunc) http.HandlerFunc {
    // return a rest api handler function that has the processing
    // where every statement before `handler.ServeHTTP(resp, req)` is concidered as middleware processing
	return func(resp http.ResponseWriter, req *http.Request) {
        // here we're setting some required http headers so that fetch can work properly
		resp.Header().Set("Access-Control-Allow-Origin", "*")
		resp.Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Authorization")
		resp.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		resp.Header().Set("Content-Type", "application/json")

		if req.Method != http.MethodOptions {
			handler.ServeHTTP(resp, req)
		}
	}
}
```
