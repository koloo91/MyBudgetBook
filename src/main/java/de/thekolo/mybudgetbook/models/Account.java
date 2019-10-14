package de.thekolo.mybudgetbook.models;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.PrePersist;
import javax.persistence.PreUpdate;
import java.time.Instant;
import java.util.UUID;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
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
