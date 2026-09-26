package com.rentmanager.ai.adapter.out.health;

import org.springframework.boot.health.contributor.Health;
import org.springframework.boot.health.contributor.HealthIndicator;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Component;

@Component("pgvector")
public final class PgVectorHealthIndicator implements HealthIndicator {
    private final JdbcTemplate jdbc;

    public PgVectorHealthIndicator(final JdbcTemplate jdbc) {
        this.jdbc = jdbc;
    }

    @Override
    public Health health() {
        try {
            final Boolean installed = jdbc.queryForObject(
                    "SELECT EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'vector')", Boolean.class);
            return Boolean.TRUE.equals(installed) ? Health.up().build()
                    : Health.down().withDetail("reason", "vector extension is not installed").build();
        } catch (final Exception exception) {
            return Health.down().withDetail("reason", "PostgreSQL is unavailable").build();
        }
    }
}
