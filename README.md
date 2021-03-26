# ploggi
Simple golang application that exposes container logs as a grpc service.

The server has two rpc:

```
    rpc GetLog (Pod) returns (PodLog);
    rpc StreamLog (Pod) returns (stream PodLog);
```
[Proto file](pkg/api/ploggi/ploggi.proto)


_hack/example/ contains example clients for both rpcs


## Local development

1. Start a kind cluster with local docker registry enabled
   ```
   ./_hack/setup_kind.sh
   ```

2. Build ploggi docker image using ko
   ```
   make local-image
   ```

3. Push image to local registry
   ```
   make push-local
   ```

4. Apply manifests to kind cluster
   ```
   make apply-kind
   ```

5. Start stream client
   ```
   go run _hack/examples/ploggi_stream_client/stream.go <ploggi-pod-name> ploggi
   ```

6. Execute get client
   ```
   go run _hack/examples/ploggi_get_client/get.go <ploggi-pod-name> ploggi
   ```
   Ploggi will log every time the get rpc is executed and logs should be streamed to the stream client.