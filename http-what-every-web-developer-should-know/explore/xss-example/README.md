Example taken from <https://www.youtube.com/watch?v=oEFPFc36weY>

> Success Attack:

In textarea, we are failed to XSS because `userMessagesList.innerHTML = messageItems;` was sanitized by default

In image, fill in

`invalid.png" onerror="alert('You are hacked!')`

An alert will come out

> To defend

1. Sanitize in both end

2. Avoid use vulnerable third party package
