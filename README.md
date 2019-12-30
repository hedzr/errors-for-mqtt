# Demo Codes for `errors` package

Here is a sample about how to extend error object on [`github.com/hedzr/errors`](https://github.com/hedzr/errors) and standard library `errors`.

Minimal copy from `hedzr/errors` and stdlib `errors`:

- Duplicate these files from `errors-for-mqtt`:
  1. can.go: simple copy
  2. new.go: copy and modify the return type
  3. publics.go: simple copy
  4. wrap.go: simple copy
- Implements custom error struct youself, refer to `MqttError`
- Declares and implements custom error codes, refer to `CloseReason`


