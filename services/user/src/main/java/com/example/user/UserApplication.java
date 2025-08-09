package com.example.user;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * User Service Application
 * 
 * This service handles user management operations including:
 * - User registration and authentication
 * - User profile management
 * - User data retrieval
 */
@SpringBootApplication
public class UserApplication {
    
    public static void main(String[] args) {
        SpringApplication.run(UserApplication.class, args);
    }
}
