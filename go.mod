module github.com/tmc/langchaingo

go 1.20

require (
	github.com/google/uuid v1.3.0
	github.com/stretchr/testify v1.8.2
)

replace github.com/nomic-ai/gpt4all/gpt4all-bindings/golang => ../../nomic-ai/gpt4all/gpt4all-bindings/golang // replace when PR merged

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/google/go-cmp v0.5.9
	github.com/nomic-ai/gpt4all/gpt4all-bindings/golang v0.0.0-20230519014017-914519e772fd
	github.com/pinecone-io/go-pinecone v0.3.0
	go.starlark.net v0.0.0-20230302034142-4b1e35fe2254
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)
