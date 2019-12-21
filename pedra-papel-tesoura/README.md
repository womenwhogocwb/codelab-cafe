### Exercício Jogo Codelab Café ☕️

Vamos escrever um jogo "Pedra, Papel, Tesoura" pra ser jogado em linha de comando contra a máquina.

Você pode construir da forma que preferir. Abaixo colocamos um passo-a-passo que pode servir de guia caso você deseje.

#### Regras do Jogo:
- 📄Papel vence ⛰Pedra
- ⛰Pedra vence ✂️Tesoura
- ✂️Tesoura vence 📄Papel

#### Passos:
1) declarar um slice de `string` com os valores: "pedra", "papel", "tesoura"
2) declarar uma variável pra receber a resposta da pessoa jogadora (atribuir qualquer valor string por enquanto)
3) declarar uma variável pra receber a resposta da máquina (atribuir qualquer valor string por enquanto) 
4) comparar os valores (sua resposta vs resposta da máquina) de acordo com as regras do jogo (por exemplo, papel vence pedra)
###### dica: você pode usar if ou switch, ou a estrutura que mais fizer sentido pra você
5) adicionar verificação de empate
6) adicionar verificação do valor entrado (só devem ser aceitos os valores pedra, papel ou tesoura)
7) adicionar sorteio pra randomizar a resposta da máquina
###### checar pacote `math/rand` da biblioteca padrão de go
8) fazer o programa receber o valor do que for digitado na linha de comando pela pessoa jogadora
###### usar Scan, checar pacote `bufio` (NewScanner, Scan, Text)
9) printar mensagem dependendo do resultado da comparação :)