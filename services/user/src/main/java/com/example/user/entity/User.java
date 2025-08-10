package com.example.user.entity;

import lombok.Data;

import java.util.List;

@Data
public class User {
    private String id;
    private String email;
    private String username;
    private String firstName;
    private String lastName;
    private String phone;
    private List<String> hobbies;
}
