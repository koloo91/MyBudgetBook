package de.thekolo.mybudgetbook

import de.thekolo.mybudgetbook.repositories.AccountRepository
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.test.context.ActiveProfiles
import org.springframework.test.web.servlet.MockMvc
import spock.lang.Specification

@ActiveProfiles("test")
@SpringBootTest
@AutoConfigureMockMvc
class MybudgetbookApplicationTests extends Specification {

    @Autowired
    MockMvc mockMvc;

    @Autowired
    AccountRepository accountRepository

    def setup() {
        accountRepository.deleteAll()
    }
}
