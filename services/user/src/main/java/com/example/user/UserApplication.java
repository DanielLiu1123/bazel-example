package com.example.user;

import com.example.user.converter.UserConverter;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import java.util.List;

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

        var user = UserModel.newBuilder()
                .setId("1")
                .setEmail("test@example.com")
                .setUsername("testuser")
                .setFirstName("John")
                .setLastName("Doe")
                .addAllHobbies(List.of("Reading", "Coding"))
                .build();

        var userEntity = UserConverter.INSTANCE.modelToEntity(user);
        System.out.println(userEntity);

        user = UserConverter.INSTANCE.entityToModel(userEntity);
        System.out.println(user);
    }
}
