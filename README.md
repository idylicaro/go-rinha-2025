# 🐔 go-rinha-2025

## Descrição:

Implementação do desafio Rinha de Backend 2025.

## Solução

A solução é baseada em uma API HTTP com dois endpoints principais:

- `POST /payments`: recebe uma enumeras requisições de transações e faz o intermedio para o `payment-processor`.
- `GET /payments-summary`: retorna o resumo das consolidações de pagamentos feitos para o `payment-processor`.

? A solução precisa de um retorno instataneo para as requisições `POST /payments` e um processamento async para consolidar o pagamento encaminhado para o `payment-processor` podendo apenas contabilizar, quando receber uma respota positiva (2xx) do `payment-processor`.
?

## Arquitetura

A arquitetura foi projetada para ser simples, eficiente e confiavel.

- Load Balancer (HAProxy)
- API (Golang + Go-chi)
- Persistência (Redis)

Olhando para o escopo determinado, como precisamos salvar as agregações para o summary e também precisamos respeitar o requisito de um health-check a cada 5 seg. Vamos utilizar o redis para persistencia das agregações/count e também para a informação temporaria do health-check.

Com base nisso, para evolução futura este projeto segue o padrão Clean Architecture, separando responsabilidades em controllers (camada de apresentação), services (casos de uso/regra de negócio) e repositories (persistência de dados).
