package com.apigee.examples.grpc.service;

import com.apigee.examples.grpc.server.grpcserver.MessengerRequest;
import com.apigee.examples.grpc.server.grpcserver.MessengerResponse;
import com.apigee.examples.grpc.server.grpcserver.MessengerServiceGrpc;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;

import java.util.HashMap;
import java.util.Map;

@GrpcService
public class MessengerServiceImpl extends MessengerServiceGrpc.MessengerServiceImplBase {

    @Override
    public void getGreeting(MessengerRequest request, StreamObserver<MessengerResponse> responseObserver) {
        String msg = request.getMsg();
        MessengerResponse response = MessengerResponse.newBuilder()
            .setMsg(String.format("Gosh! You just said '%s' to me", msg))
            .build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getPirateGreeting(MessengerRequest request, StreamObserver<MessengerResponse> responseObserver) {
        String msg = request.getMsg();
        MessengerResponse response = MessengerResponse.newBuilder()
            .setMsg(String.format("Aaargh.... Avast ye matey! Dead men tell no tales like: '%s'",msg))
            .build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

}
