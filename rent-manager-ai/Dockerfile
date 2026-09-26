FROM maven:3.9.11-eclipse-temurin-21 AS build
WORKDIR /build
COPY pom.xml .
COPY src ./src
RUN --mount=type=cache,target=/root/.m2 mvn -B verify

FROM eclipse-temurin:21-jre-jammy
RUN groupadd --system app && useradd --system --gid app app \
    && mkdir -p /app/storage && chown -R app:app /app
WORKDIR /app
COPY --from=build --chown=app:app /build/target/rent-manager-ai-0.1.0-SNAPSHOT.jar app.jar
USER app
EXPOSE 8081
ENTRYPOINT ["java", "-jar", "app.jar"]
