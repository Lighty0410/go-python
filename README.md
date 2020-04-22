# Installation instructions: 
Set this environment variables if you want to use custom GRPC server/client ports 
or custom data.json file path.

For Golang:  
    `JSON_FILE_PATH`  
    `GRPC_SERVER_ADDRESS`    
    `GRPC_CLIENT_ADDRESS`    
    `HTTP_SERVER_ADDRESS`
    
For Python:  
    `GRPC_PYTHON_CLIENT_PORT`    
    `GRPC_PYTHON_SERVER_PORT`
    
Then run `make build` and `make run`

# Hours spent doing this task 
24 

# QPS 
~21000
System spec: 
i7 9700k, 16gb RAM (ddr 4 3600mhz), ubuntu 18.04LTS.
