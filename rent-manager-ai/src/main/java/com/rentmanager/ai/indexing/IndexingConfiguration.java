package com.rentmanager.ai.indexing;

import java.time.Duration;

import org.springframework.ai.embedding.EmbeddingModel;
import org.springframework.ai.ollama.OllamaEmbeddingModel;
import org.springframework.ai.ollama.api.OllamaApi;
import org.springframework.ai.ollama.api.OllamaEmbeddingOptions;
import org.springframework.ai.vectorstore.VectorStore;
import org.springframework.ai.vectorstore.pgvector.PgVectorStore;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.DependsOn;
import org.springframework.http.client.SimpleClientHttpRequestFactory;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.web.client.RestClient;

@Configuration(proxyBeanMethods = false)
public final class IndexingConfiguration {
    @Bean
    public EmbeddingModel embeddingModel(@Value("${spring.ai.ollama.base-url}") final String baseUrl,
            @Value("${app.indexing.embedding-read-timeout-seconds}") final int timeoutSeconds) {
        final var factory = new SimpleClientHttpRequestFactory();
        factory.setConnectTimeout(Duration.ofSeconds(3));
        factory.setReadTimeout(Duration.ofSeconds(timeoutSeconds));
        final var api = OllamaApi.builder().baseUrl(baseUrl)
                .restClientBuilder(RestClient.builder().requestFactory(factory)).build();
        // Explicit wiring: no chat model bean, automatic model pulls or cloud providers.
        return OllamaEmbeddingModel.builder().ollamaApi(api)
                .options(OllamaEmbeddingOptions.builder().model("nomic-embed-text").truncate(false).build())
                .build();
    }

    @Bean
    @DependsOn("flywayInitializer")
    public VectorStore vectorStore(final JdbcTemplate jdbc, final EmbeddingModel embeddingModel) {
        // This is the SAME JdbcTemplate/DataSource used by the document registry transaction.
        return PgVectorStore.builder(jdbc, embeddingModel)
                .dimensions(768)
                .distanceType(PgVectorStore.PgDistanceType.COSINE_DISTANCE)
                .indexType(PgVectorStore.PgIndexType.NONE)
                .initializeSchema(false)
                .vectorTableValidationsEnabled(true)
                .build();
    }
}
