package de.thekolo.mybudgetbook.repositories;

import de.thekolo.mybudgetbook.models.Account;
import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface AccountRepository extends CrudRepository<Account, String> {
}
