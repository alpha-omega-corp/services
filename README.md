# SERVICES

This repository is a minimalistic SDK for integrating new services to the organisation's architecture.

This ensures coherent semantics across all services. 

The application's entry point is : [api-gateway](https://github.com/alpha-omega-corp/api-gateway)

## ABOUT IT

- [database](https://github.com/alpha-omega-corp/services/blob/production/database/database.go): Creates a database handler from a dsn string
- [server](https://github.com/alpha-omega-corp/services/blob/production/server/grpc.go): Creates a tcp server and handles grpc inside it's callback function
- [httputils](https://github.com/alpha-omega-corp/services/tree/production/httputils): Facilitates http error handling

