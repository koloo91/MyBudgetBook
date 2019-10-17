package de.thekolo.mybudgetbook.repositories;

import de.thekolo.mybudgetbook.models.account.Account;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface AccountRepository extends CrudRepository<Account, String> {
    Page<Account> findAll(Pageable pageable);
}
