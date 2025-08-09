package com.example.user.grpc;

import com.example.user.GetUserRequest;
import com.example.user.GetUserResponse;
import com.example.user.UserServiceGrpc;
import com.example.user.service.UserService;
import io.grpc.stub.StreamObserver;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * gRPC Service implementation for User operations
 * TODO: Enable when proto files are properly generated
 */
@Component
public class UserGrpcService extends UserServiceGrpc.UserServiceImplBase {
    
    private static final Logger logger = LoggerFactory.getLogger(UserGrpcService.class);
    
    @Autowired
    private UserService userService;
    
    // TODO: Implement gRPC methods when proto files are generated
    public void placeholder() {
        logger.info("UserGrpcService placeholder - gRPC implementation will be added when proto files are ready");
    }

    @Override
    public void getUser(GetUserRequest request, StreamObserver<GetUserResponse> responseObserver) {
        super.getUser(request, responseObserver);
    }
}
