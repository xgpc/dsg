
protoc --go_out=. --go-grpc_out=. base.proto -I .

# 可在根目录下运行, 并生成文件到proto所在目录
# protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative ----go-grpc_opt=paths=source_relative .\service\grpcService\proto\base.proto
