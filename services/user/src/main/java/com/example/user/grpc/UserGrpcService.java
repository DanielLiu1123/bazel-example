package com.example.user.grpc;

import com.example.user.BatchDeleteUserRequest;
import com.example.user.BatchDeleteUserResponse;
import com.example.user.BatchGetUserRequest;
import com.example.user.BatchGetUserResponse;
import com.example.user.DeleteUserRequest;
import com.example.user.DeleteUserResponse;
import com.example.user.GetUserRequest;
import com.example.user.GetUserResponse;
import com.example.user.UserServiceGrpc;
import io.grpc.stub.StreamObserver;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

/**
 * gRPC Service implementation for User operations
 * TODO: Enable when proto files are properly generated
 */
@Component
@Slf4j
public class UserGrpcService extends UserServiceGrpc.UserServiceImplBase {

    @Override
    public void getUser(GetUserRequest request, StreamObserver<GetUserResponse> responseObserver) {
        log.info("Getting user: {}", request.getId());
        super.getUser(request, responseObserver);
    }

    @Override
    public void batchGetUser(BatchGetUserRequest request, StreamObserver<BatchGetUserResponse> responseObserver) {
        log.info("Getting users: {}", request.getIdsList());
        super.batchGetUser(request, responseObserver);
    }

    @Override
    public void batchDeleteUser(BatchDeleteUserRequest request, StreamObserver<BatchDeleteUserResponse> responseObserver) {
        log.info("Deleting users: {}", request.getIdsList());
        super.batchDeleteUser(request, responseObserver);
    }

    @Override
    public void deleteUser(DeleteUserRequest request, StreamObserver<DeleteUserResponse> responseObserver) {
        super.deleteUser(request, responseObserver);
    }
}
