### Instalando gRPC

Para usar as ferramentas que trabalhamos para compilar os protofiles será necessário instalar alguns pacotes no Linux ou Windows.

Execute estes comandos no seu terminal Linux:
```
sudo apt install protobuf-compiler
go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

Para finalizar, temos que adicionar a pasta "/go/bin" no PATH do Linux para que tudo que seja instalado nesta pasta esteja disponível como comandos no terminal. Adicione no final do seu ~/.bash
```
PATH="/go/bin:$PATH"
```

Execute o comando abaixo para atualizar seu terminal:

`source ~/.bashrc`
