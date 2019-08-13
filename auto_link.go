// Lute - A structured markdown engine.
// Copyright (C) 2019-present, b3log.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lute

import (
	"bytes"
	"strings"
)

func (t *Tree) parseGfmAutoLink(tokens items) (ret Node) {
	index := bytes.Index(tokens, []byte("www."))
	if 0 > index {
		return nil
	}
	if 0 < index {
		t.context.pos += index
		return &Text{tokens: tokens[:index]}
	}

	length := len(tokens)
	var i int
	var token byte
	for ; i < length; i++ {
		token = tokens[i]
		if isWhitespace(token) || itemLess == token {
			break
		}
	}

	url := tokens[:i]
	length = len(url)
	var j int
	for ; j < length; j++ {
		token = url[j]
		if itemSlash == token {
			break
		}
	}
	domain := url[:j]
	if !t.isValidDomain(domain) {
		t.context.pos += i
		return &Text{tokens: url}
	}

	path := url[j:]
	length = len(path)
	if 0 < length {
		lastToken := path[length-1]
		if isASCIIPunct(lastToken) {
			path = path[:length-1]
		}
	}

	dest := items("http://")
	domainPath := append(domain, path...)
	dest = append(dest, domainPath...)
	ret = &Link{&BaseNode{typ: NodeLink}, encodeDestination(fromItems(dest)), ""}
	ret.AppendChild(ret, &Text{tokens: domainPath})
	t.context.pos += i
	return
}

// isValidDomain 校验 GFM 规范自动链接规则中定义的合法域名。
// https://github.github.com/gfm/#valid-domain
func (t *Tree) isValidDomain(domain items) bool {
	segments := bytes.Split(domain, []byte("."))
	length := len(segments)
	if 2 > length { // 域名至少被 . 分隔为两部分，小于两部分的话不合法
		return false
	}

	var token byte
	for i := 0; i < length; i++ {
		segment := segments[i]
		segLen := len(segment)
		for j := 0; j < segLen; j++ {
			token = segment[j]
			if !isASCIILetterNumHyphen(token) {
				return false
			}
			if 2 < i && (i == length-2 || i == length-1) {
				// 最后两个部分不能包含 _
				if itemUnderscore == token {
					return false
				}
			}
		}
	}
	return true
}

func (t *Tree) parseAutoEmailLink(tokens items) (ret Node) {
	tokens = tokens[1:]
	var dest string
	var token byte
	length := len(tokens)
	passed := 0
	i := 0
	at := false
	for ; i < length; i++ {
		token = tokens[i]
		dest += string(token)
		passed++
		if '@' == token {
			at = true
			break
		}

		if !isASCIILetterNumHyphen(token) && !strings.Contains(".!#$%&'*+/=?^_`{|}~", string(token)) {
			return nil
		}
	}

	if 1 > i || !at {
		return nil
	}

	domainPart := tokens[i+1:]
	length = len(domainPart)
	i = 0
	closed := false
	for ; i < length; i++ {
		token = domainPart[i]
		passed++
		if itemGreater == token {
			closed = true
			break
		}
		dest += string(token)
		if !isASCIILetterNumHyphen(token) && itemDot != token {
			return nil
		}
		if 63 < i {
			return nil
		}
	}

	if 1 > i || !closed {
		return nil
	}

	t.context.pos += passed + 1
	ret = &Link{&BaseNode{typ: NodeLink}, "mailto:" + dest, ""}
	ret.AppendChild(ret, &Text{tokens: toItems(dest)})

	return
}

func (t *Tree) parseAutolink(tokens items) (ret Node) {
	schemed := false
	scheme := ""
	dest := ""
	var token byte
	i := t.context.pos + 1
	for ; i < len(tokens) && itemGreater != tokens[i]; i++ {
		token = tokens[i]
		if itemSpace == token {
			return nil
		}

		dest += string(token)
		if !schemed {
			if itemColon != token {
				scheme += string(token)
			} else {
				schemed = true
			}
		}
	}
	if !schemed || 3 > len(scheme) {
		return nil
	}

	ret = &Link{&BaseNode{typ: NodeLink}, encodeDestination(dest), ""}
	if itemGreater != tokens[i] {
		return nil
	}

	t.context.pos = 1 + i
	ret.AppendChild(ret, &Text{tokens: toItems(dest)})

	return
}
