drop table if exists USER;

/*==============================================================*/
/* Table: USER                                                  */
/*==============================================================*/
create table USER
(
   USER_ID              BIGINT(18) not null auto_increment comment '用户ID',
   USER_NAME            VARCHAR(32) comment '用户姓名',
   PASSWD               VARCHAR(64) comment '密码',
   MOBILE_NUM           VARCHAR(20) comment '手机号码',
   EMAIL                VARCHAR(64) comment '电子邮件',
   STATE                NUMERIC(1) comment '状态',
   CREATE_TIME          DATETIME comment '创建时间',
   primary key (USER_ID)
);


drop table if exists ARTICLE_CLASS;

/*==============================================================*/
/* Table: ARTICLE_CLASS                                         */
/*==============================================================*/
create table ARTICLE_CLASS
(
   CLASS_ID             BIGINT(18) not null auto_increment comment '类型ID',
   CLASS_NAME           VARCHAR(64) comment '类型名称',
   DECRIPTION           VARCHAR(128) comment '描述',
   primary key (CLASS_ID)
);

alter table ARTICLE_CLASS comment '文章类型';


drop table if exists ARTICLE;

/*==============================================================*/
/* Table: ARTICLE                                               */
/*==============================================================*/
create table ARTICLE
(
   ARTICLE_ID           BIGINT(18) not null auto_increment,
   CLASS_ID             BIGINT(18) comment '类型ID',
   TITLE                VARCHAR(255),
   CONTENT              VARCHAR(255),
   BAIDU_DOWN           VARCHAR(255),
   BAIDU_DOWN_PWD       VARCHAR(255),
   TIANYI_DOWN          VARCHAR(255),
   TIANYI_DOWN_PWD      VARCHAR(255),
   STATE                NUMERIC(6),
   primary key (ARTICLE_ID)
);

alter table ARTICLE comment '文章';

alter table ARTICLE add constraint FK_Reference_5 foreign key (CLASS_ID)
      references ARTICLE_CLASS (CLASS_ID) on delete restrict on update restrict;