package de.thekolo.mybudgetbook.models.account;

public class AccountMapper {
    public static Account toEntity(AccountDto dto) {
        return Account.builder()
                .name(dto.getName())
                .build();
    }

    public static AccountDto toDto(Account entity) {
        return AccountDto.builder()
                .id(entity.getId())
                .name(entity.getName())
                .created(entity.getCreated())
                .updated(entity.getUpdated())
                .build();
    }
}
