# Desafio-golang
Api criada para o disafio de Golang da Americanas SA!

# Tecnologias usadas na construção da API
A API foi criada em Golang e com banco de dados MongoDB. O banco de dados foi criado diretamente na  [MongoDB Atlas](https://www.mongodb.com/pt-br/cloud/atlas/efficiency)

# Endpoinsts da API
- http://localhost:8080/v1/planet/create
Será criado um novo planeta no banco de dados com as informações passadas no body. As informações são passadas em JSON conforme imagem abaixo:

![image](https://user-images.githubusercontent.com/39388688/187077473-31a10713-e959-4109-806d-f65a9d48529d.png)

OBS: Peguei as informações dos planetas diretamente da API pública do [Star Wars](https://pipedream.com/apps/swapi) e, ao criar um novo planeta, será gerado um número
sequencial, onde chamei o mesmo de "PlanetApiId". Esse PlanetApiId é usado no GetById e no Get passando o nome do planeta para trazer as informações dos filmes 
onde esse planeta apareceu. É importante que os planetas sejam cadastrados na ordem em que eles aparecem na API pública do Star Wars, pois, esse "PlanetApiId" tem o 
intuito de ser o ID do planeta na API do Star Wars.

- http://localhost:8080/v1/planet/getall
Será retornado todos os planetas criados no banco de dados.

- http://localhost:8080/v1/planet/get/Tatooine
Será retornado do banco de dados o plenta com o nome Tatooine. Nesse Endpoint será mostrado também os filmes onde esse planeta apareceu, conforme imagem abaixo:

![image](https://user-images.githubusercontent.com/39388688/187077976-d98e2b92-f4e0-4a57-9485-63f14dc960dd.png)

- http://localhost:8080/v1/planet/getbyid/1
Será retornado do banco de dados o planeta com o Id "1". Também será retornado a informação de quais filmes esse planeta apareceu.

![image](https://user-images.githubusercontent.com/39388688/187078101-c8f90de5-0302-457c-bca5-51c182f453c2.png)

- http://localhost:8080/v1/planet/delete/Tatooine
Será deletado do banco de dados o planeta com o nome de "Tatooine"
