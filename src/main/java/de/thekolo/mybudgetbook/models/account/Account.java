package de.thekolo.mybudgetbook.models.account;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.time.Instant;
import java.util.UUID;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "accounts")
public class Account {

    @Id
    private String id;

    private String name;

    private Instant created;

    private Instant updated;

    @PrePersist
    public void prePersist() {
        id = UUID.randomUUID().toString();
        created = Instant.now();
        updated = created;
    }

    @PreUpdate
    public void preUpdate() {
        updated = Instant.now();
    }
}
