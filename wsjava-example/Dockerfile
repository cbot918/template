FROM ubuntu

WORKDIR /app

RUN apt update && apt install wget -y

RUN apt install openjdk-8-jdk -y \
    && export PATH="/opt/jdk-13.0.1/bin:$PATH" >> ~/.bashrc
    # && JAVA_HOME='/opt/jdk-13.0.1' \
    # && PATH="$JAVA_HOME/bin:$PATH" \
    # && export PATH
  

RUN wget https://mirrors.estointernet.in/apache/maven/maven-3/3.6.3/binaries/apache-maven-3.6.3-bin.tar.gz \
    && tar -xvf apache-maven-3.6.3-bin.tar.gz \
    && mv apache-maven-3.6.3 /opt/ 
    # && export PATH="/opt/apache-maven-3.6.3/bin:$PATH" >> ~/.bashrc
    # && M2_HOME='/opt/apache-maven-3.6.3' \
    # && PATH="$M2_HOME/bin:$PATH" \
    # && export PATH
    


COPY . /app

EXPOSE 8080

CMD ["/opt/apache-maven-3.6.3/bin/mvn","jetty:run"]