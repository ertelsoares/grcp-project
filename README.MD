<h1>GRPC É UM FRAMEWORK GRPC CRIADO PELA GOOGLE QUE UTILIZA HTTP2 PARA TRANSPORTE</h2>
E PROTOBUFS PARA SERIALIZAÇÃO OU SEJA CONVERTE EM BINARIO PARA TRANPOSTAR E PARA LER
EM CASO DE LINGUAGENS PRECISAM DE UM SDK PARA RODAR.

NO PROTOBUF VOCÊ CONSEGUE DEFINIR:

<b>SERVIÇOS :</b> NO CASO MÉTODOS QUE VÃO RECEBER ALGO OU RETORNAR ALGO SEMELHANTE A UMA INTERFACE DEFININDO UM CONTRATO

<b>LIMITAÇÕES PROBLEMA QUANDO PRECISA ALTERAR O CONTRATO</b>

<b>MENSAGENS</b>
NO CASO CADA CAMPO TEM UM ID DEFINIDO

ELE É TIPADO E TEM QUE SEMPRE GARANTIR QUE SERÁ STRING

<h4>Comandos</h4>
Para compilar o sdk para gerar os arquivos a partir do proto digita o comando : 
<code>protoc --go_out=. --go-grpc_out=. .\user.proto</code>