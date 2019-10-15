package de.thekolo.mybudgetbook.models;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.time.Instant;

import javax.validation.constraints.NotBlank;

import com.fasterxml.jackson.annotation.JsonProperty;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class AccountDto {

    @JsonProperty("id")
    private String id;

    @JsonProperty(value = "name", required = true)
    @NotBlank(message = "Name can not be empty")
    private String name;

    @JsonProperty("created")
    private Instant created;

    @JsonProperty("updated")
    private Instant updated;
}
