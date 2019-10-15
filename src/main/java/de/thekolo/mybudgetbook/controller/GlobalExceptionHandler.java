package de.thekolo.mybudgetbook.controller;

import de.thekolo.mybudgetbook.models.error.ErrorDto;
import de.thekolo.mybudgetbook.models.error.ErrorsDto;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.HttpStatus;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.ResponseStatus;

import java.util.ArrayList;
import java.util.List;

@Slf4j
@ControllerAdvice
public class GlobalExceptionHandler {

    @ResponseBody
    @ResponseStatus(HttpStatus.BAD_REQUEST)
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ErrorsDto exceptionHandler(MethodArgumentNotValidException e) {
        log.info(e.getMessage(), e);
        List<ErrorDto> errors = new ArrayList<>();

        e.getBindingResult().getAllErrors().forEach(error -> {
            errors.add(new ErrorDto(((FieldError) error).getField(), error.getDefaultMessage()));
        });

        return new ErrorsDto(errors);
    }
}
