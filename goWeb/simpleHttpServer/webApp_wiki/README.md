# webApp_wiki
这是一个类似 wiki 的 web 服务器。浏览器可以通过像这样的 url 来访问 wiki 页面的内容： localhost:8080/view/page1.

然后会到和这个名字（page1）相同的文本文件中读取文件的内容展示在页面中；页面中包含了一个可以编辑 wiki 页面的超链接（ localhost:8080/edit/page1 ）.
编辑页面用一个文本框显示内容，用户可以修改文本并通过 Save 按钮保存到文件中；然后会在相同的页面（view/page1）中查看到被修改的内容。如果想要查看的页面
不存在（例如： localhost:8080/edit/page999 ），程序会将其跳转到一个编辑页面，这样就可以创建并保存一个新的 wiki 页面。

