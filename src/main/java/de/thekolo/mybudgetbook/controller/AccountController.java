package de.thekolo.mybudgetbook.controller;

import javax.validation.Valid;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import de.thekolo.mybudgetbook.models.Account;

@RestController
@RequestMapping("/api")
public class AccountController {

    @PostMapping()
    public void create(@Valid @RequestBody Account account) {

    }
}
