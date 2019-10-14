package de.thekolo.mybudgetbook

import de.thekolo.mybudgetbook.repositories.AccountRepository
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.test.context.ActiveProfiles
import spock.lang.Specification

@ActiveProfiles("test")
@SpringBootTest
class MybudgetbookApplicationTests extends Specification {

    @Autowired
    AccountRepository accountRepository

    def setup() {
        accountRepository.deleteAll()
    }

    def "dummy spec"() {
        when:
        true

        then:
        true
    }
}
