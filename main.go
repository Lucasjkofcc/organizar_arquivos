package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: organizador <pasta>")
		return
	}

	pasta := os.Args[1]

	arquivos, err := os.ReadDir(pasta)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	for _, arquivo := range arquivos {
		if arquivo.IsDir() {
			continue
		}
		ext := filepath.Ext(arquivo.Name())
		var destino string
		switch ext {
		case ".png", ".jpg", ".jpeg", ".gif":
			destino = "Imagens"
		case ".mp3", ".wav", ".flac":
			destino = "Músicas"
		case ".zip", ".rar", ".7zp", ".tar", ".gz":
			destino = "Compactados"
		case ".pdf", ".txt", ".docx":
			destino = "Documentos"
		case ".AppImage":
			destino = "Programas"
		case ".go", ".c", "py":
			destino = "Programação"
		default:
			destino = "Outros"
		}
		pasta_destino := filepath.Join(pasta, destino)
		os.MkdirAll(pasta_destino, 0755)
		origem := filepath.Join(pasta, arquivo.Name())
		novo := filepath.Join(pasta_destino, arquivo.Name())
		err := os.Rename(origem, novo)
		if err != nil {
			fmt.Println("Erro movendo", arquivo.Name(), ":", err)
			continue
		}
		fmt.Println("Movido:", arquivo.Name(), "->", destino)
	}
	fmt.Println("Concluido :)")
}
