package com.example.user.converter;

import com.example.user.UserModel;
import com.example.user.entity.User;
import org.mapstruct.Mapper;
import org.mapstruct.factory.Mappers;

import static org.mapstruct.CollectionMappingStrategy.ADDER_PREFERRED;
import static org.mapstruct.NullValueCheckStrategy.ALWAYS;
import static org.mapstruct.ReportingPolicy.IGNORE;

/**
 *
 *
 * @author Freeman
 * @since 2025/8/10
 */
@Mapper(
        nullValueCheckStrategy = ALWAYS,
        collectionMappingStrategy = ADDER_PREFERRED,
        unmappedTargetPolicy = IGNORE
)
public abstract class UserConverter {

    public static final UserConverter INSTANCE = Mappers.getMapper(UserConverter.class);

    public abstract User modelToEntity(UserModel user);

    public abstract UserModel entityToModel(User user);
}
