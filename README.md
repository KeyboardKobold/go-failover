This is the repository for the code for the Hauptseminar in Sommersemester 18.

Author: Christian Bodenbender

Dependencies:

http://www.gorillatoolkit.org/ for smart routing in go.

How to get started:

Build the Dockerfiles by going into the respective folder where the Dockerfile lives and execute 'docker build -t <tagname> .'
Run them by typing 'docker run <tagname>'.

Or use docker-compose build and docker-compose up from the project root instead.

To create the kubernetes deployments and kubernetes service type 'kubectl create -f <filename>'

How to generate graphviz images:
Just use the generator scripts provided. They should work if you installed graphviz to the default location.
You can call it like 'mkdot <filename>' and it will generate the pictures.

ATTENTION

As I've not gotten around to implementing Kafka + Zookeeper support, the kubernetes deployments only work with 1 pod each. Sadly the pods use the same NodePort the clients are using, so only one consumer will be fed with data.

