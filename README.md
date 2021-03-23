# voicemod-challenge

## Requirements
- docker
- docker-compose

## How to launch
First build the application
```
make build
```

And then launch it
```
make start
```

## Run checks

```
make check
```

This step requires go and golangci-lint to be installed in your system

## Docs
API specification is located at `docs/api/api.v1.yaml`  and exported to `docs/api/index.html`


## Areas of improvement
- Add acceptance & integration tests
- Improve loggin
- Better error handling
- Add metrics
- Implement http server graceful shutdown
