pb-client:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./pkg/api/ploggi/ploggi.proto

local-image:
	ko publish --local -B .

push-local:
	docker tag ko.local/ploggi:latest localhost:5000/ploggi:latest
	docker push localhost:5000/ploggi:latest

apply-kind:
	kubectl apply -f _hack/resources