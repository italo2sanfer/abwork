# TODO: Crie uma Lista 'itens' para armazenar os equipamentos:
itens = []

# TODO: Crie um loop para solicita os itens ao usuário:
# TODO: Solicite o item e armazena na variável "item":
# TODO: Adicione o item à lista "itens":
MAX_NUMBER_ITENS = 3
counter = 0
while counter < MAX_NUMBER_ITENS:
  tool = input()
  itens.append(tool)
  counter +=1

# Exibe a lista de itens
print("Lista de Equipamentos:")  
for item in itens:
    # Loop que percorre cada item na lista "itens"
    print(f"- {item}")