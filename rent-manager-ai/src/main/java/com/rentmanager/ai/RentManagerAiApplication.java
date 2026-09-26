package com.rentmanager.ai;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication(proxyBeanMethods = false)
public final class RentManagerAiApplication {
    public static void main(final String[] args) {
        SpringApplication.run(RentManagerAiApplication.class, args);
    }
}
