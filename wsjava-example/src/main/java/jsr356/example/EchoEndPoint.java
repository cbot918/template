package jsr356.example;

import java.io.IOException;
import java.util.Set;
import java.util.concurrent.CopyOnWriteArraySet;
import java.util.logging.Logger;

import javax.websocket.OnClose;
import javax.websocket.OnMessage;
import javax.websocket.OnOpen;
import javax.websocket.Session;
import javax.websocket.server.ServerEndpoint;

@ServerEndpoint(value = "/echo")
public class EchoEndPoint {

    private static final Logger log = Logger.getLogger(EchoEndPoint.class.getName());

    private static final Set<Session> sessions = new CopyOnWriteArraySet<>();

    @OnOpen
    public void onOpen(Session session) {
        log.info("{" + session.getId() +  "} connected");
        sessions.add(session);
    }

    @OnClose
    public void onClose(Session session) {
        log.info("{" + session.getId() +  "} closed");
        sessions.remove(session);
    }

    @OnMessage
    public void echo(String message) throws IOException {
        for (Session session : sessions) {
            if (session.isOpen()) {
                log.info("{" + session.getId() + "} echo '" + message + "'");
                session.getBasicRemote().sendText(message);
            }
        }
    }

}