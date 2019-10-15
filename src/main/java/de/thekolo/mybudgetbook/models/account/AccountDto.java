package de.thekolo.mybudgetbook.models.account;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.validation.constraints.NotBlank;
import java.time.Instant;

@Data
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class AccountDto {

    @JsonProperty("id")
    private String id;

    @JsonProperty(value = "name", required = true)
    @NotBlank(message = "Name is mandatory")
    private String name;

    @JsonProperty("created")
    private Instant created;

    @JsonProperty("updated")
    private Instant updated;
}
