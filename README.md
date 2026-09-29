# frontend

A gRPC service on the vikrant platform, serving
`frontend.v1.FrontendService`.

Write your code in [`internal/handlers/`](internal/handlers/). Everything
around it is already in place:

- a hardened server that protects itself from toxic clients;
- a drain that moves long-lived streams to other replicas on shutdown;
- a canonical client for your callers;
- a rootless image;
- a canary rollout;
- deny-by-default networking;
- autoscaling on in-flight RPCs;
- trace context propagated through every hop;
- one overlay per environment (dev, staging, prod);
- CI that proves all of it on every push, then publishes the image and
  pins dev to it.

## Who owns what

| Path | Owner | Changes when |
|---|---|---|
| `service.yaml`, `proto/` | you | you edit them |
| `internal/handlers/`, `cmd/`, `internal/platform/` | you, from day one | you edit them; the platform never writes here again |
| `client/` (canonical client and stubs), `deploy/base/`, `deploy/envs/*/kustomization.yaml`, `Dockerfile`, `Makefile`, `buf*.yaml`, `.github/`, `e2e/` | the platform | regenerated from `service.yaml` and the proto |
| `deploy/envs/*/image.yaml` | delivery | CI pins dev on every merge to `main`; the `promote` workflow moves a tested digest to the next environment |

To change who may call you, or what you may call, edit `service.yaml` and
push. The platform regenerates the derived files and pushes a commit to the
same branch, so your PR shows the intent and its effect together. CI's
`check-stamp` fails if a platform-owned file is edited by hand.

Environments differ only in scale and rollout pace; security and the call
graph are the same everywhere. Override scale per environment in
`service.yaml` under `spec.environments`. Each environment's Argo CD syncs
`deploy/envs/<env>` from `main`.

## Run it

```sh
make run                  # the server on :50051, metrics on :9090
make test                 # unit tests + the toxic-client conformance suite
make tools check-generated  # stubs match the proto
make image                # the rootless production image
```

## Calling this service from another one

```go
import (
	frontendv1 "github.com/smallStepGiantLeap/frontend/client/gen/frontend/v1"
	"github.com/smallStepGiantLeap/frontend/client"
)

cc, err := client.Dial(client.Target)
c := frontendv1.NewFrontendServiceClient(cc)
```

To call another service, list it under `outbound` with the methods you
need. If it has not granted them, the platform opens an access-request
pull request on its repository; once that is merged, this repository is
re-rendered and its e2e expects the calls to succeed.
