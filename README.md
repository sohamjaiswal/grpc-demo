# Grpc Demo

## Steps to run

0. install deps using `go mod tidy`
1. rename `.sample.dev.env` to `.dev.env`
2. run server using `go run ./server`
3. You can test using any grpc client or included client
   - integrated client tests using a sample of 3 names, you can change them in client src
   - apidog is a pretty good client too in my experience
