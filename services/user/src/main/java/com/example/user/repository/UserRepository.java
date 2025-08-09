package com.example.user.repository;

import com.example.user.entity.UserEntity;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.Optional;

/**
 * Repository interface for User entities
 */
@Repository
public interface UserRepository extends JpaRepository<UserEntity, String> {
    
    /**
     * Find user by email
     */
    Optional<UserEntity> findByEmail(String email);
    
    /**
     * Find user by username
     */
    Optional<UserEntity> findByUsername(String username);
    
    /**
     * Check if email exists
     */
    boolean existsByEmail(String email);
    
    /**
     * Check if username exists
     */
    boolean existsByUsername(String username);
    
    /**
     * Find users by status
     */
    Page<UserEntity> findByStatus(UserEntity.UserStatus status, Pageable pageable);
    
    /**
     * Search users by name or email
     */
    @Query("SELECT u FROM UserEntity u WHERE " +
           "LOWER(u.firstName) LIKE LOWER(CONCAT('%', :search, '%')) OR " +
           "LOWER(u.lastName) LIKE LOWER(CONCAT('%', :search, '%')) OR " +
           "LOWER(u.email) LIKE LOWER(CONCAT('%', :search, '%')) OR " +
           "LOWER(u.username) LIKE LOWER(CONCAT('%', :search, '%'))")
    Page<UserEntity> searchUsers(@Param("search") String search, Pageable pageable);
    
    /**
     * Find active users
     */
    @Query("SELECT u FROM UserEntity u WHERE u.status = 'ACTIVE'")
    Page<UserEntity> findActiveUsers(Pageable pageable);
    
    /**
     * Count users by status
     */
    long countByStatus(UserEntity.UserStatus status);
}
