package com.example.user.grpc;

import com.example.user.GetUserRequest;
import com.example.user.GetUserResponse;
import com.example.user.UserServiceGrpc;
import io.grpc.stub.StreamObserver;
import org.springframework.stereotype.Component;

/**
 * gRPC Service implementation for User operations
 * TODO: Enable when proto files are properly generated
 */
@Component
public class UserGrpcService extends UserServiceGrpc.UserServiceImplBase {

    @Override
    public void getUser(GetUserRequest request, StreamObserver<GetUserResponse> responseObserver) {
        super.getUser(request, responseObserver);
    }
}
