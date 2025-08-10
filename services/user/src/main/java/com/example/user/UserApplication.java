package com.example.user;

import com.example.user.converter.UserConverter;
import com.example.user.entity.UserEntity;
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

        var user = User.newBuilder()
                .setId("1")
                .setEmail("test@example.com")
                .setUsername("testuser")
                .setFirstName("John")
                .setLastName("Doe")
                .build();

        var userEntity = UserConverter.INSTANCE.modelToEntity(user);
        System.out.println(userEntity);

        user = UserConverter.INSTANCE.entityToModel(userEntity);
        System.out.println(user);
    }
}
