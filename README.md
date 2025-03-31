# labs-four
Projeto de conclusão de pós-graduação


for i in {1..7}; do curl -i http://localhost:8080/hello; echo; done


for i in {1..20}; do curl -i -H "API_KEY: meu-token" http://localhost:8080/hello; echo; done