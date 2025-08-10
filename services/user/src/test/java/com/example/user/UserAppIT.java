
package com.example.user;

import grpcstarter.server.GrpcServer;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest
public class UserAppIT {

    @Autowired
    GrpcServer grpcServer;

    @Test
    void contextLoads() {
        assertThat(grpcServer).isNotNull();
    }
}
