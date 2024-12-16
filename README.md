This application is a small scale application for the purposes of showcasing simple use of the RabbitMQ message broker for AMQP.
When running Rabbit MQ as a service, the application will run a goroutine (seperate thread) to immediately start consuming messages on the CafeOfRestQ
The user can input what they would like to order (coffee, temp, milk) and once done, the order will be sent to the queue.
The consumer will recieve and output the coffee order is ready.
