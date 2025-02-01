create table xf_users
(
    user_uuid  char(36)                not null comment '用户主键'
        primary key,
    username   varchar(32)             not null comment '用户名',
    email      varchar(255)            not null comment '用户邮箱',
    password   char(60)                not null comment '用户密码',
    role       char(36)                not null comment '用户所在角色',
    is_enable  boolean   default true  not null comment '是否启用用户',
    is_ban     boolean   default false not null comment '是否封禁',
    created_at timestamp default now() not null comment '创建时间',
    updated_at timestamp default now() not null on update now() comment '更新时间',
    constraint xf_users_xf_roles_role_uuid_fk
        foreign key (role) references xf_roles (role_uuid)
            on update cascade
)
    comment '用户表';

