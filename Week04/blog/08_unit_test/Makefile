setup:
	wire ./...
	make gen_proto

gen_proto:
	protoc -I ./01_example/api -I ./01_example/third_party \
	--openapiv2_out ./01_example/api --openapiv2_opt logtostderr=true --openapiv2_opt json_names_for_fields=false \
	--go_out ./01_example/api --go_opt=paths=source_relative \
	--go-gin_out ./01_example/api --go-gin_opt=paths=source_relative \
	01_example/api/product/app/v1/*.proto
	protoc-go-inject-tag -input=./01_example/api/product/app/v1/post.pb.go

mock:
	mockgen -source=./01_example/internal/domain/post.go  -destination=./01_example/internal/mock/domain/post.go