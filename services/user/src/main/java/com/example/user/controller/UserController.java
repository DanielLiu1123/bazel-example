package com.example.user.controller;

import com.example.user.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

/**
 * REST Controller for User operations
 * 
 * Provides HTTP endpoints for user management
 */
@RestController
@RequestMapping("/api/users")
public class UserController {
    
    @Autowired
    private UserService userService;
    
    @GetMapping("/{id}")
    public String getUser(@PathVariable String id) {
        return "User: " + id;
    }
    
    @PostMapping
    public String createUser(@RequestBody String userData) {
        return "Created user: " + userData;
    }
    
    @PutMapping("/{id}")
    public String updateUser(@PathVariable String id, @RequestBody String userData) {
        return "Updated user: " + id;
    }
    
    @DeleteMapping("/{id}")
    public String deleteUser(@PathVariable String id) {
        return "Deleted user: " + id;
    }
}
