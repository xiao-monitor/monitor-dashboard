create table xf_roles
(
    role_uuid  char(36)                            not null comment '角色 UUID'
        primary key,
    role_name  varchar(16)                         not null comment '角色名',
    created_at timestamp default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint xf_roles_role_name_uindex
        unique (role_name)
)
    comment '角色表';