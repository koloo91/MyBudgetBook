package de.thekolo.mybudgetbook.models.error;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ErrorDto {

    @JsonProperty("fieldName")
    private String fieldName;

    @JsonProperty("message")
    private String message;
}
