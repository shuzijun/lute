// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package test

import (
	"github.com/88250/lute/ast"
	"testing"

	"github.com/88250/lute"
)

var vditorIRBlockDOMListCommandTabTests = []*parseTest{

	{"2", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130183105-s07dvyx\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130113945-0974ppq\">foo<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130190601-wcnhone\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130191136-dl45lo9\"></li><li data-marker=\"*\" data-node-id=\"20210130190137-0n7j6j3\"><wbr>bar</li></ul></li></ul>", "<ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130183105-s07dvyx\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130113945-0974ppq\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">foo</p><ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130190601-wcnhone\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130191136-dl45lo9\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">\u200b</p><ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130190137-0n7j6j3\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\"><wbr>bar</p></li></ul></li></ul></li></ul>"},
	{"1", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20210129233054-46gy2j6\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210129233055-7696gne\">foo<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20210129235255-ja5gq0n\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210129235250-8iwmh3o\">bar</li></ul></li><li data-marker=\"*\" data-node-id=\"20210129235251-d0d17bv\"><wbr>baz</li></ul>", "<ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210129233054-46gy2j6\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210129233055-7696gne\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">foo</p><ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210129235250-8iwmh3o\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">bar</p></li><li data-marker=\"*\" data-node-id=\"20210129235251-d0d17bv\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\"><wbr>baz</p></li></ul></li></ul>"},
	{"0", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20210129215814-ztc4hye\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210129213421-p4d5yrp\">foo</li><li data-marker=\"*\" data-node-id=\"20210129213421-1qqjs0a\"><wbr><br><ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20210129230307-4oqo3do\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210129215550-bt8jdxi\">ccc</li></ul></li></ul>", "<ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210129215814-ztc4hye\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210129213421-p4d5yrp\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">foo</p><ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210129230307-4oqo3do\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210129213421-1qqjs0a\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">\u200b<wbr></p></li><li data-marker=\"*\" data-node-id=\"20210129215550-bt8jdxi\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">ccc</p></li></ul></li></ul>"},
}

func TestVditorIRBlockDOMListCommand(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.SetVditorIR(true)
	luteEngine.ParseOptions.Mark = true
	luteEngine.ParseOptions.BlockRef = true
	luteEngine.SetKramdownIAL(true)
	luteEngine.ParseOptions.SuperBlock = true
	luteEngine.SetLinkBase("/siyuan/0/测试笔记/")
	luteEngine.SetAutoSpace(false)
	luteEngine.SetSub(true)
	luteEngine.SetSup(true)
	luteEngine.SetGitConflict(true)

	ast.Testing = true
	for _, test := range vditorIRBlockDOMListCommandTabTests {
		html := luteEngine.VditorIRBlockDOMListCommand(test.from, "tab")

		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
	ast.Testing = false
}

var vditorIRBlockDOMListCommandSTabTests = []*parseTest{

	{"1", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130183105-s07dvyx\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130113945-0974ppq\">foo<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130190138-rh6fpw6\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130190139-yny17o1\"><wbr></li><li data-marker=\"*\" data-node-id=\"20210130190137-0n7j6j3\">bar</li></ul></li></ul>", "<ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130183105-s07dvyx\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130113945-0974ppq\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">foo</p></li><li data-marker=\"*\" data-node-id=\"20210130190139-yny17o1\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">\u200b<wbr></p><ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130190138-rh6fpw6\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130190137-0n7j6j3\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">bar</p></li></ul></li></ul>"},
	{"0", "<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20210129233054-46gy2j6\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210129233055-7696gne\">foo<ul data-tight=\"true\" data-marker=\"*\" data-block=\"0\" data-node-id=\"20210129233106-nen167c\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210129233106-d3oiwuc\"><wbr>bar</li></ul></li></ul>", "<ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210129233054-46gy2j6\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210129233055-7696gne\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">foo</p></li><li data-marker=\"*\" data-node-id=\"20210129233106-d3oiwuc\"><p data-block=\"0\" data-node-id=\"20210129233106-nen167c\" data-type=\"p\">\u200b<wbr>bar</p></li></ul>"},
}

func TestVditorIRBlockDOMListCommandSTab(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.SetVditorIR(true)
	luteEngine.ParseOptions.Mark = true
	luteEngine.ParseOptions.BlockRef = true
	luteEngine.SetKramdownIAL(true)
	luteEngine.ParseOptions.SuperBlock = true
	luteEngine.SetLinkBase("/siyuan/0/测试笔记/")
	luteEngine.SetAutoSpace(false)
	luteEngine.SetSub(true)
	luteEngine.SetSup(true)
	luteEngine.SetGitConflict(true)

	ast.Testing = true
	for _, test := range vditorIRBlockDOMListCommandSTabTests {
		html := luteEngine.VditorIRBlockDOMListCommand(test.from, "stab")

		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
	ast.Testing = false
}


var vditorIRBlockDOMListCommandEnterTests = []*parseTest{

	{"0", "<ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130210502-h5p7mhf\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130210503-p89seth\"><p data-block=\"0\" data-node-id=\"20210130210506-jfsx7bc\" data-type=\"p\">foo<wbr></p><ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130210503-nqslo5q\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130210503-4jj2eqq\"><p data-block=\"0\" data-node-id=\"20210130210506-u1x76c9\" data-type=\"p\">bar</p></li></ul></li></ul>", "<ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130210502-h5p7mhf\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130210503-p89seth\"><p data-block=\"0\" data-node-id=\"20210130210506-jfsx7bc\" data-type=\"p\">foo<wbr></p></li><li data-marker=\"*\" data-node-id=\"20060102150405-1a2b3c4\"><p data-block=\"0\" data-node-id=\"20060102150405-1a2b3c4\" data-type=\"p\">\u200b<wbr></p><ul data-marker=\"*\" data-block=\"0\" data-node-id=\"20210130210503-nqslo5q\" data-type=\"ul\"><li data-marker=\"*\" data-node-id=\"20210130210503-4jj2eqq\"><p data-block=\"0\" data-node-id=\"20210130210506-u1x76c9\" data-type=\"p\">bar</p></li></ul></li></ul>"},
}

func TestVditorIRBlockDOMListCommandEnter(t *testing.T) {
	luteEngine := lute.New()
	luteEngine.SetVditorIR(true)
	luteEngine.ParseOptions.Mark = true
	luteEngine.ParseOptions.BlockRef = true
	luteEngine.SetKramdownIAL(true)
	luteEngine.ParseOptions.SuperBlock = true
	luteEngine.SetLinkBase("/siyuan/0/测试笔记/")
	luteEngine.SetAutoSpace(false)
	luteEngine.SetSub(true)
	luteEngine.SetSup(true)
	luteEngine.SetGitConflict(true)

	ast.Testing = true
	for _, test := range vditorIRBlockDOMListCommandEnterTests {
		html := luteEngine.VditorIRBlockDOMListCommand(test.from, "enter")

		if test.to != html {
			t.Fatalf("test case [%s] failed\nexpected\n\t%q\ngot\n\t%q\noriginal html\n\t%q", test.name, test.to, html, test.from)
		}
	}
	ast.Testing = false
}
