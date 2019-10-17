package de.thekolo.mybudgetbook.services;

import de.thekolo.mybudgetbook.exceptions.DuplicateAccountNameException;
import de.thekolo.mybudgetbook.models.account.Account;
import de.thekolo.mybudgetbook.repositories.AccountRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;


@Service
@RequiredArgsConstructor
public class AccountService {

    private final AccountRepository accountRepository;

    public Account create(Account account) {
        try {
            return accountRepository.save(account);
        } catch (DataIntegrityViolationException e) {
            throw new DuplicateAccountNameException(account.getName());
        }
    }

    public Page<Account> get(Pageable pageable) {
        return accountRepository.findAll(pageable);
    }
}
