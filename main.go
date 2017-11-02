package main

import (
	"fmt"
	"reflect"

	"github.com/kerrigell/wadb/ast"
	"github.com/kerrigell/wadb/parser"
	"github.com/kerrigell/wadb/plan"
)

func main() {
	sql := `ALTER TABLE  hotel_item 
	ADD COLUMN  guarantee_type   smallint(6) NOT NULL DEFAULT 0 COMMENT '担保类型[0:无担保；1: 信用卡]' AFTER  is_immediate ,
	ADD COLUMN  is_guarantee   tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否但保[0:不担保；1：担保]' AFTER  guarantee_type ;
	create table movie_poster_detect_log
	(
	   id                   int(11) not null auto_increment comment '主键ID',
	   member_id            varchar(64) not null comment '会员Id',
	   image_name           varchar(128) not null comment '文件名',
	   city_id              int not null comment '地市ID',
	   create_time          timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
	   update_time          timestamp not null default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
	   creator_id           varchar(64) not null default '' comment '创建者ID',
	   operator_id          varchar(64) not null default '' comment '操作者ID',
	   primary key (id)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='电影海报识别业务日志表';
	
	
	`
	fmt.Println("test")
	pr := parser.New()
	nodes, err := pr.Parse(sql, "", "")
	// err := plan.Validate(node, false)
	println(err)
	for _, n := range nodes {
		switch t := n.(type) {
		// case ast.StmtNode:
		// 	println("stmt node")
		case *ast.AlterTableStmt:
			println("alter table")
			err := plan.Validate(n, true)
			fmt.Println(err)
		case *ast.SelectStmt:
			println("select stmt")
		//... etc
		default:
			_ = t
			println("unknown")
		}

		fmt.Println(reflect.TypeOf(n).String())
		fmt.Println("test")
		println(n)
	}

	fmt.Println("end")
}
