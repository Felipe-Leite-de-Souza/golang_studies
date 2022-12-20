# Aplicação linha de comando

Uma aplicação pequena, criada para demonstrar recursos da linhagem golang.
Tem como objetivo buscar ip's e os nomes dos servidores

Para buscar os ip's de um site utilize o seguinte comando:
```bash
    ./aplicacaoLinhaDeComando ip --host {site}
```

Exemplo:
```bash
    ./aplicacaoLinhaDeComando ip --host amazon.com.br
```

Para buscar o nome de um site utilize o seguinte comando:
```bash
    ./aplicacaoLinhaDeComando servidores --host {site}
```

Exemplo:
```bash
    ./aplicacaoLinhaDeComando servidores --host amazon.com.br
```