go test -v -coverprofile test/coverage/cover.out ./... && \
go tool cover -html=test/coverage/cover.out -o test/coverage/cover.html && \
open test/coverage/cover.html