# Europe Railway

## Descrição

O projeto **Europe Railway** é um estudo sobre arquitetura orientada a eventos, real-time, GPS e o protocolo MQTT. A ideia é criar um sistema onde maquinistas podem aceitar trabalhos para viagens de trem entre cidades europeias, como uma rota entre Holanda e Espanha. A inspiração principal é o jogo *Euro Truck Simulator*, mas aplicado ao universo ferroviário.

## Brainstorm

A ideia do projeto consiste nos seguintes pontos:

* **Cadastro de Rotas:** As concessionárias de trem poderão cadastrar rotas de viagem em aberto, especificando o trajeto e o valor pago pelo serviço.
* **Disponibilidade para Maquinistas:** Maquinistas freelancers poderão visualizar as rotas disponíveis e se candidatar para realizá-las.
* **Sistema de Eventos:** Ao se candidatar, um evento será disparado para a concessionária.
* **Aprovação:** A concessionária poderá aceitar ou recusar a candidatura do maquinista.
* **Simulação da Viagem:** O projeto irá simular a viagem em tempo real para que seja possível acompanhar o trajeto e o funcionamento do sistema.
