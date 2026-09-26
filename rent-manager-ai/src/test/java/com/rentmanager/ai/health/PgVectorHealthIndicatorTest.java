package com.rentmanager.ai.health;

import org.junit.jupiter.api.Test;
import org.springframework.boot.health.contributor.Status;
import org.springframework.dao.DataAccessResourceFailureException;
import org.springframework.jdbc.core.JdbcTemplate;

import static org.assertj.core.api.Assertions.assertThat;
import static org.mockito.ArgumentMatchers.anyString;
import static org.mockito.ArgumentMatchers.eq;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;

final class PgVectorHealthIndicatorTest {
    private final JdbcTemplate jdbc = mock(JdbcTemplate.class);
    private final PgVectorHealthIndicator indicator = new PgVectorHealthIndicator(jdbc);

    @Test
    void reportsUpOnlyWhenExtensionIsInstalled() {
        when(jdbc.queryForObject(anyString(), eq(Boolean.class))).thenReturn(true);
        assertThat(indicator.health().getStatus()).isEqualTo(Status.UP);
        when(jdbc.queryForObject(anyString(), eq(Boolean.class))).thenReturn(false);
        assertThat(indicator.health().getStatus()).isEqualTo(Status.DOWN);
    }

    @Test
    void reportsDatabaseFailureWithoutLeakingConnectionDetails() {
        when(jdbc.queryForObject(anyString(), eq(Boolean.class)))
                .thenThrow(new DataAccessResourceFailureException("sensitive connection details"));
        final var health = indicator.health();
        assertThat(health.getStatus()).isEqualTo(Status.DOWN);
        assertThat(health.getDetails().toString()).doesNotContain("sensitive");
    }
}
