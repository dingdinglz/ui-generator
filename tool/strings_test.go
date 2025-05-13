package tool

import (
	"fmt"
	"testing"
)

func TestStringBetween(t *testing.T) {
	fmt.Println(StringBetween("312{3{test}213123", "{", "}"))

	fmt.Println(StringBetween("312{3{test}}213123", "{", "}}"))

	fmt.Println(StringBetweenContain(`基于您提供的"Cook Your Life"生活手册应用需求，我将设计一个高保真原型图。首先，让我分析产品并制定设计方案。

我理解这是一个提供生活技能指南的百科全书式应用，用户可以在上面查找和学习各种生活技能。现在我将列出需要创建的页面文件列表：

{"data":[{"name":"index.html","description":"展示所有页面的入口"},{"name":"home.html","description":"应用首页，展示分类和推荐内容"},{"name":"category.html","description":"技能分类页面"},{"name":"skill_detail.html","description":"技能详情页面"},{"name":"search.html","description":"搜索页面"},{"name":"bookmark.html","description":"收藏页面"},{"name":"profile.html","description":"个人中心页面"},{"name":"settings.html","description":"设置页面"}]}`, "{", "}"))
}
