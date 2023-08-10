# gRPC Golang

A Golang repository with some study projects using gRPC

### Greet (w/ SSL) and Calculator projects
It has some examples of simple RPC, server-streaming RPC, client-streaming RPC, and bidirectional-streaming RPC communication types.

### Blog
A simple example of a CRUD project using gRPC with client/server communication.

## Running the projects
This repository uses a `Makefile` file to build all the project's binaries. Just run `make 'project_name'` to create the binaries, for example. `make blog`.
The blog project also uses a Docker Compose file. It can be easily initialized by running `docker compose up`. Make sure you have [Docker](https://www.docker.com/) properly configured on your machine first.

To run the server, use the following command: `./bin/blog/server`. Same with the client `./bin/blog/client`.

Be sure to build the project with the `make` command every time you update a file, this way the changes will be reflected in the binary files.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details
