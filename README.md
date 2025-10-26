# test-auto-version

## Sobre o Projeto

Este é um projeto de **teste e estudo** criado para explorar e validar a funcionalidade da ferramenta [release-please](https://github.com/googleapis/release-please) do Google para versionamento automático de software.

**Este não é um projeto comercial.** O objetivo é aprender e experimentar com automação de releases usando Conventional Commits e GitHub Actions.

## O que é Release Please?

Release Please é uma ferramenta do Google que automatiza a criação de releases baseada em commits convencionais (Conventional Commits). Ela:

- Analisa o histórico de commits
- Gera CHANGELOGs automaticamente
- Cria Pull Requests de release
- Gerencia versionamento semântico
- Cria tags e releases no GitHub

## Estrutura do Projeto

Este projeto contém uma aplicação simples em Go com um serviço de calculadora gRPC, usado para gerar commits e testar o fluxo de versionamento automático.

### Tecnologias

- **Go** - Linguagem de programação
- **gRPC** - Protocol Buffers e comunicação RPC
- **GitHub Actions** - CI/CD pipeline
- **Release Please** - Automação de releases

## Workflows GitHub Actions

- **`.github/workflows/.ci.yaml`** - Pipeline de CI com testes e build
- **`.github/workflows/release.yaml`** - Automação de releases com release-please

## Como Funciona

1. Commits seguem o padrão [Conventional Commits](https://www.conventionalcommits.org/)
   - `feat:` - Nova funcionalidade (minor bump)
   - `fix:` - Correção de bug (patch bump)
   - `feat!:` ou `BREAKING CHANGE:` - Mudança breaking (major bump)

2. Ao fazer push na branch `develop`, o release-please:
   - Analisa os commits desde a última release
   - Cria um PR com CHANGELOG atualizado e nova versão

3. Ao fazer merge do PR de release:
   - Tag e Release são criados automaticamente no GitHub

## Objetivo Educacional

Este repositório serve como laboratório para:
- Entender o funcionamento do release-please
- Praticar Conventional Commits
- Experimentar com automação de versionamento
- Aprender sobre CI/CD com GitHub Actions
