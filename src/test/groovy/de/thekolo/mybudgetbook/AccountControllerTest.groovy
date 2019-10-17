package de.thekolo.mybudgetbook

import de.thekolo.mybudgetbook.models.account.Account
import groovy.json.JsonOutput
import org.springframework.http.HttpHeaders
import org.springframework.http.MediaType
import org.springframework.test.web.servlet.ResultActions
import org.springframework.test.web.servlet.result.MockMvcResultHandlers

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.post
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status

class AccountControllerTest extends MybudgetbookApplicationTests {

    def "should be able create an account"() {
        given:
        def body = [name: "Mein Konto"]

        when:
        ResultActions result = mockMvc.perform(post("/api/accounts")
                .header(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE)
                .content(JsonOutput.toJson(body)))

        then:
        result.andExpect(status().isCreated())
                .andExpect(jsonPath("\$.id").exists())
                .andExpect(jsonPath("\$.name").value("Mein Konto"))
                .andExpect(jsonPath("\$.created").exists())
                .andExpect(jsonPath("\$.updated").exists())

        and:
        accountRepository.findAll().size() == 1
        Account account = accountRepository.findAll()[0]
        account.id != null
        account.name == "Mein Konto"
        account.created != null
        account.updated != null
    }

    def "should not be able create an account with empty name"() {
        given:
        def body = [name: ""]

        when:
        ResultActions result = mockMvc.perform(post("/api/accounts")
                .header(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE)
                .content(JsonOutput.toJson(body)))

        then:
        result.andDo(MockMvcResultHandlers.print())
                .andExpect(status().isBadRequest())
                .andExpect(jsonPath("\$.errors.length()").value(1))
                .andExpect(jsonPath("\$.errors[0].fieldName").value("name"))
                .andExpect(jsonPath("\$.errors[0].message").value("Name is mandatory"))

        and:
        accountRepository.findAll().size() == 0
    }

    def "should be able create an account with duplicate name"() {
        given:
        def body = [name: "Mein Konto"]

        when:
        ResultActions resultOne = mockMvc.perform(post("/api/accounts")
                .header(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE)
                .content(JsonOutput.toJson(body)))

        then:
        resultOne.andExpect(status().isCreated())
                .andExpect(jsonPath("\$.id").exists())
                .andExpect(jsonPath("\$.name").value("Mein Konto"))
                .andExpect(jsonPath("\$.created").exists())
                .andExpect(jsonPath("\$.updated").exists())

        and:
        accountRepository.findAll().size() == 1
        Account account = accountRepository.findAll()[0]
        account.id != null
        account.name == "Mein Konto"
        account.created != null
        account.updated != null

        when:
        ResultActions resultTwo = mockMvc.perform(post("/api/accounts")
                .header(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE)
                .content(JsonOutput.toJson(body)))

        then:
        resultTwo.andExpect(status().isBadRequest())
                .andExpect(jsonPath("\$.errors.length()").value(1))
                .andExpect(jsonPath("\$.errors[0].fieldName").value(""))
                .andExpect(jsonPath("\$.errors[0].message").value("Account with name 'Mein Konto' already exists"))
    }

    def "should return a paged list of accounts"() {
        given:
        (0..1).forEach {
            accountRepository.save(Account.builder().name("$it").build())
        }

        when:
        ResultActions result = mockMvc.perform(get("/api/accounts")
                .header(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE))

        then:
        result.andDo(MockMvcResultHandlers.print()).andExpect(status().isOk())
                .andExpect(jsonPath("\$.content.length()").value(2))
                .andExpect(jsonPath("\$.totalElements").value(2))
                .andExpect(jsonPath("\$.size").value(20))
                .andExpect(jsonPath("\$.totalPages").value(1))
                .andExpect(jsonPath("\$.content[0].id").exists())
                .andExpect(jsonPath("\$.content[0].name").value("0"))
                .andExpect(jsonPath("\$.content[0].created").exists())
                .andExpect(jsonPath("\$.content[0].updated").exists())
                .andExpect(jsonPath("\$.content[1].id").exists())
                .andExpect(jsonPath("\$.content[1].name").value("1"))
                .andExpect(jsonPath("\$.content[1].created").exists())
                .andExpect(jsonPath("\$.content[1].updated").exists())
    }

}
