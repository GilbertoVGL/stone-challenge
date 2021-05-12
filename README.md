# Desafio Técnico Stone

O desafio é criar uma API de integração com a API Rest do Github. É esperado que o desafio seja entregue em até 2 semanas.

## API

### Documentação do Github

[Aqui.](https://docs.github.com/en/rest)

### Regras gerais

- Usar formato JSON para leitura e escrita.
- A documentação da API faz parte do entregável do desafio.
- O código pode estar em qualquer linguagem.
- Pode-se utilizar qualquer package ou framework.
- Sua aplicação precisa ser executada em containers.
- Utilização de GIT para versionamento é obrigatório.

**_Indicação de material de estudo no final_**

### Rotas

As entregas são separadas em tier, você **não precisa** cumprir todos os tiers para que avaliemos o seu desafio, cumpra o máximo que conseguir dentro do prazo combinado, é recomendado mas não obrigatório implementar os tiers na ordem que são apresentados.
Se ficar em dúvida entre entregar mais funcionalidades com menos qualidade ou menos funcionalidade com mais qualidade, **escolha mais qualidade**. (:

#### Tier  1

1. Recebendo **apenas** o user do GitHub como parâmetro, retornar o repositório com mais stars desse user.
  - Exemplo de rota: `GET /users/{user}/popular-repository`
2. Dado um user do GitHub e um repositório, retornar a Issue **aberta** que possui mais comentários.
  - Exemplo de rota: `GET /users/{user}/repositories/{repository}/popular-issue`
3. Dado um user do GitHub e um repositório, retornar todos os Pull Request **abertos** que ainda não foram interagidos, isso é, não possuem nenhum comentário, não foram aprovados e nem rejeitados.
  - Exemplo de rota: `GET /users/{user}/repositories/{repository}/open-pull-requests`

#### Tier 2

1. Implementação de uma camada de cache, para as rotas do tier 1. Assim, se um request idêntico for feito em seguida, **não realizar nenhum request** para a API do GitHub e retornar o valor do cache.
2. Dado um user do GitHub, repositório, pull request e comment ID, copiar o conteúdo do comment para um segundo user, repositório e pull request enviado na requisição. **O comment original deve ser apagado após isso**.
3. Dado um user do GitHub e repositório, criar um arquivo .gitignore no repositório seguindo o template da linguagem predominante. **Se o repositório já possuir um arquivo .gitignore sobreescreva-o**.

#### Tier 3

1. Implementar uma stack de testes de integração, usando um servidor de Mock que funcione como retorno da API do GitHub. Para esses testes, a API precisa estar de pé **mas todos os requests devem ser feitos para o servidor de Mock** em vez de para a API do GitHub, validando as regras de negócio dos tiers 1 e 2.

## Bônus

Configurar uma integração contínua que executa testes e faz validações pertinentes no repositório do projeto.

## Forma de entrega

Link para repositório (GitHub/BitBucket/GitLab) público ou privado, se for privado, solicitar no email os usuários que serão necessários terem acesso para a revisão.

## Critérios de Avaliação

O desafio será avaliado através de quatro critérios.

### Entrega

- O resultado final está completo para ser executado?
- O resultado final atende ao que se propõe fazer?
- O resultado final atende totalmente aos requisitos propostos?
- O código possui algum controle de dependências?

### Boas Práticas

- O código está de acordo com o guia de estilo da linguagem?
- O código está bem estruturado?
- O código está fluente na linguagem?
- O código está coberto com testes unitários ?

### Documentação

- O código foi entregue com um arquivo de README claro de como se guiar?
- O código possui comentários pertinentes?
- Os commits são pequenos e consistentes?
- As mensagens de commit são claras?

### Código Limpo

- O código possibilita expansão para novas funcionalidades?
- O código é fácil de compreender?
- O código possui testes?

## Material de Estudo

### Rest

- [Desing RESTful API's](https://hackernoon.com/restful-api-design-step-by-step-guide-2f2c9f9fcdbf)
- [HTTP Status Code](https://kinsta.com/pt/blog/lista-codigos-status-http/)

### Docker

- [Docker na prática](https://blog.geekhunter.com.br/docker-na-pratica-como-construir-uma-aplicacao/)

### Outros

- [SOLID](https://www.youtube.com/watch?v=rtmFCcjEgEw)
- [KISS](http://people.apache.org/~fhanik/kiss.html)
- [GitHub Actions](https://github.com/features/actions)
- [Mock Server](https://www.mock-server.com/mock_server/getting_started.html)