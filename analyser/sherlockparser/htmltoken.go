package sherlockparser

import "strings"

/*
HTMLToken interface.
*/
type HTMLToken interface {
	Type() TokenType
	RawContent() string
	AddToRawContent(toAdd string)
	ToString() string
}

//TokenType represents HTML token types.
type TokenType int

/*
Used to enumerate TokenTypes of HTML-Elements.
*/
const (
	StartTag       TokenType = 0
	EndTag         TokenType = 1
	SelfClosingTag TokenType = 2
	PlainText      TokenType = 3
)

/*
tokenize returns a pointer to a HTMLToken-slice which is generated based on the input html stored in the HTMLTree.
*/
func (tree *HTMLTree) tokenize() []HTMLToken {
	element := ""
	classified := []HTMLToken{}
	textTags := []string{"style", "script", "textarea", "title"}
	var lastTagType string
	var lastTokenType TokenType
	for i := 0; i < len(tree.htmlRaw); i++ {
		if _, found := find(textTags, lastTagType); found && lastTokenType == StartTag {
			nextElement := ""
			addedTextToken := false
			for l := i; l < len(tree.htmlRaw); l++ {
				if tree.htmlRaw[l] == '>' {
					tag := ""
					for k := len(nextElement) - 1; k >= 0; k-- {
						if nextElement[k] == '<' {
							newTag := tree.handleTag(tag)

							if newTag.Type() == EndTag && newTag.TagType() == lastTagType {
								i = l
								l = len(tree.htmlRaw)
								if addedTextToken {
									classified[len(classified)-1].AddToRawContent(nextElement[:k])
								} else {
									textToken := nextElement[:k]
									if textToken != "" {
										classified = append(classified, &TextToken{
											tokenType:  PlainText,
											rawContent: nextElement[:k],
										})
									}
								}
								lastTagType = newTag.TagType()
								lastTokenType = newTag.Type()
								classified = append(classified, newTag)
								break
							} else {
								nextElement += string(tree.htmlRaw[l])
								classified = append(classified, &TextToken{
									tokenType:  PlainText,
									rawContent: nextElement,
								})
								addedTextToken = true
								nextElement = ""
								break
							}
						} else {
							tag = string(nextElement[k]) + tag
						}
					}
				} else {
					nextElement += string(tree.htmlRaw[l])
				}
			}
			continue
		}

		if tree.htmlRaw[i] == '<' {
			if element != "" {
				element = strings.TrimSpace(element)
				if element != "" {
					classified = append(classified, &TextToken{
						tokenType:  PlainText,
						rawContent: element,
					})
				}
			}
			element = ""

			tag := ""
			if tree.htmlRaw[i+1] == '!' { // if charsequence <! found, assume comment or doctype and fastforward until >
				for l := i + 1; l < len(tree.htmlRaw); l++ {
					if tree.htmlRaw[l] == '>' {
						i = l
						break
					}
				}
			} else {
				for l := i + 1; l < len(tree.htmlRaw); l++ {
					if tree.htmlRaw[l] == '>' {
						i = l
						newTag := tree.handleTag(tag)
						lastTagType = newTag.tagType
						lastTokenType = newTag.tokenType
						classified = append(classified, newTag)
						break
					} else {
						tag += string(tree.htmlRaw[l])
					}
				}
			}
		} else {
			toAdd := string(tree.htmlRaw[i])
			element += toAdd
		}
	}
	return classified
}

/*
handleTag returns a pointer to a TagToken extracted from the input string.
*/
func (tree *HTMLTree) handleTag(tag string) *TagToken {
	var token *TagToken
	tagRaw := strings.TrimSpace(tag)
	voidElements := []string{"area", "base", "br", "col", "embed", "hr",
		"img", "input", "link", "meta", "param", "source", "track", "wbr"}
	tagType := strings.Split(tagRaw, " ")[0]

	_, contained := find(voidElements, tagType)
	switch {
	case contained:
		token = &TagToken{
			tokenType:  SelfClosingTag,
			tagType:    tagType,
			rawContent: tagRaw,
		}
	case len(tagRaw) > 0 && tagRaw[0] == '/':
		token = &TagToken{
			tokenType:  EndTag,
			tagType:    tagType[1:],
			rawContent: tagRaw,
		}
	default:
		token = &TagToken{
			tokenType:  StartTag,
			tagType:    tagType,
			rawContent: tagRaw,
		}
	}

	return token
}

/*
find finds a string in a string slice. Returns index of string (-1 if not found) and bool if string was found.
*/
func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
