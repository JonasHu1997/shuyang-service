CREATE TABLE `sms_log` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `phone` varchar(20) NOT NULL DEFAULT '' comment '手机号',
  `content_meta` varchar(400) not null DEFAULT '' comment '短信内容元数据',
  `provider_id` int(10) NOT NULL DEFAULT '0'  comment '服务商ID',
  `app_id` int(10) NOT NULL DEFAULT '0'  comment '调用方ID',
  `ctime` int(10) NOT NULL DEFAULT '0' comment '发送时间',
  `status` tinyint(1) UNSIGNED NOT NULL default '0' comment '发送状态：0未发送，1发送成功，2发送失败，3发送中',
  `response` varchar(200) not null default '' comment '服务商返回消息',
  PRIMARY KEY (`id`),
  KEY `idx_app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='短信记录表';

CREATE TABLE `sms_tpl` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `app_id` int(10) NOT NULL DEFAULT '0' comment '模板所属应用id, 0为所有团队共有',
  `name` varchar(100) NOT NULL DEFAULT '' comment '名称',
  `remark` varchar(100) NOT NULL DEFAULT '' comment '备注',
  `tpl` varchar(200) NOT NULL DEFAULT '' comment '模板',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' comment '1激活，0停用',
  PRIMARY KEY (`id`),
  KEY `idx_app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='短信模板';


CREATE TABLE `provider` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' comment '服务商名称',
  `remark` varchar(200) NOT NULL DEFAULT '' comment '备注',
  `conf_meta` varchar(200) NOT NULL DEFAULT '' comment '服务商配置',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' comment '1激活，0停用',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='服务商表';

CREATE TABLE `app` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' comment '调用方名称',
  `remark` varchar(200) NOT NULL DEFAULT '' comment '备注',
  `auth_id` int(10) NOT NULL DEFAULT '0' comment '唯一认证id',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' comment '1激活，0停用',
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='调用方表';

CREATE TABLE `app_provider` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `app_id` int(10) NOT NULL DEFAULT '0' comment '调用方id',
  `provider_id` int(10) NOT NULL DEFAULT '0' comment '渠道id',
  `is_default` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' comment '0非默认，1默认',
  PRIMARY KEY(`id`),
  UNIQUE KEY `idx_app_provider` (`app_id`, `provider_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 comment='调用方渠道绑定';

insert into `app` set name='测试调用方', remark='testttest....', auth_id=1 ;
insert into `provider` set name='测试渠道商',remark='testttest...', conf_meta='{"api_key":"xxxxx", "api_secret":"xxxxx"}';
insert into `app_provider` set app_id=1, provider_id=1, is_default=1;
insert into `sms_tpl` set app_id='0', name='测试模板', tpl='尊敬的{{.name}}，您的xxx为{{.code}}，请查收。';