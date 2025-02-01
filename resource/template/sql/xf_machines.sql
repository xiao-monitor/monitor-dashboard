create table xf_machines
(
    machine_uuid        char(36)                             not null comment '机器主键'
        primary key,
    machine_name        varchar(64)                          not null comment '机器备注名字',
    machine_description text                                 null comment '机器描述',
    machine_tags        json                                 null comment '机器标签',
    machine_region      varchar(100)                         null,
    host_system         varchar(32)                          not null comment '机器系统(Linux,Windows,MacOS)',
    host_platform       varchar(64)                          not null comment '机器平台(Ubuntu,Centos)',
    host_cpu            varchar(255)                         null comment '机器 CPU',
    host_mem            bigint unsigned                      null comment '机器运存大小',
    host_swap           bigint unsigned                      null comment '机器虚拟内存大小',
    host_virtualization varchar(64)                          null comment '机器虚拟化',
    host_boot           timestamp                            not null comment '机器启动时间',
    is_register         tinyint(1) default 0                 not null comment '是否注册',
    created_at          timestamp  default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at          timestamp  default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '修改时间'
)
    comment '机器';

