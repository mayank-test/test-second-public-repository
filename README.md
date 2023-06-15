# Service Template
The repository could be used like a basic layout of Nightfall services.

## How to Use
The service_template can be used OOTB, you just need to change the app name and ports. In addition, you will need to make a choice between different servers (gRPC, REST HTTP) and based on that choice you will then need to remove different code snippet from main.go, `/api` package, and the `/config` files.

### main.go
Different server interface are started in parallel with make (REST and gRPC), select the relevant providers for your service [here](/main.go#L42) and the appropriate invocation [here](/main.go#L56).

### `/api`
Remove the api package that you don't need, either [gprc](/api/grpc/)
 or [rest](/api/rest/)

### `/config`
Remove the configuration for the server type that is not needed in all of the respective config files.

REST has top level field of `httpServer` and `router` : [base](/config/base.json#L3), [development](/config/development.json#L2) , and [production](/config/production.json#L2)

gRPC config has the top level field of `grpcServer`: [base](/config/base.json#L27), [development](/config/development.json#L28) , and [production](/config/production.json#L21)



# Files in the root `/` of the repository
### `main.go`
Contains the `main()` function for the service executable. We use [dependency injection library](https://github.com/uber-go/fx) from Uber to construct the service internal structures.
### `go.mod`
We use [Go modules](https://github.com/golang/go/wiki/Modules) for Go-program dependencies.

**NOTE** When creating the new repo, you should not copy the file from the template, but just start to use go modules for the new repo(for example `my_new_service` repo):
```
go mod init github.com/watchtowerai/my_new_service
```

### `Makefile`
The `Makefile` is provided for building purposes and it supports different the service building commands like `clean`, `test`, `build` etc.
### `Dockerfile`
The `Dockerfile` contains instructions for building Docker container. The file requires some environment variables like `GIT_USER`, `GITHUB_OAUTH_TOKEN` and `SERVICE_NAME` which comes from the `Makefile`
**NOTE** If you are not creating a grpc service, you may remove both references to `grpc_health_probe`, as it will not be needed.
**NOTE** If you are creating a http service, you may wish to un-comment the datadog labels on the container, to get go-expvar data and a healthcheck
### `.gitignore`
Juts standard `.gitignore` for a Golang package
### `.arcconfig`
[Arcanist](https://secure.phabricator.com/book/phabricator/article/arcanist/) is a change control tool that we use for commiting changes into our git database
### `.arclint`
Arcanist lint configuration.

## Service Application Directories
The directories names have singular form, not plural one just to improve the code redability. For example if we have a `User` entity the code will look like:
```
import "gihtub.com/watchtowerai/service_template/entity"
...
func foo() {
    var user entity.User // but not entities.User
}
```
### `/api`
The `/api` folder intends for the source code which provides an implementation of a specific interface (REST, gRPC, GraphQL etc.) for accessing the service. It is recommended that a specific implementation reside in the appropriate sub-folder like `/api/rest` would contain some RESTFull API specific functions (`gin.Router`, an http server etc.)

### `/client`
Usually include things like database drivers, launch darkly, bugsnag, third party api's, etc. These files typically export the client initialize into an fx module so fx can add it to the dependency injection graph

### `/config`
We use [Viper](https://github.com/spf13/viper) for config (anything non secret related). We have configs that are created in a hierarchical structure. Ever `config` folder has a `base.json` as well as a `[env].json` file which viper merges field values on top of the base by looking at our GO_ENV variable (set to development by default). You can also set a `user.json` which is merged on top of base and the `[env].json` for user specific development configs. user.json is gitignored.

The config service code ([Viper](https://github.com/spf13/viper) for example) should be probably placed in common libraries (see [QA](#where-is-Logger) )

### `/entity`
The model part of MVC. Usually contains abstracted "classes" that interact directly with the database. For example an entity associated with users in our database.

### `/service`
Contains business logic for the service. In the folder there is an abstraction model which is described by some publicly available structures and interfaces, that could be used by a concrete API implementation for providing the functionality to external clients.

## CircleCI Configuration
Replace all instances of service_template in docker-compose.ci.yml and .circleci/config.yml with your service name. GITHUB_OAUTH_TOKEN, GIT_USER, and other environment variables that apply to all workflows can be configured on CircleCI under watchtowerai => settings => context under the group "build". Project-specific environment variables can be configured in the project's settings.

## ReviewDog Configuration
Add the app to the repo here: https://github.com/apps/reviewdog/installations/6991369.
Then go to `reviewdog.app/gh/watchtowerai/REPO_NAME_HERE` to get the token, and add it to the project's settings on CircleCI.


## Q&A
### Where is Logger?
The very common code like `Logger` or `Config`, that are used by almost any service, can be placed in the common libraries repository(`watchtorwer_go_libraries`) and an interface methods could be used out of there.
### How configuration works?
Config logic is in `watchtower_go_lib`, please take a look at `https://github.com/watchtowerai/watchtower_go_libraries/tree/master/pkg/config/README.md`
