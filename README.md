How to run this program
=======================
To run this program, an installation of Go is required. To run the script to generate the certificates, the `openssl` and make command is required. The following steps will guide you through the process of running the program.

First of run the following command to make sure files are correct:
```bash
go mod tidy
```
Start of by by navigating to the root of the project in a terminal (I use wsl for windows, and so the paths provided might not work in things like cmd or powershell).
Run the following command to generate the certificates:
```bash
make cert
```

1. Open a terminal and navigate to the root directory of the project.
2. Run the following command to start the server:
```bash
go run server/hospital.go
```
3. Open 3 new terminals and run the following commands (one in each terminal) to start the clients (the input value can be changed to any integer value):
```bash
go run client/patient.go -name=Alice -port=:9000 -input=1
```
```bash
go run client/patient.go -name=Bob -port=:9001 -input=2
```
```bash
go run client/patient.go -name=Charlie -port=:9002 -input=3
```

4. The patients will now follow protocol, and through the print statements, it will be shown that they each use secret sharing, share their secrets with each other, compute their final aggregate share, and sends it to the hospital. The hospital will then compute the aggregate value of the patients' secrets and print it to the terminal.