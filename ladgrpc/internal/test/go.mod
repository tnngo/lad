module github.com/tnngo/lad/ladgrpc/internal/test

go 1.17

require (
<<<<<<< HEAD:ladgrpc/internal/test/go.mod
	github.com/stretchr/testify v1.8.0
	github.com/tnngo/lad v0.0.0-00010101000000-000000000000
=======
	github.com/stretchr/testify v1.8.1
	go.uber.org/zap v1.16.0
>>>>>>> 87577d85d58b6d92d0158967df29303d04d30e36:zapgrpc/internal/test/go.mod
	google.golang.org/grpc v1.42.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/tnngo/lad => ../../..
