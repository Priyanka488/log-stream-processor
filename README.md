

# Log Stream Processor

Ingress -------> Handler ----> Processor -----> Filter ----> Egress (DB)
(Source)

- TCP server
- GRPC server
- API server



ingress <----channel-----> handler


Context


Done() {channel}
cancel()