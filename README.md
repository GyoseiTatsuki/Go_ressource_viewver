# Go Resource Viewer

## Description

Ce projet se compose de deux composants principaux : un **client** et un **serveur**. Il permet de surveiller les ressources système (CPU, RAM, et disque) des machines et de les afficher sur une page web.

### Composants

- **Client** :  
  Fournit une API qui retourne les ressources utilisées (CPU, RAM, disque) au format JSON à l'adresse `@IP:8080/stats`.  
  Exemple de réponse JSON :  
  ```json
  {
    "CPUUsage": 0.26036823507920076,
    "MemoryUsage": 25.94676120264016,
    "DiskUsage": 36.80057865965248
  }
- **Serveur** :  
- Récupère les données des ressources des clients (via leurs adresses IP) et les affiche sur une page web accessible à l'adresse `@IP:8081`.

## Mise en Place

### 1. Lancer les clients

1.  Choisissez le système d'exploitation (Linux ou Windows) correspondant à la machine cible.
2.  Exécutez la commande suivante dans le répertoire du client :
    
    bash
    
    Copier le code
    
    `go run main.go` 
    
3.  Répétez cette étape pour toutes les machines que vous souhaitez surveiller.

### 2. Lancer le serveur

1.  Accédez au répertoire du serveur.
2.  Exécutez la commande suivante :
    
    bash
    
    Copier le code
    ```
    go run main.go` 
    ```
3.  Dans le terminal, entrez les adresses IP des clients que vous souhaitez surveiller.  
    Le serveur web démarrera ensuite sur le port `8081`.

> **Note :** Par défaut, la page web du serveur s'actualise toutes les secondes. Vous pouvez modifier cet intervalle à la ligne 59 du fichier `main.go`.
## Création des Modules

Voici les commandes pour configurer les modules Go nécessaires :

### Client Windows

```go mod init Go_ressource_viewver/client/windows
go get github.com/gin-gonic/gin
go get github.com/shirou/gopsutil/cpu
go get github.com/shirou/gopsutil/mem
go get github.com/shirou/gopsutil/disk
```
### Client Linux

```
go mod init Go_ressource_viewver/client/linux
go get github.com/gin-gonic/gin
go get github.com/shirou/gopsutil/v3/cpu
go get github.com/shirou/gopsutil/v3/mem
go get github.com/shirou/gopsutil/v3/disk
```
## Aides
**ChatGPT** : Pour l'assistance dans la recherche et l'utilisation des modules nécessaires pour les clients.
