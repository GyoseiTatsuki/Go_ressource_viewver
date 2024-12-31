# Go_ressource_viewver
Idée de base :
  - le client :
      - un petit programme installé sur une machine, il contient une api web qui renvoie une fois intérogée un json     
        contenant les ressources utilisée sur cette machine en question.
  
  - le serveur : héberge un site web qui va récupérer les ressources des différentes machines.
    
  - l'utilisateur : Ce connecte via une interface web sur le serveur est peu ajouter les
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

Comment fonctionne le projet :
Il y a deux composants principal :
  - le client :
      Il contient une API qui renvoie les ressources utilisée ( CPU, RAM et disque) et les  renvoie sur @ip:8080/stats au format json :
    {"CPUUsage":0.26036823507920076,"MemoryUsage":25.94676120264016,"DiskUsage":36.80057865965248}

  - le serveur :
      Son rôle est de récupérer les ressources de ou des ip qu'on lui renseigne pour les afficher sur une page web : @ip:8081

Comment le mettre en place :
  Dans un premier temps il faut lancer les clients, pour ça il faut
