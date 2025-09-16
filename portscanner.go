package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Alvo e range de portas
	target := "scanme.nmap.org" // você pode mudar para um IP ou domínio
	startPort := 1
	endPort := 1024

	fmt.Printf("🔎 Escaneando %s (%d-%d)...\n", target, startPort, endPort)

	for port := startPort; port <= endPort; port++ {
		address := fmt.Sprintf("%s:%d", target, port)

		// Timeout para não travar
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		if err != nil {
			continue // porta fechada
		}
		conn.Close()

		fmt.Printf("✅ Porta aberta: %d\n", port)
	}
}
