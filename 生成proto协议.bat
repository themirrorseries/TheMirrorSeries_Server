protoc ./proto/Types.txt --go_out=./protofiles/
protoc ./proto/Types.txt --csharp_out=./protofiles/

protoc ./proto/LoginTypes.txt --go_out=./protofiles/
protoc ./proto/LoginTypes.txt --csharp_out=./protofiles/
protoc ./proto/MatchTypes.txt --go_out=./protofiles/
protoc ./proto/MatchTypes.txt --csharp_out=./protofiles/
protoc ./proto/UserTypes.txt --go_out=./protofiles/
protoc ./proto/UserTypes.txt --csharp_out=./protofiles/

protoc ./proto/dto/AnyDTO.txt --go_out=./protofiles/
protoc ./proto/dto/AnyDTO.txt --csharp_out=./protofiles/