```
service DeploymentSrv {
	rpc Create(DeploymentCreateRequest) returns (DeploymentCreateResponse) {
		option (grpcauth.simple) = {
			path: "prj/{project_id.name}/dpl",
			verbs: [ perms.WRITE ]
		}
	};
}
```
