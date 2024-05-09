# go-auth

## Criar API [x]

## Criar conexão banco [x]

  Criamos uma interface chamada database -> A interface possui a assinatura de métodos padrões para conectar em bancos
  e fazer consultas

  Nós criamos um adapter de banco postgres que usará essa interface, para personalizar os métodos e conectar em cenários de banco postgres

  Os repositórios, usam a interface database, que pode ser tanto um adapter postgres ou qualqeur outro adaptaer
  quem chamar esse repository que injetará a dependencia de qual banco usar, deixando o código mais modlular e reutilizável 


## Criar JWT BEARER logic[]
  JÁ CRIAMOS A CONFERENICA NO BANCO CASO O USER EXISTA
  FALTA: 
    [x] REOTNRAR FALSE QUANDO NÃO EXISTIR
    [x] REOTNRAR FALSE QUANDO NÃO EXISTIR e a senha não for válida
    [x] retornar o JWT quando o user existir e a senha for válida

## Criar sign-in
  [x] 

## Criar token que expira
  ## -> o jwt já faz isso, podemos só validar e-mail e senha e retornar um jwt para o usuário
  [x] 

## ROTAS:
  ### SIGNIN[x]
  ### SIGNUP[x]


