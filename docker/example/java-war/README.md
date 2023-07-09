# Java War

1. fetcher dir

    - `docker build -t jamtur01/fetcher .`

2. fetch url's war

    - `docker run -it --name sample jamtur01/fetcher https://tomcat.apache.org/tomcat-7.0-doc/appdev/sample/sample.war`

3. tomcat7 dir

    - `docker build -t jamtur01/tomcat7 .`

4. Running war file

    - `docker run --name sample_app --volumes-from sample -d -P jamtur01/tomcat7`

    - `docker port sample_app`

5. Building on top of our Tomcat application server

    - ```bash
      sudo apt-get update
      sudo apt-get install ruby make ruby-dev
      sudo gem install --no-document tprov
      sudo tprov
      curl localhost:4567
      ```

    - TProv is pretty simple, to demonstrate how powerful Docker can be used as a tool
    for building applications and services. [source code](https://github.com/turnbullpress/dockerbook-code/blob/master/code/6/tomcat/tprov/lib/tprov/app.rb)
