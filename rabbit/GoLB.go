package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io/ioutil"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	var result map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println(err)
	}
	json.Unmarshal([]byte(body), &result)
	fmt.Println(result)

	/**
	*rabbit
	*/
	conn, err := amqp.Dial("amqp://guest:guest@138.68.230.101:5672/")
	failOnError(err, "Fallo RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Fallo al abrir el canal de conexion ")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"proyecto2", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Fallo al crear la cola")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "application/json",
			Body:        []byte(body),
		})
	failOnError(err, "Fallo al publicar el mensaje")
	/**
	*rabbit
	*/
	fmt.Fprintf(w, "%s", "termino")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Server Running on port: 8090")
	router.HandleFunc("/", info).Methods("POST")
    handler := cors.Default().Handler(router)
    http.ListenAndServe(":8090", handler)
	
	log.Fatal(http.ListenAndServe(":8090", router))
}