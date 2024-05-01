# Sample gRPC with sqlc

REST isn't the only answer when creating distributed systems. [GRPC][grpc] is a
popular alternative.

It's built on top of [protocol buffers][protobuf], a performant binary format
for data serialization.

We also will sample here another query builder: [sqlc][sqlc]

## Fair warning: code generation

The gRPC generates client/server code, so you don't have to worry much about the
distributed system part.

_Famous last words_.

## Requirements

- go 1.22
- gRPC 1.63
- protobuf 1.34
- protoc compiler 3
- go-sqlite3 driver 1.14
- sqlc query generator 1.26 

## Project setup

Create the skeleton:

```bash
# cd 0017-sample-grpc
mkdir -p ./{sample-grpc-client,sample-grpc-server,protobuf}
cd sample-grpc-client
go mod init sample-grpc-client
go get -u google.golang.org/protobuf
go get -u google.golang.org/grpc
cd ..
cd sample-grpc-server
go mod init sample-grpc-server
go get -u google.golang.org/protobuf
go get -u google.golang.org/grpc
cd ..
touch ./{sample-grpc-client,sample-grpc-server}/main.go
touch ./protos/todo.proto
mkdir -p ./sample-grpc-server/db
```

Then you need to install [the protobuffer compiler][protoc] and golang plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

[Install sqlc][sqlc-install] after that and init server project:

```bash
sqlc init
touch db/schema.sql
touch db/queries.sql
```

Modify the generated sqlc.yaml file to this:

```yml
---
version: "2"
sql:
  - engine: "sqlite"
    queries: "queries.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "gen"
        out: "db/gen"
```

Add The schema in `sample-grp-server/db/schema.sql` file:

```sql
create table if not exists todos (
    id integer not null primary key,
    description text not null,
    done boolean not null default false,
    created timestamp not null default CURRENT_TIMESTAMP,
    updated timestamp not null default CURRENT_TIMESTAMP
);
```

Create The queries in `sample-grp-server/db/queries.sql` file:

```sql
-- name: List :many
select * from todos where lower(description) like lower(concat('%',?,'%'))

-- name: Find :one
select * from todos where id = ?

-- the other queries
```

And finally [configure the project][sqlc-configure], so we can generate
database client code:

```bash
# cd 0017-sample-grpc
cd sample-grpc-server
sqlc generate
```

## How to build

Besides hit `sqlc generate` every time a query is modified, you need to generate
grpc code as well.

Projects are independent of each other, in fact client/server could be written
in many different languages.

But you will need the code generation step for both projects:

```bash
# compile for server
protoc --go_opt=Mprotos/todo.proto=sample-grpc-server/protos --go_out=. \
  --go-grpc_opt=Mprotos/todo.proto=sample-grpc-server/protos --go-grpc_out=. \
  protos/todo.proto
# comile for client
protoc --go_opt=Mprotos/todo.proto=sample-grpc-client/protos --go_out=. \
  --go-grpc_opt=Mprotos/todo.proto=sample-grpc-client/protos --go-grpc_out=. \
  protos/todo.proto
```

Once client and server gets created, complete the client and server code as
described in [this tutorial][basics]. For example, our protobuf definition
defines a `service` called **TodoService**:

```protobuf
// ... other definitions 

service TodoService {
  rpc List (TodoRequest) returns (TodoResponse);
  rpc Insert (TodoRequest) returns (TodoResponse);
  rpc Find (TodoRequest) returns (TodoResponse);
  rpc Update (TodoRequest) returns (TodoResponse);
  rpc Delete (TodoRequest) returns (TodoResponse);
}
```

This definition will produce, in generated code, a client that can be built with
`protos.NewTodoServiceClient` and in a [server interface][server-interface] that
you will need to implement:

```go
package main

import (
	"context"
	pb "sample-grpc-server/protos"
)

// TodoServer - receiver for protos.TodoServiceServer implementation
type TodoServer struct {
	pb.UnimplementedTodoServiceServer
}

func (s *TodoServer) List(ctx context.Context, request *pb.TodoRequest) (*pb.TodoResponse, error) {
	//...
}

func (s *TodoServer) Insert(ctx context.Context, request *pb.TodoRequest) (*pb.TodoResponse, error) {
	//...
}

func (s *TodoServer) Find(ctx context.Context, request *pb.TodoRequest) (*pb.TodoResponse, error) {
	//...
}

func (s *TodoServer) Update(ctx context.Context, request *pb.TodoRequest) (*pb.TodoResponse, error) {
	//...
}

func (s *TodoServer) Delete(ctx context.Context, request *pb.TodoRequest) (*pb.TodoResponse, error) {
	//...
}
```

Once you implement the stub, do the `go build` in each folder:

```bash
# cd 0017-sample-grpc
cd sample-grpc-client ; go build ; cd ..
cd sample-grpc-server ; go build ; cd ..
```

## How to run

In one terminal:

```bash
# cd 0017-sample-grpc
cd sample-grpc-server ; go run .
```

In another:

```bash
# cd 0017-sample-grpc
cd sample-grpc-client ; go run .
```

## Noteworthy

- I firmly still believe that code generators are nice foot guns, but boy it's
  running already.
- The sqlc generator is far more pleasant that what i expected.
- The grpc generator for go is a little clumsy. I mean, `--go-grpc_out`, what on
  earth people who designed this where thinking?
- Server interface implementation is a breeze thanks to intellij tooling. Define
  a struct, pass it to `RegisterTodoServiceServer` and let the IDE offer to do
  the interface implementation. Neat.
- One drawback is the domain fragmentation: both sql and grpc define a Todo
  struct. It's important to pay attention and be sure which one is being used.

## Should generated code get versioned or not?

This is a major issue and there is no final, correct answer.

- The tools to codegen might not be present in the future.
- Outdated generated code might not be compatible with new libraries.
- Build workflow might not get access to code generation tools.

Lots of things could go wrong, but get this: usually we don't version build
artifacts, binaries, we don't put derivatives on git. We do serve those things,
those artifacts, by other channels, but not on git.

Therefore, the instructions on how to generate code must be versioned, not the
byproduct of it.

Your pipelines must be able to invoke those tools and always generate the latest
code, it helps to keep things solid and fresh.

[grpc]: https://grpc.io/docs/what-is-grpc/introduction/
[sqlc]: https://docs.sqlc.dev/en/stable/index.html
[protobuf]: https://protobuf.dev/getting-started/gotutorial/
[protoc]: https://grpc.io/docs/protoc-installation/
[sqlc-install]: https://docs.sqlc.dev/en/stable/overview/install.html
[sqlc-configure]: https://docs.sqlc.dev/en/stable/tutorials/getting-started-sqlite.html
[basics]: https://grpc.io/docs/languages/go/basics/
[server-interface]: https://github.com/grpc/grpc-go/blob/master/examples/route_guide/routeguide/route_guide_grpc.pb.go#L193
