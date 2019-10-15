package de.thekolo.mybudgetbook

import org.springframework.test.web.servlet.request.MockMvcRequestBuilders

class AccountControllerTest extends MybudgetbookApplicationTests {
    def "should create an account"() {
        mockMvc.perform(MockMvcRequestBuilders.post()
                .content())
    }
}
