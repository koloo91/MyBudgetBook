package de.thekolo.mybudgetbook.controller;

import javax.validation.Valid;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import de.thekolo.mybudgetbook.models.Account;
import de.thekolo.mybudgetbook.models.AccountDto;

@RestController
@RequestMapping("/api")
public class AccountController {

    @PostMapping()
    public ResponseEntity<AccountDto> create(@Valid @RequestBody AccountDto accountDto) {
        return ResponseEntity.created(null).body(accountDto);
    }
}
