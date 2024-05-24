generate-sso:
	- protoc -I proto proto/sso/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative

generate-question:
	- protoc -I proto proto/question-service/*.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative

generate-answer-go:
	-protoc -I proto proto/answer/*.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go --go-grpc_opt=paths=source_relative

generate-answer-python:
	python3 -m grpc_tools.protoc -I proto proto/answer/*.proto --python_out=./gen/python --pyi_out=. --grpc_python_out=./gen/python
