# 初学问题

* 为什么单调队列在生成中被压缩，最后取数据时却不会数量不够
    1. [https://labuladong.online/algo/data-structure/monotonic-queue/#%E4%B8%80%E3%80%81%E6%90%AD%E5%BB%BA%E8%A7%A3%E9%A2%98%E6%A1%86%E6%9E%B6](https://labuladong.online/algo/data-structure/monotonic-queue/#%E4%B8%80%E3%80%81%E6%90%AD%E5%BB%BA%E8%A7%A3%E9%A2%98%E6%A1%86%E6%9E%B6)
    2. 因为在Queue移除元素的时候，并非直接移除，而是先比较当前元素和Queue中的首个元素是否相等，再决定移除。所以不会影响