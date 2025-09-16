# Scanner-de-Portas-TCP
Scanner de Portas TCP


# Port Scanner (Go) 🐹

Um scanner de portas TCP simples, escrito em **Go**, pensado como projeto de aprendizado para **Cybersecurity** (pentest básico). Ele verifica um intervalo de portas em um alvo e lista as portas abertas.

> **Aviso legal:** Use esta ferramenta apenas em alvos que você tem permissão explícita para escanear. Escanear sistemas sem autorização pode ser ilegal e antiético.

---

## Funcionalidades

* Escaneamento de portas TCP em um intervalo definido
* Timeout configurável para conexões
* Código simples e didático, ótimo para aprender `net.DialTimeout` e conceitos básicos de I/O em Go

> Observação: a versão inicial é sequencial (varredura síncrona). Recomenda-se adicionar concorrência com goroutines para melhorar a velocidade.

---

## Requisitos

* Go 1.18+ instalado ([https://golang.org/dl/](https://golang.org/dl/))
* Conexão de rede

---

## Estrutura do projeto

```
portscanner/
├── portscanner.go   # código-fonte principal
└── README.md        # este arquivo
```

---

## Código de exemplo

O arquivo `portscanner.go` contém um scanner simples. Exemplo (resumido):

```go
package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    target := "scanme.nmap.org"
    startPort := 1
    endPort := 1024

    fmt.Printf("🔎 Escaneando %s (%d-%d)...\n", target, startPort, endPort)

    for port := startPort; port <= endPort; port++ {
        address := fmt.Sprintf("%s:%d", target, port)
        conn, err := net.DialTimeout(
```
