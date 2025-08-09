package com.example.user.service;

import org.springframework.stereotype.Service;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import java.util.UUID;
import java.util.concurrent.ConcurrentHashMap;
import java.util.Map;

/**
 * User Service
 *
 * Business logic for user operations
 */
@Service
public class UserService {

    private static final Logger logger = LoggerFactory.getLogger(UserService.class);

    // In-memory storage for demo purposes
    private final Map<String, Object> users = new ConcurrentHashMap<>();

    public Object getUserById(String id) {
        logger.info("Retrieving user with ID: {}", id);
        return users.get(id);
    }

    public Object createUser(Object request) {
        logger.info("Creating user");

        String userId = UUID.randomUUID().toString();
        users.put(userId, request);
        logger.info("Created user with ID: {}", userId);
        return request;
    }

    public Object updateUser(Object request) {
        logger.info("Updating user");
        return request;
    }

    public boolean deleteUser(String id) {
        logger.info("Deleting user with ID: {}", id);
        Object removed = users.remove(id);
        return removed != null;
    }
}
