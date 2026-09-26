package com.rentmanager.ai.health;

import java.io.IOException;
import java.net.InetSocketAddress;

import com.sun.net.httpserver.HttpServer;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.boot.health.contributor.Status;

import static org.assertj.core.api.Assertions.assertThat;

final class OllamaHealthIndicatorTest {
    private HttpServer server;
    private OllamaHealthIndicator indicator;

    @BeforeEach
    void startServer() throws IOException {
        server = HttpServer.create(new InetSocketAddress("127.0.0.1", 0), 0);
        server.start();
        indicator = new OllamaHealthIndicator("http://127.0.0.1:" + server.getAddress().getPort());
    }

    @AfterEach
    void stopServer() {
        server.stop(0);
    }

    @Test
    void checksModelListEndpointWithoutInvokingAModel() {
        server.createContext("/api/tags", exchange -> {
            exchange.sendResponseHeaders(200, -1);
            exchange.close();
        });
        assertThat(indicator.health().getStatus()).isEqualTo(Status.UP);
    }

    @Test
    void reportsUnhealthyOnServerError() {
        server.createContext("/api/tags", exchange -> {
            exchange.sendResponseHeaders(503, -1);
            exchange.close();
        });
        assertThat(indicator.health().getStatus()).isEqualTo(Status.DOWN);
    }

    @Test
    void reportsUnhealthyOnConnectionFailure() {
        server.stop(0);
        assertThat(indicator.health().getStatus()).isEqualTo(Status.DOWN);
    }
}
