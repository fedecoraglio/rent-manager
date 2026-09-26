package com.rentmanager.ai.health;

import java.time.Duration;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.health.contributor.Health;
import org.springframework.boot.health.contributor.HealthIndicator;
import org.springframework.http.client.SimpleClientHttpRequestFactory;
import org.springframework.stereotype.Component;
import org.springframework.web.client.RestClient;

@Component("ollama")
public final class OllamaHealthIndicator implements HealthIndicator {
    private final RestClient client;

    public OllamaHealthIndicator(@Value("${spring.ai.ollama.base-url}") final String baseUrl) {
        final var factory = new SimpleClientHttpRequestFactory();
        factory.setConnectTimeout(Duration.ofSeconds(2));
        factory.setReadTimeout(Duration.ofSeconds(3));
        this.client = RestClient.builder().baseUrl(baseUrl).requestFactory(factory).build();
    }

    @Override
    public Health health() {
        try {
            // Connectivity only: never load a model or generate tokens during a health probe.
            client.get().uri("/api/tags").retrieve().toBodilessEntity();
            return Health.up().build();
        } catch (final Exception exception) {
            return Health.down().withDetail("reason", "Ollama is unavailable").build();
        }
    }
}
