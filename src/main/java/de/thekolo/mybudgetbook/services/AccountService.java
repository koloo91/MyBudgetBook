package de.thekolo.mybudgetbook.services;

import de.thekolo.mybudgetbook.models.account.Account;
import de.thekolo.mybudgetbook.repositories.AccountRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class AccountService {

    private final AccountRepository accountRepository;

    public Account create(Account account) {
        return accountRepository.save(account);
    }
}
