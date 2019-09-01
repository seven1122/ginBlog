CREATE DATABASE IF NOT EXISTS blog CHARSET=utf8;


CREATE TABLE `blog`.`blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `add_dt` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `update_dt` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `deleted_dt` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';

CREATE TABLE `blog`.`blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text COMMENT '文章内容',
  `view_count` int(10) unsigned default '0' COMMENT '查看数',
  `comment_count` int(10) unsigned default '0' COMMENT '评论数',
  `praise_count` int(10) unsigned default '0' COMMENT '点赞数',
  `add_dt` int(10) unsigned DEFAULT '0' COMMENT '添加时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `update_dt` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `deleted_dt` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';


CREATE TABLE `blog`.`blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');