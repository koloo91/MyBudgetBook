package de.thekolo.mybudgetbook.models.error;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.ArrayList;
import java.util.List;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class ErrorsDto {

    @JsonProperty("errors")
    private List<ErrorDto> errors = new ArrayList<>();
}
