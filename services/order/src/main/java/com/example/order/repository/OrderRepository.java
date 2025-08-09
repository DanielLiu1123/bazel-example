package com.example.order.repository;

import com.example.order.entity.OrderEntity;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.time.LocalDateTime;
import java.util.List;

/**
 * Repository interface for Order entities
 */
@Repository
public interface OrderRepository extends JpaRepository<OrderEntity, String> {
    
    /**
     * Find orders by user ID
     */
    Page<OrderEntity> findByUserId(String userId, Pageable pageable);
    
    /**
     * Find orders by status
     */
    Page<OrderEntity> findByStatus(OrderEntity.OrderStatus status, Pageable pageable);
    
    /**
     * Find orders by user ID and status
     */
    Page<OrderEntity> findByUserIdAndStatus(String userId, OrderEntity.OrderStatus status, Pageable pageable);
    
    /**
     * Find orders created between dates
     */
    Page<OrderEntity> findByCreatedAtBetween(LocalDateTime startDate, LocalDateTime endDate, Pageable pageable);
    
    /**
     * Find orders by user ID created between dates
     */
    Page<OrderEntity> findByUserIdAndCreatedAtBetween(String userId, LocalDateTime startDate, LocalDateTime endDate, Pageable pageable);
    
    /**
     * Count orders by status
     */
    long countByStatus(OrderEntity.OrderStatus status);
    
    /**
     * Count orders by user ID
     */
    long countByUserId(String userId);
    
    /**
     * Find recent orders for a user
     */
    @Query("SELECT o FROM OrderEntity o WHERE o.userId = :userId ORDER BY o.createdAt DESC")
    List<OrderEntity> findRecentOrdersByUserId(@Param("userId") String userId, Pageable pageable);
    
    /**
     * Find orders with total amount greater than specified value
     */
    @Query("SELECT o FROM OrderEntity o WHERE o.totalAmount > :amount")
    Page<OrderEntity> findOrdersWithAmountGreaterThan(@Param("amount") java.math.BigDecimal amount, Pageable pageable);
    
    /**
     * Get order statistics
     */
    @Query("SELECT o.status, COUNT(o) FROM OrderEntity o GROUP BY o.status")
    List<Object[]> getOrderStatistics();
}
