# 海报

>需求中经常会有推广，需要在朋友圈分享图片，图片上有活动的二维码，用户的头像，昵称。

在golang中可以使用`image`包来实现绘图，[go-qrcode](http://github.com/skip2/go-qrcode)绘制二维码，[freetype](http://github.com/golang/freetype)实现图上绘制文字。

思路：
- 准备底图
- 绘制二维码
- 将二维码在底图指定位置绘制
- 获取微信头像
- 绘制微信头像原型缩略图
- 将缩略图在底图置顶位置绘制
- 获取微信昵称
- 将昵称在底图指定位置绘制