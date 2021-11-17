## Migração da AWS para Digital Ocean

Código de migração principal: `migrate.go`
Código de migração auxiliar: `saidas.go`

O código principal realiza migração de todos os assets contidos em `assets.json`. O código auxiliar foi criado por uma inconsistência no nome de uma propriedade do asset `saida`.

Para migrar: `go run migrate.go`. ATENÇÃO: cheque os IPs configurados no código.

O arquivo `assets.json` e o arquivo `saidas.json` foram obtidos pelo Fauxton diretamente do couchDB do peer.

