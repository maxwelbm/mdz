#!/bin/bash

set -e

# Configurações do Cloudsmith
CLOUDSMITH_REPO="mdz/mdz" # Substitua pela sua organização e repositório
API_KEY="${CLOUDSMITH_API_KEY:?Erro: A variável CLOUDSMITH_API_KEY não está definida.}"

# Gere os pacotes com Goreleaser
echo "Gerando pacotes com Goreleaser..."
goreleaser release --snapshot --clean

# Verifica se a pasta dist/ foi criada
if [ ! -d "dist" ]; then
    echo "Erro: Pasta dist/ não encontrada. Certifique-se de que o Goreleaser gerou os pacotes corretamente."
    exit 1
fi

# Faz o upload dos pacotes para o Cloudsmith
echo "Fazendo upload dos pacotes para o Cloudsmith..."
for file in dist/*.deb dist/*.rpm dist/*.tar.gz; do
    if [ -f "$file" ]; then
        # Determina o tipo do pacote baseado na extensão
        if [[ "$file" == *.deb ]]; then
            PACKAGE_TYPE="deb"
        elif [[ "$file" == *.rpm ]]; then
            PACKAGE_TYPE="rpm"
        elif [[ "$file" == *.tar.gz ]]; then
            PACKAGE_TYPE="generic"
        else
            echo "Tipo de pacote não suportado: $file"
            continue
        fi

        echo "Enviando: $file (tipo: $PACKAGE_TYPE)"
        cloudsmith push "$PACKAGE_TYPE" "$CLOUDSMITH_REPO" "$file"
    fi
done

echo "Upload concluído com sucesso!"
