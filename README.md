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
        conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
        if err != nil {
            continue
        }
        conn.Close()
        fmt.Printf("✅ Porta aberta: %d\n", port)
    }
}
```

---

## Instalar e executar

1. Clone ou copie o código para uma pasta:

```bash
git clone <repo-ou-copia-do-codigo>
cd portscanner
```

2. Compilar (build):

```bash
go build portscanner.go
```

3. Executar:

```bash
./portscanner
```

> Para rodar sem compilar:
>
> ```bash
> go run portscanner.go
> ```

---

## Sugestões de melhorias (próximos passos)

1. **Concorrência com goroutines** — paralelizar o escaneamento para acelerar o processo.
2. **Linha de comando (flags)** — permitir configurar alvo, intervalos de portas, timeout e número máximo de goroutines via `flag`.
3. **Detecção de banners** — ler respostas dos serviços nas portas abertas para identificar serviços/versionamento.
4. **Saída em diferentes formatos** — JSON, CSV ou relatório em texto para integração com outras ferramentas.
5. **Rate limiting / Throttling** — evitar sobrecarregar redes ou sistemas alvo.
6. **Modo seguro / whitelist** — checar se alvo está na lista de permissões antes de escanear.

---

## Exemplo de uso avançado (sugestão)

```bash
# build com nome custom
go build -o portscan portscanner.go

# executar (exemplo) - alvo e range como flags (se implementado)
./portscan -target 192.168.0.10 -start 1 -end 65535 -timeout 300
```

---
