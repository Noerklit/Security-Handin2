go run server/hospital.go
go run client/patient.go -name=Alice -port=:9000 -input=1
go run client/patient.go -name=Bob -port=:9001 -input=2
go run client/patient.go -name=Charlie -port=:9002 -input=3