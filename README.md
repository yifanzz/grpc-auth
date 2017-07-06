```
service PingService {
	rpc Ping(PingRequest) returns (PingResponse) {
		option (grpcauth.simple) = {
			path: "resource/ping/{.int_value}",
			verbs: [ perms.WRITE ]
		}
	};
}
```
