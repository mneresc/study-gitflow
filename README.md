# GIT

## Git flow

![GIT FLOW](https://lh3.googleusercontent.com/70jaEZnESXQ6SssU5uI4yO62JBz6xq2sNrrz8bW_ap2CuWUaQlbKs3j6NyRJnvcvYwAugkW8WzNJX21dZ2SMd9O_1TTpKZT-FsBkYSPy4rUSpJSo2C-WPTaLc2jQ8ancyj1TetXQ "GIT Flow")


## Pull Request/Template PR

* Mudar branch padrao para develop
* Proteger main
* [Posso adicionar um template de Pull Request](./.github/PULL_REQUEST_TEMPLATE.md)
* [Code Owner: Pode colocar quem e responsavel por determinado formato de arquivo, podendo ser pessoas ou grupos](./.github/codeowners)
* Commitlint + Husky [link](https://commitlint.js.org/guides/local-setup.html) [instalacao](https://dev.to/vitordevsp/padronizacao-de-commit-com-commitlint-husky-e-commitizen-3g1n)
* 

```bash

yarn add @commitlint/config-conventional @commitlint/cli -D

```
* commisar - Validacao de complience de commit
* commitzen - linha de comando para ajudar a commitar

## Code Review

Plugin do VSCode

## CI

* Script
    * Teste
    * Linter
    * Verificacao de qualidade
    * Verificacao de seguranca
    * Geracao de artefatos para deploy
    * Identificacao da proxima versao gerada
    * Geracao de tags e releases

* Workflow
    * Event: on:push
    * Filter: branches: - master
    * Envoriment: runs-on:ubuntu
    * Actions: steps: 
            -action/run-build # actions  
            - npm run test    # comando
* Active status check 
    * https://github.com/mneresc/study-gitflow/settings/rules/new?target=branch
    * Require status checks to pass
    * Adjust ci to use PR trigger

    ```yml
    on: 
    pull_request:
        branches:
        - develop
    ```

* [Action to build docker container](https://github.com/marketplace/actions/build-and-push-docker-images)

* Sonar
    * Run sonar on docker: docker run --name sonarqube-custom -p 9000:9000 sonarqube:10.6-community

    ´´´bash
    sonar-scanner \
    -Dsonar.projectKey=gotest \
    -Dsonar.sources=. \
    -Dsonar.host.url=http://localhost:9000 \
    -Dsonar.token=sqp_f008795e4eb500ca65bb42855613bd39b878b29f

    # or on sonar properties path

    sonar-scanner 

    # or

    sonar-scanner -Dsonar.projectProperties=sonar-project.properties
    ´´´
    * Generate coverage with go
    ´´´bash
        go go test --coverprofile=coverage.out
    ´´´
[PAREI] 4. Integração contínua-20240320T194553Z-002 - 5. CI com Docker - aula 

