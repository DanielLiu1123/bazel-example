package com.example.order;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * Order Service Application
 * 
 * This service handles order management operations including:
 * - Order creation and processing
 * - Order status management
 * - Order retrieval and listing
 */
@SpringBootApplication
public class OrderApplication {
    
    public static void main(String[] args) {
        SpringApplication.run(OrderApplication.class, args);
    }
}
