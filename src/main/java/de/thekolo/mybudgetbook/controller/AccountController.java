package de.thekolo.mybudgetbook.controller;

import de.thekolo.mybudgetbook.models.account.Account;
import de.thekolo.mybudgetbook.models.account.AccountDto;
import de.thekolo.mybudgetbook.models.account.AccountMapper;
import de.thekolo.mybudgetbook.services.AccountService;
import lombok.RequiredArgsConstructor;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;

@RestController
@RequestMapping("/api/accounts")
@RequiredArgsConstructor
public class AccountController {

    private final AccountService accountService;

    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE, produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<AccountDto> createAccount(@Valid @RequestBody AccountDto accountDto) {
        Account account = AccountMapper.toEntity(accountDto);
        account = accountService.create(account);

        return ResponseEntity.status(HttpStatus.CREATED)
                .body(AccountMapper.toDto(account));
    }

    @GetMapping(produces = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<Page<AccountDto>> getAccounts(Pageable pageable) {

        return ResponseEntity.status(HttpStatus.OK)
                .body(accountService.get(pageable)
                        .map(AccountMapper::toDto));
    }
}
