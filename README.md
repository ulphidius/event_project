# Event Project

## Description

Ubiquitous Language:

- event (Evénement publicitaire)
- statistics (Statistiques sur les événements publicitaires)
- type (Type d'événement publicitaire)
    - Impression
    - Click
    - Visible

### Prérequis

- Go 1.13+
- Go modules

### Lancer le projet

Les fichiers de paramétrage sont les fichier **config.*.json**.
**config.json** est utilisé pour lancer l'application normalement et **config.test.json** est utilisé pour le lancement des tests.

#### Docker

```bash
docker pull mariadb
docker build --tag event_project .
docker network create event_network

# Les paramètres pour la création de la base de données et pour l'ouverture des ports dépendent des valeurs en configuration
docker run --interactive --tty --detach --name event-project-db --env MARIADB_DATABASE=event --env MARIADB_USER=event --env MARIADB_PASSWORD=secret --env MARIADB_ROOT_PASSWORD=root --network=event_network --publish 8306:3306 mariadb
docker run --interactive --tty --network=event_network --publish 8080:8080 --name event-project-api event_project
```

#### Local

```bash
go run main.go
```

### Utilisation

Deux routes sont disponibles:

- POST /event
- GET /statistics

La route **/event** permet la création en base des **event**.
**event** attend comme donnée:
```json
// timestamp doit être une valeur numérique supérieure à zéro correspondant au temps écoulé depuis la création de Unix
// type doit être une de ces trois valeurs: 
// - Impression
// - Click
// - Visible
{
    "timestamp": 0,
    "type": "Impression"
}
```
La route **/statistics** permet la récupération des **statistics**.
**statistics** attend en paramètre:
```bash
# type permet de choisir quel statistics nous souhaitons obtenir (par type ou par os). La valeur attendue doit être soit os soit type
# min et max correspond à un intervalle de temps qui sert de filtre pour les données
/statistics?type=type&min=1420070400&max=1577836800
```

#### Requêtes

```bash
curl -H "Content-Type: application/json" -X POST -d '{"timestamp": 0, "type": "Impression"}' http://localhost:8080/event
curl "http://127.0.0.1:8080/statistics?type=type&min=1420070400&max=1577836800"
```
