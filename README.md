# Go_ressource_viewver
Idée de base :
  - le client :
      - un petit programme installé sur une machine, il contient une api web qui renvoie une fois intérogée un json     
        contenant les ressources utilisée sur cette machine en question.
  
  - le serveur : héberge un site web qui va récupérer les ressources des différentes machines.
    
  - l'utilisateur : Ce connecte via une interface web sur le serveur est peu ajouter les machines qu'il veut (entre de gros guillemet "suppervisé"

commandes clients windows :
go mod init Go_ressource_viewver/client/windows
go get github.com/gin-gonic/gin
go get github.com/shirou/gopsutil/cpu
go get github.com/shirou/gopsutil/mem
go get github.com/shirou/gopsutil/disk

Pour client linux :
go mod init Go_ressource_viewver/client/linux
go get github.com/gin-gonic/gin
go get github.com/shirou/gopsutil/v3/mem
go get github.com/shirou/gopsutil/v3/disk
go get github.com/shirou/gopsutil/v3/cpu
