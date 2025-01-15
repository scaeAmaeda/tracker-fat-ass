package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/sqlite"
)

type Meal struct {
	Identifiant int
	Date        string
	Moment      string
	Is300g      int
	IsNoSugar   int
}

func ConnectDB(dbCible string) *sql.DB {
	db, err := sql.Open("sqlite", dbCible)
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la base de données : %v", err)
	}
	// Vérifie que la connexion est opérationnelle
	if err := db.Ping(); err != nil {
		log.Fatalf("Impossible de se connecter à la base : %v", err)
	}
	return db
}

func GetMeals(db *sql.DB) []Meal {
	// query de tout sur la base
	rows, err := db.Query("SELECT identifiant, date, moment, is300g, isnosugar FROM Meals")
	if err != nil {
		log.Fatalf("Erreur lors du select : %v", err)
	}
	defer rows.Close()
	// initialisation du retour
	AllMeals := []Meal{}
	//boucle sur chaque ligne, qui est ajoutée au retour
	for rows.Next() {
		var currentLine Meal
		var identifiant int
		var date string
		var moment string
		var is300g int
		var isNoSugar int
		err = rows.Scan(&identifiant, &date, &moment, &is300g, &isNoSugar)
		if err != nil {
			log.Fatalf("Erreur lors de la lecture des lignes : %v", err)
		}
		currentLine.Identifiant = identifiant
		currentLine.Date = date
		currentLine.Moment = moment
		currentLine.Is300g = is300g
		currentLine.IsNoSugar = isNoSugar
		AllMeals = append(AllMeals, currentLine)
	}
	return AllMeals
}

func AddMeal(db *sql.DB, newOne Meal) error {
	// Préparer la requête d'insertion
	stmt, err := db.Prepare("INSERT INTO Meals (date, moment, is300g, isnosugar) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("erreur lors de la préparation de la requête INSERT : %w", err)
	}
	defer stmt.Close()

	date := newOne.Date
	moment := newOne.Moment
	is300g := newOne.Is300g
	isnosugar := newOne.IsNoSugar

	// Exécuter la requête avec les paramètres
	_, err = stmt.Exec(date, moment, is300g, isnosugar)
	if err != nil {
		return fmt.Errorf("erreur lors de l'exécution de la requête INSERT : %w", err)
	}
	fmt.Println("Nouvel enregistrement inséré avec succès !")
	return nil
}
