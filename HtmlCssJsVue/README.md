# HTML, CSS, JavaScript, Vue

This is a learning note all from **[Mozilla](https://developer.mozilla.org/zh-CN/docs/Learn/Getting_started_with_the_web/Dealing_with_files)**, A very detail and high quality tutorial, highly recommended to read this.

Start: 23/07/2022

End: 31/07/2022



## VS Code Tips

`<Tips> `shortcut to create html template in vsCode --- `!`

![image-20220627191526858](Images/image-20220627191526858.png)



`<Tips> `customize shortcut to input auto template like above mentioned

> 1. File -> Preferences -> Configure User Snippets
>
> 2. new Global Snippets file...
>
> 3. name it eg. vue-html.code-snippets
>
>    ```html
>    {
>    	"vue htm":{
>    		"scope": "html",
>    		"prefix": "vh", //this is the shortcut
>    		"body": [
>    			"<!DOCTYPE html>",
>    			"<html lang=\"en\">",
>    			"",
>    			"<head>",
>    			"    <meta charset=\"UTF-8\">",
>    			"    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">",
>    			"    <meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">",
>    			"    <title>Document</title>",
>    			"    <script src=\"./lib/vue.js\"></script>",
>    			"</head>",
>    			"",
>    			"<body>",
>    			"    <div id=\"app\">",
>    			"",
>    			"    </div>",
>    			"",
>    			"    <script>",
>    			"        new Vue({",
>    			"            el: '#app',",
>    			"            data: {",
>    			"                $1",
>    			"            },",
>    			"        })",
>    			"    </script>",
>    			"</body>",
>    			"",
>    			"</html>"			
>    		],
>    		"description": "my vue template in html"
>    	}
>    }
>    ```
>
> 

# HTML

>  Hypertext Markup Language

## Work Space and File Management

> Create a work space (a folder) to manage all your projects. eg web-projects

This work space will keeps all your projects



> Naming your file and folder

Do: 

- '-' to separate the words
- small capital

Don't:

- use capital letter
- '_' underlined, as some search engines will not recognize underscores and will drop pages that contain them.



> Structure

1. `index.html`
   - The first page that user visit in the website
2. `images` folder 
   - Keeps all the images
3. `styles` folder
   - Keeps all css
4. `scripts` folder
   - Keeps all JavaScript



## Structure

```html
<!DOCTYPE html> <!-- Document Type. Ensure the document can be read -->
<html lang="en"> <!-- Enclosed all the information of the page, so called root elemen. 
			lang attribute is strongly recommended to be there as it is much more efficient for search engine to search -->
  <head> <!-- Is invisible to user, it include the keywords for search engine, the page description, CSS style, and encoding declaration -->
    <base href=""> <!-- A base for followed href to start from. Tips: only allowed one in the page -->
    ---------------------------------
    meta
    <meta charset="utf-8"> <!-- meta element，describe your webpage. Also able to be searched in search engine -->
    <meta name="author" content="Pokemon">
    <meta name="description" content="ABCDEFG">` <!-- when you search ABCDEFG this will listed in search engine! -->
    <!-- og: <= facebook meta protocol open graph data, when you connect MDN with facebook, then the link will show following image, description and title-->
    <meta property="og:image" content="https://developer.mozilla.org/static/img/opengraph-logo.png">
    <meta property="og:description" content="The Mozilla.....">
    <meta property="og:title" content="Mozilla Developer Network">
    <meta name="viewport" content=",initial-scale=1"> <!-- set the width of the page -->
    ---------------------------------
    link
    <!-- rel attribute specifies the relationship between the current document and the linked document/resource -->
    <link rel="icon" href="...favicon.ico" type="image/x-icon"> <!-- Tips: use .ico which is compatible with IE6, however most browser supports png and gif -->
    <!-- To use css file --> 
    <link rel="stylesheet" href="my-css-file.css">`
    <link href='https://fonts.googleapis.com/css?family=Rock+Salt' rel='stylesheet' type='text/css'
    ---------------------------------
    <title>My test page</title> <!-- Topic of the page, it is the title show on browser tab, also the bookmarks description -->
    ---------------------------------
    style
    <style>
    	....  
    </style>
    ---------------------------------
    script
    <!-- it is not necessary to be inner of head. it is to use .js. 'defer' is preferred as to confirm the script is executed when a page has finished parsing. Otherwise there will be error if visiting an unexist element.--> 
    <!-- Please do have end tag </script> --><
    <script src="my-js-file.js" defer></script>
    <script defer>....</script>
  </head>
  <body>
    <img src="" alt="My test image">
  </body>
</html>
```



## Semantic Elements

![https://www.w3schools.com/html/html5_semantic_elements.asp](Images/image-20220723235348167.png)



```html
<body>
    <header></header>
    <main>
    	<section></section> <!-- for functions to click like mini map, table of content --> 
        <article></article> 
        <aside></aside> 
        <div></div>
    </main>
    <footer></footer>
</body>
```



## Paragraph and Sentences

```html
<p></p>  ---- paragraph
<span></span> <!-- can use lang attb to identify the language used -->
<br> ---- next line
<hr> ---- a horizontal line
```

```html
Headings
<h1></h1> ---- strongly recommend ONLY one per page
<h2></h2> ---- not recommended >3 in one page
<h3></h3> ---- not recommended >3 in one page
<h4></h4> ---- not recommended >3 in one page
<h5></h5> ---- not recommended
<h6></h6> ---- not recommended
```

```html
<b></b>
<i></i>
<u></u>

- 斜体：外国文字，分类名称，技术术语，一种思想……
- 粗体：关键字，产品名称，引导句……
- 下划线(to be different from link, use with css: text-decoration-style:wavy[example])：专有名词，拼写错误……
```

Tips: you may apply attribute tabindex="-1" to enable focusable for the element that cannot be focus.

### Other

|                               | Attribute / Usage                                            | Description                                                  |
| ----------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `<dl><dt><dd></dd></dt></dl>` | eg dt: word dd: meaning  eg dt: question dd: answer          | Description List                                             |
| `<blockquote></blockquote>`   | `<blockquote cite="https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote"> ... </blockquote>` | a block, normally to cite a resource. Browser will have it indented |
| `<code></code>`               |                                                              | code pattern                                                 |
| `<q></q>`                     | `cite="link"`                                                | for quote ...> add " .... "                                  |
| `<cite></cite>`               | Tips: can be used inner of `<a>`                             | default : italic                                             |
| `<abbr></abbr>`               | `title`="it is a title when your cursor stayed" `<abbr title="Hypertext Markup Language">HTML</abbr>` | for short term / abbreviation                                |
| `<address></address>`         | `<address>` `<p>Chris Mills, Manchester, The Grim North, UK</p>` `</address>` | for contact html creator. Default font italic.               |
| `<sup> & <sub>`               |                                                              | sup = up sub = down                                          |
| `<pre></pre>`                 |                                                              | it show exactly what you code, or else browser will ignore those spaces. |
| `<var></var>`                 |                                                              | variable                                                     |
| `<kbd></kbd>`                 |                                                              | keyboard button                                              |
| `<samp></samp>`               |                                                              | 用于标记计算机程序的输出。 Sample Output Element             |
| `<time></time>`               | `datetime`="yyyy-MM-dd" `datetime`="yyyy-MM" `datetime`="MM-dd" `datetime`="HH-mm" `datetime`="yyyy-MM-ddTHH-mm" `datetime`="yyyy-MM-ddTHH-mm+XX:XX" <= offset `<time datetime="2016-01-20">20 January 2016</time>` | time format for computer to recognize                        |
| `<canvas></canvas>`           | `width` = '..'<br />`height`= '..'<br />innerHTML <= displayed when canvas is not supported in this browser |                                                              |
| `<progress></progress>`       | `max` ="100"<br />`value` = "75"<br />content is the fallback when the browser does not support this tag |                                                              |
| `<meter></meter>`             | `min` = "0"<br />`max` = "100"<br />`value` = "70"<br />`low` = "33"<br />`high` = “66”<br />`optimum` = “50”<br /> content is the fallback when the browser does not support this tag | when it is preferred part => green<br />average => yellow<br />worst => red |



## Link

```html
<a href="contact.html">Contact Page</a>
<a href="contact.html#Mailing_address"></a> <!-- go to id Mailing_address --> 
<a href="an API" download="abc.txt">
<a href="mailTo:abc@gmail.com">my email</a>
<a href="mailTo:abc@gmail.com?
         cc=def@gmail.com&
         bcc=ghi@gmail.com&
         subject=The%20subject">my email</a>

-------------------
Attributes
href="https://....."  <!-- to where [hypertext reference] -->
title="infor...." <!-- when cursor stay on it -->
target="_blank" <!-- open in new tab -->
download= "..." <!-- to set download name -->
```



### Target Attribute

- `_self`: 载入结果到当前浏览上下文中。（该值是元素的默认值）。
- `_blank`: 载入结果到一个新的未命名的浏览上下文。
- `_parent`: 载入结果到父级浏览上下文（如果当前页是内联框）。如果没有父级结构，该选项的行为和`_self`一样。
- `_top`: 载入结果到顶级浏览上下文（该浏览上下文是当前上下文的最顶级上下文）。如果没有父级，该选项的行为和_self 一样。



### Absolute vs Relative URL

- Absolute: including protocol and => eg http://....com
- Relative : start from the file of this page



## List

```html
Ordered list:
<ol>
    <li></li>
    <li></li>
</ol>
---------------------------
Unordered list:
<ul>
    <li></li>
    <li></li>
</ul>

---------------------------
Attrubutes:
ol
--
start="number to start"
reversed

li
--
value="number to fix"
```





## Images

```html
<img/>
-----------------
Attributes:
src="your path to img"
alt="text to show when img not shown" (user who visually impaired, screen reader will read this)
width="400
height="341"
title="text when cursor stayed"
```



### Tips

> If the image has no meaning, use css



### Resolution

> srcset format: `<filename>` `[space]` `<width(noted that it is 'w')`
>
> sizes format: `<condition>` `[space]` `image plug size`
>
> Procedure: 
>
> 1. get the size of device
> 2. check the conditions
> 3. get the plug size
> 4. choose whichever closest between plug size and srcset width
> 5. if none, go for src

```html
<img srcset="elva-fairy-320w.jpg 320w, 
             elva-fairy-480w.jpg 480w,
             elva-fairy-800w.jpg 800w"
     sizes="(max-width: 320px) 280px, 
            (max-width: 480px) 440px,
            800px"
     src="elva-fairy-800w.jpg" alt="Elva dressed as a fairy">
```



### Case : Mobile , Desktop

```html
<picture>
  <source media="(max-width: 799px)" srcset="elva-480w-close-portrait.jpg">
  <source media="(min-width: 800px)" srcset="elva-800w.jpg">
  <img src="elva-800w.jpg" alt="Chris standing up holding his daughter Elva">
</picture>
```



### Figure with caption

```html
<figure>
    <img src=".....">
    <figcaption>This is a dinosour</figcaption>
</figure>
```



## Video

```html

<video controls width="250">
    <source src="/media/cc0-videos/flower.webm"
            type="video/webm">
    <source src="/media/cc0-videos/flower.mp4"
            type="video/mp4">
    Sorry, your browser doesn't support embedded videos.
</video>
---------------------
Attributes:
src
type="video/mpr"` <!-- other: ../webm-->
controls <!-- for user to control video -->
width="400"
height="400"
autoplay (not recommended)
loop (not recommended)
muted
poster="...url" <!-- the image before play -->
preload="none / auto / metadata" <!-- to preload after page load? (for big file) -->
```



## Audio

> Refer to Video Section

```html
src
type="audio/mp3" <!-- other: audio/ogg -->
controls <!-- for user to control audio -->
autoplay (not recommended)
loop (not recommended)
muted
preload="none / auto / metadata" <!-- to preload after page load? (for big file) -->
```





## Table

> thead, tfoot, tbody is for structure purpose, in case you need to apply css on it.

```html
<table>
  <caption>This is a table</caption>
  <colgroup>
    <col>
    <col style="background-color: yellow">
  </colgroup>
  <thead>
      <tr>
        <th>Data 1</th>
        <th>Data 2</th>
      </tr>
  </thead>
  <tfoot>
      <tr>
        <tr>Total</tr>
        <tr></tr>
      </tr>
  </tfoot>
  <tr>
    <td>Calcutta</td>
    <td>Orange</td>
  </tr>
  <tr>
    <td>Robots</td>
    <td>Jazz</td>
  </tr>
</table>
-----------------
Attributes:
th
scope="col / row / colgroup / rowgroup" <!-- For person who is visually impaired -->


th / tr
colspan="number"
rowspan="number"

col
style="css style"
span="number" <= in case you need 2 colums with same style


```



## Iframe

```html
<iframe .....>
</iframe>

----------------------
Attributes
allowfullscreen
frameborder="1 / 0" <!-- default 1, have border (not recommended, css can have better implementation eg border: none -->
src="url"
width="x"
height="y"
sandbox <!-- more secure, to prevent being attacked by hacker -->
```



## WebVTT File (.vtt)

> In case of the situation people cannot listen / watch the audio or video at the moment.
>
> WebVTT File is the format to show the info for different duration
>
> SEO can track it

### 3 Types: 

1. subtitles: translation for different nation
2. captions: the information of video
3. timed descriptions : change the words into audio, for people who are visually impaired.

### .vtt

```
WEBVTT

1
00:00:22.230 --> 00:00:24.606
第一段字幕

2
00:00:30.739 --> 00:00:34.074
第二段

  ...
```

### `<track>`

```html
<video controls>
    <source src="example.mp4" type="video/mp4">
    <source src="example.webm" type="video/webm">
    <track kind="subtitles" src="subtitles_en.vtt" srclang="en"> #after source
</video>
```



## Malicious Technique

- Clickjacking

  > The page cver the iframe with transparent layer and lead user to click and enter malicious page.



## Debugging

> [Markup Validation Service by W3C](https://validator.w3.org/#validate_by_upload) can help on debugging which return a error report if any
>
> F12 in your browser to check code





# CSS

This is a learning notes from [Mozilla](https://developer.mozilla.org/zh-CN/docs/Learn/CSS/First_steps/Getting_started)

Start: 24/07/2022

End: 



> Example of key and value

```css
Key

all: unset
background-color: pink
background-image: url("../../media/examples/lizard.png")
border: 1px solid black
border-radius: 3px
border-bottom: 4mm ridge rgba(211, 220, 50, .6)
box-shadow : 
caption-side: bottom
color: pink
font-family: "Gill Sans", Verdana, sans-serif 
font-size: 12px
font-weight: bold | 100 | normal
letter-spacing: 1px
list-style-type: none
margin: 30px auto /*this is the distance of border toward outer, auto => center the item*/ 
max-width: 33em
min-width:
min-height: 
max-height:
object-fit: cover | contains | fill (for img inside a box)
opacity: 0.6
padding: 5px /*this is the distance of border to your content */
text-align: center
text-decoration: green wavy underline;
text-shadow: 1px 1px 1px black
transform: rotate(0.8turn) | translate(x, y) | scale(number)
transform-origin: center
transition: all 0.3s cubic-bezier(0.25, 0.25, 0.56, 2);
width: 100px /*look like block width*/

```

```css
value function

width: calc(90% - 30px)
transform: rotate(0.8turn)
```



## Selector

```css
* 					---- all
tagselector
.classSelector
#idselector
li.classselector 	---- apply to class in li only
```



### || Attribute Selectors

- [attr]
- [attr=value] 
- [attr~=value] : at least one is stated value eg p[class~="special"]
- [attr|=value] : start with value and followed with '-' or exactly value eg div[lang|="zh"] 
- [attr^=value] : start with value eg li[class^="box-"]
- [attr$=value] : end with value eg li[class$="-box"]
- [attr\*=value] : wherever there is value in string eg li[class\*=box]
- [attr=value i] : i means non case sensitive



### || Pseudo Class

> :pseudo-class-name

```css
p:first-child 		---- is firstchild and p
:last-child
:nth-child(5)		---- is fifthchild
:nth-child(2n+1)	---- 1,3,5,7,9..
:nth-child(3n+4)	---- 4,7,10,13..
:nth-child(-n+3)	---- first 3 child
:only-child 		---- apply to the only child

:active				---- when click
:hover
:link 				---- yet visited
:visited 			---- visited

:focus				---- when user use keyboard
:default			---- for checkbox / radio / select's option
:checked			---- for checkbox / radio
:indeterminate		---- neither check and uncheck (when page first load)
:invalid /:out-of-renage	---- if the input pattern is not match
:valid
:read-only
:read-write
:required
:optional
:enabled
:disabled
:placeholder-shown

:checked			---- when radio / checkbox checked
```



### || Pseudo Element

> ::pseudo-element-name

```css
::first-line		---- the first line (even after resolution changed)
::first-letter
::selection			---- selected word(s) (when double clicked)
::before			---- can used with content(key) to insert information before element (like icon), is invisible to screen reader
::after
```



```css
Example of ::before /::after
.box::before {
    content: "";
    display: block;
    width: 100px;
    height: 100px;
    background-color: rebeccapurple;
    border: 1px solid black;
} 

.box::after {
    content: " ➥"
}
```



### || Combinator

#### |||| Descendant Combinator

> including child, grandson ... 

```css
body article p 		---- 
```

#### |||| Child Combinator

> child only

```css
article > p
```

#### |||| Adjacent sibling combinator

> second element only if it *immediately* follows the first element

```css
p + img
h1 + h2 + p 		---- apply to h1 and same level p right after h2 which after h1, if any. This will not apply to h1 and h2
```

#### |||| General sibling combinator

> same parent and select siblings after itself

```css
h1 ~ p
```



## Box (Display) | Margin | Padding | Border

> Block box : heading, p, ul
>
> Inline box : a, span, strong. em
>
> ----------------------
>
> use display: inline | block | flex (block outer) | inline-flex (inline outer) | inline-block

![image-20220725143255639](Images/image-20220725143255639.png)

> content: width, height (not including padding and border)
>
> Tips: 
>
> if want include => use key box-sizing: border-box

```css
Tips to use in all:
html {
    box-sizing: border-box;
}
*, *::before, *::after {
    box-sizing: inherit;
}
```



### || Block Box

- Width and height workable
- padding, margin, border can push away other element
- Box will change line

### || Inline Box

- Width and height not workable
- for left and right padding, margin, border, it is functionable but will push away the other inline box
- for top and bottom padding, margin, border, it is functionable and will not push away the other inline box
- Box will not change line

> Tips : if do not wish to overlap top and bottom elements, hope to set width and height : use display: inline-block



### || Margin

```css
margin
margin-top
margin-right
margin-bottom
margin-left

-------------------
margin: 0 auto /* center */
```

- will overlapped instead of plus

  ```css
  eg
  1 div with margin-bottom: 50px
  vs
  1 div with margin-top: 30px
  ```

  

### || Border

```css
border
border-top | right | bottom | left
border-width | style | color
border-top-width ...
```



### || Padding

```css
padding
padding-top | right | bottom | left
```



### || Unit

```css
Absolute length
---------------
cm
mm
Q
in
pc
pt
px

Relative length
---------------
em			---- parent's font-size
ex 			---- char 'x' height
ch			---- char '0' width
rem			---- root element's font-size
lh			---- line-height
vw			---- 1% of width of viewscreen 
vh			---- ditto but height
vmin		---- 1% of min side of viewscreen
vmax		---- ditto but max
%
```



## Application to Elements

### || Values

#### |||| Value and Units

> `<Integer>` - no decimal
>
> `<number>` - can have decimal
>
> `<dimension>` - is number, with unit eg deg, s, px
>
> `<percentage>` - %



#### |||| Color

> rgb -> red green blue
>
> rgba -> red green blue alpha
>
> hsl -> hue (0-360), saturation (0-100%) , lightness ( 0-100%)
>
> hsla -> + alpha (opacity 0.0-1)



#### |||| Position

> top
>
> left
>
> bottom
>
> right
>
> center
>
> background-position: _horizontal_ _vertical_



### || Background

```css
background: red url(bg-graphic.png) 10px 10px repeat-x fixed;
```

```css
background-attachment: fixed | scroll | local (when scroll, what happen)
background-color: red;
background-image: url(bg-graphic.png);
background-position: 10px 10px | top center
background-repeat: repeat-x | repeat-y | no-repeat | repeat
background-size: 100% | x y | cover(cover all box with maintain right x y) | contain (either right or bottom hit the max shall have blank space at one side )
```

> Gradient generate website: https://cssgradient.io/



### || Table

```css
table-layout: fixed
border-collapse: collapse
```



### || Font

```css
font-family: "Gill Sans", Verdana, sans-serif  
font-style: italic | normal | oblique
font-weight: light | normal | bold .. 		---- thickness
text-transform: none | uppercase | lowercase | capitalize (first letter)
text-decoration: none | underline | overline | line-through 
text-shadow: 4px 4px 5px red /* horizontal-offset vertical-offset blur-radius(default 0) color(default black) */ /* can be multi shadow, seperator ',' */
text-align: left | right | center | justify
line-height: 1.5 /* font-size x value */
letter-spacing: 2px
word-spacing: 4px

```

> Font Compile
>
> sequence: font-style, font-variant, font-weight, font-stretch, font-size '/' line-height, font-family
>
> ```css
> font: italic normal bold normal 3em/1.5 Helvetica, Arial, sans-serif;
> ```
>
> 



> font-family: "Gill Sans", Verdana, sans-serif  
>
> ----------------------------------------------------
>
> if gill sans not available => Verdana 
>
> *serif, sans-serif, monospace, cursive, fantasy: meaning refer to https://developer.mozilla.org/zh-CN/docs/Learn/CSS/Styling_text/Fundamentals*

> font-size for browser is default: 16px
>
> `<h1>` => 2em

```css
unit
----
Absolute
========
px

Relative
========
em
rem				---- follow % of <html> => can set <html> 10px, then do necessary counting using rem



```



#### |||| Overflow

> How to deal with overflow

```css
overflow: auto (only scroll bar visible when needed) | hidden | scroll
overflow-y: scroll
overflow-x: scroll
```



#### |||| Writing Mode (Logical Properties)

```css
writing-mode: horizontal-tb | vertical-rl | vertical-lr
-----------------------
X width | ok inline-size 
X height | ok block-size
```

![image-20220725164746361](Images/image-20220725164746361.png)

```css
Example:

border-inline-end
padding-inline-end
margin-block-start
```





### || List

> for ul and ol

```css
list-style-type: upper-roman | lower-alpha (default disc)
list-style-position: inside (refer https://developer.mozilla.org/zh-CN/docs/Learn/CSS/Styling_text/Styling_lists)
list-style-image: url(star.png) /* not recommended, prefer use background style with height and width */
------------------
compile
=======
line-style: square url(...) inside
```



### || Link

> Default: underline, unvisited: blue, visited: purple, when point: small hand, when focus (tab): border, when click(active) : red

```css
color:
cursor: cell | copy | crosshair | grab ...
outline: thick double #32a1ce
```

```css
Follow sequence!
----------------
a {

}


a:link {

}

a:visited {

}

a:focus {

}

a:hover {

}

a:active {

}
```



### || Input

```css
input {
    appearance: none; /* allow the input with styling restriction to adjust height and font-size freely */
}
```

Case to delete 'x' button when focus in Safari

```css
input[type="search"]:not(:focus, :active)::-webkit-search-cancel-button { display: none; }
```



## Layout

>Normal flow	--- no control on layout (default)

> float: left | right | none | inherit	--- let an element float to left in same line
>
> display: flex | grid | table 	--- set layout
>
> position: relative | absolute | fixed | sticky	--- default static, set the position of the element



### || Display: Flex



![image-20220726133957882](Images/image-20220726133957882.png)

Image from Mozilla, [mdn web docs](https://developer.mozilla.org/zh-CN/docs/Learn/CSS/CSS_layout/Flexbox)



```css
flex-direction: row --- default row
flex-wrap: wrap --- when overflow => next line
/* combine: flex : row wrap */
align-items: stretch | center --- 
/* do items y-axis */
/* default stretch(stretch to full height of container) */
/*center(remain content height but positioned to mid of container) */
justify-content: space-around | space-between | center | flex-end | flex-start /* do items x -axis */ 
/* space-around->split balancely with padding most l&r, space-between same but without padding
```

>  Items in flex

```css
flex: 1 | 200px --- items will be stretch to full width of container
flex: 1 200px --- 1 portion of total width, with min width 200px
align-self: flex-end
order: 1 /* bigger than appear later */
order: -1
-------------
1 is flex-grow
200px is flex-basis
```



### || Display: Grid

```css
grid-template-columns: 1fr 1fr 1fr  
/* 1fr is  1/available space */
grid-template-columns: 200px 1fr 1fr 
grid-template-columns: repear(3, 1fr)
grid-template-rows: 100px 100px
grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
grid-auto-rows: minmax(100px, auto) /*if <100px height = 100px, if > 100px height = adjust)
grid-gap: 5px /* please write also gap, as this is deprecated */
gap: 5px

grid-template-areas: /* use '.' if the cell is empty*/
"header header"
"sidebar content"
"footer footer"
```

![image-20220726131005535](Images/image-20220726131005535.png)

>  Items in grid

```css
grid-column: 2 / 4 --- the number is grid-line, so it means col span 2 & 3
grid-row: 1
grid-area: header /* relate it in container's grid-template-areas */
```



### || Float

> Be noted float does not affect the result of screenreader

```css
float: left
float: right
clear: both | left | right
```

> Margin problem solution, instead to add clear: both on the element, create one element class clearfix

```css
.clearfix {
  clear: both;
}
```







### || Position

```css
position: 
relative --- move from currunt position 
absolute --- absolute position from webpage's container
fixed --- absolute position from view (eg menu on screen) | sticky	--- fixed when the element appeared, static when user dint scroll to that element yet
static --- deault, do nothing

with
----
top: x unit
left: x unit
bottom: x unit
right: x unit
z-index: 1 /* the higher, the top */

```

> Tips : absolute
>
> if you hope to position from the body, then in your body { position: relative }

### || Display: Table

> Seldom, used for browser that do not support grid and flex

> Items in grid

```css
display: table-row
display: table-cell
display: table-caption
```



### || Multi column layout

```css
column-width: 200px --- one column minimum width 200px
column-count: 3
column-gap: 20px
column-rule: 4px dotted red --- line between columns

break-inside: avoid --- prevent a block to be divided into 2 different columns
page-break-inside: avoid --- together as for older browser
```

![image-20220726132832857](Images/image-20220726132832857.png)

![image-20220726132848225](Images/image-20220726132848225.png)



## '@' 

### || @import - import other css

```CSS
@import 'style3.css' /* import another css file */

@media (min-width: 30em) { /*in case media requirement met */
@media screen and (min-width: 800px) {
    body {
		key: value
    }
}

```



### || @media 媒介查询

> media-type:
>
> 1. all (default)
> 2. print
> 3. screen
> 4. speech

> media-feature-rule:
>
> 1. min-width
> 2. max-width
> 3. width
> 4. orientation : landscape | portrait

```css
@media media-type and (media-feature-rule) {
  /* CSS rules go here */
}

---------
Example:
---------
@media (min-width: 30em) { /*in case media requirement met */
@media screen and (min-width: 800px) {
@media screen and (min-width: 400px) and (orientation: landscape) {
@media screen and (min-width: 400px), screen and (orientation: landscape) {
@media not all and (orientation: landscape) { /* while not landscape */
    body {
		key: value
    }
}
  
```



### || @Support  - older browser

> For runnable in older browser

```css
@supports (display: grid) {
  .item {
      width: auto;
  }
}
/* if support, then inside the curly bracket works.
```





## How CSS Work

1. Load Html
2. Parse Html
3. Create Dom Tree
4. Load CSS
5. Parse Html
6. Attach Style to Dom Nodes
7. Display



## Dispute of Style

> 	Priority from high to low
>
> 1. 4 digits marks : when tag with style="..."
> 2. 3 digits marks: ID Selector
> 3. 2 digits marks: class Selector, Attribute selector, Pseudo-class
> 4. 1 digit marks: element Selector, Pseudo-element
>
> eg 0001, 1000,0100, 0022



### || Inheritance

> width, margin, padding, and borders will not be inherited
>
> form element too.

```css
button,
input,
select,
textarea {
  font-family : inherit;
  font-size : 100%;
} 
```



#### |||| Inheritance control

> every key can use below value to control inheritance
>
> 1. inherit
> 2. initial (browser default)
> 3. revert
> 4. revert-layer
> 5. unset



### || !important 

> border: none !important;

This will be the top priority.

Not recommended to use, as it change the layer operation, which make debugging hard.



## Other Tips

### ||  Scrollable when using window.innerWidth and window.innerHeight

> because of there is a margin from `<body>`

```css
<style>
  body {
    margin: 0;
    overflow: hidden;
  }
</style>
```



### || CSS Standard



```css
no
---
.box { background-color: #567895; }
h2 { background-color: black; color: white; }
---------------------
yes
---
.box {
  background-color: #567895;
}

h2 {
  background-color: black;
  color: white;
}
```



```css
Type of logic
-------------
/* || General styles */

body { ... }

h1, h2, h3, h4 { ... }

ul { ... }

blockquote { ... }


/* || Typography */

...

/* || Header and Main Navigation */

...

/* || UTILITIES */

.nobullets {
  list-style: none;
  margin: 0;
  padding: 0;
}

/* || SITEWIDE */

.main-nav { ... }

.logo { ... }

...

/* || STORE PAGES */

.product-listing { ... }

.product-box { ... }

```



### || Organize CSS with a modular architecture

> - OOCSS
> - BEM (Block Element Modifier) : when you noticed '--' '__' 



### || CSS pre-processor

> Sass

```css
$base-color: #c6538c;

.alert {
  border: 1px solid $base-color;
}
```

after compile:

```css
.alert {
  border: 1px solid #c6538c;
}
```



# JavaScript

## Load JavaScript

Method 1

> async - use when your script does not need webpage to parse  and independent.
>
> defer - otherwise, eg script1 needs jQuery.js

```html
<script src="script.js" async></script>
```

Method 2

```html
<script>
    // 函数：创建一个新的段落并添加至 HTML body 底部。
    function createParagraph() {
      let para = document.createElement('p');
      para.textContent = '你点了这个按钮！';
      document.body.appendChild(para);
    }

    /*
      1. 取得页面上所有按钮的引用并将它们置于一个数组中。
      2. 通过一个循环为每个按钮添加一个点击事件的监听器。
      当按钮被点击时，调用 createParagraph() 函数。
    */

    const buttons = document.querySelectorAll('button');

    for (let i = 0; i < buttons.length; i++) {
      buttons[i].addEventListener('click', createParagraph);
    }
</script>
```



## Method of Working a function

Method 1

```js
function createParagraph() {
    let para = document.createElement('p');
    para.textContent = '你点了这个按钮！';
    document.body.appendChild(para);
}

const buttons = document.querySelectorAll('button');

for (let i = 0; i < buttons.length; i++) {
    buttons[i].addEventListener('click', createParagraph);
}
```

body:

```html
<button>点我呀</button>
```



Method 2 :

```js
function createParagraph() {
    let para = document.createElement('p');
    para.textContent = '你点了这个按钮！';
    document.body.appendChild(para);
}

const buttons = document.querySelectorAll('button');

for (let i = 0; i < buttons.length; i++) {
    buttons[i].onclick = createParagraph;
}
```

body:

```html
<button>点我呀</button>
```



Method 3: (Not recommended, mixed html with javaScript)

```js
function createParagraph() {
    let para = document.createElement('p');
    para.textContent = '你点了这个按钮！';
    document.body.appendChild(para);
}
```

body: 

```html
<button onclick="createParagraph()">点我呀</button>
```



## Variable , Operators, Types

### || Variable

- let : newer version than var, recommended to replace var. let does not support existence of two same variable name
- var
- const: constant

> Good Practice:
>
> - Camel Case
> - Do not start with _ as have some special meaning

### || Operators

> \+ \- \* / % **(5\*\*5 === 5<sup>5</sup> === Math.pow(5,5)(old))   
>
> ++ -- +=  -= *= /=
>
> ===	!==	<	> <= >=



### || Types

> typeof [variable]

1. Number 	-  *eg 30, 2.456*
2. String   - *eg 'Hello'* ........ Tips: actually it is an array : string[0] => first letter
3. Boolean  - *true | false*(false, undefined, null, 0, NaN, "")
4. Array  - *eg [10,15,40]*
5. Object  -  *eg {name: 'Spot', breed: 'Dalmatian'*

```js
num to string
-------------
myNum.toString()

string to num
-------------
Number(myStr)
```

#### |||| String methods

```js
string.length
string[index]
string.indexOf('abc') // return -1 when found nothing ==> use !== -1 as Java string.contains("abc")
string.includes('abc')
string.slice(0,3) //return 0,1,2 => java: .substring(0,3)
string.slice(2) //from index 2
string.toLowerCase()
string.toUpperCase()
string.replace('abc','def') //only first met
string.replaceAll('abc','def')
string.split(',') // return array
```



#### |||| Array method

```js
array.join(',') //join array and form string
array.toString() //ditto but only ','
array.push('abc')
array.push('def',13) // let newLen = array.push('def,13') ** return array length
array.pop() //return removed item
array.unshift('abc') //same as push but from front
array.shift() // same as pop but from front
array.splice(index, how many to delete)
array.length
array.find(item => item.isDone)
array.findIndex(item => item.isDone)
array.filter(item => {item.name.includes("abc")})
array.forEach(function () {..})


```



## Conditional and Repetition

- if-else
- switch - case
- Ternary Conditional Operator (test ? if true : if false)
- for
- while
- do while

Same usage as Java



## Useful  Utils Methods

```js
console.log()
alert()
prompt() 	// ask question and receipt input
setInterval(function().. , delay in ms) // return ID
clearInterval(id)


document
--------
document.querySelector("..") //eg 'div' //Trending
document.querySelectorAll("..") //return array
document.getElementById("..")
document.getElementsByClassName("..") //return array
document.getElementsByTagName('p') 
document.createElement("div")
document.createTextNode("abcdefg") // sentence without tag 
location.reload() // refresh the webpage

Node
----
let x = URL.createObjectURL(blob) //to src

###HTML // Be noted  can be detected by search engine
.setAttribute('class','...'); // will override
.removeAttribute('class')
.className = "..."
.classList.add('active')
.classList.remove('active')
.classList.contains('active')
.textContent = "..."
.href = 'https:/...'
.appendChild(param)
.cloneNode(NodeA)
.removeChild(nodeYouWantToDelete); // ** a.parentNode.removeChild(a);
.parentNode
while (section.firstChild) {
      section.removeChild(section.firstChild);
  }


input.value = "..." 
select.value = "..."
input.focus() 	// focus on this input

let img = new Image()
img.src = 'firefox.png'

###CSS
.style.backgroundColor = "..." // Camel Case
.style.color = '...'
.style.pading = '...'
.style.widht = '...'
.style.textAlign = 'center'
.style.visibility = 'visible'


###validity
const input = document.getElementById("age") // button, fieldset, input, output, select, textArea
if(input.validity.badInput) {...}
badInput => string in number input
patternMismatch => pattern in input not met
rangeUnderflow => not met min requirement
rangeOverflow => not met max requirement
stepMismatch => not complied to (value - min) % step = 0
tooLong | tooShort
typeMismatch => not comply to type eg email / url in type of input
valid

form element
------------
checkValidity() // return true/false. true if have no problem
reportValidity()
setCustomValidity(message)

###Video
if(video.paused) ...
video.play()
video.pause()

Check
-------
isNaN(num) // if it is Not a number (when string = true)

Math
----
Math.floor(Math.random() * 100);   // 0 - 99
Math.floor(Math.random() * array.length)
Math.tan(degToRad(degrees))
function degToRad(degrees) {
  return degrees * Math.PI / 180;
}

Window
------
window.innerWidth //not function
window.innerHeight
window.setTimeout( function, millisecond) // window.setTimeout( () => {alert(..)}, 1000)


canvas
------
let ctx = canvas.getContext('2d')
ctx.beginPath()
ctx.translate(width/2, height/2); // move pen to middle of screen
ctx.fillStyle = "red" // set color
//ctx.strokeStyle = this.color; <= this
//ctx.lineWidth = 3; <= with tis will create O shape circle (a hole)
ctx.arc(x, y, <radius>, start [eg degToRad(0)], end eg 2 * Math.PI or [degToRad(0)], true = anticlockwise / false = clockwise) //draw circle. x & y is the center of circle .start and end is based on r. 360 degree is 2 * Math.PI. (since 2 PI r == perimeter, 360 degree by r is: 2 PI r / r = 2 PI) // it start from 3 o'clock // if not a full circle,  ctx.lineTo(x,y) must be after this line and before fill(), otherwise the circle will be weird.
ctx.fillRect(0, 0, width, height)
ctx.clearRect(x, y, width, height)
ctx.strokeRect(x , y, width, height)
ctx.moveTo(x, y) // without drawing action
ctx.lineTo(x, y)
ctx.closePath() // draw a straigth line and go back to the start
ctx.font = '36px arial'
ctx.strokeText('Canvas text', x, y)
ctx.fillText('Canvas text', 50, 150);
ctx.fill() // fill and end drawing
ctx.stroke()
ctx.rotate(rad) // rotate the canvas clockwise
ctx.save() // save the fillStyle....
ctx.restore()  // restore to stage of last saved
ctx.onmousedown = function(){}
ctx.onmouseup = function(){}
ctx.onmousemove = function(){}

## below showing to draw a triangle
ctx.beginPath() // start drawing
ctx.moveTo(50, 50);
// 绘制路径
ctx.lineTo(150, 50);
var triHeight = 50 * Math.tan(degToRad(60)); // degToRed is a customized function // tan(degree) = height / width
ctx.lineTo(100, 50+triHeight); 
ctx.lineTo(50, 50);
ctx.fill(); // fill the inner side of triangle
// advance refer me: https://developer.mozilla.org/zh-CN/docs/Web/API/Canvas_API/Tutorial/Drawing_shapes

window.requestAnimationFrame(loop); // run the loop function in time frame //return an id // or requestAnimationFrame(loop) // 60 frame per second highest
window.cancelAnimatonFrame(id)
```

> btn.onclick = displayMessage  is not same as btn.onclick = displayMessage()
>
> the first will not trigger when page load, the second will. So, if there is a case that displayMessage with parameters, it must be = function() {displayMessage(x,y)};

image in canvas

```js
var image = new Image();
image.src = 'firefox.png';

image.onload = function() {
  ctx.drawImage(image, x, y);
}

--------------
detail version
--------------
drawImage(image, dx, dy)
drawImage(image, dx, dy, dWidth, dHeight)
drawImage(image, sx, sy, sWidth, sHeight, dx, dy, dWidth, dHeight)

```

![image-20220729150523898](Images/image-20220729150523898.png)

Image from [Mozilla](https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/drawImage), demostrating the canvas.drawImage function

**Tips: for nodeList to use forEach(array function)**

```js
NodeList.prototype.forEach = function (callback) {
  Array.prototype.forEach.call(this, callback);
}
```





## Event

[Events listed here](https://developer.mozilla.org/zh-CN/docs/Web/Events)

> using two .onclick on same element, the latter will override the previous,
>
> using addEventListener(...) , there can be more than one functions. 

```js
Most Element
--------------
.onclick = function() ..
.ondblclick = function() .. //double click
.onfocus = function() ..
.onblur = function() ..	//normally when wrong input
.onchange = function() {} // for select tag
.onmouseover = function() ..
.onmouseout = function() ..
.addEventListener('click', function() ...) // it is DOM level 2 events
.removeEventListener('click', function() ...)
form.onsubmit = function(){}

Video Element
-------------
.onplay = function() ..
.addEventListener('timeupdate', function()..)

Window Element
--------------
window.onload = function() ...
window.onkeydown = function() .. //when keyboard key is pressed
window.onkeyup = function() .. //ditto but loosened
window.onresize = function() ..

Canvas Element
--------------
let ctx = ....
ctx.onmousedown = function(){}
ctx.onmouseup = function(){}
ctx.onmousemove = function(){}
```



### || Event Object

> e / evt / event -> allow the function use the element

```js
const divs = document.querySelectorAll('div');

for (let i = 0; i < divs.length; i++) {
  divs[i].onclick = function(e) {  //HERE
    e.target.style.backgroundColor = bgChange();
  }
}
```

```js
e.target
e.preventDefault()
```



### || Prevent event

> Form Example

```js
const form = document.querySelector('form');
const fname = document.getElementById('fname');
const lname = document.getElementById('lname');
const submit = document.getElementById('submit');
const para = document.querySelector('p');

form.onsubmit = function(e) {
  if (fname.value === '' || lname.value === '') { //Too weak
    e.preventDefault();		//HERE
    para.textContent = 'You need to fill in both names!';
  }
}
```



### || Event Capturing and Event Bubbling Case Study

> Problem Example : [show-video-box.html](https://mdn.github.io/learning-area/javascript/building-blocks/events/show-video-box.html)

```js
Use stopPropagation() to solve
------------------------------
videoBox.onclick = function() {
  videoBox.setAttribute('class', 'hidden');
};

video.onclick = function(e) {
  e.stopPropagation(); //HERE
  video.play();
}
```





## Object

```js
person = {
    name: "hello",
    interest: ['music', 'skiing'],
    interest2: { a: 'music', b: 'skiing'},
    bio: function() {alert(this.name + " hello")} // this is pointing to the object that run this function
}

-------
console
-------
person.name or person['name']
person.interest[0] or person['interest'][0]
person.interest.a or person['interest']['a']
person.bio()
// By using person[x], the benefit is x can be a variable

------
Set
------
person.name = 'byebye'
person.bio = function() {..}
```

### || Class

> Example of Class

```js
class Professor extends Person { // can extendss, like Java

  teaches;

  constructor(name, teaches) { // different with Java, constructor instead of public Professor(...)
    super(name);
    this.teaches = teaches;
  }

  introduceSelf() {
    console.log(`My name is ${this.name}, and I will be your ${this.teaches} professor.`);
  }

  grade(paper) {
    const grade = Math.floor(Math.random() * (5 - 1) + 1);
    console.log(grade);
  }

}
```

> Private field / function

```js
class Test {
	#teachers; //when use in class : this.#teachers = ...

    #privateMethod() {....}
}
```



### || Prototype

> Every object has one prototype object, as a template. It has the functions and fields of itself which to be inherited. One object will inherit these template layer by layer, this kind of relationship called -- prototype chain. These field and functions are inside their prototype property
>
> ```js
> .prototype // template to its children
> .__proto__ // upper level
> ```
>
> 
>
> ```js
> class TestClass{...}
> let a = new TestClass();
> -----
> when
> -----
> TestClass.prototype.x = function(){}
> 
> a.x() // no problem at all	
> ```

Example 1

```js
function doSomething(){}
doSomething.prototype.foo = "bar";// add a property onto the prototype

var doSomeInstancing = new doSomething(); //please be noted doSomething() will be invoked
doSomeInstancing.prop = "some value";
console.log("doSomeInstancing.prop:      " + doSomeInstancing.prop); // ".... some value"
console.log("doSomeInstancing.foo:       " + doSomeInstancing.foo); // ".... bar"
console.log("doSomething.prop:           " + doSomething.prop); // ".. undefined"
console.log("doSomething.foo:            " + doSomething.foo); // ".. undefined"
console.log("doSomething.prototype.prop: " + doSomething.prototype.prop); // ".. undefined"
console.log("doSomething.prototype.foo:  " + doSomething.prototype.foo); // "bar"
```

Tips: Object.prototype.watch() , ...valueOf() will be inherit, however, Object.is(), Object.keys() cannot be inherit. let c = Object.create(b) <= c.prototype is b



## Json

> like Object, but not fully
>
> - Only field, no function
> - must be " " instead of ' '
> - Json can be array, which mean [{...} , {....}]

```js
JSON.parse(string) // convert string to Json
JSON.stringify(Json) //conver Json to string
```



## Asynchronized

### || XMLHttpRequest.onload()

> can also use eventlistener to start new thread:: 
>
> ```js
> const xhr = new XMLHttpRequest();
> xhr.addEventListener('loadend', () => {
>     log.textContent = `${log.textContent}完成！状态码：${xhr.status}`;
> });
> xhr.open('GET', 'https://raw.githubusercontent.com/mdn/content/main/files/en-us/_wikihistory.json');
> xhr.send();
> log.textContent = `${log.textContent}请求已发起\n`;});
> ```
>
> 



### || Promise *

> Is an object show current the thread status : Promise {` <state>`: "pending" }
>
> - pending (return pending hen using fetch())
> - fulfilled (promise completed, then() is called)
> - rejected (promise failed, catch() is called)
>
> Tips: settled is fulfilled or rejected

```js
const fetchPromise = fetch('https://mdn.github.io/learning-area/javascript/apis/fetching-data/can-store/products.json');

console.log(fetchPromise); // Promise { <state>: "pending" }

fetchPromise.then( response => {
  console.log(`已收到响应：${response.status}`);
});

console.log("已发送请求……");
```

Promise Chain with Catch Exception

```js
const fetchPromise = fetch('https://mdn.github.io/learning-area/javascript/apis/fetching-data/can-store/products.json');

fetchPromise
  .then( response => {
    if (!response.ok) {
      throw new Error(`HTTP 请求错误：${response.status}`);
    }
    return response.json();
  })
  .then( json => {
    console.log(json[0].name);
  })
  .catch( error => {
    console.error(`无法获取产品列表：${error}`);
  });
```

Multi Promise

> .then() is called only all response is fulfilled. 
>
> .catch() is called if any responses is rejected

```js
Promise.all([fetchPromise1, fetchPromise2, fetchPromise3])
  .then( responses => {
    for (const response of responses) { //collection
      console.log(`${response.url}：${response.status}`);
    }
  })
  .catch( error => {
    console.error(`获取失败：${error}`)
  });
```

> .then() is called if any response is fulfilled.
>
> .catch() is called if all responses are rejected

```js
Promise.any([fetchPromise1, fetchPromise2, fetchPromise3])
  .then( response => {
    console.log(`${response.url}：${response.status}`);
  })
  .catch( error => {
    console.error(`获取失败：${error}`)
  });
```

Bad example:

```js
function doStep1(init, callback) {
  const result = init + 1;
  callback(result);
}
function doStep2(init, callback) {
  const result = init + 2;
  callback(result);
}
function doStep3(init, callback) {
  const result = init + 3;
  callback(result);
}
function doOperation() {
  doStep1(0, result1 => {
    doStep2(result1, result2 => {
      doStep3(result2, result3 => {
        console.log(`结果：${result3}`);
      });
    });
  });
}
doOperation();
```





### || Promise Constructor

> use to wrap the functions that is not returning Promise
>
> - resolve <= is a function called when it is fulfilled
> - reject <= is a function called when it is rejected, or an error caught

```js
function alarm(person, delay) {
  return new Promise((resolve, reject) => {
    if (delay < 0) {
      throw new Error('Alarm delay must not be negative');
    }
    window.setTimeout(() => {
      resolve(`Wake up, ${person}!`); 
    }, delay);
  });
}

button.addEventListener('click', () => {
  alarm(name.value, delay.value)
    .then(message => output.textContent = message) //Wakeup, ...!
    .catch(error => output.textContent = `Couldn't set alarm: ${error}`);
});
--------------------------
async and await
--------------------------
button.addEventListener('click', async () => { // async
  try {
    const message = await alarm(name.value, delay.value);
    output.textContent = message;
  }
  catch (error) {
    output.textContent = `Couldn't set alarm: ${error}`;
  }
});
```





### || Async and await *

> Much more easier. asynchronized function
>
> It return a promise

```js
async function fetchProducts() {
  try {
    // 在这一行之后，我们的函数将等待 `fetch()` 调用完成
    // 调用 `fetch()` 将返回一个“响应”或抛出一个错误
    const response = await fetch('https://mdn.github.io/learning-area/javascript/apis/fetching-data/can-store/products.json');
    if (!response.ok) {
      throw new Error(`HTTP 请求错误：${response.status}`);
    }
    // 在这一行之后，我们的函数将等待 `response.json()` 的调用完成
    // `response.json()` 调用将返回 JSON 对象或抛出一个错误
    const json = await response.json();
    return json;
  }
  catch(error) {
    console.error(`无法获取产品列表：${error}`);
  }
}

//const json = fetchProducts(); <== WRONG!
//console.log(json[0].name);   // json 是一个 Promise 对象，因此这句代码无法正常工作
const jsonPromise = fetchProducts();
jsonPromise.then((json) => console.log(json[0].name));
```



### || Workers

> Main.js and worker.js shall be seperated, worker shall not access DOM as there is the risk that multiples threads modify Main js variable.

Types:

- dedicated workers
- shared workers
- service workers

Example (of dedicated worker)

main.js

```js
const worker = new Worker('./generate.js');

document.querySelector('#generate').addEventListener('click', () => {
  const quota = document.querySelector('#quota').value;
  worker.postMessage({ // post one message to worker
    command: 'generate', // message.data.command = 'generate'
    quota: quota	//message.data.quota = quota
  });
});

worker.addEventListener('message', message => { // when worker return a message
  /// message.data = 
  document.querySelector('#output').textContent = `Finished generating ${message.data} primes!`;
});

document.querySelector('#reload').addEventListener('click', () => {
  document.querySelector('#user-input').value = 'Try typing in here immediately after pressing "Generate primes"';
  document.location.reload();
});
```

generate.js

```js
addEventListener("message", message => { // listening
  if (message.data.command === 'generate') { 
    generatePrimes(message.data.quota);
  }
});

// your function
function generatePrimes(quota) { 
    ...
	postMessage(result);
}
```



## Web API

### || DOM

Refer Above: Useful Utils Method

> Able to change HTML and CSS using document.command
>
> Able to get Window informations.

![image-20220728215238125](Images/image-20220728215238125.png)

Picture refer to [Mozilla](https://developer.mozilla.org/zh-CN/docs/Learn/JavaScript/Client-side_web_APIs/Manipulating_documents), demostrating the document, window, and navigator (navigator => to get user infor such as prefer language, geography location ...)

### ||  AJAX - Asynchronized javascript and XML

> Asynchronized to get and send data between client and server
>
> 'AJAX' term for the technique is still use although currently we use Json as a result to transfer data between client and server instead of XML.

#### 1. XMLHttpRequest (XHR)

```js
var requestURL = 'https://mdn.github.io/learning-area/javascript/oojs/json/superheroes.json';
var request = new XMLHttpRequest();
request.open('GET', requestURL);
request.responseType = 'json'; // | text => JSON.parse(..)
request.send();

request.onload = function(){
  var superHeroes = request.response; // get the string from the response
  populateHeader(superHeroes);
  showHeroes(superHeroes);
}
```

XHR Ready States

- 0 => unsent
- 1 => opened
- 2 => headers_received (send() method was called)
- 3 => loading
- 4 => done

#### 2. Fetch API

```js
fetch(url).then(function(response) { 
  response.text().then(function(text) { // .text() as same meaning in XHR.reponseType='text'
    poemDisplay.textContent = text;
  });
});

--
or
--
response.text()
response.json()
response.blob() // 二进制大对象 eg video / img
---
fetch(url).then(function(response) { // same as .then( response => )
  if (response.ok) {return response.text()} 
  else {console.log('Network request failed with response ' + response.status + ': ' + response.statusText);}
}).then(function(text) {
  poemDisplay.textContent = text;
});
```

Tips: to create ImageURL

```js
fetch(url)
      .then( response => {
        if (!response.ok) {
          throw new Error(`HTTP error: ${response.status}`);
        }
        return response.blob();
      })
      .then( blob => {const objectURL = URL.createObjectURL(blob); ...} )
      .catch( err => console.error(`Fetch problem: ${err.message}`) );
```



### ||  Third Party API

#### Google Map Javascript API

> Below example is deprecated. kindly search online for new version

```js
navigator.geolocation.getCurrentPosition(success, error, options)
```

```js
navigator.geolocation.getCurrentPosition(function(position) { // get user position
  var latlng = new google.maps.LatLng(position.coords.latitude,position.coords.longitude);
  var myOptions = {
    zoom: 8,
    center: latlng,
    mapTypeId: google.maps.MapTypeId.TERRAIN,
    disableDefaultUI: true
  }
  var map = new google.maps.Map(document.querySelector("#map_canvas"), myOptions);
});
```



Another type of third party API is RESTful API, which send the request to server by using xhr or fetch(), and server return the data.



### || 3D API

> WebGL API. 
>
> Js library  to assist: Three.js, PlayCanvas, Babylon.js



### || Web Storage API

> Session Storage 
>
> - last as long as browser / tab is open
>
> Cookie Storage
>
> - no expiration time until cleared through js, or clearing the browser cache

```js
localStorage / sessionStorage
.key(index eg 1, 2, 3) // return the key name, if none => null
.getItem('key')
.setItem('key', 'value')
.removeItem('key')
.clear()
```



### || IndexedDB

> Save the complicated data -- Detail refer https://developer.mozilla.org/zh-CN/docs/Learn/JavaScript/Client-side_web_APIs/Client-side_storage



### || Cache API and Service Worker

> Store responses from server in client side, so that user can visit website without network 
>
> Detail refer to https://developer.mozilla.org/zh-CN/docs/Learn/JavaScript/Client-side_web_APIs/Client-side_storage also



## Reference

1. [Balls Collision Game](https://roy-tian.github.io/learning-area/javascript/oojs/assessment/)

   > Knowledge:
   >
   > - Extends using:
   >
   > ```js
   > // 定义 Ball 构造器，继承自 Shape
   > 
   > function Ball(x, y, velX, velY, exists, color, size) {
   >   Shape.call(this, x, y, velX, velY, exists); //用来调用所有者对象作为参数的方法
   > 
   >   this.color = color;
   >   this.size = size;
   > }
   > 
   > 
   > Ball.prototype = Object.create(Shape.prototype); //用于创建一个新对象，使用现有的对象来作为新创建对象的原型（prototype）。
   > Ball.prototype.constructor = Ball;
   > ```
   >
   > - canvas API

2. [Rotating Alice](https://developer.mozilla.org/zh-CN/docs/Learn/JavaScript/Asynchronous/Sequencing_animations)

   > Knowledge:
   >
   > - Web Animations API
   >
   >   ```js
   >   const aliceTumbling = [
   >     { transform: 'rotate(0) scale(1)' }, //from
   >     { transform: 'rotate(360deg) scale(0)' } //to
   >   ];
   >         
   >   const aliceTiming = {
   >     duration: 2000,
   >     iterations: 1, // for 1 time
   >     fill: 'forwards' // after 2000ms, the fill is the result, not the begin image size
   >   }
   >         
   >   const alice1 = document.querySelector("#alice1");
   >         
   >   alice1.animate(aliceTumbling, aliceTiming); // alice will rotate and miss in 2000ms
   >   ```
   >
   >   

3. Set up a local test server

> Refer [Mozilla](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/set_up_a_local_testing_server)

Below command is for my reference only ==> port 8000

```shell
python -m http.server
```





# Form

> Guide on web form usability -- https://www.smashingmagazine.com/2011/11/extensive-guide-web-form-usability/
>
> .... more articles on designing a good form for user, refer https://developer.mozilla.org/zh-CN/docs/Learn/Forms/Your_first_form for articles links.

```html
attributes
----------
action = "url"
method = "post | get"
enctype = "multipart/form-data" <!-- what kind of data, default: application/x-www-form-urlencoded -->
novalidate <!-- no validate message offered by browser, this was used inorder to have our customized message --> 

-----------
sub element
-----------

=================================================================================
fieldset
--------
<!-- Q. When to use? -->
<!-- 1. multiple choice question (radio / checkboxes) -->
<!-- 2. several question relating to the same topic -->
disable <!-- disable the field inside this fieldset -->

legend <!-- a title of the fieldset -->
------


label <!-- important, screen reader to read content-->
------
for = "say" <!-- <label for="say">What greeting do you want to say?</label> ==> this for maps the component(eg input)'s id, which allow user to click on label to focus on the component-->
<!-- <input name="say" id="say" value="Hi"> -->

input !!no end tag  <!-- refer me https://developer.mozilla.org/zh-CN/docs/Web/HTML/Element/Input -->
-----
name = '..' <!-- key -->
id = '..'
value = '..'
type = ".." 
<!-- file -->
<!-- text (default), email (only email format), password, number, tel (will pop out number keyboard in some device), url (ditto but for url), searcg -->
<!-- color, date, datetime-local, month, time, week -->
<!-- image (same action with submit, this will send server the pos.x and pos.y from this top left of img), submit, reset, hidden, button, checkbox, radio (shall have same name and different id and value, ditto to checkbox) --> 
<!-- range -->
multiple <!-- for file / email <== seperate by ,-->
accept = '..' <!-- acceot what kind of format , for file , eg ".jpg,.png,.doc", audio/*, video/*, image.*-->
autofocus <!-- when the page load, focus here -->
checked <!-- radio, checkbox  has pseudo class of :default eventhough is diselect after that-->
disabled
max = '2013-06-01' <!-- for num / date related -->
min = '...' <!-- for num / date related -->
maxlength = '..'
minlength = '..'
pattern = '..' <!-- password、text、tel, eg [0-9]{3} -->
placeholder = '..' <!-- show when the component is empty, is not a value -->
readonly
required
step = '1' <!-- default: 1, applicable to number only. the step for spinner buttons to incr or desc. if intend to have float, then eg step='0.01'. -->
list = 'datalist's id'

Autocomplete box
----------------
datalist
--------
id = 'input list attb'

option
------


=================================================================================
textArea
--------
id = '..'
name = '..'
cols = '..' <!-- the width of text control, default 20 -->
rows = '..' <!-- height, represented by row in textArea, default 2-->
wrap = 'soft | hard | none' <!-- default soft. none -> no line break, became horizontally scrollable. -->
resize = 'both | horizontal | vertical | none' <!-- able to resize -->

=================================================================================
select
------
id = '..'
name = '..'
multiple <!-- be noted that the select box will not be displayed as dropped down, you should set size attb also -->
size ='5' <!-- row to display -->

optgroup !!no end tag  <!-- This is not selectable in the selections, how ever it display the title of the selections --> 
--------
label='fruits'

option
------
value = ".." <!-- if no value is provided, the content is used as the value -->
selected


=================================================================================
button
------
type = ".." <!-- submit (default), reset, button (for js) -->
<!-- different between button & input type="submit" => button is allowed to have html content -->
```

> Example of request using get
>
> ```
> GET /?say=Hi&to=Mom HTTP/2.0
> Host: foo.com  <= the website your visitting
> ```
>
> 
>
> Example of request using post
>
> ```
> POST / HTTP/2.0
> Host: foo.com
> Content-Type: application/x-www-form-urlencoded
> Content-Length: 13
> 
> say=Hi&to=Mom
> ```



**Example of Field Set and Legend**

```html
form>
  <fieldset>
    <legend>Choose your favorite monster</legend>

    <input type="radio" id="kraken" name="monster">  
    <label for="kraken">Kraken</label><br/> <!-- reader will read "Choose your favorite monster. Kraken, radio button"
	...
  </fieldset>
</form>

```

**Example of Input with Required Tips**

```html
<div>
  <label for="username">Name: <abbr title="required">*</abbr></label>
  <input id="username" type="text" name="username">
</div>
```

**Example of Accept**

```html
<input type="file" name="file" id="file" accept="image/*" multiple>
<input type="file" accept="image/*;capture=camera">
<input type="file" accept="video/*;capture=camcorder">
<input type="file" accept="audio/*;capture=microphone">
```

**Example of Autocomplete box**

```html
<label for="myFruit">What's your favorite fruit?</label>
<input type="text" name="myFruit" id="myFruit" list="mySuggestion">
<datalist id="mySuggestion">
  <option>Apple</option>
  <option>Banana</option>
  <option>Blackberry</option>
  <option>Blueberry</option>
  <option>Lemon</option>
  <option>Lychee</option>
  <option>Peach</option>
  <option>Pear</option>
</datalist>
```

**Example of select**

```html
<select id="groups" name="groups">
  <optgroup label="fruits">
    <option>Banana</option>
    <option selected>Cherry</option>
    <option>Lemon</option>
  </optgroup>
  <optgroup label="vegetables">
    <option>Carrot</option>
    <option>Eggplant</option>
    <option>Potato</option>
  </optgroup>
</select>
```

**Example of showing value of type='range'**

html

```html
<input type="range" name="price" id="price" min="50000" max="500000" step="100" value="250000">
<output class="price-output" for="price"></output
```

js

```js
const price = document.querySelector('#price');
const output = document.querySelector('.price-output');

output.textContent = price.value;

price.addEventListener('input', function() {
  output.textContent = price.value;
});
```



**Solving some form component does not inherit the font family, does not consistent in size**

```css
button, input, select, textarea {
  font-family: inherit;
  font-size: 100%;
}
input, textarea, select, button {
  width : 150px;
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}
```

**Legend Position**

```css
fieldset {
  position: relative;
}

legend {
  position: absolute;
  bottom: 0;
  right: 0;
}
```

**Example of Custom Validity Message**

```js
const email = document.getElementById("mail");

email.addEventListener("input", function (event) {
  if (email.validity.typeMismatch) {
    email.setCustomValidity("I am expecting an e-mail address!");
    email.reportValidity();
  } else {
    email.setCustomValidity("");
  }
});
```





## Regular Expression

> Please refer : https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Regular_Expressions

javascript

```js
const re = /ab+c/;
or
const re = new RegExp('ab+c');
```

pattern in input tag

```html
<input id="choose" name="i_like" required pattern="[Bb]anana|[Cc]herry">
```

Tips:

\* <= no or more

? <= no or 1

## Safety

### || XSS - Cross Site Scripting

> register a "script" which run a script to send your cookies, localStorage ... data to hacker server when you visited this registration, and

### || CSRF - Cross-Site Request Forgery

> sending a request using victim authorization. Eg. when the victim opens a vulnerable website and then a malicious website which send a post request to the vulnerable website.

### || SQL Injection

> happened in database, which hacker play tricks on the input and get the information of other user

### ...



**What server can do?**

- Delete `<script>` `<iframe>` `<object>` when processing data.
- Filter data from client, never trust user





# Vue

> A JavaScript Framework to make the development much more efficient. But bear in mind that it cost extra space, and other bad side. You should not use framework if the JavaScript is not that much.

## Installation

- Node.js 8.11+ installed
- npm or yarn.
- `npm install --global @vue/cli`

-------------------

### Create Project

Project directory terminal

1. `vue create <project name>`
2. Manually select features <== to config the TypeScript, linting, vue-router, testing ...
3. To select : space;  (select babel, linter)
4. vue 3.0 or 2.0?
5. linter setting: Eslint with error prevention only
6. Lint feature => lint on save
7. Keep config in "dedicated config files"

Tips 

Problem: vue.ps1 cannot be loaded because running scripts is disabled on this system

```shell
Set-ExecutionPolicy -ExecutionPolicy Unrestricted -Scope CurrentUser
```

### Run project

`npm run serve` <== in package.json ： vue-cli-service serve

## Project Structure

![image-20220730211612975](Images/image-20220730211612975.png)

1. .eslintrc.js 	---- eslint config file, manage your coding rules
2. babel.config.js ---- babel config file, degrade the new code to old code in order to run on older browser
3. .browserslistrc ----  Browserslist config, to manage which browser to be supported.
4. public ---- contains files that do not processed by Webpack during build, exclude index.html
5. src ---- core coding of vue
   - main.js --- entry of the application
     - initialize Vue application and link it to index.html's element
     - register the global components here
     - keep other vue libraries.
   - App.vue --- the root component
   - components --- a directory to keep the customized components
   - assets  ---- css, images...

## .vue file standard format

```vue
<template></template>
<script></script>
<style></style>
```

### Template

html here



### Script

JavaScript here

### Style

css.

it can be `style scoped` if you want the style restricted in this component only





## **Export and Import Components**

> Be noted that component ToDoItem.vue is always camel case and start with capital letter, while in App.vue it will become` <to-do-item>` (suggest) or `<ToDoItem>`.

components/ToDoItem.vue

```vue
<template> 
  <div>
    <input type="checkbox" id="todo-item" checked="false" />
    <label for="todo-item">My Todo Item</label>
  </div>
</template>
<script>
  export default {};
</script>
```

App.vue

```vue
<template>
  <div id="app">
    <h1>To-Do List</h1>
    <ul>
      <li><to-do-item></to-do-item></li>
    </ul>
  </div>
</template>

<script>
import ToDoItem from './components/ToDoItem.vue'
export default {
  name: 'App',
  components: {
    ToDoItem
  }
}
</script>

<style>
</style>
```



Export

```vue
<template> 
  <div>
    <input type="checkbox" id="todo-item" checked="false" />
    <label for="todo-item">My Todo Item</label>
  </div>
</template>
<script>
  export default {};
</script>
```

Import

```vue
<script>
import ToDoItem from './components/ToDoItem.vue'
export default {
  name: 'App',
  components: {
    ToDoItem //*****
  }
}
</script>
```





## Props

> From my understanding, with using props:
>
> - able to carried the element attribute from App.vue to the ToDoItem.vue
> - a one-way data binding, component shall never modify its own prop as it will be challenging to debug as the value might be passed to multiple children

Format

```js
export default {
    props: {
        label: {required: true, type: String},
        done: {default: false, type: Boolean},
        id: {required: true, type: String}
    },
}
```

**Example**

components/ToDoItem.vue

> Read from App.vue, and transfer back to App.vu

```vue
<template> 
  <div>
    <input type="checkbox" :id="id"  :checked="isDone" />
    <label for="todo-item">{{label}}</label>
  </div>
</template>
<script>

  export default {
    props: {
        label: {required: true, type: String},
        done: {default: false, type: Boolean},
        id: {required: true, type: String}
    },
    data() {
        return {
            isDone: this.done // a prop value should never change, that's why we use this instead of vbind 'done'.
        };
    }
  };
</script>
```

 App.vue

```vue
<li v-for="item in ToDoItems" :key="item.id"><to-do-item :label="item.label" :done="item.done" :id="item.id"></to-do-item></li>
```



## Data

```vue
data() {
	return {
		key: value,
		key: value
	}
}
```

Tips: 

to call in javascript

`this.key`

to combine the string

```
`hello ${this.key}`
or
'hello' + this.key
```



## Methods

```script
export default {
	methods: {
		onSubmit() {
		
		},
		...
	}
}
```

Tips:

To add

```js
array.push({...:, ..:, ..:})
```

To delete

```js
array = array.filter(item => item !== itemInParam)
array.splice(itemIndex, how many following to delete)
```

To find

```js
array.find(item => ...)
```



## Computed

> perform calculation and return the data.  A computed property will only re-evaluate when some of its reactive dependencies have changed eg list value. This is the biggest difference between computed and a method
>
> Best practice:
>
> - to a list that seldom change

```html
<h1>{{computed data}}</h1>
```

```js
export default {
    computed: {
        return this.lists.length
    }
}
```



## Watch

> watch the data.
>
> if the data changed, then this will be triggered

Example 

```js
export default {
	data() {
        return {
            question: '',
        }
    },
    watch: {
        question(imNewData, imOldData) {
            console(imNewData + imOldData)
        }
        --
        or
        --
        question: 'IamAAsyncMethodToFetchData'
    }
}
```



## Event Modifier

`<form @submit.prevent="onSubmit"` <= prevent redirect and trigger onSubmit

| Vue      | JavaScript                                                   | Description                                                  |
| -------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| .stop    | Event.stopPropagation()                                      | Stop capturing and bubbling phases                           |
| .prevent | Event.preventDefault()                                       |                                                              |
| .self    |                                                              | Triggers the handler only if the event was dispatched from this exact element. |
| {.key}   | .addEventListener('*key*', function listener() { /* do something */ }) | trigger when the key are press eg page-down<br />https://developer.mozilla.org/en-US/docs/Web/API/UI_Events/Keyboard_event_key_values |
| .native  |                                                              | listen for a native event on root element                    |
| .once    |                                                              | only trigger one time                                        |
| .left    |                                                              | mouse left click                                             |
| .right   |                                                              | mouse right click                                            |
| .middle  |                                                              | mouse middle click                                           |
| .passive | .addEventListener('touchmove', function listener() { /* do something */ }, { passive: true }) | when true, preventDefault() will be never called (will warn in console if do so) => it can solve touchmove |



## $emit

> to emit method from the component to another component.
>
> 1st argument : method name
>
> 2nd ++ argument : datas to emit

Tips: $event is used when your emitted methods has its own param : @from-component="mymethod(ownparam, $event)"

Method1

```js
export default {
    methods: {
        onSubmit() {
            this.$emit('hello', "nihao")
        }
    }
}
--------------------------------------------------
to
-------------------------------------------------
<that-component @hello="(msg) => yourData = msg"></that-component>

```





Method 2

```js
export default {
    methods: {
        onSubmit() {
            this.$emit('hello', 1,2,3)
        }
    }
}
--------------------------------------------------
to
-------------------------------------------------
<that-component @hello="hellohere"></that-component>
...

export default {
	methods: {
        hellohere(no1, no2, no3) {
            console.log(no1 + no2 + no3)
        }
    }
}
```



Method 3

```vue
export default {
    methods: {
        onSubmit() {
            this.$emit('hello', abc)
        }
    }
}
--------------------------------------------------
to
-------------------------------------------------
<that-component @hello="hellohere(item.myid, $event)"></that-component>
...

export default {
	methods: {
        hellohere(myId, abc) {
            console.log(no1 + no2 + no3)
        }
    }
}
```

## $refs and ref

> `ref` to register this element to component $ref.
>
> we can do something like focus() with this
>
> **Be careful that there is some delay in reflecting in DOM. Please consider play this with $nextTick or mounted() if the DOM is affected** 

Tips: 

`<input ref="element1">`

when calling in methods:

`const reference = this.$refs.element1`

`reference.focus()`



## this.$nextTick(function)

> the function will be execute when the DOM finished updating

Example 

focusing to a ref

```js
focusOnEditButton() {
    this.$nextTick(() => {
        const editButtonRef = this.$refs.editButton
        editButtonRef.focus()
    })
}
```



## Using CSS

### External CSS File

- put your css in src/assets/

- import in src/main.js

  ```js
  import './assets/reset.css'
  ```

### Global Style to Single File Components

- put your css in the .vue

- this style can be applied to imported components

  ```vue
  <style>
  	...
  </style>
  ```

### Style in only component

```css
<style scoped>
	...
</style>
```



## Lifecycles

| Stage           | Description                                          |
| --------------- | ---------------------------------------------------- |
| beforeCreate()  | data and event yet ready                             |
| create()        | component is initialized. the VDOM yet prepared      |
| beforeMount()   | templated is compiled but yet rendered to actual DOM |
| mounted()       | component is mounted to DOM, can access refs here    |
| beforeUpdate()  | data in component is updated but yet rendered to DOM |
| updated()       |                                                      |
| beforeDestroy() | before components is removed from DOM                |
| destroyed()     |                                                      |
| activated()     |                                                      |
| deactivated()   |                                                      |

Example to use mounted with $ref

```js
export default {
	mounted() {
		const labelInput = this.$refs.labelInput
    	labelInput.focus()
	}
}
```



## v-

###  **v-bind** 

```vue
v-bind:attribute="expression"
:attribute="expression"
```

> `one way binding -- display text` (shortcut :) 

```html
<div id="app">
    <h1 v-bind:title="message"> 
        {{content || "no data here"}} <!-- no data here if there is no message here -->
    </h1>
</div>
<script src="vue.min.js"></script>
<script>
    new Vue({
        el: '#app',
        data: {
            content: 'im title',
            message: 'page loded on ' + new Date().toLocaleString()
        }
    })
</script>
```

![image-20220627200908071](Images/image-20220627200908071.png)

### v-model 

> For all input, includes checkboxes, radio, select
>
> v-model.trim="hello"
>
> v-model.lazy.trim="hello" <= lazy means only reflect once the input lost focus or submitted, as it is no need to bind when every single key down

>  `two ways binding -- when input will reflect to other v-bind `

```html
<div id="app">
    <input type="text" v-bind:value="searchMap.keyWord"/>
    <input type="text" v-model="searchMap.keyWord"/>
    <input type="text" v-model="searchMap.keyWord"/>
    <h1>{{searchMap.keyWord}}</h1>
</div>
<script src="vue.min.js"></script>
<script>
    new Vue({
        el: '#app',
        data: {
            searchMap: {
                keyWord: '尚硅谷'
            }
        }
    })
</script>
```

_below example: change in middle affected all_

![image-20220627202014302](Images/image-20220627202014302.png)



### **v-on:click **/ @click

```vue
v-on:argument.modifiers="value"
@submit.prevent="OnSubmit"
```

example:

@submit

@change (like for checkboxes...)

@click

> `bind event`

```html
<div id="app">
    <button v-on:click="search()">search</button> #"search" also can but will crash with data if have same name
    <button @click="f1()">f1</button>
</div>
<script src="vue.min.js"></script>
<script>
    new Vue({
        el: '#app',
        methods: {
            search() {
                console.log('search...')
            },
            f1() {
                console.log("f1...")
            }
        }
    })
</script>
```

![image-20220627203259999](Images/image-20220627203259999.png)



### v-if ***

> IF the condition is not met, then this element will not shown
>
> 懒加载， 不适合频繁切换，而v-show是不管初始条件如何都会被渲染

```html
<div id="app">
    <input type="checkbox" v-model="ok">test</input>
<h1 v-if="ok">Yes</h1>
<h1 v-else>No</h1>
</div>
<script src="vue.min.js"></script>
<script>
    new Vue({
        el: '#app',
        data: {
            ok: true
        }
    })
</script>
```

![image-20220627205746118](Images/image-20220627205746118.png)![image-20220627205802710](Images/image-20220627205802710.png)

### v-show

```html
<h1 v-show="ok">Yes</h1>
<h1 v-show="!ok">No</h1>
```



### v-for ***

> There shall be key attribute, this shall be unique id. it is used so that there is no need to recreate the element when if it is the same element (it means they are re-useable),  whenever the list changed

#### x in Number

```html
<ul>
    <li v-for="n in 5">{{n}}</li>
</ul>
<ol>
    <li v-for="(n, index) in 5">{{n}} -- {{index}}</li>
</ol>
```

![image-20220627210442582](Images/image-20220627210442582.png)



#### x in yourObjects or computedObject

```html
<div id="app">
    <table border="1">
        <tr v-for="(user, index) in userlist">
            <td>{{index}}</td>
            <td>{{user.id}}</td>
            <td>{{user.username}}</td>
            <td>{{user.age}}</td>
        </tr>
    </table>
</div>
<script src="vue.min.js"></script>
<script>
    new Vue({
        el: '#app',
        data: {
            userlist: [
                {id: 1, username: 'helen', age: 18},
                {id: 2, username: 'ben', age: 19},
                {id: 3, username: 'tom', age: 20}
            ]
        }
    })
</script>
```



![image-20220627211002953](Images/image-20220627211002953.png)



## Router

 ```vue
<div id="app">
    <p>
        <router-link to="/">index</router-link>
        <router-link to="/student">student</router-link>
        <router-link to="/teacher">teacher</router-link>
    </p>
    <router-view></router-view>
</div>
<script src="vue.min.js"></script>
<script src="vue-router.js"></script>
<script>
    const Welcome = { template: '<div>Welcome</div>' }
    const Student = { template: '<div>Student</div>' }
    const Teacher = { template: '<div>Teacher</div>' }

    const routes = [
        { path: '/', redirect: '/welcome' },
        { path: '/welcome', component: Welcome },
        { path: '/student', component: Student },
        { path: '/teacher', component: Teacher }
    ]

    const router = new VueRouter({
        routes //eq to routes: routes
    })

    const app = new Vue({
        el: '#app',
        router
    })
</script>
 ```

![image-20220627224635545](Images/image-20220627224635545.png)



## Other Useful Components

- lodash

  > to generate a unique id by using uniqueId()
  >
  > `npm install --save lodash.uniqueid`

  ```js
  
  import uniqueId from 'lodash.uniqueid'
  export default {
    data() {
      return {
        id: uniqueId('todo-') // 
      };
    }
  };
  ```

  
